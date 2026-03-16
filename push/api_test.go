package push

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

func TestBroadcastSend(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	err := a.BroadcastSend("fromUser1", "RC:TxtMsg", `{"content":"hello"}`, "push content", "push data")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	c := mc.LastCall()
	assertEqual(t, c.Method, "Post")
	assertEqual(t, c.Path, "/message/broadcast.json")
	assertEqual(t, c.Params["fromUserId"], "fromUser1")
	assertEqual(t, c.Params["objectName"], "RC:TxtMsg")
	assertEqual(t, c.Params["content"], `{"content":"hello"}`)
	assertEqual(t, c.Params["pushContent"], "push content")
	assertEqual(t, c.Params["pushData"], "push data")
}

func TestSystemSend(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	err := a.SystemSend("fromUser1", "toUser1", "RC:TxtMsg", `{"content":"hi"}`, "push", "data")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	c := mc.LastCall()
	assertEqual(t, c.Method, "Post")
	assertEqual(t, c.Path, "/message/system/publish.json")
	assertEqual(t, c.Params["fromUserId"], "fromUser1")
	assertEqual(t, c.Params["toUserId"], "toUser1")
	assertEqual(t, c.Params["objectName"], "RC:TxtMsg")
	assertEqual(t, c.Params["content"], `{"content":"hi"}`)
}

func TestTagPush(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	req := &TagPushReq{
		PlatForm:   []PlatForm{IOSPlatForm, AndroidPlatForm},
		FromUserID: "fromUser1",
		Audience: Audience{
			Tag: []string{"tag1"},
		},
		Message: &Message{
			Content:    `{"content":"hello"}`,
			ObjectName: "RC:TxtMsg",
		},
		Notification: Notification{
			Alert: "test alert",
		},
	}
	resp, err := a.TagPush(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Method, "PostJSON")
	assertEqual(t, c.Path, "/push/tag.json")
	if c.Body == nil {
		t.Fatal("expected non-nil body for PostJSON call")
	}
}

func TestPushQueryTask(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	resp, err := a.PushQueryTask("task123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Method, "Post")
	assertEqual(t, c.Path, "/push/query/task.json")
	assertEqual(t, c.Params["taskId"], "task123")
}

func TestPushDelete(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	err := a.PushDelete("task123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	c := mc.LastCall()
	assertEqual(t, c.Method, "Post")
	assertEqual(t, c.Path, "/push/delete.json")
	assertEqual(t, c.Params["taskId"], "task123")
}

func TestPushQueryStatus(t *testing.T) {
	mc := testutil.NewMockClient()
	a := NewAPI(mc)
	resp, err := a.PushQueryStatus("task456")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)
	c := mc.LastCall()
	assertEqual(t, c.Path, "/push/query/status.json")
	assertEqual(t, c.Params["taskId"], "task456")
}

func TestBroadcastSend_ErrorPropagation(t *testing.T) {
	mc := testutil.NewMockClient()
	expectedErr := errors.New("network error")
	mc.PostFunc = func(path string, params map[string]string, resp any) error {
		return expectedErr
	}
	a := NewAPI(mc)
	err := a.BroadcastSend("fromUser1", "RC:TxtMsg", `{"content":"hello"}`, "push", "data")
	if !errors.Is(err, expectedErr) {
		t.Errorf("expected error %v, got %v", expectedErr, err)
	}
}

func TestTagPush_ErrorPropagation(t *testing.T) {
	mc := testutil.NewMockClient()
	expectedErr := errors.New("api error")
	mc.PostJSONFunc = func(path string, body any, resp any) error {
		return expectedErr
	}
	a := NewAPI(mc)
	req := &TagPushReq{
		PlatForm: []PlatForm{IOSPlatForm},
		Audience: Audience{IsToAll: true},
	}
	resp, err := a.TagPush(req)
	if resp != nil {
		t.Errorf("expected nil resp on error, got %v", resp)
	}
	if !errors.Is(err, expectedErr) {
		t.Errorf("expected error %v, got %v", expectedErr, err)
	}
}
