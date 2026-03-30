package friend

import (
	"encoding/json"
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

	resp, err := a.Add("user1", "friend1", AddOptTypeByVerifyLevel, "hello")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/friend/add.json")
	assertEqual(t, call.Params["userId"], "user1")
	assertEqual(t, call.Params["targetId"], "friend1")
	assertEqual(t, call.Params["optType"], "1")
	assertEqual(t, call.Params["extra"], "hello")
}

func TestAdd_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("network error")
	}
	a := NewAPI(mock)

	_, err := a.Add("user1", "friend1", AddOptTypeByVerifyLevel, "hello")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	assertEqual(t, err.Error(), "network error")
}

func TestAdd_OmitOptionalParams(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.Add("user1", "friend1", 0, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Params["userId"], "user1")
	assertEqual(t, call.Params["targetId"], "friend1")
	if _, ok := call.Params["optType"]; ok {
		t.Fatal("expected optType to be omitted when zero")
	}
	if _, ok := call.Params["extra"]; ok {
		t.Fatal("expected extra to be omitted when empty")
	}
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
	assertEqual(t, call.Path, "/friend/get.json")
	assertEqual(t, call.Params["userId"], "user1")
	if _, ok := call.Params["pageToken"]; ok {
		t.Fatal("expected pageToken to be omitted by default")
	}
	if _, ok := call.Params["size"]; ok {
		t.Fatal("expected size to be omitted by default")
	}
	if _, ok := call.Params["order"]; ok {
		t.Fatal("expected order to be omitted by default")
	}
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

func TestQuery_ResponseCompatFields(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return json.Unmarshal([]byte(`{
			"code": 200,
			"pageToken": "next-token",
			"totalCount": 1,
			"friendList": [{
				"userId": "friend1",
				"name": "Alice",
				"remarkName": "best friend",
				"friendExtProfile": "{\"tag\":\"vip\"}",
				"time": 1710000000
			}]
		}`), resp)
	}
	a := NewAPI(mock)

	resp, err := a.Query("user1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	assertEqual(t, resp.PageToken, "next-token")
	assertEqual(t, resp.TotalCount, 1)
	assertEqual(t, len(resp.FriendList), 1)
	assertEqual(t, resp.FriendList[0].UserId, "friend1")
	assertEqual(t, resp.FriendList[0].RemarkName, "best friend")
	assertEqual(t, resp.FriendList[0].Remark, "best friend")
	assertEqual(t, resp.FriendList[0].FriendId, "friend1")
	assertEqual(t, len(resp.Friends), 1)
	assertEqual(t, resp.Friends[0].UserId, "friend1")
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
	assertEqual(t, call.Path, "/friend/check.json")
	assertEqual(t, call.Params["userId"], "user1")
	assertEqual(t, call.Params["targetIds"], "friend1")
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

func TestQueryWithPage(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.QueryWithPage("user1", "next-token", 50, QueryOrderDesc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/friend/get.json")
	assertEqual(t, call.Params["userId"], "user1")
	assertEqual(t, call.Params["pageToken"], "next-token")
	assertEqual(t, call.Params["size"], "50")
	assertEqual(t, call.Params["order"], "1")
}

func TestCheck(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.Check("user1", []string{"friend1", "friend2"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/friend/check.json")
	assertEqual(t, call.Params["userId"], "user1")
	assertEqual(t, call.Params["targetIds"], "friend1,friend2")
}

func TestCheck_Response(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return json.Unmarshal([]byte(`{
			"code": 200,
			"results": [
				{"userId": "friend1", "result": 1},
				{"userId": "friend2", "result": 0}
			]
		}`), resp)
	}
	a := NewAPI(mock)

	resp, err := a.Check("user1", []string{"friend1", "friend2"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	assertEqual(t, resp.Code, 200)
	assertEqual(t, len(resp.Results), 2)
	assertEqual(t, resp.Results[0].UserId, "friend1")
	assertEqual(t, resp.Results[0].Result, 1)
	assertEqual(t, resp.Results[1].UserId, "friend2")
	assertEqual(t, resp.Results[1].Result, 0)
}

func TestCheck_EmptyTargetIds(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	_, err := a.Check("user1", nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	assertEqual(t, err.Error(), "rongcloud: targetIds is required")
	if len(mock.Calls) != 0 {
		t.Fatal("expected no request to be sent")
	}
}

func TestCheck_TooManyTargetIds(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	targetIds := make([]string, 101)
	for i := range targetIds {
		targetIds[i] = "friend"
	}

	_, err := a.Check("user1", targetIds)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	assertEqual(t, err.Error(), "rongcloud: targetIds must not exceed 100")
	if len(mock.Calls) != 0 {
		t.Fatal("expected no request to be sent")
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
