package userblock

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
// BlacklistAdd
// ---------------------------------------------------------------------------

func TestBlacklistAdd(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	err := a.BlacklistAdd("user1", "blackUser1")
	assertNoError(t, err)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/user/blacklist/add.json")
	assertEqual(t, call.Params["userId"], "user1")
	assertEqual(t, call.Params["blackUserId"], "blackUser1")
}

func TestBlacklistAdd_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("blacklist add failed")
	}
	a := NewAPI(mock)

	err := a.BlacklistAdd("user1", "blackUser1")
	assertError(t, err)
}

// ---------------------------------------------------------------------------
// BlacklistRemove
// ---------------------------------------------------------------------------

func TestBlacklistRemove(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	err := a.BlacklistRemove("user1", "blackUser1")
	assertNoError(t, err)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/user/blacklist/remove.json")
	assertEqual(t, call.Params["userId"], "user1")
	assertEqual(t, call.Params["blackUserId"], "blackUser1")
}

func TestBlacklistRemove_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("blacklist remove failed")
	}
	a := NewAPI(mock)

	err := a.BlacklistRemove("user1", "blackUser1")
	assertError(t, err)
}

// ---------------------------------------------------------------------------
// BlacklistQuery
// ---------------------------------------------------------------------------

func TestBlacklistQuery(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.BlacklistQuery("user1")
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/user/blacklist/query.json")
	assertEqual(t, call.Params["userId"], "user1")
}

func TestBlacklistQuery_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("blacklist query failed")
	}
	a := NewAPI(mock)

	resp, err := a.BlacklistQuery("user1")
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// WhitelistAdd
// ---------------------------------------------------------------------------

func TestWhitelistAdd(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	err := a.WhitelistAdd("user1", "whiteUser1")
	assertNoError(t, err)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/user/whitelist/add.json")
	assertEqual(t, call.Params["userId"], "user1")
	assertEqual(t, call.Params["whiteUserId"], "whiteUser1")
}

func TestWhitelistAdd_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("whitelist add failed")
	}
	a := NewAPI(mock)

	err := a.WhitelistAdd("user1", "whiteUser1")
	assertError(t, err)
}

// ---------------------------------------------------------------------------
// MsgFilterAdd
// ---------------------------------------------------------------------------

func TestMsgFilterAdd(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	err := a.MsgFilterAdd("user1", "1", "target1,target2")
	assertNoError(t, err)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/user/msgfilter/add.json")
	assertEqual(t, call.Params["userId"], "user1")
	assertEqual(t, call.Params["type"], "1")
	assertEqual(t, call.Params["targetIds"], "target1,target2")
}

func TestMsgFilterAdd_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("msg filter add failed")
	}
	a := NewAPI(mock)

	err := a.MsgFilterAdd("user1", "1", "target1")
	assertError(t, err)
}
