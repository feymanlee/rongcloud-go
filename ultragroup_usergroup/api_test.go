package ultragroupusergroup

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

	req := &CreateReq{GroupId: "group1", UserGroupId: "ug1"}
	resp, err := a.Create(req)
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "PostJSON")
	assertEqual(t, call.Path, "/ultragroup/usergroup/create.json")
	body := call.Body.(*CreateReq)
	assertEqual(t, body.GroupId, "group1")
	assertEqual(t, body.UserGroupId, "ug1")
}

func TestCreate_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostJSONFunc = func(path string, body any, resp any) error {
		return errors.New("create failed")
	}
	a := NewAPI(mock)

	resp, err := a.Create(&CreateReq{GroupId: "group1", UserGroupId: "ug1"})
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

	req := &DeleteReq{GroupId: "group1", UserGroupId: "ug1"}
	resp, err := a.Delete(req)
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "PostJSON")
	assertEqual(t, call.Path, "/ultragroup/usergroup/del.json")
	body := call.Body.(*DeleteReq)
	assertEqual(t, body.GroupId, "group1")
	assertEqual(t, body.UserGroupId, "ug1")
}

func TestDelete_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostJSONFunc = func(path string, body any, resp any) error {
		return errors.New("delete failed")
	}
	a := NewAPI(mock)

	resp, err := a.Delete(&DeleteReq{GroupId: "group1", UserGroupId: "ug1"})
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

	req := &QueryReq{GroupId: "group1"}
	resp, err := a.Query(req)
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "PostJSON")
	assertEqual(t, call.Path, "/ultragroup/usergroup/query.json")
	body := call.Body.(*QueryReq)
	assertEqual(t, body.GroupId, "group1")
}

func TestQuery_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostJSONFunc = func(path string, body any, resp any) error {
		return errors.New("query failed")
	}
	a := NewAPI(mock)

	resp, err := a.Query(&QueryReq{GroupId: "group1"})
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// MemberAdd
// ---------------------------------------------------------------------------

func TestMemberAdd(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	req := &MemberAddReq{GroupId: "group1", UserGroupId: "ug1", UserIds: []string{"u1", "u2"}}
	resp, err := a.MemberAdd(req)
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "PostJSON")
	assertEqual(t, call.Path, "/ultragroup/usergroup/member/add.json")
	body := call.Body.(*MemberAddReq)
	assertEqual(t, body.GroupId, "group1")
	assertEqual(t, body.UserGroupId, "ug1")
	assertEqual(t, len(body.UserIds), 2)
	assertEqual(t, body.UserIds[0], "u1")
}

func TestMemberAdd_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostJSONFunc = func(path string, body any, resp any) error {
		return errors.New("member add failed")
	}
	a := NewAPI(mock)

	resp, err := a.MemberAdd(&MemberAddReq{GroupId: "group1", UserGroupId: "ug1", UserIds: []string{"u1"}})
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}

// ---------------------------------------------------------------------------
// ChannelBind
// ---------------------------------------------------------------------------

func TestChannelBind(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	req := &ChannelBindReq{GroupId: "group1", BusChannel: "ch1", UserGroupId: "ug1"}
	resp, err := a.ChannelBind(req)
	assertNoError(t, err)
	assertNotNil(t, resp)
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "PostJSON")
	assertEqual(t, call.Path, "/ultragroup/usergroup/channel/bindv2.json")
	body := call.Body.(*ChannelBindReq)
	assertEqual(t, body.GroupId, "group1")
	assertEqual(t, body.BusChannel, "ch1")
	assertEqual(t, body.UserGroupId, "ug1")
}

func TestChannelBind_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostJSONFunc = func(path string, body any, resp any) error {
		return errors.New("bind failed")
	}
	a := NewAPI(mock)

	resp, err := a.ChannelBind(&ChannelBindReq{GroupId: "group1", BusChannel: "ch1", UserGroupId: "ug1"})
	assertError(t, err)
	if resp != nil {
		t.Error("expected nil response on error")
	}
}
