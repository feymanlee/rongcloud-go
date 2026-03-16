package pushperiod

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

	err := a.Set("user1", "22:00:00", "360", "1")
	assertNoError(t, err)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/user/blockPushPeriod/set.json")
	assertEqual(t, call.Params["userId"], "user1")
	assertEqual(t, call.Params["startTime"], "22:00:00")
	assertEqual(t, call.Params["period"], "360")
	assertEqual(t, call.Params["level"], "1")
}

func TestSet_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("set failed")
	}
	a := NewAPI(mock)

	err := a.Set("user1", "22:00:00", "360", "1")
	assertError(t, err)
}

// ---------------------------------------------------------------------------
// Get
// ---------------------------------------------------------------------------

func TestGet(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.Get("user1")
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/user/blockPushPeriod/get.json")
	assertEqual(t, call.Params["userId"], "user1")
}

func TestGet_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("get failed")
	}
	a := NewAPI(mock)

	resp, err := a.Get("user1")
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// Delete
// ---------------------------------------------------------------------------

func TestDelete(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	err := a.Delete("user1")
	assertNoError(t, err)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/user/blockPushPeriod/delete.json")
	assertEqual(t, call.Params["userId"], "user1")
}

func TestDelete_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("delete failed")
	}
	a := NewAPI(mock)

	err := a.Delete("user1")
	assertError(t, err)
}
