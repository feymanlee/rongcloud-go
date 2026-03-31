package embeddedconsole

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
)

func assertEqual(t *testing.T, got, want any) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestGetAccessToken(t *testing.T) {
	var gotMethod string
	var gotContentType string
	var gotReq GetAccessTokenReq
	var decodeErr error

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotContentType = r.Header.Get("Content-Type")

		if err := json.NewDecoder(r.Body).Decode(&gotReq); err != nil {
			decodeErr = err
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(GetAccessTokenResp{
			Code: successCode,
			Data: "https://embed-console.rongcloud.cn/embed?token=test-token",
		})
	}))
	defer server.Close()

	a := &api{
		accessKey: "app-key-1",
		endpoint:  server.URL,
		client:    httptestClient(),
	}

	resp, err := a.GetAccessToken(PageCodeUserManage, 3600)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if decodeErr != nil {
		t.Fatalf("decode request: %v", decodeErr)
	}

	assertEqual(t, gotMethod, http.MethodPost)
	assertEqual(t, gotContentType, "application/json")
	assertEqual(t, gotReq.AccessKey, "app-key-1")
	assertEqual(t, gotReq.PageCode, PageCodeUserManage)
	assertEqual(t, gotReq.UsefulLife, int64(3600))
	assertEqual(t, resp.Code, successCode)
	assertEqual(t, resp.Data, "https://embed-console.rongcloud.cn/embed?token=test-token")
}

func TestGetAccessToken_BusinessError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(GetAccessTokenResp{
			Code: 10001,
		})
	}))
	defer server.Close()

	a := &api{
		accessKey: "app-key-1",
		endpoint:  server.URL,
		client:    httptestClient(),
	}

	_, err := a.GetAccessToken("bad_page", 3600)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	assertEqual(t, err.Error(), "rongcloud: code=10001, message=")
}

func TestGetAccessToken_HTTPError(t *testing.T) {
	a := &api{
		accessKey: "app-key-1",
		endpoint:  "http://127.0.0.1:1",
		client:    httptestClient(),
	}

	_, err := a.GetAccessToken(PageCodeUserManage, 3600)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestGetAccessToken_EmptyAccessKey(t *testing.T) {
	a := &api{
		client: httptestClient(),
	}

	_, err := a.GetAccessToken(PageCodeUserManage, 3600)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	assertEqual(t, err.Error(), "rongcloud: accessKey is required")
}

func TestGetAccessToken_EmptyPageCode(t *testing.T) {
	a := &api{
		accessKey: "app-key-1",
		client:    httptestClient(),
	}

	_, err := a.GetAccessToken("", 3600)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	assertEqual(t, err.Error(), "rongcloud: pageCode is required")
}

func TestGetAccessToken_InvalidUsefulLife(t *testing.T) {
	a := &api{
		accessKey: "app-key-1",
		client:    httptestClient(),
	}

	_, err := a.GetAccessToken(PageCodeUserManage, 0)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	assertEqual(t, err.Error(), "rongcloud: usefulLife must be between 1 and 31536000")
}

func TestNewAPI_DefaultTimeout(t *testing.T) {
	a, ok := NewAPI("app-key-1", 0).(*api)
	if !ok {
		t.Fatal("expected concrete api type")
	}

	assertEqual(t, a.accessKey, "app-key-1")
	assertEqual(t, a.endpoint, defaultEndpoint)
	assertEqual(t, a.client.GetClient().Timeout, defaultTimeout)
}

func httptestClient() *resty.Client {
	return resty.New().SetTimeout(time.Second)
}

func TestGetAccessToken_HTTPStatusWithJSONError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(GetAccessTokenResp{
			Code: 10002,
		})
	}))
	defer server.Close()

	a := &api{
		accessKey: "bad-key",
		endpoint:  server.URL,
		client:    httptestClient(),
	}

	_, err := a.GetAccessToken(PageCodeUserManage, 3600)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	assertEqual(t, err.Error(), "rongcloud: code=10002, message=")
}

func TestGetAccessToken_InvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("not-json"))
	}))
	defer server.Close()

	a := &api{
		accessKey: "app-key-1",
		endpoint:  server.URL,
		client:    httptestClient(),
	}

	_, err := a.GetAccessToken(PageCodeUserManage, 3600)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	var syntaxErr *json.SyntaxError
	if !errors.As(err, &syntaxErr) {
		t.Fatalf("expected json syntax error, got %T", err)
	}
}
