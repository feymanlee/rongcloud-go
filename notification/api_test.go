package notification

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

	err := a.Set("1", "req1", "target1", "1")
	assertNoError(t, err)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/conversation/notification/set.json")
	assertEqual(t, call.Params["conversationType"], "1")
	assertEqual(t, call.Params["requestId"], "req1")
	assertEqual(t, call.Params["targetId"], "target1")
	assertEqual(t, call.Params["isMuted"], "1")
}

func TestSet_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("set failed")
	}
	a := NewAPI(mock)

	err := a.Set("1", "req1", "target1", "1")
	assertError(t, err)
}

// ---------------------------------------------------------------------------
// Get
// ---------------------------------------------------------------------------

func TestGet(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.Get("1", "req1", "target1")
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/conversation/notification/get.json")
	assertEqual(t, call.Params["conversationType"], "1")
	assertEqual(t, call.Params["requestId"], "req1")
	assertEqual(t, call.Params["targetId"], "target1")
}

func TestGet_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("get failed")
	}
	a := NewAPI(mock)

	resp, err := a.Get("1", "req1", "target1")
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// GlobalSet
// ---------------------------------------------------------------------------

func TestGlobalSet(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	err := a.GlobalSet("req1", "1", "1")
	assertNoError(t, err)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/conversation/notification/global/set.json")
	assertEqual(t, call.Params["requestId"], "req1")
	assertEqual(t, call.Params["conversationType"], "1")
	assertEqual(t, call.Params["isMuted"], "1")
}

func TestGlobalSet_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("global set failed")
	}
	a := NewAPI(mock)

	err := a.GlobalSet("req1", "1", "1")
	assertError(t, err)
}

// ---------------------------------------------------------------------------
// GlobalGet
// ---------------------------------------------------------------------------

func TestGlobalGet(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.GlobalGet("req1", "1")
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/conversation/notification/global/get.json")
	assertEqual(t, call.Params["requestId"], "req1")
	assertEqual(t, call.Params["conversationType"], "1")
}

func TestGlobalGet_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("global get failed")
	}
	a := NewAPI(mock)

	resp, err := a.GlobalGet("req1", "1")
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}
