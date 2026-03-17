package message

import "encoding/json"

// ObjectName 常量定义
const (
	ObjectNameTxtMsg       = "RC:TxtMsg"       // 文字消息
	ObjectNameImgMsg       = "RC:ImgMsg"       // 图片消息
	ObjectNameStreamMsg    = "RC:StreamMsg"    // 流式消息
	ObjectNameGIFMsg       = "RC:GIFMsg"       // GIF 图片消息
	ObjectNameHQVCMsg      = "RC:HQVCMsg"      // 高清语音消息
	ObjectNameFileMsg      = "RC:FileMsg"      // 文件消息
	ObjectNameSightMsg     = "RC:SightMsg"     // 小视频消息
	ObjectNameLBSMsg       = "RC:LBSMsg"       // 位置消息
	ObjectNameReferenceMsg = "RC:ReferenceMsg" // 引用消息
	ObjectNameCombineMsg   = "RC:CombineMsg"   // 合并转发消息
	ObjectNameImgTextMsg   = "RC:ImgTextMsg"   // 图文消息
)

// MentionedInfo @ 消息信息
type MentionedInfo struct {
	Type             int      `json:"type"`                       // @ 类型：1 全部，2 指定用户
	UserIdList       []string `json:"userIdList,omitempty"`       // 被 @ 的用户 ID 列表
	MentionedContent string   `json:"mentionedContent,omitempty"` // 自定义 @ 推送内容
}

// UserInfo 发送者用户信息
type UserInfo struct {
	Id       string `json:"id,omitempty"`       // 用户 ID
	Name     string `json:"name,omitempty"`     // 用户昵称
	Portrait string `json:"portrait,omitempty"` // 用户头像
	Extra    string `json:"extra,omitempty"`    // 扩展信息
}

// TxtMsg 文字消息
type TxtMsg struct {
	Content       string         `json:"content"`                 // 文字内容，包括表情
	MentionedInfo *MentionedInfo `json:"mentionedInfo,omitempty"` // @ 消息信息
	User          *UserInfo      `json:"user,omitempty"`          // 发送者信息
	Extra         string         `json:"extra,omitempty"`         // 扩展信息
}

func (m TxtMsg) String() string { b, _ := json.Marshal(m); return string(b) }

// ImgMsg 图片消息
type ImgMsg struct {
	Content  string    `json:"content"`            // Base64 缩略图（建议 5KB，最大 10KB）
	Name     string    `json:"name,omitempty"`     // 文件名
	ImageUri string    `json:"imageUri"`           // 图片远程地址
	User     *UserInfo `json:"user,omitempty"`     // 发送者信息
	Extra    string    `json:"extra,omitempty"`    // 扩展信息
}

func (m ImgMsg) String() string { b, _ := json.Marshal(m); return string(b) }

// StreamMsg 流式消息
type StreamMsg struct {
	Content        string         `json:"content"`                 // 流式数据片段
	Seq            int64          `json:"seq"`                     // 序列号（>0，递增）
	Complete       bool           `json:"complete"`                // 流结束标志
	CompleteReason int            `json:"completeReason,omitempty"` // 自定义结束错误码
	Type           string         `json:"type,omitempty"`          // 流类型：text/markdown/html（默认 text）
	MessageUID     string         `json:"messageUID,omitempty"`    // 流式消息 ID（续传时使用）
	MentionedInfo  *MentionedInfo `json:"mentionedInfo,omitempty"` // @ 消息信息（仅首包）
	User           *UserInfo      `json:"user,omitempty"`          // 发送者信息（仅首包）
	Extra          string         `json:"extra,omitempty"`         // 扩展信息（仅首包）
}

func (m StreamMsg) String() string { b, _ := json.Marshal(m); return string(b) }

// GIFMsg GIF 图片消息
type GIFMsg struct {
	GifDataSize int       `json:"gifDataSize"`        // GIF 文件大小（字节）
	Name        string    `json:"name,omitempty"`     // 文件名
	RemoteUrl   string    `json:"remoteUrl"`          // GIF 远程地址
	Width       int       `json:"width"`              // 宽度（像素）
	Height      int       `json:"height"`             // 高度（像素）
	User        *UserInfo `json:"user,omitempty"`     // 发送者信息
	Extra       string    `json:"extra,omitempty"`    // 扩展信息
}

func (m GIFMsg) String() string { b, _ := json.Marshal(m); return string(b) }

// HQVCMsg 高清语音消息
type HQVCMsg struct {
	Name      string    `json:"name,omitempty"`     // 文件名
	RemoteUrl string    `json:"remoteUrl"`          // 语音远程地址（AAC 格式）
	Duration  int       `json:"duration"`           // 语音时长（秒），最长 60 秒
	User      *UserInfo `json:"user,omitempty"`     // 发送者信息
	Extra     string    `json:"extra,omitempty"`    // 扩展信息
}

func (m HQVCMsg) String() string { b, _ := json.Marshal(m); return string(b) }

// FileMsg 文件消息
type FileMsg struct {
	Name    string    `json:"name,omitempty"`     // 文件名
	Size    string    `json:"size"`               // 文件大小（字节）
	Type    string    `json:"type"`               // 文件类型
	FileUrl string    `json:"fileUrl"`            // 文件远程地址
	User    *UserInfo `json:"user,omitempty"`     // 发送者信息
	Extra   string    `json:"extra,omitempty"`    // 扩展信息
}

func (m FileMsg) String() string { b, _ := json.Marshal(m); return string(b) }

// SightMsg 小视频消息
type SightMsg struct {
	SightUrl string    `json:"sightUrl"`           // 视频远程地址
	Content  string    `json:"content"`            // Base64 缩略图（建议 5KB，最大 10KB）
	Duration int       `json:"duration"`           // 视频时长（秒），默认最长 120 秒
	Size     string    `json:"size"`               // 视频大小（字节）
	Name     string    `json:"name"`               // 视频文件名（MP4 格式）
	User     *UserInfo `json:"user,omitempty"`     // 发送者信息
	Extra    string    `json:"extra,omitempty"`    // 扩展信息
}

func (m SightMsg) String() string { b, _ := json.Marshal(m); return string(b) }

// LBSMsg 位置消息
type LBSMsg struct {
	Content   string    `json:"content"`            // 位置缩略图（Base64 编码的 JPG）
	Latitude  float64   `json:"latitude"`           // 纬度
	Longitude float64   `json:"longitude"`          // 经度
	Poi       string    `json:"poi"`                // 兴趣点信息
	User      *UserInfo `json:"user,omitempty"`     // 发送者信息
	Extra     string    `json:"extra,omitempty"`    // 扩展信息
}

func (m LBSMsg) String() string { b, _ := json.Marshal(m); return string(b) }

// ReferenceMsg 引用消息（回复）
type ReferenceMsg struct {
	Content        string         `json:"content"`                 // 回复文字内容
	ReferMsgUserId string         `json:"referMsgUserId"`          // 被引用消息的发送者 ID
	ReferMsg       string         `json:"referMsg"`                // 被引用消息的 JSON 结构
	ObjName        string         `json:"objName"`                 // 被引用消息类型
	MentionedInfo  *MentionedInfo `json:"mentionedInfo,omitempty"` // @ 消息信息
	User           *UserInfo      `json:"user,omitempty"`          // 发送者信息
	Extra          string         `json:"extra,omitempty"`         // 扩展信息
}

func (m ReferenceMsg) String() string { b, _ := json.Marshal(m); return string(b) }

// CombineMsg 合并转发消息
type CombineMsg struct {
	RemoteUrl        string `json:"remoteUrl"`        // 合并消息 HTML 文件远程地址
	ConversationType int    `json:"conversationType"` // 会话类型：1 单聊，3 群聊
	NameList         string `json:"nameList"`         // 前 4 条消息发送者名称
	SummaryList      string `json:"summaryList"`      // 前 4 条消息摘要
}

func (m CombineMsg) String() string { b, _ := json.Marshal(m); return string(b) }

// ImgTextMsg 图文消息
type ImgTextMsg struct {
	Title    string    `json:"title"`              // 消息标题
	Content  string    `json:"content"`            // 文字内容
	ImageUri string    `json:"imageUri"`           // 图片地址（120x120 像素）
	Url      string    `json:"url"`                // 跳转链接
	User     *UserInfo `json:"user,omitempty"`     // 发送者信息
	Extra    string    `json:"extra,omitempty"`    // 扩展信息
}

func (m ImgTextMsg) String() string { b, _ := json.Marshal(m); return string(b) }
