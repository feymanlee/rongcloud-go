package message

import "github.com/feymanlee/rongcloud-go/internal/types"

// SendResp 发送消息响应
type SendResp struct {
	types.BaseResp
	MessageUID string `json:"messageUID,omitempty"`
}

// RecallResp 撤回消息响应
type RecallResp struct {
	types.BaseResp
}

// HistoryResp 历史消息查询响应
type HistoryResp struct {
	types.BaseResp
	URL string `json:"url"`
}

// HistoryDeleteResp 历史消息删除响应
type HistoryDeleteResp struct {
	types.BaseResp
}

// SendPrivateReq 发送单聊消息请求参数
type SendPrivateReq struct {
	FromUserId           string   // 发送人用户 ID
	ToUserId             []string // 接收用户 ID，最多 1000 人
	ObjectName           string   // 消息类型
	Content              string   // 消息内容（JSON 字符串），最大 128KB
	PushContent          string   // 推送显示内容
	PushData             string   // 推送附加信息
	PushExt              string   // 推送通知配置（JSON 字符串），如标题、模板 ID、渠道 ID 等
	DisablePush          bool     // 是否为静默消息，true 不发送离线推送通知
	Count                int      // iOS 角标数，仅单个接收人且 iOS 端有效
	ContentAvailable     int      // iOS 静默推送，1 开启，0 关闭（默认 0）
	IsPersisted          *int     // 是否在历史消息云端存储，0 不存储，1 存储（默认 1）
	IsIncludeSender      *int     // 是否同步给发送人客户端，1 同步，0 不同步（默认 0）
	DisableUpdateLastMsg bool     // 是否不更新会话最后一条消息
	Expansion            bool     // 是否开启消息扩展，默认 false
	ExtraContent         string   // 消息扩展自定义 KV 对（JSON 字符串），最多 100 个
	NeedReadReceipt      int      // 是否需要已读回执，1 需要，0 不需要（默认 0）
	VerifyBlacklist      int      // 是否过滤接收人黑名单，0 不过滤，1 过滤（默认 0）
}

// SendGroupReq 发送群组消息请求参数
type SendGroupReq struct {
	FromUserId           string   // 发送人用户 ID
	ToGroupId            []string // 接收群组 ID，最多 3 个（定向消息时仅支持 1 个）
	ToUserId             []string // 定向消息的接收成员 ID（仅单个群组时有效）
	ObjectName           string   // 消息类型
	Content              string   // 消息内容（JSON 字符串），最大 128KB
	PushContent          string   // 推送显示内容
	PushData             string   // 推送附加信息
	PushExt              string   // 推送通知配置（JSON 字符串），如标题、模板 ID、渠道 ID 等
	DisablePush          bool     // 是否为静默消息，true 不发送离线推送通知
	IsIncludeSender      *int     // 是否同步给发送人客户端，1 同步，0 不同步（默认 0）
	IsPersisted          *int     // 是否在历史消息云端存储，0 不存储，1 存储（默认 1）
	IsMentioned          int      // 是否为 @ 消息，1 是，0 不是
	ContentAvailable     int      // iOS 静默推送，1 开启，0 关闭（默认 0）
	Expansion            bool     // 是否开启消息扩展，默认 false
	ExtraContent         string   // 消息扩展自定义 KV 对（JSON 字符串），最多 100 个
	DisableUpdateLastMsg bool     // 是否不更新会话最后一条消息
	NeedReadReceipt      int      // 是否需要已读回执，1 需要，0 不需要（默认 0）
}

// SendChatroomReq 发送聊天室消息请求参数
type SendChatroomReq struct {
	FromUserId   string   // 发送人用户 ID
	ToChatroomId []string // 接收聊天室 ID
	ObjectName   string   // 消息类型
	Content      string   // 消息内容（JSON 字符串）
}

// SendSystemReq 发送系统消息请求参数
type SendSystemReq struct {
	FromUserId  string   // 发送人用户 ID
	ToUserId    []string // 接收用户 ID，最多 100 人
	ObjectName  string   // 消息类型
	Content     string   // 消息内容（JSON 字符串）
	PushContent string   // 推送显示内容
	PushData    string   // 推送附加信息
}

// SendBroadcastReq 发送广播消息请求参数
type SendBroadcastReq struct {
	FromUserId string // 发送人用户 ID
	ObjectName string // 消息类型
	Content    string // 消息内容（JSON 字符串）
}

// RecallReq 撤回消息请求参数
type RecallReq struct {
	FromUserId       string // 发送人用户 ID
	TargetId         string // 目标 ID
	MessageUID       string // 消息唯一标识
	SentTime         int    // 消息发送时间
	ConversationType int    // 会话类型：1 单聊，3 群聊
}

// TemplateMsgContent 模板消息内容
type TemplateMsgContent struct {
	TargetId    string            `json:"targetId"`
	Data        map[string]string `json:"data"`
	PushContent string            `json:"pushContent"`
	PushData    string            `json:"pushData"`
}

// SendPrivateTemplateReq 发送单聊模板消息请求参数
type SendPrivateTemplateReq struct {
	FromUserId  string               `json:"fromUserId"`
	ObjectName  string               `json:"objectName"`
	Content     string               `json:"content"`
	ToUserId    []string             `json:"toUserId"`
	Values      []map[string]string  `json:"values"`
	PushContent []string             `json:"pushContent"`
	PushData    []string             `json:"pushData"`
}

// SendStatusMessageReq 发送群组状态消息请求参数
type SendStatusMessageReq struct {
	FromUserId string   // 发送人用户 ID
	ToGroupId  []string // 接收群组 ID
	ObjectName string   // 消息类型
	Content    string   // 消息内容（JSON 字符串）
}

// PrivateRecallReq 单聊消息撤回请求参数
type PrivateRecallReq struct {
	FromUserId string // 发送人用户 ID
	TargetId   string // 接收人用户 ID
	MessageUID string // 消息唯一标识
	SentTime   int    // 消息发送时间
}

// GroupRecallReq 群聊消息撤回请求参数
type GroupRecallReq struct {
	FromUserId string // 发送人用户 ID
	TargetId   string // 目标群组 ID
	MessageUID string // 消息唯一标识
	SentTime   int    // 消息发送时间
}

// ExpansionSetReq 消息扩展设置请求参数
type ExpansionSetReq struct {
	MsgUID           string // 消息唯一标识
	UserId           string // 操作用户 ID
	ConversationType string // 会话类型
	TargetId         string // 目标 ID
	ExtraKeyVal      string // 扩展内容（JSON 字符串）
	IsSyncSender     int    // 是否同步给发送者
}

// ExpansionSetResp 消息扩展设置响应
type ExpansionSetResp struct {
	types.BaseResp
}

// ExpansionQueryResp 消息扩展查询响应
type ExpansionQueryResp struct {
	types.BaseResp
	ExtraContent map[string]map[string]interface{} `json:"extraContent"`
}

// ExpansionItem 消息扩展项
type ExpansionItem struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	Timestamp int64  `json:"timestamp"`
}
