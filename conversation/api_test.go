package conversation

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

func TestMute(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.Mute(1, "reqUser1", "targetUser1", 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/conversation/notification/set.json")
	assertEqual(t, call.Params["conversationType"], "1")
	assertEqual(t, call.Params["requestId"], "reqUser1")
	assertEqual(t, call.Params["targetId"], "targetUser1")
	assertEqual(t, call.Params["isMuted"], "1")
}

func TestMute_Unmute(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.Mute(3, "reqUser2", "groupTarget", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Params["conversationType"], "3")
	assertEqual(t, call.Params["requestId"], "reqUser2")
	assertEqual(t, call.Params["targetId"], "groupTarget")
	assertEqual(t, call.Params["isMuted"], "0")
}

func TestMute_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("network error")
	}
	a := NewAPI(mock)

	_, err := a.Mute(1, "reqUser1", "targetUser1", 1)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	assertEqual(t, err.Error(), "network error")
}
