package chatroompriority

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
// Add
// ---------------------------------------------------------------------------

func TestAdd(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.Add("RC:TxtMsg")
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/chatroom/message/priority/add.json")
	assertEqual(t, call.Params["objectName"], "RC:TxtMsg")
}

func TestAdd_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("add failed")
	}
	a := NewAPI(mock)

	resp, err := a.Add("RC:TxtMsg")
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

	resp, err := a.Remove("RC:TxtMsg")
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/chatroom/message/priority/remove.json")
	assertEqual(t, call.Params["objectName"], "RC:TxtMsg")
}

func TestRemove_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("remove failed")
	}
	a := NewAPI(mock)

	resp, err := a.Remove("RC:TxtMsg")
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

	resp, err := a.Query()
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/chatroom/message/priority/query.json")
	assertEqual(t, len(call.Params), 0)
}

func TestQuery_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("query failed")
	}
	a := NewAPI(mock)

	resp, err := a.Query()
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}
