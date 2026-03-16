package chatroom

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
	resp, err := a.Create(map[string]string{"cr1": "Room1"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Method, "Post")
	assertEqual(t, c.Path, "/chatroom/create.json")
	assertEqual(t, c.Params["chatroom[cr1]"], "Room1")
}

func TestDestroy(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	resp, err := a.Destroy("cr1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Path, "/chatroom/destroy.json")
	assertEqual(t, c.Params["chatroomId"], "cr1")
}

func TestQuery(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	resp, err := a.Query("cr1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Path, "/chatroom/query.json")
	assertEqual(t, c.Params["chatroomId"], "cr1")
}

func TestQueryMembers(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	resp, err := a.QueryMembers("cr1", 50, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Path, "/chatroom/user/query.json")
	assertEqual(t, c.Params["chatroomId"], "cr1")
	assertEqual(t, c.Params["count"], "50")
	assertEqual(t, c.Params["order"], "1")
}

func TestIsExist(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	resp, err := a.IsExist("cr1", "u1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Path, "/chatroom/user/exist.json")
	assertEqual(t, c.Params["chatroomId"], "cr1")
	assertEqual(t, c.Params["userId"], "u1")
}

func TestKeepaliveAdd(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	resp, err := a.KeepaliveAdd("cr1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Path, "/chatroom/keepalive/add.json")
	assertEqual(t, c.Params["chatroomId"], "cr1")
}

func TestBlockAdd(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	resp, err := a.BlockAdd("cr1", "u1", 60)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Path, "/chatroom/user/block/add.json")
	assertEqual(t, c.Params["chatroomId"], "cr1")
	assertEqual(t, c.Params["userId"], "u1")
	assertEqual(t, c.Params["minute"], "60")
}

func TestCreate_ErrorPropagation(t *testing.T) {
	mc := testutil.NewMockClient()
	expectedErr := errors.New("network error")
	mc.PostFunc = func(path string, params map[string]string, resp any) error {
		return expectedErr
	}
	a := NewAPI(mc)
	resp, err := a.Create(map[string]string{"cr1": "Room1"})
	if resp != nil {
		t.Errorf("expected nil resp on error, got %v", resp)
	}
	if !errors.Is(err, expectedErr) {
		t.Errorf("expected error %v, got %v", expectedErr, err)
	}
}
