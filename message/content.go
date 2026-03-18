package message

import "encoding/json"

type ObjectName string

// ObjectName 常量定义 —— 用户内容类消息
const (
	ObjectNameTxtMsg       ObjectName = "RC:TxtMsg"       // 文字消息
	ObjectNameImgMsg       ObjectName = "RC:ImgMsg"       // 图片消息
	ObjectNameStreamMsg    ObjectName = "RC:StreamMsg"    // 流式消息
	ObjectNameGIFMsg       ObjectName = "RC:GIFMsg"       // GIF 图片消息
	ObjectNameHQVCMsg      ObjectName = "RC:HQVCMsg"      // 高清语音消息
	ObjectNameFileMsg      ObjectName = "RC:FileMsg"      // 文件消息
	ObjectNameSightMsg     ObjectName = "RC:SightMsg"     // 小视频消息
	ObjectNameLBSMsg       ObjectName = "RC:LBSMsg"       // 位置消息
	ObjectNameReferenceMsg ObjectName = "RC:ReferenceMsg" // 引用消息
	ObjectNameCombineMsg   ObjectName = "RC:CombineMsg"   // 合并转发消息
	ObjectNameImgTextMsg   ObjectName = "RC:ImgTextMsg"   // 图文消息
)

// ObjectName 常量定义 —— 通知类消息
const (
	ObjectNameRcNtf      ObjectName = "RC:RcNtf"      // 撤回通知消息
	ObjectNameContactNtf ObjectName = "RC:ContactNtf" // 联系人（好友）通知消息
	ObjectNameProfileNtf ObjectName = "RC:ProfileNtf" // 资料通知消息
	ObjectNameInfoNtf    ObjectName = "RC:InfoNtf"    // 提示条通知消息
	ObjectNameGrpNtf     ObjectName = "RC:GrpNtf"     // 群组通知消息
	ObjectNameCmdNtf     ObjectName = "RC:CmdNtf"     // 命令提醒消息
)

// ObjectName 常量定义 —— 信令类消息
const (
	ObjectNameCmdMsg        ObjectName = "RC:CmdMsg"        // 命令消息
	ObjectNameRcCmd         ObjectName = "RC:RcCmd"         // 撤回命令消息
	ObjectNameReadNtf       ObjectName = "RC:ReadNtf"       // 单聊已读回执消息
	ObjectNameRRReqMsg      ObjectName = "RC:RRReqMsg"      // 群聊已读回执请求消息
	ObjectNameRRRspMsg      ObjectName = "RC:RRRspMsg"      // 群聊已读回执响应消息
	ObjectNameSRSMsg        ObjectName = "RC:SRSMsg"        // 多端已读状态同步消息
	ObjectNameChrmKVNotiMsg ObjectName = "RC:chrmKVNotiMsg" // 聊天室属性通知消息
	ObjectNameMsgExMsg      ObjectName = "RC:MsgExMsg"      // 消息扩展功能消息
	ObjectNameVCAccept      ObjectName = "RC:VCAccept"      // 实时音视频接受信令
	ObjectNameVCHangup      ObjectName = "RC:VCHangup"      // 实时音视频挂断信令
	ObjectNameVCInvite      ObjectName = "RC:VCInvite"      // 实时音视频邀请信令
	ObjectNameVCModifyMedia ObjectName = "RC:VCModifyMedia" // 实时音视频切换信令
	ObjectNameVCModifyMem   ObjectName = "RC:VCModifyMem"   // 实时音视频成员变化信令
	ObjectNameVCRinging     ObjectName = "RC:VCRinging"     // 实时音视频响铃信令
)

// ObjectName 常量定义 —— 状态类消息
const (
	ObjectNameTypSts ObjectName = "RC:TypSts" // 正在输入状态消息
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
	Content  string    `json:"content"`         // Base64 缩略图（建议 5KB，最大 10KB）
	Name     string    `json:"name,omitempty"`  // 文件名
	ImageUri string    `json:"imageUri"`        // 图片远程地址
	User     *UserInfo `json:"user,omitempty"`  // 发送者信息
	Extra    string    `json:"extra,omitempty"` // 扩展信息
}

func (m ImgMsg) String() string { b, _ := json.Marshal(m); return string(b) }

// StreamMsg 流式消息
type StreamMsg struct {
	Content        string         `json:"content"`                  // 流式数据片段
	Seq            int64          `json:"seq"`                      // 序列号（>0，递增）
	Complete       bool           `json:"complete"`                 // 流结束标志
	CompleteReason int            `json:"completeReason,omitempty"` // 自定义结束错误码
	Type           string         `json:"type,omitempty"`           // 流类型：text/markdown/html（默认 text）
	MessageUID     string         `json:"messageUID,omitempty"`     // 流式消息 ID（续传时使用）
	MentionedInfo  *MentionedInfo `json:"mentionedInfo,omitempty"`  // @ 消息信息（仅首包）
	User           *UserInfo      `json:"user,omitempty"`           // 发送者信息（仅首包）
	Extra          string         `json:"extra,omitempty"`          // 扩展信息（仅首包）
}

func (m StreamMsg) String() string { b, _ := json.Marshal(m); return string(b) }

// GIFMsg GIF 图片消息
type GIFMsg struct {
	GifDataSize int       `json:"gifDataSize"`     // GIF 文件大小（字节）
	Name        string    `json:"name,omitempty"`  // 文件名
	RemoteUrl   string    `json:"remoteUrl"`       // GIF 远程地址
	Width       int       `json:"width"`           // 宽度（像素）
	Height      int       `json:"height"`          // 高度（像素）
	User        *UserInfo `json:"user,omitempty"`  // 发送者信息
	Extra       string    `json:"extra,omitempty"` // 扩展信息
}

func (m GIFMsg) String() string { b, _ := json.Marshal(m); return string(b) }

// HQVCMsg 高清语音消息
type HQVCMsg struct {
	Name      string    `json:"name,omitempty"`  // 文件名
	RemoteUrl string    `json:"remoteUrl"`       // 语音远程地址（AAC 格式）
	Duration  int       `json:"duration"`        // 语音时长（秒），最长 60 秒
	User      *UserInfo `json:"user,omitempty"`  // 发送者信息
	Extra     string    `json:"extra,omitempty"` // 扩展信息
}

func (m HQVCMsg) String() string { b, _ := json.Marshal(m); return string(b) }

// FileMsg 文件消息
type FileMsg struct {
	Name    string    `json:"name,omitempty"`  // 文件名
	Size    string    `json:"size"`            // 文件大小（字节）
	Type    string    `json:"type"`            // 文件类型
	FileUrl string    `json:"fileUrl"`         // 文件远程地址
	User    *UserInfo `json:"user,omitempty"`  // 发送者信息
	Extra   string    `json:"extra,omitempty"` // 扩展信息
}

func (m FileMsg) String() string { b, _ := json.Marshal(m); return string(b) }

// SightMsg 小视频消息
type SightMsg struct {
	SightUrl string    `json:"sightUrl"`        // 视频远程地址
	Content  string    `json:"content"`         // Base64 缩略图（建议 5KB，最大 10KB）
	Duration int       `json:"duration"`        // 视频时长（秒），默认最长 120 秒
	Size     string    `json:"size"`            // 视频大小（字节）
	Name     string    `json:"name"`            // 视频文件名（MP4 格式）
	User     *UserInfo `json:"user,omitempty"`  // 发送者信息
	Extra    string    `json:"extra,omitempty"` // 扩展信息
}

func (m SightMsg) String() string { b, _ := json.Marshal(m); return string(b) }

// LBSMsg 位置消息
type LBSMsg struct {
	Content   string    `json:"content"`         // 位置缩略图（Base64 编码的 JPG）
	Latitude  float64   `json:"latitude"`        // 纬度
	Longitude float64   `json:"longitude"`       // 经度
	Poi       string    `json:"poi"`             // 兴趣点信息
	User      *UserInfo `json:"user,omitempty"`  // 发送者信息
	Extra     string    `json:"extra,omitempty"` // 扩展信息
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
	Title    string    `json:"title"`           // 消息标题
	Content  string    `json:"content"`         // 文字内容
	ImageUri string    `json:"imageUri"`        // 图片地址（120x120 像素）
	Url      string    `json:"url"`             // 跳转链接
	User     *UserInfo `json:"user,omitempty"`  // 发送者信息
	Extra    string    `json:"extra,omitempty"` // 扩展信息
}

func (m ImgTextMsg) String() string { b, _ := json.Marshal(m); return string(b) }

// ---------------------------------------------------------------------------
// 通知类消息
// ---------------------------------------------------------------------------

// RcNtf 撤回通知消息
type RcNtf struct {
	OperatorId             string    `json:"operatorId,omitempty"`             // 执行撤回的用户 ID
	RecallTime             int64     `json:"recallTime"`                       // 被撤回消息的发送时间（毫秒）
	OriginalObjectName     string    `json:"originalObjectName"`               // 被撤回消息的消息类型
	OriginalMessageContent string    `json:"originalMessageContent,omitempty"` // 被撤回消息的内容
	RecallContent          string    `json:"recallContent,omitempty"`          // 撤回提示文本
	RecallActionTime       int64     `json:"recallActionTime,omitempty"`       // 撤回操作时间（毫秒）
	Admin                  bool      `json:"admin"`                            // 是否为管理员操作
	Delete                 bool      `json:"delete"`                           // 移动端是否删除原消息记录
	User                   *UserInfo `json:"user,omitempty"`                   // 发送者信息
	Extra                  string    `json:"extra,omitempty"`                  // 扩展信息
}

func (m RcNtf) String() string { b, _ := json.Marshal(m); return string(b) }

// ContactNtf 联系人（好友）通知消息
type ContactNtf struct {
	Operation    string `json:"operation"`       // 操作命令：Request、AcceptResponse、RejectResponse
	SourceUserId string `json:"sourceUserId"`    // 发起通知的用户 ID
	TargetUserId string `json:"targetUserId"`    // 接收通知的用户 ID
	Message      string `json:"message"`         // 请求或响应消息内容
	Extra        string `json:"extra,omitempty"` // 扩展信息
}

func (m ContactNtf) String() string { b, _ := json.Marshal(m); return string(b) }

// ProfileNtf 资料通知消息
type ProfileNtf struct {
	Operation string `json:"operation"`       // 资料通知操作类型
	Data      string `json:"data"`            // 操作的数据
	Extra     string `json:"extra,omitempty"` // 扩展信息
}

func (m ProfileNtf) String() string { b, _ := json.Marshal(m); return string(b) }

// InfoNtf 提示条通知消息
type InfoNtf struct {
	Message string `json:"message"`         // 提示条消息内容
	Extra   string `json:"extra,omitempty"` // 扩展信息
}

func (m InfoNtf) String() string { b, _ := json.Marshal(m); return string(b) }

// GrpNtf 群组通知消息
type GrpNtf struct {
	OperatorUserId string `json:"operatorUserId"`  // 操作人用户 ID
	Operation      string `json:"operation"`       // 操作名称：Create、Rename、Add、Kicked、Quit、Dismiss
	Data           string `json:"data"`            // 操作数据
	Message        string `json:"message"`         // 消息内容
	Extra          string `json:"extra,omitempty"` // 扩展信息
}

func (m GrpNtf) String() string { b, _ := json.Marshal(m); return string(b) }

// CmdNtf 命令提醒消息
type CmdNtf struct {
	Name string `json:"name"`           // 命令名称
	Data string `json:"data,omitempty"` // 命令数据
}

func (m CmdNtf) String() string { b, _ := json.Marshal(m); return string(b) }

// ---------------------------------------------------------------------------
// 信令类消息
// ---------------------------------------------------------------------------

// CmdMsg 命令消息
type CmdMsg struct {
	Name string `json:"name"`           // 命令名称
	Data string `json:"data,omitempty"` // 命令数据
}

func (m CmdMsg) String() string { b, _ := json.Marshal(m); return string(b) }

// RcCmd 撤回命令消息
type RcCmd struct {
	MessageUId string `json:"messageUId"` // 被撤回消息的唯一标识
	SentTime   int64  `json:"sentTime"`   // 被撤回消息的发送时间（毫秒）
}

func (m RcCmd) String() string { b, _ := json.Marshal(m); return string(b) }

// ReadNtf 单聊已读回执消息
type ReadNtf struct {
	MessageUId      string `json:"messageUId"`      // 最后一条已读消息的唯一标识
	LastMessageSend int64  `json:"lastMessageSend"` // 最后一条已读消息的发送时间（毫秒）
	Type            int    `json:"type"`            // 会话类型
}

func (m ReadNtf) String() string { b, _ := json.Marshal(m); return string(b) }

// RRReqMsg 群聊已读回执请求消息
type RRReqMsg struct {
	MessageUId string `json:"messageUId"` // 请求已读回执的消息唯一标识
}

func (m RRReqMsg) String() string { b, _ := json.Marshal(m); return string(b) }

// RRRspMsg 群聊已读回执响应消息
type RRRspMsg struct {
	ReceiptMessageDic map[string]int64 `json:"receiptMessageDic"` // 已读消息 UID 与已读时间的映射
}

func (m RRRspMsg) String() string { b, _ := json.Marshal(m); return string(b) }

// SRSMsg 多端已读状态同步消息
type SRSMsg struct {
	LastMessageSend int64 `json:"lastMessageSend"` // 最后一条已读消息的发送时间（毫秒）
}

func (m SRSMsg) String() string { b, _ := json.Marshal(m); return string(b) }

// ChrmKVNotiMsg 聊天室属性通知消息
type ChrmKVNotiMsg struct {
	Type  int               `json:"type"`            // 操作类型：1 设置，2 删除
	Key   string            `json:"key,omitempty"`   // 属性名
	Value string            `json:"value,omitempty"` // 属性值
	Extra map[string]string `json:"extra,omitempty"` // 扩展信息
}

func (m ChrmKVNotiMsg) String() string { b, _ := json.Marshal(m); return string(b) }

// MsgExMsg 消息扩展功能消息
type MsgExMsg struct {
	MessageUId string            `json:"messageUId"` // 扩展的消息唯一标识
	ExtraKey   map[string]string `json:"extraKey"`   // 扩展的 KV 对
}

func (m MsgExMsg) String() string { b, _ := json.Marshal(m); return string(b) }

// VCInvite 实时音视频邀请信令
type VCInvite struct {
	CallId    string `json:"callId"`          // 通话 ID
	MediaType string `json:"mediaType"`       // 媒体类型：audio、video
	Extra     string `json:"extra,omitempty"` // 扩展信息
}

func (m VCInvite) String() string { b, _ := json.Marshal(m); return string(b) }

// VCAccept 实时音视频接受信令
type VCAccept struct {
	CallId    string `json:"callId"`          // 通话 ID
	MediaType string `json:"mediaType"`       // 媒体类型：audio、video
	Extra     string `json:"extra,omitempty"` // 扩展信息
}

func (m VCAccept) String() string { b, _ := json.Marshal(m); return string(b) }

// VCRinging 实时音视频响铃信令
type VCRinging struct {
	CallId string `json:"callId"`          // 通话 ID
	Extra  string `json:"extra,omitempty"` // 扩展信息
}

func (m VCRinging) String() string { b, _ := json.Marshal(m); return string(b) }

// VCHangup 实时音视频挂断信令
type VCHangup struct {
	CallId string `json:"callId"`          // 通话 ID
	Reason int    `json:"reason"`          // 挂断原因
	Extra  string `json:"extra,omitempty"` // 扩展信息
}

func (m VCHangup) String() string { b, _ := json.Marshal(m); return string(b) }

// VCModifyMedia 实时音视频切换信令
type VCModifyMedia struct {
	CallId    string `json:"callId"`          // 通话 ID
	MediaType string `json:"mediaType"`       // 切换后的媒体类型：audio、video
	Extra     string `json:"extra,omitempty"` // 扩展信息
}

func (m VCModifyMedia) String() string { b, _ := json.Marshal(m); return string(b) }

// VCModifyMem 实时音视频成员变化信令
type VCModifyMem struct {
	CallId       string   `json:"callId"`                 // 通话 ID
	ModifyMemIds []string `json:"modifyMemIds,omitempty"` // 变化的成员 ID 列表
	Extra        string   `json:"extra,omitempty"`        // 扩展信息
}

func (m VCModifyMem) String() string { b, _ := json.Marshal(m); return string(b) }

// ---------------------------------------------------------------------------
// 状态类消息
// ---------------------------------------------------------------------------

// TypSts 正在输入状态消息
type TypSts struct {
	TypingContentType string `json:"typingContentType"` // 正在输入的消息类型
}

func (m TypSts) String() string { b, _ := json.Marshal(m); return string(b) }
