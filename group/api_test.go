package group

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

func TestCreate(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	resp, err := a.Create([]string{"u1", "u2"}, "g1", "TestGroup")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Method, "Post")
	assertEqual(t, c.Path, "/group/create.json")
	assertEqual(t, c.Params["userId"], "u1,u2")
	assertEqual(t, c.Params["groupId"], "g1")
	assertEqual(t, c.Params["groupName"], "TestGroup")
}

func TestJoin(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	resp, err := a.Join([]string{"u1", "u3"}, "g1", "TestGroup")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Method, "Post")
	assertEqual(t, c.Path, "/group/join.json")
	assertEqual(t, c.Params["userId"], "u1,u3")
	assertEqual(t, c.Params["groupId"], "g1")
}

func TestQuit(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	resp, err := a.Quit([]string{"u1"}, "g1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Path, "/group/quit.json")
	assertEqual(t, c.Params["userId"], "u1")
	assertEqual(t, c.Params["groupId"], "g1")
}

func TestDismiss(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	resp, err := a.Dismiss("u1", "g1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Path, "/group/dismiss.json")
	assertEqual(t, c.Params["userId"], "u1")
	assertEqual(t, c.Params["groupId"], "g1")
}

func TestRefresh(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	resp, err := a.Refresh("g1", "NewName")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Path, "/group/refresh.json")
	assertEqual(t, c.Params["groupId"], "g1")
	assertEqual(t, c.Params["groupName"], "NewName")
}

func TestQueryUser(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	resp, err := a.QueryUser("u1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Path, "/user/group/query.json")
	assertEqual(t, c.Params["userId"], "u1")
}

func TestQueryMembers(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	resp, err := a.QueryMembers("g1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Path, "/group/user/query.json")
	assertEqual(t, c.Params["groupId"], "g1")
}

func TestEntrustCreate(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	resp, err := a.EntrustCreate(EntrustCreateReq{
		GroupID: "g1",
		Name:    "TestGroup",
		Owner:   "u1",
		UserIDs: "u1,u2",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Path, "/entrust/group/create.json")
	assertEqual(t, c.Params["groupId"], "g1")
	assertEqual(t, c.Params["name"], "TestGroup")
	assertEqual(t, c.Params["owner"], "u1")
	assertEqual(t, c.Params["userIds"], "u1,u2")
}

func TestEntrustDismiss(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	resp, err := a.EntrustDismiss("g1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Path, "/entrust/group/dismiss.json")
	assertEqual(t, c.Params["groupId"], "g1")
}

func TestEntrustJoin(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	resp, err := a.EntrustJoin("g1", "u1,u2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Path, "/entrust/group/join.json")
	assertEqual(t, c.Params["groupId"], "g1")
	assertEqual(t, c.Params["userIds"], "u1,u2")
}

func TestCreate_ErrorPropagation(t *testing.T) {
	mc := testutil.NewMockClient()
	expectedErr := errors.New("network error")
	mc.PostFunc = func(path string, params map[string]string, resp any) error {
		return expectedErr
	}
	a := NewAPI(mc)
	resp, err := a.Create([]string{"u1"}, "g1", "TestGroup")
	if resp != nil {
		t.Errorf("expected nil resp on error, got %v", resp)
	}
	if !errors.Is(err, expectedErr) {
		t.Errorf("expected error %v, got %v", expectedErr, err)
	}
}
