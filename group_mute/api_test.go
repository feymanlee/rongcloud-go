package groupmute

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
	resp, err := a.MuteAdd("g1", "u1", "30")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Method, "Post")
	assertEqual(t, c.Path, "/group/ban/add.json")
	assertEqual(t, c.Params["groupId"], "g1")
	assertEqual(t, c.Params["userId"], "u1")
	assertEqual(t, c.Params["minute"], "30")
}

func TestMuteRemove(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	resp, err := a.MuteRemove("g1", "u1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Path, "/group/ban/remove.json")
	assertEqual(t, c.Params["groupId"], "g1")
	assertEqual(t, c.Params["userId"], "u1")
}

func TestMuteQuery(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	resp, err := a.MuteQuery("g1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Path, "/group/ban/query.json")
	assertEqual(t, c.Params["groupId"], "g1")
}

func TestMuteAllAdd(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	resp, err := a.MuteAllAdd("g1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Path, "/group/ban/all/add.json")
	assertEqual(t, c.Params["groupId"], "g1")
}

func TestMuteAllRemove(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	resp, err := a.MuteAllRemove("g1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Path, "/group/ban/all/remove.json")
	assertEqual(t, c.Params["groupId"], "g1")
}

func TestMuteAllQuery(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	resp, err := a.MuteAllQuery()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Path, "/group/ban/all/query.json")
}

func TestWhitelistAdd(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	resp, err := a.WhitelistAdd("g1", "u1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Path, "/group/ban/whitelist/add.json")
	assertEqual(t, c.Params["groupId"], "g1")
	assertEqual(t, c.Params["userId"], "u1")
}

func TestWhitelistRemove(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	resp, err := a.WhitelistRemove("g1", "u1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Path, "/group/ban/whitelist/remove.json")
	assertEqual(t, c.Params["groupId"], "g1")
	assertEqual(t, c.Params["userId"], "u1")
}

func TestWhitelistQuery(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	resp, err := a.WhitelistQuery("g1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Path, "/group/ban/whitelist/query.json")
	assertEqual(t, c.Params["groupId"], "g1")
}

func TestMuteAdd_ErrorPropagation(t *testing.T) {
	mc := testutil.NewMockClient()
	expectedErr := errors.New("network error")
	mc.PostFunc = func(path string, params map[string]string, resp any) error {
		return expectedErr
	}
	a := NewAPI(mc)
	resp, err := a.MuteAdd("g1", "u1", "30")
	if resp != nil {
		t.Errorf("expected nil resp on error, got %v", resp)
	}
	if !errors.Is(err, expectedErr) {
		t.Errorf("expected error %v, got %v", expectedErr, err)
	}
}
