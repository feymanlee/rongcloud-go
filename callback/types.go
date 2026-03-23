// Package callback 提供融云回调消息的结构体定义和处理功能
package callback

// MessageRouteCallback 消息路由回调（全量消息路由）
type MessageRouteCallback struct {
	FromUserId     string   `json:"fromUserId"`     // 发送用户 ID
	ToUserId       string   `json:"toUserId"`       // 目标 ID（targetId）
	ObjectName     string   `json:"objectName"`     // 消息类型，如 RC:TxtMsg、RC:ImgMsg
	Content        string   `json:"content"`        // 发送的消息内容
	ChannelType    string   `json:"channelType"`    // 会话类型：PERSON、GROUP、TEMPGROUP 等
	MsgTimestamp   string   `json:"msgTimestamp"`   // 服务端收到消息的时间（毫秒）
	MsgUID         string   `json:"msgUID"`         // 消息唯一标识
	OriginalMsgUID string   `json:"originalMsgUID"` // 原始消息 ID，仅超级群有效
	SensitiveType  int      `json:"sensitiveType"`  // 是否含敏感信息：0=不含，1=含屏蔽敏感词，2=含替换敏感词
	Source         string   `json:"source"`         // 消息发送源头：iOS、Android、HarmonyOS 等
	BusChannel     string   `json:"busChannel"`     // 超级群频道 ID，未使用时为空
	GroupUserIds   []string `json:"groupUserIds"`   // 群定向消息的接收成员用户 ID 数组
	AIGenerated    bool     `json:"aiGenerated"`    // 是否为 AI 生成消息
}

// UserOnlineStatusCallback 用户在线状态回调
type UserOnlineStatusCallback struct {
	UserID    string `json:"userid"`    // 用户 ID
	Status    string `json:"status"`    // 状态变更事件：0=上线，1=离线，2=登出
	OS        string `json:"os"`        // 操作系统：iOS、Android、HarmonyOS、Websocket、PC、MiniProgram
	Time      int64  `json:"time"`      // 发生时间（毫秒）
	ClientIP  string `json:"clientIp"`  // 用户当前 IP 地址及端口
	SessionID string `json:"sessionId"` // 连接唯一 ID，用于区分多端在线
}

// AuditResultCallback 审核结果回调
type AuditResultCallback struct {
	Result          int    `json:"result"`          // 审核是否通过：10000=通过，10001=不通过
	Content         string `json:"content"`         // 审核的消息内容（JSON 字符串）
	MsgUID          string `json:"msgUID"`          // 消息 ID
	ServiceProvider string `json:"serviceProvider"` // 审核渠道商标识
	ResultDetail    string `json:"resultDetail"`    // 审核结果详情（JSON 字符串）
}

// AuditContent 审核内容解析后的结构
type AuditContent struct {
	AppKey           string                 `json:"appKey"`           // App Key
	FromUserId       string                 `json:"fromUserId"`       // 发送者 ID
	TargetId         string                 `json:"targetId"`         // 接收者 ID
	ToUserIds        string                 `json:"toUserIds"`        // 群成员 ID 列表（英文逗号分隔）
	ConversationType string                 `json:"conversationType"` // 会话类型：PERSON/GROUP/TEMPGROUP/ULTRAGROUP
	ObjectName       string                 `json:"objectName"`       // 消息类型
	Message          string                 `json:"message"`          // 消息体内容
	ExtraContent     map[string]interface{} `json:"extraContent"`     // 消息扩展内容
	ClientOs         string                 `json:"clientOs"`         // 客户端类型
	MessageTime      int64                  `json:"messageTime"`      // 服务端收到消息的时间（毫秒）
	MessageId        string                 `json:"messageId"`        // 消息 ID
}

// ChatroomStatusCallback 聊天室状态回调
type ChatroomStatusCallback struct {
	ChatroomId string `json:"chatroomId"` // 聊天室 ID
	Status     int    `json:"status"`     // 状态：0=创建，1=销毁
	Time       int64  `json:"time"`       // 发生时间（毫秒）
}

// ChatroomKVCallback 聊天室 KV 属性回调
type ChatroomKVCallback struct {
	ChatroomId string `json:"chatroomId"` // 聊天室 ID
	UserId     string `json:"userId"`     // 操作用户 ID
	Key        string `json:"key"`        // 属性名
	Value      string `json:"value"`      // 属性值
	Time       int64  `json:"time"`       // 发生时间（毫秒）
}

// UserDeactivationCallback 用户注销/激活回调
type UserDeactivationCallback struct {
	UserId string `json:"userId"` // 用户 ID
	Type   int    `json:"type"`   // 类型：0=注销，1=激活
	Time   int64  `json:"time"`   // 发生时间（毫秒）
}

// MessageOperationCallback 消息操作状态同步回调（消息撤回/删除）
type MessageOperationCallback struct {
	EventType        int                           `json:"eventType"`        // 事件类型：1=消息撤回，2=消息删除
	FromUserId       string                        `json:"fromUserId"`       // 操作人用户 ID
	OptTime          int64                         `json:"optTime"`          // 操作时间戳（毫秒）
	Source           string                        `json:"source"`           // 操作来源：Android、iOS、WebSocket、Server 等
	ConversationInfo MessageOperationConversation  `json:"conversationInfo"` // 会话信息
	OriginalMsgInfo  MessageOperationOriginalMsg   `json:"originalMsgInfo"`  // 原始消息信息
	RecallMsgInfo    *MessageOperationRecallInfo   `json:"recallMsgInfo,omitempty"`   // 撤回消息特有信息（eventType=1时）
	DeleteMsgInfo    *MessageOperationDeleteInfo   `json:"deleteMsgInfo,omitempty"`   // 删除消息特有信息（eventType=2时）
}

// MessageOperationConversation 消息操作回调中的会话信息
type MessageOperationConversation struct {
	ConversationType int    `json:"conversationType"` // 会话类型
	TargetId         string `json:"targetId"`         // 会话 ID
	BusChannel       string `json:"busChannel"`       // 频道 ID（可选）
}

// MessageOperationOriginalMsg 消息操作回调中的原始消息信息
type MessageOperationOriginalMsg struct {
	MessageId    string `json:"messageId"`    // 原始消息 ID
	MessageTime  int64  `json:"messageTime"`  // 原始消息时间
}

// MessageOperationRecallInfo 消息撤回特有信息
type MessageOperationRecallInfo struct {
	IsDelete int    `json:"isDelete"` // 接收方是否本地删除：0=否，1=是
	IsAdmin  int    `json:"isAdmin"`  // 是否管理员操作：0=否，1=是
	Extra    string `json:"extra"`    // 扩展信息
}

// MessageOperationDeleteInfo 消息删除特有信息
type MessageOperationDeleteInfo struct {
	DeleteType   int   `json:"deleteType"`   // 删除类型：1=指定消息删除，2=按时间戳删除
	MsgTimestamp int64 `json:"msgTimestamp"` // 时间戳（deleteType=2时有效）
}

// MessageCallbackService 消息回调服务（自定义条件消息抄送）
// 注意：此回调使用 application/x-www-form-urlencoded 格式，且 appKey 在请求体中
type MessageCallbackService struct {
	AppKey         string `json:"appKey"`         // 应用 App Key
	FromUserId     string `json:"fromUserId"`     // 发送用户 ID
	TargetId       string `json:"targetId"`       // 目标会话 ID
	ToUserIds      string `json:"toUserIds"`      // 群成员 ID 列表（逗号分隔）
	MsgType        string `json:"msgType"`        // 消息类型标识
	Content        string `json:"content"`        // JSON 结构的消息内容
	PushContent    string `json:"pushContent"`    // 推送通知栏显示内容
	DisablePush    bool   `json:"disablePush"`    // 是否为静默消息
	PushExt        string `json:"pushExt"`        // 推送通知配置
	Expansion      bool   `json:"expansion"`      // 是否为可扩展消息
	ExtraContent   string `json:"extraContent"`   // 消息扩展内容（JSON 字符串）
	ChannelType    string `json:"channelType"`    // 会话类型：PERSON/PERSONS/GROUP/TEMPGROUP/ULTRAGROUP
	MsgTimeStamp   string `json:"msgTimeStamp"`   // 服务器时间戳（毫秒）
	MessageId      string `json:"messageId"`      // 消息唯一标识
	OriginalMsgUID string `json:"originalMsgUID"` // 原始消息 ID（超级群有效）
	OS             string `json:"os"`             // 消息来源：iOS/Android/HarmonyOS/Websocket/MiniProgram/PC/Server
	BusChannel     string `json:"busChannel"`     // 超级群频道 ID
	ClientIp       string `json:"clientIp"`       // 用户 IP 地址及端口
	AiGenerated    bool   `json:"aiGenerated"`    // 是否为 AI 生成消息
}

// BotMessageCallback 机器人消息回调
type BotMessageCallback struct {
	Type      string                 `json:"type"`      // 回调事件类型
	Timestamp int64                  `json:"timestamp"` // 触发时间（Unix 毫秒）
	Bot       BotInfo                `json:"bot"`       // 机器人信息
	Data      map[string]interface{} `json:"data"`      // 事件特定数据
}

// BotInfo 机器人信息
type BotInfo struct {
	UserId      string                 `json:"userId"`      // 机器人用户 ID
	Name        string                 `json:"name"`        // 机器人名称
	ProfileUrl  string                 `json:"profileUrl"`  // 机器人头像 URL
	Type        string                 `json:"type"`        // 机器人类型
	Metadata    map[string]interface{} `json:"metadata"`    // 机器人元数据
}

// 机器人消息事件类型常量
const (
	// 消息事件
	BotEventMessagePrivate              = "message:private"               // 私聊消息
	BotEventMessageGroupMentioned       = "message:group_mentioned"       // 群组@消息
	BotEventMessagePrivateRecall        = "message:private_recall"        // 私聊消息撤回
	BotEventMessageGroupMentionedRecall = "message:group_mentioned_recall" // 群组@消息撤回
	BotEventMessagePrivateRead          = "message:private_read"          // 私聊已读回执
	BotEventMessageGroupRead            = "message:group_read"            // 群组已读回执
	// 群组事件
	BotEventGroupBotJoin   = "group:bot_join"   // 机器人被加入群组
	BotEventGroupBotLeft   = "group:bot_left"   // 机器人被移出群组
	BotEventGroupDismiss   = "group:dismiss"    // 群组解散
	BotEventGroupUserJoin  = "group:user_join"  // 其他用户加入群组
	BotEventGroupUserLeft  = "group:user_left"  // 其他用户离开群组
)

// CallbackParams 从请求中提取的回调参数
type CallbackParams struct {
	AppKey    string // 应用和环境的 App Key
	Nonce     string // 随机数，不超过 18 个字符
	Timestamp string // 时间戳（毫秒）
	Signature string // 数据签名
}
