package ultragroup

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
// Create
// ---------------------------------------------------------------------------

func TestCreate(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.Create("user1", "group1", "TestGroup")
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/ultragroup/create.json")
	assertEqual(t, call.Params["userId"], "user1")
	assertEqual(t, call.Params["groupId"], "group1")
	assertEqual(t, call.Params["groupName"], "TestGroup")
}

func TestCreate_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("create failed")
	}
	a := NewAPI(mock)

	resp, err := a.Create("user1", "group1", "TestGroup")
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// Dismiss
// ---------------------------------------------------------------------------

func TestDismiss(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.Dismiss("group1")
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/ultragroup/dis.json")
	assertEqual(t, call.Params["groupId"], "group1")
}

func TestDismiss_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("dismiss failed")
	}
	a := NewAPI(mock)

	resp, err := a.Dismiss("group1")
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// Join
// ---------------------------------------------------------------------------

func TestJoin(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.Join("user1", "group1")
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/ultragroup/join.json")
	assertEqual(t, call.Params["userId"], "user1")
	assertEqual(t, call.Params["groupId"], "group1")
}

func TestJoin_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("join failed")
	}
	a := NewAPI(mock)

	resp, err := a.Join("user1", "group1")
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// Quit
// ---------------------------------------------------------------------------

func TestQuit(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.Quit("user1", "group1")
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/ultragroup/quit.json")
	assertEqual(t, call.Params["userId"], "user1")
	assertEqual(t, call.Params["groupId"], "group1")
}

func TestQuit_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("quit failed")
	}
	a := NewAPI(mock)

	resp, err := a.Quit("user1", "group1")
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// HisMsgPublish
// ---------------------------------------------------------------------------

func TestHisMsgPublish(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.HisMsgPublish("user1", "group1", "RC:TxtMsg", `{"content":"hello"}`)
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/message/ultragroup/publish.json")
	assertEqual(t, call.Params["fromUserId"], "user1")
	assertEqual(t, call.Params["toGroupId"], "group1")
	assertEqual(t, call.Params["objectName"], "RC:TxtMsg")
	assertEqual(t, call.Params["content"], `{"content":"hello"}`)
}

func TestHisMsgPublish_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("publish failed")
	}
	a := NewAPI(mock)

	resp, err := a.HisMsgPublish("user1", "group1", "RC:TxtMsg", `{"content":"hello"}`)
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// ChannelCreate
// ---------------------------------------------------------------------------

func TestChannelCreate(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.ChannelCreate("group1", "channel1", "1")
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/ultragroup/channel/create.json")
	assertEqual(t, call.Params["groupId"], "group1")
	assertEqual(t, call.Params["busChannel"], "channel1")
	assertEqual(t, call.Params["type"], "1")
}

func TestChannelCreate_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("channel create failed")
	}
	a := NewAPI(mock)

	resp, err := a.ChannelCreate("group1", "channel1", "1")
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}
