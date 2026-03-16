package core

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/feymanlee/rongcloud-go/internal/types"
)

const (
	testAppKey    = "test-app-key"
	testAppSecret = "test-app-secret"
)

// testResp is a response struct that implements BaseRespInterface via embedding.
type testResp struct {
	types.BaseResp
	Result string `json:"result,omitempty"`
}

// newTestServer creates an httptest server that validates auth headers and content-type,
// then returns a configurable JSON response.
func newTestServer(t *testing.T, wantContentType string, respCode int, respMsg string) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate auth headers are present
		for _, h := range []string{"App-Key", "Nonce", "Timestamp", "Signature"} {
			if r.Header.Get(h) == "" {
				t.Errorf("missing header %s", h)
				http.Error(w, "missing header", http.StatusBadRequest)
				return
			}
		}

		// Validate App-Key
		if got := r.Header.Get("App-Key"); got != testAppKey {
			t.Errorf("App-Key: expected %q, got %q", testAppKey, got)
		}

		// Validate signature
		nonce := r.Header.Get("Nonce")
		timestamp := r.Header.Get("Timestamp")
		h := sha1.New()
		io.WriteString(h, testAppSecret+nonce+timestamp)
		expectedSig := fmt.Sprintf("%x", h.Sum(nil))
		if got := r.Header.Get("Signature"); got != expectedSig {
			t.Errorf("Signature: expected %q, got %q", expectedSig, got)
		}

		// Validate Content-Type
		if ct := r.Header.Get("Content-Type"); ct != wantContentType {
			t.Errorf("Content-Type: expected %q, got %q", wantContentType, ct)
		}

		w.Header().Set("Content-Type", "application/json")
		resp := map[string]any{"code": respCode}
		if respMsg != "" {
			resp["errorMessage"] = respMsg
		}
		if respCode == 200 {
			resp["result"] = "ok"
		}
		json.NewEncoder(w).Encode(resp)
	}))
}

func newTestClient(serverURL string) *client {
	c := NewClient(&Options{
		AppKey:    testAppKey,
		AppSecret: testAppSecret,
		Region:    Region{PrimaryDomain: serverURL, BackupDomain: serverURL},
		Timeout:   5 * time.Second,
	})
	return c.(*client)
}

// --- NewClient tests ---

func TestNewClientOptions(t *testing.T) {
	c := NewClient(&Options{
		AppKey:    "key1",
		AppSecret: "secret1",
		Region:    RegionBeijing,
		Timeout:   3 * time.Second,
	}).(*client)

	if c.appKey != "key1" {
		t.Errorf("appKey: expected 'key1', got %q", c.appKey)
	}
	if c.appSecret != "secret1" {
		t.Errorf("appSecret: expected 'secret1', got %q", c.appSecret)
	}
	if c.primaryDomain != RegionBeijing.PrimaryDomain {
		t.Errorf("primaryDomain: expected %q, got %q", RegionBeijing.PrimaryDomain, c.primaryDomain)
	}
	if c.backupDomain != RegionBeijing.BackupDomain {
		t.Errorf("backupDomain: expected %q, got %q", RegionBeijing.BackupDomain, c.backupDomain)
	}
	if c.currentURI != RegionBeijing.PrimaryDomain {
		t.Errorf("currentURI: expected primary domain, got %q", c.currentURI)
	}
}

func TestNewClientDefaultTimeout(t *testing.T) {
	c := NewClient(&Options{
		AppKey:    "key1",
		AppSecret: "secret1",
		Region:    RegionBeijing,
		Timeout:   0, // should use default
	}).(*client)

	// The resty client should have the default timeout.
	// We can't easily inspect resty's timeout directly, but we verify
	// the client was created without error and the changeURIDuration is set.
	if c.changeURIDuration != defaultChangeURIDuration {
		t.Errorf("changeURIDuration: expected %v, got %v", defaultChangeURIDuration, c.changeURIDuration)
	}
}

// --- Post (form-encoded) tests ---

func TestPostFormSuccess(t *testing.T) {
	var receivedBody string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate Content-Type
		if ct := r.Header.Get("Content-Type"); ct != "application/x-www-form-urlencoded" {
			t.Errorf("Content-Type: expected application/x-www-form-urlencoded, got %q", ct)
		}
		// Validate auth headers
		for _, h := range []string{"App-Key", "Nonce", "Timestamp", "Signature"} {
			if r.Header.Get(h) == "" {
				t.Errorf("missing header %s", h)
			}
		}
		// Read form data
		r.ParseForm()
		receivedBody = r.FormValue("userId")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"code": 200, "result": "ok"})
	}))
	defer server.Close()

	c := newTestClient(server.URL)
	var resp testResp
	err := c.Post("/test", map[string]string{"userId": "user123"}, &resp)
	if err != nil {
		t.Fatalf("Post returned error: %v", err)
	}
	if receivedBody != "user123" {
		t.Errorf("form data userId: expected 'user123', got %q", receivedBody)
	}
	if resp.Code != 200 {
		t.Errorf("resp.Code: expected 200, got %d", resp.Code)
	}
	if resp.Result != "ok" {
		t.Errorf("resp.Result: expected 'ok', got %q", resp.Result)
	}
}

func TestPostFormSignature(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nonce := r.Header.Get("Nonce")
		timestamp := r.Header.Get("Timestamp")
		sig := r.Header.Get("Signature")

		h := sha1.New()
		io.WriteString(h, testAppSecret+nonce+timestamp)
		expected := fmt.Sprintf("%x", h.Sum(nil))

		if sig != expected {
			t.Errorf("Signature mismatch: expected %q, got %q", expected, sig)
			http.Error(w, "bad sig", http.StatusUnauthorized)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"code": 200})
	}))
	defer server.Close()

	c := newTestClient(server.URL)
	var resp testResp
	err := c.Post("/test", nil, &resp)
	if err != nil {
		t.Fatalf("Post returned error: %v", err)
	}
}

func TestPostFormErrorResponse(t *testing.T) {
	server := newTestServer(t, "application/x-www-form-urlencoded", 1001, "test error")
	defer server.Close()

	c := newTestClient(server.URL)
	var resp testResp
	err := c.Post("/test", nil, &resp)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	apiErr, ok := err.(Error)
	if !ok {
		t.Fatalf("expected Error type, got %T", err)
	}
	if apiErr.Code() != 1001 {
		t.Errorf("error code: expected 1001, got %d", apiErr.Code())
	}
	if apiErr.Message() != "test error" {
		t.Errorf("error message: expected 'test error', got %q", apiErr.Message())
	}
}

// --- PostJSON tests ---

func TestPostJSONSuccess(t *testing.T) {
	type reqBody struct {
		UserID string `json:"userId"`
		Name   string `json:"name"`
	}

	var received reqBody
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate Content-Type
		if ct := r.Header.Get("Content-Type"); ct != "application/json" {
			t.Errorf("Content-Type: expected application/json, got %q", ct)
		}
		// Validate auth headers
		for _, h := range []string{"App-Key", "Nonce", "Timestamp", "Signature"} {
			if r.Header.Get(h) == "" {
				t.Errorf("missing header %s", h)
			}
		}
		// Read JSON body
		body, _ := io.ReadAll(r.Body)
		json.Unmarshal(body, &received)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"code": 200, "result": "ok"})
	}))
	defer server.Close()

	c := newTestClient(server.URL)
	var resp testResp
	err := c.PostJSON("/test", reqBody{UserID: "u1", Name: "Alice"}, &resp)
	if err != nil {
		t.Fatalf("PostJSON returned error: %v", err)
	}
	if received.UserID != "u1" {
		t.Errorf("body userId: expected 'u1', got %q", received.UserID)
	}
	if received.Name != "Alice" {
		t.Errorf("body name: expected 'Alice', got %q", received.Name)
	}
	if resp.Code != 200 {
		t.Errorf("resp.Code: expected 200, got %d", resp.Code)
	}
	if resp.Result != "ok" {
		t.Errorf("resp.Result: expected 'ok', got %q", resp.Result)
	}
}

func TestPostJSONErrorResponse(t *testing.T) {
	server := newTestServer(t, "application/json", 1001, "json error")
	defer server.Close()

	c := newTestClient(server.URL)
	var resp testResp
	err := c.PostJSON("/test", map[string]string{"key": "val"}, &resp)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	apiErr, ok := err.(Error)
	if !ok {
		t.Fatalf("expected Error type, got %T", err)
	}
	if apiErr.Code() != 1001 {
		t.Errorf("error code: expected 1001, got %d", apiErr.Code())
	}
	if apiErr.Message() != "json error" {
		t.Errorf("error message: expected 'json error', got %q", apiErr.Message())
	}
}

// --- checkResp tests ---

func TestCheckRespSuccess(t *testing.T) {
	c := &client{}
	resp := &testResp{}
	resp.Code = 200
	err := c.checkResp(resp)
	if err != nil {
		t.Errorf("expected nil error for code 200, got %v", err)
	}
}

func TestCheckRespError(t *testing.T) {
	c := &client{}
	resp := &testResp{}
	resp.Code = 1001
	resp.ErrorMessage = "bad request"

	err := c.checkResp(resp)
	if err == nil {
		t.Fatal("expected error for non-200 code, got nil")
	}

	apiErr, ok := err.(Error)
	if !ok {
		t.Fatalf("expected Error type, got %T", err)
	}
	if apiErr.Code() != 1001 {
		t.Errorf("error code: expected 1001, got %d", apiErr.Code())
	}
	if apiErr.Message() != "bad request" {
		t.Errorf("error message: expected 'bad request', got %q", apiErr.Message())
	}
}

func TestCheckRespNonBaseRespInterface(t *testing.T) {
	c := &client{}
	// Pass a plain struct that does not implement BaseRespInterface
	resp := struct{ Foo string }{Foo: "bar"}
	err := c.checkResp(resp)
	if err != nil {
		t.Errorf("expected nil error for non-BaseRespInterface, got %v", err)
	}
}

// --- changeURI (domain failover) tests ---

func TestChangeURISwitchesToBackup(t *testing.T) {
	c := &client{
		primaryDomain:     "https://primary.example.com",
		backupDomain:      "https://backup.example.com",
		currentURI:        "https://primary.example.com",
		changeURIDuration: 0, // no cooldown for test
	}

	c.changeURI()
	if c.currentURI != "https://backup.example.com" {
		t.Errorf("expected backup domain, got %q", c.currentURI)
	}
}

func TestChangeURISwitchesBackToPrimary(t *testing.T) {
	c := &client{
		primaryDomain:     "https://primary.example.com",
		backupDomain:      "https://backup.example.com",
		currentURI:        "https://backup.example.com",
		changeURIDuration: 0, // no cooldown for test
	}

	c.changeURI()
	if c.currentURI != "https://primary.example.com" {
		t.Errorf("expected primary domain, got %q", c.currentURI)
	}
}

func TestChangeURIRespectsCooldown(t *testing.T) {
	c := &client{
		primaryDomain:     "https://primary.example.com",
		backupDomain:      "https://backup.example.com",
		currentURI:        "https://primary.example.com",
		changeURIDuration: 30 * time.Second,
		lastChangeURITime: time.Now(), // just changed
	}

	c.changeURI()
	// Should NOT have switched because cooldown hasn't elapsed
	if c.currentURI != "https://primary.example.com" {
		t.Errorf("expected primary domain (cooldown not elapsed), got %q", c.currentURI)
	}
}

func TestChangeURIAfterCooldown(t *testing.T) {
	c := &client{
		primaryDomain:     "https://primary.example.com",
		backupDomain:      "https://backup.example.com",
		currentURI:        "https://primary.example.com",
		changeURIDuration: 1 * time.Millisecond,
		lastChangeURITime: time.Now().Add(-1 * time.Second), // well past cooldown
	}

	c.changeURI()
	if c.currentURI != "https://backup.example.com" {
		t.Errorf("expected backup domain after cooldown, got %q", c.currentURI)
	}
}

func TestPostTriggersFailoverOnServerError(t *testing.T) {
	// Primary server returns 500
	primary := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{"code": 500, "errorMessage": "server error"})
	}))
	defer primary.Close()

	// Backup server returns 200
	backup := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"code": 200, "result": "from-backup"})
	}))
	defer backup.Close()

	c := NewClient(&Options{
		AppKey:    testAppKey,
		AppSecret: testAppSecret,
		Region:    Region{PrimaryDomain: primary.URL, BackupDomain: backup.URL},
		Timeout:   5 * time.Second,
	}).(*client)
	c.changeURIDuration = 0 // no cooldown

	// First request to primary should trigger failover
	var resp1 testResp
	_ = c.Post("/test", nil, &resp1)

	if c.currentURI != backup.URL {
		t.Errorf("expected currentURI to be backup %q, got %q", backup.URL, c.currentURI)
	}

	// Second request should go to backup and succeed
	var resp2 testResp
	err := c.Post("/test", nil, &resp2)
	if err != nil {
		t.Fatalf("expected success from backup, got error: %v", err)
	}
	if resp2.Result != "from-backup" {
		t.Errorf("expected result 'from-backup', got %q", resp2.Result)
	}
}
