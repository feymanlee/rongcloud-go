package chatroommute

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

func TestMuteAdd(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	resp, err := a.MuteAdd("cr1", "u1", 30)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Method, "Post")
	assertEqual(t, c.Path, "/chatroom/ban/add.json")
	assertEqual(t, c.Params["chatroomId"], "cr1")
	assertEqual(t, c.Params["userId"], "u1")
	assertEqual(t, c.Params["minute"], "30")
}

func TestMuteRemove(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	resp, err := a.MuteRemove("cr1", "u1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Path, "/chatroom/ban/remove.json")
	assertEqual(t, c.Params["chatroomId"], "cr1")
	assertEqual(t, c.Params["userId"], "u1")
}

func TestMuteQuery(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	resp, err := a.MuteQuery("cr1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Path, "/chatroom/ban/query.json")
	assertEqual(t, c.Params["chatroomId"], "cr1")
}

func TestMuteAllAdd(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	resp, err := a.MuteAllAdd("cr1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Path, "/chatroom/ban/all/add.json")
	assertEqual(t, c.Params["chatroomId"], "cr1")
}

func TestGlobalMuteAdd(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	resp, err := a.GlobalMuteAdd("u1", 60)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Path, "/chatroom/ban/global/add.json")
	assertEqual(t, c.Params["userId"], "u1")
	assertEqual(t, c.Params["minute"], "60")
}

func TestWhitelistAdd(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	resp, err := a.WhitelistAdd("cr1", "u1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Path, "/chatroom/ban/whitelist/add.json")
	assertEqual(t, c.Params["chatroomId"], "cr1")
	assertEqual(t, c.Params["userId"], "u1")
}

func TestGlobalMuteQuery(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	resp, err := a.GlobalMuteQuery()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Path, "/chatroom/ban/global/query.json")
}

func TestMuteAdd_ErrorPropagation(t *testing.T) {
	mc := testutil.NewMockClient()
	expectedErr := errors.New("network error")
	mc.PostFunc = func(path string, params map[string]string, resp any) error {
		return expectedErr
	}
	a := NewAPI(mc)
	resp, err := a.MuteAdd("cr1", "u1", 30)
	if resp != nil {
		t.Errorf("expected nil resp on error, got %v", resp)
	}
	if !errors.Is(err, expectedErr) {
		t.Errorf("expected error %v, got %v", expectedErr, err)
	}
}
