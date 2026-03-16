package chatroomkv

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

	resp, err := a.Set("chatroom1", "user1", "color", "red", 1, "")
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/chatroom/entry/set.json")
	assertEqual(t, call.Params["chatroomId"], "chatroom1")
	assertEqual(t, call.Params["userId"], "user1")
	assertEqual(t, call.Params["key"], "color")
	assertEqual(t, call.Params["value"], "red")
	assertEqual(t, call.Params["autoDelete"], "1")
	// objectName should not be in params when empty
	if _, ok := call.Params["objectName"]; ok {
		t.Error("objectName should not be set when empty")
	}
}

func TestSet_WithObjectName(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.Set("chatroom1", "user1", "color", "red", 0, "RC:TxtMsg")
	assertNoError(t, err)
	assertNotNil(t, resp)

	call := mock.LastCall()
	assertEqual(t, call.Params["objectName"], "RC:TxtMsg")
}

func TestSet_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("set failed")
	}
	a := NewAPI(mock)

	resp, err := a.Set("chatroom1", "user1", "color", "red", 1, "")
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// Remove
// ---------------------------------------------------------------------------

func TestRemove(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.Remove("chatroom1", "user1", "color")
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/chatroom/entry/remove.json")
	assertEqual(t, call.Params["chatroomId"], "chatroom1")
	assertEqual(t, call.Params["userId"], "user1")
	assertEqual(t, call.Params["key"], "color")
}

func TestRemove_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("remove failed")
	}
	a := NewAPI(mock)

	resp, err := a.Remove("chatroom1", "user1", "color")
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// Query
// ---------------------------------------------------------------------------

func TestQuery(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.Query("chatroom1", "color")
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/chatroom/entry/query.json")
	assertEqual(t, call.Params["chatroomId"], "chatroom1")
	assertEqual(t, call.Params["key"], "color")
}

func TestQuery_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("query failed")
	}
	a := NewAPI(mock)

	resp, err := a.Query("chatroom1", "color")
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// BatchSet (PostJSON)
// ---------------------------------------------------------------------------

func TestBatchSet(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	req := &BatchSetReq{
		ChatroomID:   "chatroom1",
		AutoDelete:   1,
		EntryOwnerID: "user1",
		EntryInfo:    map[string]string{"color": "red", "size": "large"},
	}
	resp, err := a.BatchSet(req)
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "PostJSON")
	assertEqual(t, call.Path, "/chatroom/entry/batch/set.json")
	body := call.Body.(*BatchSetReq)
	assertEqual(t, body.ChatroomID, "chatroom1")
	assertEqual(t, body.AutoDelete, 1)
	assertEqual(t, body.EntryOwnerID, "user1")
	assertEqual(t, len(body.EntryInfo), 2)
}

func TestBatchSet_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostJSONFunc = func(path string, body any, resp any) error {
		return errors.New("batch set failed")
	}
	a := NewAPI(mock)

	resp, err := a.BatchSet(&BatchSetReq{ChatroomID: "chatroom1"})
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}
