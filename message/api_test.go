package message

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

func intPtr(v int) *int { return &v }

func TestSendPrivate(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.SendPrivate(&SendPrivateReq{
		FromUserId:  "fromUser1",
		ToUserId:    []string{"toUser1", "toUser2"},
		ObjectName:  "RC:TxtMsg",
		Content:     `{"content":"hello"}`,
		PushContent: "you have a message",
		PushData:    `{"key":"val"}`,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/message/private/publish.json")
	assertEqual(t, call.Params["fromUserId"], "fromUser1")
	assertEqual(t, call.Params["toUserId"], "toUser1,toUser2")
	assertEqual(t, call.Params["objectName"], "RC:TxtMsg")
	assertEqual(t, call.Params["content"], `{"content":"hello"}`)
	assertEqual(t, call.Params["pushContent"], "you have a message")
	assertEqual(t, call.Params["pushData"], `{"key":"val"}`)
	// optional fields should not be present when zero value
	if _, ok := call.Params["disablePush"]; ok {
		t.Error("disablePush should not be set")
	}
	if _, ok := call.Params["isPersisted"]; ok {
		t.Error("isPersisted should not be set when nil")
	}
}

func TestSendPrivate_AllFields(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.SendPrivate(&SendPrivateReq{
		FromUserId:           "fromUser1",
		ToUserId:             []string{"toUser1"},
		ObjectName:           "RC:TxtMsg",
		Content:              `{"content":"hello"}`,
		PushContent:          "push",
		PushData:             "data",
		PushExt:              `{"title":"hi"}`,
		DisablePush:          true,
		Count:                5,
		ContentAvailable:     1,
		IsPersisted:          intPtr(0),
		IsIncludeSender:      intPtr(1),
		DisableUpdateLastMsg: true,
		Expansion:            true,
		ExtraContent:         `{"k":"v"}`,
		NeedReadReceipt:      1,
		VerifyBlacklist:      1,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Params["pushExt"], `{"title":"hi"}`)
	assertEqual(t, call.Params["disablePush"], "true")
	assertEqual(t, call.Params["count"], "5")
	assertEqual(t, call.Params["contentAvailable"], "1")
	assertEqual(t, call.Params["isPersisted"], "0")
	assertEqual(t, call.Params["isIncludeSender"], "1")
	assertEqual(t, call.Params["disableUpdateLastMsg"], "true")
	assertEqual(t, call.Params["expansion"], "true")
	assertEqual(t, call.Params["extraContent"], `{"k":"v"}`)
	assertEqual(t, call.Params["needReadReceipt"], "1")
	assertEqual(t, call.Params["verifyBlacklist"], "1")
}

func TestSendPrivate_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("network error")
	}
	a := NewAPI(mock)

	_, err := a.SendPrivate(&SendPrivateReq{FromUserId: "u1", ToUserId: []string{"u2"}, ObjectName: "RC:TxtMsg", Content: "{}"})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	assertEqual(t, err.Error(), "network error")
}

func TestSendGroup(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.SendGroup(&SendGroupReq{
		FromUserId:  "fromUser1",
		ToGroupId:   []string{"group1", "group2"},
		ObjectName:  "RC:TxtMsg",
		Content:     `{"content":"hi"}`,
		PushContent: "push",
		PushData:    "data",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/message/group/publish.json")
	assertEqual(t, call.Params["fromUserId"], "fromUser1")
	assertEqual(t, call.Params["toGroupId"], "group1,group2")
	assertEqual(t, call.Params["objectName"], "RC:TxtMsg")
	assertEqual(t, call.Params["content"], `{"content":"hi"}`)
	// optional fields should not be present when zero value
	if _, ok := call.Params["toUserId"]; ok {
		t.Error("toUserId should not be set when empty")
	}
	if _, ok := call.Params["isMentioned"]; ok {
		t.Error("isMentioned should not be set when zero")
	}
}

func TestSendGroup_AllFields(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.SendGroup(&SendGroupReq{
		FromUserId:           "fromUser1",
		ToGroupId:            []string{"group1"},
		ToUserId:             []string{"u1", "u2"},
		ObjectName:           "RC:TxtMsg",
		Content:              `{"content":"hi"}`,
		PushContent:          "push",
		PushData:             "data",
		PushExt:              `{"title":"grp"}`,
		DisablePush:          true,
		IsIncludeSender:      intPtr(1),
		IsPersisted:          intPtr(0),
		IsMentioned:          1,
		ContentAvailable:     1,
		Expansion:            true,
		ExtraContent:         `{"k":"v"}`,
		DisableUpdateLastMsg: true,
		NeedReadReceipt:      1,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Params["toUserId"], "u1,u2")
	assertEqual(t, call.Params["pushExt"], `{"title":"grp"}`)
	assertEqual(t, call.Params["disablePush"], "true")
	assertEqual(t, call.Params["isIncludeSender"], "1")
	assertEqual(t, call.Params["isPersisted"], "0")
	assertEqual(t, call.Params["isMentioned"], "1")
	assertEqual(t, call.Params["contentAvailable"], "1")
	assertEqual(t, call.Params["expansion"], "true")
	assertEqual(t, call.Params["extraContent"], `{"k":"v"}`)
	assertEqual(t, call.Params["disableUpdateLastMsg"], "true")
	assertEqual(t, call.Params["needReadReceipt"], "1")
}

func TestSendGroup_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("fail")
	}
	a := NewAPI(mock)

	_, err := a.SendGroup(&SendGroupReq{FromUserId: "u1", ToGroupId: []string{"g1"}, ObjectName: "RC:TxtMsg", Content: "{}"})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestSendChatroom(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.SendChatroom(&SendChatroomReq{
		FromUserId:   "fromUser1",
		ToChatroomId: []string{"room1", "room2"},
		ObjectName:   "RC:TxtMsg",
		Content:      `{"content":"hey"}`,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/message/chatroom/publish.json")
	assertEqual(t, call.Params["fromUserId"], "fromUser1")
	assertEqual(t, call.Params["toChatroomId"], "room1,room2")
	assertEqual(t, call.Params["objectName"], "RC:TxtMsg")
	assertEqual(t, call.Params["content"], `{"content":"hey"}`)
}

func TestSendChatroom_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("fail")
	}
	a := NewAPI(mock)

	_, err := a.SendChatroom(&SendChatroomReq{FromUserId: "u1", ToChatroomId: []string{"r1"}, ObjectName: "RC:TxtMsg", Content: "{}"})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestSendSystem(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.SendSystem(&SendSystemReq{
		FromUserId:  "sys1",
		ToUserId:    []string{"u1", "u2"},
		ObjectName:  "RC:TxtMsg",
		Content:     `{"content":"sys"}`,
		PushContent: "push",
		PushData:    "data",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/message/system/publish.json")
	assertEqual(t, call.Params["fromUserId"], "sys1")
	assertEqual(t, call.Params["toUserId"], "u1,u2")
}

func TestSendSystem_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("fail")
	}
	a := NewAPI(mock)

	_, err := a.SendSystem(&SendSystemReq{FromUserId: "u1", ToUserId: []string{"u2"}, ObjectName: "RC:TxtMsg", Content: "{}"})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestSendBroadcast(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.SendBroadcast(&SendBroadcastReq{
		FromUserId: "admin",
		ObjectName: "RC:TxtMsg",
		Content:    `{"content":"broadcast"}`,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/message/broadcast.json")
	assertEqual(t, call.Params["fromUserId"], "admin")
	assertEqual(t, call.Params["objectName"], "RC:TxtMsg")
	assertEqual(t, call.Params["content"], `{"content":"broadcast"}`)
}

func TestSendBroadcast_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("fail")
	}
	a := NewAPI(mock)

	_, err := a.SendBroadcast(&SendBroadcastReq{FromUserId: "u1", ObjectName: "RC:TxtMsg", Content: "{}"})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestRecallPrivate(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.RecallPrivate(&RecallReq{
		FromUserId: "u1",
		TargetId:   "u2",
		MessageUID: "msg-uid-123",
		SentTime:   1620000000,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/message/recall.json")
	assertEqual(t, call.Params["fromUserId"], "u1")
	assertEqual(t, call.Params["targetId"], "u2")
	assertEqual(t, call.Params["messageUID"], "msg-uid-123")
	assertEqual(t, call.Params["sentTime"], "1620000000")
	assertEqual(t, call.Params["conversationType"], "1")
}

func TestRecallPrivate_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("fail")
	}
	a := NewAPI(mock)

	_, err := a.RecallPrivate(&RecallReq{FromUserId: "u1", TargetId: "u2", MessageUID: "m1", SentTime: 100})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestRecallGroup(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.RecallGroup(&RecallReq{
		FromUserId: "u1",
		TargetId:   "g1",
		MessageUID: "msg-uid-456",
		SentTime:   1620000001,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/message/recall.json")
	assertEqual(t, call.Params["conversationType"], "3")
	assertEqual(t, call.Params["fromUserId"], "u1")
	assertEqual(t, call.Params["targetId"], "g1")
	assertEqual(t, call.Params["messageUID"], "msg-uid-456")
	assertEqual(t, call.Params["sentTime"], "1620000001")
}

func TestRecallGroup_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("fail")
	}
	a := NewAPI(mock)

	_, err := a.RecallGroup(&RecallReq{FromUserId: "u1", TargetId: "g1", MessageUID: "m1", SentTime: 100})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestHistoryQuery(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.HistoryQuery("2023010101")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/message/history.json")
	assertEqual(t, call.Params["date"], "2023010101")
}

func TestHistoryQuery_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("fail")
	}
	a := NewAPI(mock)

	_, err := a.HistoryQuery("2023010101")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestHistoryDelete(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.HistoryDelete("2023010101")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/message/history/delete.json")
	assertEqual(t, call.Params["date"], "2023010101")
}

func TestHistoryDelete_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("fail")
	}
	a := NewAPI(mock)

	_, err := a.HistoryDelete("2023010101")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestSendPrivateTemplate(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	req := &SendPrivateTemplateReq{
		FromUserId:  "tmplUser",
		ObjectName:  "RC:TxtMsg",
		Content:     `{"content":"{name}"}`,
		ToUserId:    []string{"u1", "u2"},
		Values:      []map[string]string{{"name": "Alice"}, {"name": "Bob"}},
		PushContent: []string{"push1", "push2"},
		PushData:    []string{"data1", "data2"},
	}
	resp, err := a.SendPrivateTemplate(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "PostJSON")
	assertEqual(t, call.Path, "/message/private/publish_template.json")
	body, ok := call.Body.(*SendPrivateTemplateReq)
	if !ok {
		t.Fatalf("expected body to be *SendPrivateTemplateReq, got %T", call.Body)
	}
	assertEqual(t, body.FromUserId, "tmplUser")
	assertEqual(t, body.ObjectName, ObjectName("RC:TxtMsg"))
}

func TestSendPrivateTemplate_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostJSONFunc = func(path string, body any, resp any) error {
		return errors.New("json error")
	}
	a := NewAPI(mock)

	_, err := a.SendPrivateTemplate(&SendPrivateTemplateReq{FromUserId: "u1"})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	assertEqual(t, err.Error(), "json error")
}

func TestSendStatusMessage(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.SendStatusMessage(&SendStatusMessageReq{
		FromUserId: "u1",
		ToGroupId:  []string{"g1", "g2"},
		ObjectName: "RC:StsMsg",
		Content:    `{"status":"typing"}`,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/statusmessage/group/publish.json")
	assertEqual(t, call.Params["fromUserId"], "u1")
	assertEqual(t, call.Params["toGroupId"], "g1,g2")
	assertEqual(t, call.Params["objectName"], "RC:StsMsg")
	assertEqual(t, call.Params["content"], `{"status":"typing"}`)
}

func TestSendStatusMessage_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("fail")
	}
	a := NewAPI(mock)

	_, err := a.SendStatusMessage(&SendStatusMessageReq{FromUserId: "u1", ToGroupId: []string{"g1"}, ObjectName: "RC:StsMsg", Content: "{}"})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestPrivateRecallMessage(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.PrivateRecallMessage(&PrivateRecallReq{
		FromUserId: "u1",
		TargetId:   "u2",
		MessageUID: "msg-uid-789",
		SentTime:   1630000000,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/message/private/recall.json")
	assertEqual(t, call.Params["fromUserId"], "u1")
	assertEqual(t, call.Params["targetId"], "u2")
	assertEqual(t, call.Params["messageUID"], "msg-uid-789")
	assertEqual(t, call.Params["sentTime"], "1630000000")
}

func TestPrivateRecallMessage_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("fail")
	}
	a := NewAPI(mock)

	_, err := a.PrivateRecallMessage(&PrivateRecallReq{FromUserId: "u1", TargetId: "u2", MessageUID: "m1", SentTime: 100})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestGroupRecallMessage(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.GroupRecallMessage(&GroupRecallReq{
		FromUserId: "u1",
		TargetId:   "g1",
		MessageUID: "msg-uid-abc",
		SentTime:   1640000000,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/message/group/recall.json")
	assertEqual(t, call.Params["fromUserId"], "u1")
	assertEqual(t, call.Params["targetId"], "g1")
	assertEqual(t, call.Params["messageUID"], "msg-uid-abc")
	assertEqual(t, call.Params["sentTime"], "1640000000")
}

func TestGroupRecallMessage_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("fail")
	}
	a := NewAPI(mock)

	_, err := a.GroupRecallMessage(&GroupRecallReq{FromUserId: "u1", TargetId: "g1", MessageUID: "m1", SentTime: 100})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestExpansionSet(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.ExpansionSet(&ExpansionSetReq{
		MsgUID:           "msg-uid-exp",
		UserId:           "u1",
		ConversationType: "1",
		TargetId:         "u2",
		ExtraKeyVal:      `{"key1":"val1"}`,
		IsSyncSender:     1,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/message/expansion/set.json")
	assertEqual(t, call.Params["msgUID"], "msg-uid-exp")
	assertEqual(t, call.Params["userId"], "u1")
	assertEqual(t, call.Params["conversationType"], "1")
	assertEqual(t, call.Params["targetId"], "u2")
	assertEqual(t, call.Params["extraKeyVal"], `{"key1":"val1"}`)
	assertEqual(t, call.Params["isSyncSender"], "1")
}

func TestExpansionSet_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("fail")
	}
	a := NewAPI(mock)

	_, err := a.ExpansionSet(&ExpansionSetReq{MsgUID: "m1", UserId: "u1", ConversationType: "1", TargetId: "t1", ExtraKeyVal: "{}", IsSyncSender: 0})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestExpansionQuery(t *testing.T) {
	mock := testutil.NewMockClient()
	a := NewAPI(mock)

	resp, err := a.ExpansionQuery("msg-uid-query", 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, resp.Code, 200)

	call := mock.LastCall()
	assertEqual(t, call.Method, "Post")
	assertEqual(t, call.Path, "/message/expansion/query.json")
	assertEqual(t, call.Params["msgUID"], "msg-uid-query")
	assertEqual(t, call.Params["pageNo"], "1")
}

func TestExpansionQuery_Error(t *testing.T) {
	mock := testutil.NewMockClient()
	mock.PostFunc = func(path string, params map[string]string, resp any) error {
		return errors.New("fail")
	}
	a := NewAPI(mock)

	_, err := a.ExpansionQuery("m1", 1)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
