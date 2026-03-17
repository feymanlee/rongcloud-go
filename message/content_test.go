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
