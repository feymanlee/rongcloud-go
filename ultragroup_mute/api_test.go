package ultragroupmute

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
// MuteAdd
// ---------------------------------------------------------------------------

func TestMuteAdd(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.MuteAdd("group1", []string{"u1", "u2"}, "channel1")
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/ultragroup/ban/add.json")
	assertEqual(t, call.Params["groupId"], "group1")
	assertEqual(t, call.Params["userIds"], "u1,u2")
	assertEqual(t, call.Params["busChannel"], "channel1")
}

func TestMuteAdd_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("mute add failed")
	}
	a := NewAPI(mock)

	resp, err := a.MuteAdd("group1", []string{"u1"}, "channel1")
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// MuteRemove
// ---------------------------------------------------------------------------

func TestMuteRemove(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.MuteRemove("group1", []string{"u1"}, "channel1")
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/ultragroup/ban/remove.json")
	assertEqual(t, call.Params["groupId"], "group1")
	assertEqual(t, call.Params["userIds"], "u1")
	assertEqual(t, call.Params["busChannel"], "channel1")
}

func TestMuteRemove_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("mute remove failed")
	}
	a := NewAPI(mock)

	resp, err := a.MuteRemove("group1", []string{"u1"}, "channel1")
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// MuteQuery
// ---------------------------------------------------------------------------

func TestMuteQuery(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.MuteQuery("group1", "channel1")
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/ultragroup/ban/query.json")
	assertEqual(t, call.Params["groupId"], "group1")
	assertEqual(t, call.Params["busChannel"], "channel1")
}

func TestMuteQuery_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("mute query failed")
	}
	a := NewAPI(mock)

	resp, err := a.MuteQuery("group1", "channel1")
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

	resp, err := a.WhitelistAdd("group1", []string{"u1", "u2"}, "channel1")
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/ultragroup/ban/whitelist/add.json")
	assertEqual(t, call.Params["groupId"], "group1")
	assertEqual(t, call.Params["userIds"], "u1,u2")
	assertEqual(t, call.Params["busChannel"], "channel1")
}

func TestWhitelistAdd_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("whitelist add failed")
	}
	a := NewAPI(mock)

	resp, err := a.WhitelistAdd("group1", []string{"u1"}, "channel1")
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}
