package ultragroupchannel

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

	resp, err := a.Create("group1", "channel1")
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/ultragroup/channel/private/create.json")
	assertEqual(t, call.Params["groupId"], "group1")
	assertEqual(t, call.Params["busChannel"], "channel1")
}

func TestCreate_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("create failed")
	}
	a := NewAPI(mock)

	resp, err := a.Create("group1", "channel1")
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

	resp, err := a.Dismiss("group1", "channel1")
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/ultragroup/channel/private/del.json")
	assertEqual(t, call.Params["groupId"], "group1")
	assertEqual(t, call.Params["busChannel"], "channel1")
}

func TestDismiss_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("dismiss failed")
	}
	a := NewAPI(mock)

	resp, err := a.Dismiss("group1", "channel1")
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// MembersAdd
// ---------------------------------------------------------------------------

func TestMembersAdd(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.MembersAdd("group1", "channel1", []string{"u1", "u2"})
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/ultragroup/channel/private/users/add.json")
	assertEqual(t, call.Params["groupId"], "group1")
	assertEqual(t, call.Params["busChannel"], "channel1")
	assertEqual(t, call.Params["userIds"], "u1,u2")
}

func TestMembersAdd_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("add failed")
	}
	a := NewAPI(mock)

	resp, err := a.MembersAdd("group1", "channel1", []string{"u1"})
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// MembersRemove
// ---------------------------------------------------------------------------

func TestMembersRemove(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.MembersRemove("group1", "channel1", []string{"u1", "u2"})
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/ultragroup/channel/private/users/del.json")
	assertEqual(t, call.Params["groupId"], "group1")
	assertEqual(t, call.Params["busChannel"], "channel1")
	assertEqual(t, call.Params["userIds"], "u1,u2")
}

func TestMembersRemove_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("remove failed")
	}
	a := NewAPI(mock)

	resp, err := a.MembersRemove("group1", "channel1", []string{"u1"})
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// MembersQuery
// ---------------------------------------------------------------------------

func TestMembersQuery(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.MembersQuery("group1", "channel1", 1, 20)
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/ultragroup/channel/private/users/query.json")
	assertEqual(t, call.Params["groupId"], "group1")
	assertEqual(t, call.Params["busChannel"], "channel1")
	assertEqual(t, call.Params["page"], "1")
	assertEqual(t, call.Params["pageSize"], "20")
}

func TestMembersQuery_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("query failed")
	}
	a := NewAPI(mock)

	resp, err := a.MembersQuery("group1", "channel1", 1, 20)
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}
