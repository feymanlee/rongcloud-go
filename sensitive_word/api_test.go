package sensitiveword

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

	err := a.Add("badword", "***")
	assertNoError(t, err)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/sensitiveword/add.json")
	assertEqual(t, call.Params["word"], "badword")
	assertEqual(t, call.Params["replaceWord"], "***")
}

func TestAdd_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("add failed")
	}
	a := NewAPI(mock)

	err := a.Add("badword", "***")
	assertError(t, err)
}

// ---------------------------------------------------------------------------
// BatchDelete
// ---------------------------------------------------------------------------

func TestBatchDelete(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	err := a.BatchDelete("word1,word2")
	assertNoError(t, err)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/sensitiveword/batch/delete.json")
	assertEqual(t, call.Params["words"], "word1,word2")
}

func TestBatchDelete_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("batch delete failed")
	}
	a := NewAPI(mock)

	err := a.BatchDelete("word1")
	assertError(t, err)
}

// ---------------------------------------------------------------------------
// List
// ---------------------------------------------------------------------------

func TestList(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.List("1")
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/sensitiveword/list.json")
	assertEqual(t, call.Params["type"], "1")
}

func TestList_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("list failed")
	}
	a := NewAPI(mock)

	resp, err := a.List("1")
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

	err := a.Delete("badword")
	assertNoError(t, err)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/sensitiveword/delete.json")
	assertEqual(t, call.Params["word"], "badword")
}

func TestDelete_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("delete failed")
	}
	a := NewAPI(mock)

	err := a.Delete("badword")
	assertError(t, err)
}
