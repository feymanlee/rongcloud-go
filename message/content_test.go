package message

import (
	"encoding/json"
	"testing"
)

func TestTxtMsg_String(t *testing.T) {
	msg := TxtMsg{
		Content: "hello world",
		MentionedInfo: &MentionedInfo{
			Type:       2,
			UserIdList: []string{"u1", "u2"},
		},
		User:  &UserInfo{Id: "sender1", Name: "Alice"},
		Extra: "ext",
	}
	s := msg.String()
	var got TxtMsg
	if err := json.Unmarshal([]byte(s), &got); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	assertEqual(t, got.Content, "hello world")
	assertEqual(t, got.MentionedInfo.Type, 2)
	assertEqual(t, len(got.MentionedInfo.UserIdList), 2)
	assertEqual(t, got.User.Id, "sender1")
	assertEqual(t, got.Extra, "ext")
}

func TestImgMsg_String(t *testing.T) {
	msg := ImgMsg{
		Content:  "base64data",
		Name:     "photo.jpg",
		ImageUri: "https://example.com/photo.jpg",
	}
	s := msg.String()
	var got ImgMsg
	if err := json.Unmarshal([]byte(s), &got); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	assertEqual(t, got.Content, "base64data")
	assertEqual(t, got.Name, "photo.jpg")
	assertEqual(t, got.ImageUri, "https://example.com/photo.jpg")
}

func TestStreamMsg_String(t *testing.T) {
	msg := StreamMsg{
		Content:  "chunk1",
		Seq:      1,
		Complete: false,
		Type:     "markdown",
	}
	s := msg.String()
	var got StreamMsg
	if err := json.Unmarshal([]byte(s), &got); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	assertEqual(t, got.Content, "chunk1")
	assertEqual(t, got.Seq, int64(1))
	assertEqual(t, got.Complete, false)
	assertEqual(t, got.Type, "markdown")
}

func TestGIFMsg_String(t *testing.T) {
	msg := GIFMsg{
		GifDataSize: 102400,
		Name:        "funny.gif",
		RemoteUrl:   "https://example.com/funny.gif",
		Width:       320,
		Height:      240,
	}
	s := msg.String()
	var got GIFMsg
	if err := json.Unmarshal([]byte(s), &got); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	assertEqual(t, got.GifDataSize, 102400)
	assertEqual(t, got.Width, 320)
	assertEqual(t, got.Height, 240)
	assertEqual(t, got.RemoteUrl, "https://example.com/funny.gif")
}

func TestHQVCMsg_String(t *testing.T) {
	msg := HQVCMsg{
		Name:      "voice.aac",
		RemoteUrl: "https://example.com/voice.aac",
		Duration:  30,
	}
	s := msg.String()
	var got HQVCMsg
	if err := json.Unmarshal([]byte(s), &got); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	assertEqual(t, got.Duration, 30)
	assertEqual(t, got.RemoteUrl, "https://example.com/voice.aac")
}

func TestFileMsg_String(t *testing.T) {
	msg := FileMsg{
		Name:    "doc.pdf",
		Size:    "204800",
		Type:    "pdf",
		FileUrl: "https://example.com/doc.pdf",
	}
	s := msg.String()
	var got FileMsg
	if err := json.Unmarshal([]byte(s), &got); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	assertEqual(t, got.Name, "doc.pdf")
	assertEqual(t, got.Size, "204800")
	assertEqual(t, got.Type, "pdf")
	assertEqual(t, got.FileUrl, "https://example.com/doc.pdf")
}

func TestSightMsg_String(t *testing.T) {
	msg := SightMsg{
		SightUrl: "https://example.com/video.mp4",
		Content:  "thumb-base64",
		Duration: 60,
		Size:     "5242880",
		Name:     "video.mp4",
	}
	s := msg.String()
	var got SightMsg
	if err := json.Unmarshal([]byte(s), &got); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	assertEqual(t, got.SightUrl, "https://example.com/video.mp4")
	assertEqual(t, got.Duration, 60)
	assertEqual(t, got.Size, "5242880")
}

func TestLBSMsg_String(t *testing.T) {
	msg := LBSMsg{
		Content:   "thumb-base64",
		Latitude:  39.9042,
		Longitude: 116.4074,
		Poi:       "北京天安门",
	}
	s := msg.String()
	var got LBSMsg
	if err := json.Unmarshal([]byte(s), &got); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	assertEqual(t, got.Latitude, 39.9042)
	assertEqual(t, got.Longitude, 116.4074)
	assertEqual(t, got.Poi, "北京天安门")
}

func TestReferenceMsg_String(t *testing.T) {
	msg := ReferenceMsg{
		Content:        "reply content",
		ReferMsgUserId: "u1",
		ReferMsg:       `{"content":"original"}`,
		ObjName:        "RC:TxtMsg",
		MentionedInfo: &MentionedInfo{
			Type:       1,
		},
	}
	s := msg.String()
	var got ReferenceMsg
	if err := json.Unmarshal([]byte(s), &got); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	assertEqual(t, got.Content, "reply content")
	assertEqual(t, got.ReferMsgUserId, "u1")
	assertEqual(t, got.ObjName, "RC:TxtMsg")
	assertEqual(t, got.MentionedInfo.Type, 1)
}

func TestCombineMsg_String(t *testing.T) {
	msg := CombineMsg{
		RemoteUrl:        "https://example.com/combine.html",
		ConversationType: 3,
		NameList:         "Alice,Bob",
		SummaryList:      "hello,world",
	}
	s := msg.String()
	var got CombineMsg
	if err := json.Unmarshal([]byte(s), &got); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	assertEqual(t, got.ConversationType, 3)
	assertEqual(t, got.RemoteUrl, "https://example.com/combine.html")
}

func TestImgTextMsg_String(t *testing.T) {
	msg := ImgTextMsg{
		Title:    "News Title",
		Content:  "News content",
		ImageUri: "https://example.com/img.jpg",
		Url:      "https://example.com/article",
	}
	s := msg.String()
	var got ImgTextMsg
	if err := json.Unmarshal([]byte(s), &got); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	assertEqual(t, got.Title, "News Title")
	assertEqual(t, got.Content, "News content")
	assertEqual(t, got.Url, "https://example.com/article")
}

// ---------------------------------------------------------------------------
// 通知类消息测试
// ---------------------------------------------------------------------------

func TestRcNtf_String(t *testing.T) {
	msg := RcNtf{
		OperatorId:         "admin1",
		RecallTime:         1620000000000,
		OriginalObjectName: "RC:TxtMsg",
		Admin:              true,
		Delete:             false,
	}
	s := msg.String()
	var got RcNtf
	if err := json.Unmarshal([]byte(s), &got); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	assertEqual(t, got.OperatorId, "admin1")
	assertEqual(t, got.RecallTime, int64(1620000000000))
	assertEqual(t, got.OriginalObjectName, "RC:TxtMsg")
	assertEqual(t, got.Admin, true)
	assertEqual(t, got.Delete, false)
}

func TestContactNtf_String(t *testing.T) {
	msg := ContactNtf{
		Operation:    "Request",
		SourceUserId: "u1",
		TargetUserId: "u2",
		Message:      "let's be friends",
	}
	s := msg.String()
	var got ContactNtf
	if err := json.Unmarshal([]byte(s), &got); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	assertEqual(t, got.Operation, "Request")
	assertEqual(t, got.SourceUserId, "u1")
	assertEqual(t, got.TargetUserId, "u2")
	assertEqual(t, got.Message, "let's be friends")
}

func TestProfileNtf_String(t *testing.T) {
	msg := ProfileNtf{Operation: "Update", Data: `{"name":"new"}`}
	s := msg.String()
	var got ProfileNtf
	if err := json.Unmarshal([]byte(s), &got); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	assertEqual(t, got.Operation, "Update")
	assertEqual(t, got.Data, `{"name":"new"}`)
}

func TestInfoNtf_String(t *testing.T) {
	msg := InfoNtf{Message: "you joined the group"}
	s := msg.String()
	var got InfoNtf
	if err := json.Unmarshal([]byte(s), &got); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	assertEqual(t, got.Message, "you joined the group")
}

func TestGrpNtf_String(t *testing.T) {
	msg := GrpNtf{
		OperatorUserId: "admin",
		Operation:      "Add",
		Data:           `{"targetUserIds":["u1"]}`,
		Message:        "admin added u1",
	}
	s := msg.String()
	var got GrpNtf
	if err := json.Unmarshal([]byte(s), &got); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	assertEqual(t, got.OperatorUserId, "admin")
	assertEqual(t, got.Operation, "Add")
}

func TestCmdNtf_String(t *testing.T) {
	msg := CmdNtf{Name: "refreshUI", Data: `{"page":"home"}`}
	s := msg.String()
	var got CmdNtf
	if err := json.Unmarshal([]byte(s), &got); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	assertEqual(t, got.Name, "refreshUI")
	assertEqual(t, got.Data, `{"page":"home"}`)
}

// ---------------------------------------------------------------------------
// 信令类消息测试
// ---------------------------------------------------------------------------

func TestCmdMsg_String(t *testing.T) {
	msg := CmdMsg{Name: "clearCache", Data: `{"scope":"all"}`}
	s := msg.String()
	var got CmdMsg
	if err := json.Unmarshal([]byte(s), &got); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	assertEqual(t, got.Name, "clearCache")
	assertEqual(t, got.Data, `{"scope":"all"}`)
}

func TestRcCmd_String(t *testing.T) {
	msg := RcCmd{MessageUId: "msg-123", SentTime: 1620000000000}
	s := msg.String()
	var got RcCmd
	if err := json.Unmarshal([]byte(s), &got); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	assertEqual(t, got.MessageUId, "msg-123")
	assertEqual(t, got.SentTime, int64(1620000000000))
}

func TestReadNtf_String(t *testing.T) {
	msg := ReadNtf{MessageUId: "msg-456", LastMessageSend: 1620000000000, Type: 1}
	s := msg.String()
	var got ReadNtf
	if err := json.Unmarshal([]byte(s), &got); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	assertEqual(t, got.MessageUId, "msg-456")
	assertEqual(t, got.Type, 1)
}

func TestRRReqMsg_String(t *testing.T) {
	msg := RRReqMsg{MessageUId: "msg-789"}
	s := msg.String()
	var got RRReqMsg
	if err := json.Unmarshal([]byte(s), &got); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	assertEqual(t, got.MessageUId, "msg-789")
}

func TestRRRspMsg_String(t *testing.T) {
	msg := RRRspMsg{ReceiptMessageDic: map[string]int64{"msg-1": 1620000000000}}
	s := msg.String()
	var got RRRspMsg
	if err := json.Unmarshal([]byte(s), &got); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	assertEqual(t, got.ReceiptMessageDic["msg-1"], int64(1620000000000))
}

func TestChrmKVNotiMsg_String(t *testing.T) {
	msg := ChrmKVNotiMsg{Type: 1, Key: "theme", Value: "dark"}
	s := msg.String()
	var got ChrmKVNotiMsg
	if err := json.Unmarshal([]byte(s), &got); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	assertEqual(t, got.Type, 1)
	assertEqual(t, got.Key, "theme")
	assertEqual(t, got.Value, "dark")
}

func TestVCInvite_String(t *testing.T) {
	msg := VCInvite{CallId: "call-1", MediaType: "video"}
	s := msg.String()
	var got VCInvite
	if err := json.Unmarshal([]byte(s), &got); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	assertEqual(t, got.CallId, "call-1")
	assertEqual(t, got.MediaType, "video")
}

func TestVCHangup_String(t *testing.T) {
	msg := VCHangup{CallId: "call-1", Reason: 1}
	s := msg.String()
	var got VCHangup
	if err := json.Unmarshal([]byte(s), &got); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	assertEqual(t, got.CallId, "call-1")
	assertEqual(t, got.Reason, 1)
}

// ---------------------------------------------------------------------------
// 状态类消息测试
// ---------------------------------------------------------------------------

func TestTypSts_String(t *testing.T) {
	msg := TypSts{TypingContentType: "RC:TxtMsg"}
	s := msg.String()
	var got TypSts
	if err := json.Unmarshal([]byte(s), &got); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	assertEqual(t, got.TypingContentType, "RC:TxtMsg")
}

func TestTxtMsg_String_OmitsEmpty(t *testing.T) {
	msg := TxtMsg{Content: "simple"}
	s := msg.String()
	var raw map[string]any
	if err := json.Unmarshal([]byte(s), &raw); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if _, ok := raw["user"]; ok {
		t.Error("user should be omitted when nil")
	}
	if _, ok := raw["mentionedInfo"]; ok {
		t.Error("mentionedInfo should be omitted when nil")
	}
	if _, ok := raw["extra"]; ok {
		t.Error("extra should be omitted when empty")
	}
}
