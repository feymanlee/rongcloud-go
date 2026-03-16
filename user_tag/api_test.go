package usertag

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

// ---------------------------------------------------------------------------
// TagSet
// ---------------------------------------------------------------------------

func TestTagSet(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	err := a.TagSet(&SetReq{UserID: "user1", Tags: []string{"tag1", "tag2"}})
	assertNoError(t, err)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/user/tag/set.json")
	assertEqual(t, call.Params["userId"], "user1")
	// tags is JSON-encoded
	assertEqual(t, call.Params["tags"], `["tag1","tag2"]`)
}

func TestTagSet_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("tag set failed")
	}
	a := NewAPI(mock)

	err := a.TagSet(&SetReq{UserID: "user1", Tags: []string{"tag1"}})
	assertError(t, err)
}

// ---------------------------------------------------------------------------
// TagBatchSet
// ---------------------------------------------------------------------------

func TestTagBatchSet(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	err := a.TagBatchSet(&BatchSetReq{UserIDs: []string{"u1", "u2"}, Tags: []string{"tag1"}})
	assertNoError(t, err)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/user/tag/batch/set.json")
	assertEqual(t, call.Params["userIds"], `["u1","u2"]`)
	assertEqual(t, call.Params["tags"], `["tag1"]`)
}

func TestTagBatchSet_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("batch set failed")
	}
	a := NewAPI(mock)

	err := a.TagBatchSet(&BatchSetReq{UserIDs: []string{"u1"}, Tags: []string{"tag1"}})
	assertError(t, err)
}

// ---------------------------------------------------------------------------
// TagGet
// ---------------------------------------------------------------------------

func TestTagGet(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.TagGet(&GetReq{UserIDs: []string{"u1", "u2"}})
	assertNoError(t, err)
	if resp == nil {
		t.Fatal("expected non-nil response")
	}
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/user/tag/get.json")
	assertEqual(t, call.Params["userIds"], `["u1","u2"]`)
}

func TestTagGet_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("tag get failed")
	}
	a := NewAPI(mock)

	resp, err := a.TagGet(&GetReq{UserIDs: []string{"u1"}})
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}
