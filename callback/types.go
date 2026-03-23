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

// CallbackParams 从请求中提取的回调参数
type CallbackParams struct {
	AppKey    string // 应用和环境的 App Key
	Nonce     string // 随机数，不超过 18 个字符
	Timestamp string // 时间戳（毫秒）
	Signature string // 数据签名
}
