package userprofile

import (
	"errors"
	"testing"

	"github.com/feymanlee/rongcloud-go/internal/testutil"
)

func assertEqual(t *testing.T, got, want any) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func assertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func assertNotNil(t *testing.T, v any) {
	t.Helper()
	if v == nil {
		t.Fatal("expected non-nil value")
	}
}

// ---------------------------------------------------------------------------
// Set
// ---------------------------------------------------------------------------

func TestSet(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	gender := 1
	req := &SetReq{
		UserID: "user1",
		UserProfile: &UserProfile{
			Name:   "Alice",
			Gender: &gender,
		},
		UserExtProfile: map[string]string{
			"ext_city": "Beijing",
		},
	}
	err := a.Set(req)
	assertNoError(t, err)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/user/profile/set.json")
	assertEqual(t, call.Params["userId"], "user1")
	// userProfile should be a JSON string
	if call.Params["userProfile"] == "" {
		t.Error("userProfile should not be empty")
	}
	if call.Params["userExtProfile"] == "" {
		t.Error("userExtProfile should not be empty")
	}
}

func TestSet_OnlyProfile(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	req := &SetReq{
		UserID:      "user1",
		UserProfile: &UserProfile{Name: "Bob"},
	}
	err := a.Set(req)
	assertNoError(t, err)

	call := mock.LastCall()
	if _, ok := call.Params["userExtProfile"]; ok {
		t.Error("userExtProfile should not be set when empty")
	}
}

func TestSet_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("set failed")
	}
	a := NewAPI(mock)

	err := a.Set(&SetReq{UserID: "user1"})
	assertError(t, err)
}

// ---------------------------------------------------------------------------
// Get
// ---------------------------------------------------------------------------

func TestGet(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	req := &GetReq{UserID: "user1", Keys: []string{"name"}}
	resp, err := a.Get(req)
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "PostJSON")
	assertEqual(t, call.Path, "/user/profile/get.json")
	body := call.Body.(*GetReq)
	assertEqual(t, body.UserID, "user1")
	assertEqual(t, len(body.Keys), 1)
}

func TestGet_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostJSONFunc = func(path string, body any, resp any) error {
		return errors.New("get failed")
	}
	a := NewAPI(mock)

	resp, err := a.Get(&GetReq{UserID: "user1"})
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// BatchGet
// ---------------------------------------------------------------------------

func TestBatchGet(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	req := &BatchGetReq{UserIDs: []string{"u1", "u2"}, Keys: []string{"name"}}
	resp, err := a.BatchGet(req)
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "PostJSON")
	assertEqual(t, call.Path, "/user/profile/batch/get.json")
	body := call.Body.(*BatchGetReq)
	assertEqual(t, len(body.UserIDs), 2)
}

func TestBatchGet_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostJSONFunc = func(path string, body any, resp any) error {
		return errors.New("batch get failed")
	}
	a := NewAPI(mock)

	resp, err := a.BatchGet(&BatchGetReq{UserIDs: []string{"u1"}})
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// BatchQuery
// ---------------------------------------------------------------------------

func TestBatchQuery(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	req := &BatchQueryReq{UserIDs: []string{"u1", "u2", "u3"}}
	resp, err := a.BatchQuery(req)
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/user/profile/batch/query.json")
	assertEqual(t, call.Params["userId"], "u1,u2,u3")
}

func TestBatchQuery_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("batch query failed")
	}
	a := NewAPI(mock)

	resp, err := a.BatchQuery(&BatchQueryReq{UserIDs: []string{"u1"}})
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// Clean
// ---------------------------------------------------------------------------

func TestClean(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	err := a.Clean(&CleanReq{UserIDs: []string{"u1", "u2"}})
	assertNoError(t, err)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/user/profile/clean.json")
	assertEqual(t, call.Params["userId"], "u1,u2")
}

func TestClean_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("clean failed")
	}
	a := NewAPI(mock)

	err := a.Clean(&CleanReq{UserIDs: []string{"u1"}})
	assertError(t, err)
}

// ---------------------------------------------------------------------------
// Query
// ---------------------------------------------------------------------------

func TestQuery(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.Query(&QueryReq{Page: 2, Size: 50, Order: 1})
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/user/profile/query.json")
	assertEqual(t, call.Params["page"], "2")
	assertEqual(t, call.Params["size"], "50")
	assertEqual(t, call.Params["order"], "1")
}

func TestQuery_Defaults(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	_, err := a.Query(&QueryReq{})
	assertNoError(t, err)

	call := mock.LastCall()
	if _, ok := call.Params["page"]; ok {
		t.Error("page should not be set when zero")
	}
	if _, ok := call.Params["size"]; ok {
		t.Error("size should not be set when zero")
	}
	if _, ok := call.Params["order"]; ok {
		t.Error("order should not be set when zero")
	}
}

func TestQuery_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("query failed")
	}
	a := NewAPI(mock)

	resp, err := a.Query(&QueryReq{Page: 1})
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// CleanExpansion
// ---------------------------------------------------------------------------

func TestCleanExpansion(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	req := &CleanExpansionReq{UserID: "user1", Keys: []string{"key1"}}
	err := a.CleanExpansion(req)
	assertNoError(t, err)

	call := mock.LastCall()
	assertEqual(t, call.Method, "PostJSON")
	assertEqual(t, call.Path, "/user/profile/expansion/clean.json")
	body := call.Body.(*CleanExpansionReq)
	assertEqual(t, body.UserID, "user1")
	assertEqual(t, len(body.Keys), 1)
}

func TestCleanExpansion_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostJSONFunc = func(path string, body any, resp any) error {
		return errors.New("clean failed")
	}
	a := NewAPI(mock)

	err := a.CleanExpansion(&CleanExpansionReq{UserID: "user1"})
	assertError(t, err)
}
