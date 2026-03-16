package chatroomwhitelist

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
// MsgAdd
// ---------------------------------------------------------------------------

func TestMsgAdd(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.MsgAdd("RC:TxtMsg")
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/chatroom/whitelist/add.json")
	assertEqual(t, call.Params["objectnames"], "RC:TxtMsg")
}

func TestMsgAdd_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("msg add failed")
	}
	a := NewAPI(mock)

	resp, err := a.MsgAdd("RC:TxtMsg")
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// MsgRemove
// ---------------------------------------------------------------------------

func TestMsgRemove(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.MsgRemove("RC:TxtMsg")
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/chatroom/whitelist/delete.json")
	assertEqual(t, call.Params["objectnames"], "RC:TxtMsg")
}

func TestMsgRemove_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("msg remove failed")
	}
	a := NewAPI(mock)

	resp, err := a.MsgRemove("RC:TxtMsg")
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// UserAdd
// ---------------------------------------------------------------------------

func TestUserAdd(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.UserAdd("chatroom1", "user1")
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/chatroom/user/whitelist/add.json")
	assertEqual(t, call.Params["chatroomId"], "chatroom1")
	assertEqual(t, call.Params["userId"], "user1")
}

func TestUserAdd_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("user add failed")
	}
	a := NewAPI(mock)

	resp, err := a.UserAdd("chatroom1", "user1")
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// UserQuery
// ---------------------------------------------------------------------------

func TestUserQuery(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.UserQuery("chatroom1")
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/chatroom/user/whitelist/query.json")
	assertEqual(t, call.Params["chatroomId"], "chatroom1")
}

func TestUserQuery_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("user query failed")
	}
	a := NewAPI(mock)

	resp, err := a.UserQuery("chatroom1")
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}
