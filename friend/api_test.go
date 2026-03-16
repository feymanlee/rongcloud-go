package friend

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

func TestAdd(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.Add("user1", "friend1", "hello", "agreed")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/friend/add.json")
	assertEqual(t, call.Params["userId"], "user1")
	assertEqual(t, call.Params["friendId"], "friend1")
	assertEqual(t, call.Params["message"], "hello")
	assertEqual(t, call.Params["status"], "agreed")
}

func TestAdd_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("network error")
	}
	a := NewAPI(mock)

	_, err := a.Add("user1", "friend1", "hello", "agreed")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	assertEqual(t, err.Error(), "network error")
}

func TestRemove(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.Remove("user1", "friend1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/friend/remove.json")
	assertEqual(t, call.Params["userId"], "user1")
	assertEqual(t, call.Params["friendId"], "friend1")
}

func TestRemove_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("fail")
	}
	a := NewAPI(mock)

	_, err := a.Remove("user1", "friend1")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestBatchRemove(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.BatchRemove("user1", []string{"f1", "f2", "f3"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/friend/batch/remove.json")
	assertEqual(t, call.Params["userId"], "user1")
	assertEqual(t, call.Params["friendIds"], "f1,f2,f3")
}

func TestBatchRemove_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("fail")
	}
	a := NewAPI(mock)

	_, err := a.BatchRemove("user1", []string{"f1"})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestQuery(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.Query("user1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/friend/query.json")
	assertEqual(t, call.Params["userId"], "user1")
}

func TestQuery_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("fail")
	}
	a := NewAPI(mock)

	_, err := a.Query("user1")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestQueryByFriendId(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.QueryByFriendId("user1", "friend1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/friend/query.json")
	assertEqual(t, call.Params["userId"], "user1")
	assertEqual(t, call.Params["friendId"], "friend1")
}

func TestQueryByFriendId_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("fail")
	}
	a := NewAPI(mock)

	_, err := a.QueryByFriendId("user1", "friend1")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestSetRemark(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.SetRemark("user1", "friend1", "best friend")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/friend/set_remark.json")
	assertEqual(t, call.Params["userId"], "user1")
	assertEqual(t, call.Params["friendId"], "friend1")
	assertEqual(t, call.Params["remark"], "best friend")
}

func TestSetRemark_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("fail")
	}
	a := NewAPI(mock)

	_, err := a.SetRemark("user1", "friend1", "remark")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestDirectionFriendQuery(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.DirectionFriendQuery("user1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/friend/direction_friend/query.json")
	assertEqual(t, call.Params["userId"], "user1")
}

func TestDirectionFriendQuery_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("fail")
	}
	a := NewAPI(mock)

	_, err := a.DirectionFriendQuery("user1")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestGetBlacklist(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.GetBlacklist("user1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/friend/blacklist/query.json")
	assertEqual(t, call.Params["userId"], "user1")
}

func TestGetBlacklist_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("fail")
	}
	a := NewAPI(mock)

	_, err := a.GetBlacklist("user1")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
