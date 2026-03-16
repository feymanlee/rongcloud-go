package user

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/feymanlee/rongcloud-go/internal/testutil"
)

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

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
// GetToken
// ---------------------------------------------------------------------------

func TestGetToken(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.GetToken("uid001", "name1", "https://example.com/avatar.png")
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/user/getToken.json")
	assertEqual(t, call.Params["userId"], "uid001")
	assertEqual(t, call.Params["name"], "name1")
	assertEqual(t, call.Params["portraitUri"], "https://example.com/avatar.png")
}

func TestGetToken_EmptyPortraitURI(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.GetToken("uid001", "name1", "")
	assertNoError(t, err)
	assertNotNil(t, resp)

	call := mock.LastCall()
	assertEqual(t, call.Path, "/user/getToken.json")
	if _, ok := call.Params["portraitUri"]; ok {
		t.Error("portraitUri should not be set when empty")
	}
}

func TestGetToken_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("network error")
	}
	a := NewAPI(mock)

	resp, err := a.GetToken("uid001", "name1", "")
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
	assertEqual(t, err.Error(), "network error")
}

func TestGetToken_CustomResponse(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		data := []byte(`{"code":200,"token":"test-token-abc","userId":"uid001"}`)
		return json.Unmarshal(data, resp)
	}
	a := NewAPI(mock)

	resp, err := a.GetToken("uid001", "name1", "")
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Token, "test-token-abc")
	assertEqual(t, resp.UserID, "uid001")
	assertEqual(t, resp.Code, 200)
}

// ---------------------------------------------------------------------------
// Update
// ---------------------------------------------------------------------------

func TestUpdate(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.Update("uid002", "newName", "https://example.com/new.png")
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/user/refresh.json")
	assertEqual(t, call.Params["userId"], "uid002")
	assertEqual(t, call.Params["name"], "newName")
	assertEqual(t, call.Params["portraitUri"], "https://example.com/new.png")
}

func TestUpdate_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("timeout")
	}
	a := NewAPI(mock)

	resp, err := a.Update("uid002", "newName", "https://example.com/new.png")
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// UserInfoGet
// ---------------------------------------------------------------------------

func TestUserInfoGet(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.UserInfoGet("uid003")
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/user/info.json")
	assertEqual(t, call.Params["userId"], "uid003")
}

func TestUserInfoGet_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("server error")
	}
	a := NewAPI(mock)

	resp, err := a.UserInfoGet("uid003")
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// TokenExpire
// ---------------------------------------------------------------------------

func TestTokenExpire(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.TokenExpire("uid004", 1700000000)
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/user/token/expire.json")
	assertEqual(t, call.Params["userId"], "uid004")
	assertEqual(t, call.Params["time"], "1700000000")
}

func TestTokenExpire_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("expire failed")
	}
	a := NewAPI(mock)

	resp, err := a.TokenExpire("uid004", 1700000000)
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// BlockAdd
// ---------------------------------------------------------------------------

func TestBlockAdd(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.BlockAdd("uid005", 60)
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/user/block.json")
	assertEqual(t, call.Params["userId"], "uid005")
	assertEqual(t, call.Params["minute"], "60")
}

func TestBlockAdd_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("block failed")
	}
	a := NewAPI(mock)

	resp, err := a.BlockAdd("uid005", 60)
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// BlockRemove
// ---------------------------------------------------------------------------

func TestBlockRemove(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.BlockRemove("uid006")
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/user/unblock.json")
	assertEqual(t, call.Params["userId"], "uid006")
}

func TestBlockRemove_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("unblock failed")
	}
	a := NewAPI(mock)

	resp, err := a.BlockRemove("uid006")
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// BlockQuery
// ---------------------------------------------------------------------------

func TestBlockQuery(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.BlockQuery()
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/user/block/query.json")
	assertEqual(t, len(call.Params), 0)
}

func TestBlockQuery_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("query failed")
	}
	a := NewAPI(mock)

	resp, err := a.BlockQuery()
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// OnlineStatusCheck
// ---------------------------------------------------------------------------

func TestOnlineStatusCheck(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.OnlineStatusCheck("uid007")
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/user/checkOnline.json")
	assertEqual(t, call.Params["userId"], "uid007")
}

func TestOnlineStatusCheck_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("check failed")
	}
	a := NewAPI(mock)

	resp, err := a.OnlineStatusCheck("uid007")
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// Ban
// ---------------------------------------------------------------------------

func TestBan(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.Ban("uid008", 120)
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/user/ban.json")
	assertEqual(t, call.Params["userId"], "uid008")
	assertEqual(t, call.Params["minute"], "120")
}

func TestBan_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("ban failed")
	}
	a := NewAPI(mock)

	resp, err := a.Ban("uid008", 120)
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// BanQuery
// ---------------------------------------------------------------------------

func TestBanQuery(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.BanQuery()
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/user/ban/query.json")
	assertEqual(t, len(call.Params), 0)
}

func TestBanQuery_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("ban query failed")
	}
	a := NewAPI(mock)

	resp, err := a.BanQuery()
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// UnBan
// ---------------------------------------------------------------------------

func TestUnBan(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.UnBan("uid009")
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/user/unban.json")
	assertEqual(t, call.Params["userId"], "uid009")
}

func TestUnBan_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("unban failed")
	}
	a := NewAPI(mock)

	resp, err := a.UnBan("uid009")
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// Deactivate
// ---------------------------------------------------------------------------

func TestDeactivate(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.Deactivate([]string{"uid010", "uid011", "uid012"})
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/user/deactivate.json")
	assertEqual(t, call.Params["userId"], "uid010,uid011,uid012")
}

func TestDeactivate_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("deactivate failed")
	}
	a := NewAPI(mock)

	resp, err := a.Deactivate([]string{"uid010"})
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// DeactivateQuery
// ---------------------------------------------------------------------------

func TestDeactivateQuery(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.DeactivateQuery(1, 50)
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/user/deactivate/query.json")
	assertEqual(t, call.Params["pageNo"], "1")
	assertEqual(t, call.Params["pageSize"], "50")
}

func TestDeactivateQuery_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("deactivate query failed")
	}
	a := NewAPI(mock)

	resp, err := a.DeactivateQuery(1, 50)
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// Reactivate
// ---------------------------------------------------------------------------

func TestReactivate(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.Reactivate([]string{"uid013", "uid014"})
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/user/reactivate.json")
	assertEqual(t, call.Params["userId"], "uid013,uid014")
}

func TestReactivate_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("reactivate failed")
	}
	a := NewAPI(mock)

	resp, err := a.Reactivate([]string{"uid013"})
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}
