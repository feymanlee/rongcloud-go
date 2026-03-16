package push

import "github.com/feymanlee/rongcloud-go/internal/types"

// PlatForm 推送平台类型
type PlatForm string

const (
	// IOSPlatForm iOS 平台
	IOSPlatForm PlatForm = "ios"
	// AndroidPlatForm Android 平台
	AndroidPlatForm PlatForm = "android"
)

// Audience 推送条件
type Audience struct {
	Tag         []string `json:"tag,omitempty"`         // 用户标签，最多 20 个，标签之间为 AND 关系
	TagOr       []string `json:"tag_or,omitempty"`      // 用户标签，最多 20 个，标签之间为 OR 关系
	UserID      []string `json:"userid,omitempty"`      // 用户 ID，最多 1000 个
	IsToAll     bool     `json:"is_to_all"`             // 是否推送给所有用户
	PackageName string   `json:"packageName,omitempty"` // 应用包名
}

// Message 广播消息内容
type Message struct {
	Content              string `json:"content"`                        // 消息内容
	ObjectName           string `json:"objectName"`                     // 消息类型
	DisableUpdateLastMsg bool   `json:"disableUpdateLastMsg,omitempty"` // 是否禁止更新最后一条消息
}

// IOSPush iOS 平台推送设置
type IOSPush struct {
	Title            string      `json:"title,omitempty"`            // 推送标题，iOS 8.2 及以上支持
	ContentAvailable int         `json:"contentAvailable,omitempty"` // 静默推送，1 开启，0 关闭
	Alert            string      `json:"alert,omitempty"`            // 推送消息内容
	Extras           interface{} `json:"extras,omitempty"`           // 附加信息
	Badge            int         `json:"badge,omitempty"`            // 应用角标数
	Category         string      `json:"category,omitempty"`         // iOS 富推送类型
	RichMediaURI     string      `json:"richMediaUri,omitempty"`     // iOS 富推送内容 URL
	ThreadID         string      `json:"threadId,omitempty"`         // iOS 通知分组 ID
	ApnsCollapseID   string      `json:"apns-collapse-id,omitempty"` // iOS 消息合并 ID
}

// AndroidPush Android 平台推送设置
type AndroidPush struct {
	Alert          string      `json:"alert,omitempty"`          // 推送消息内容
	Extras         interface{} `json:"extras,omitempty"`         // 附加信息
	ChannelID      string      `json:"channelId,omitempty"`      // 厂商推送通道 ID
	Importance     string      `json:"importance,omitempty"`     // 华为通知优先级
	Image          string      `json:"image,omitempty"`          // 华为推送自定义图标 URL
	LargeIconURI   string      `json:"large_icon_uri,omitempty"` // 小米推送自定义图标 URL
	Classification string      `json:"classification,omitempty"` // vivo 推送通道类型
}

// Notification 按操作系统类型推送消息内容
type Notification struct {
	Alert   string      `json:"alert"`             // 默认推送消息内容
	IOS     IOSPush     `json:"ios,omitempty"`     // iOS 平台推送设置
	Android AndroidPush `json:"android,omitempty"` // Android 平台推送设置
}

// PushUserResp 推送用户响应
type PushUserResp struct {
	types.BaseResp
	ID string `json:"id,omitempty"` // 推送唯一标识
}

// PushCustomResp 自定义推送响应
type PushCustomResp struct {
	types.BaseResp
	ID string `json:"id,omitempty"` // 推送唯一标识
}

// BroadcastSendResp 广播消息发送响应
type BroadcastSendResp struct {
	types.BaseResp
}

// BroadcastRecallResp 广播消息撤回响应
type BroadcastRecallResp struct {
	types.BaseResp
}

// BroadcastPushResp 广播推送响应
type BroadcastPushResp struct {
	types.BaseResp
	ID string `json:"id,omitempty"` // 推送唯一标识
}

// SystemSendResp 系统消息发送响应
type SystemSendResp struct {
	types.BaseResp
}

// SystemSendTemplateReq 系统消息模板发送请求
type SystemSendTemplateReq struct {
	FromUserID      string            `json:"fromUserId"`               // 发送人用户 ID
	ToUserIDs       []string          `json:"toUserId"`                 // 接收人用户 ID 列表
	ObjectName      string            `json:"objectName"`               // 消息类型
	Values          []map[string]string `json:"values"`                 // 模板变量值
	Content         string            `json:"content"`                  // 消息内容模板
	PushContent     []string          `json:"pushContent,omitempty"`    // 推送内容
	PushData        []string          `json:"pushData,omitempty"`       // 推送附加数据
	VerifyBlacklist int               `json:"verifyBlacklist,omitempty"` // 是否过滤黑名单
}

// SystemSendTemplateResp 系统消息模板发送响应
type SystemSendTemplateResp struct {
	types.BaseResp
}

// TagPushReq 标签推送请求
type TagPushReq struct {
	PlatForm     []PlatForm   `json:"platform"`               // 目标平台
	FromUserID   string       `json:"fromuserid,omitempty"`   // 发送人用户 ID
	Audience     Audience     `json:"audience"`               // 推送条件
	Message      *Message     `json:"message,omitempty"`      // 消息内容
	Notification Notification `json:"notification,omitempty"` // 推送通知内容
}

// TagPushResp 标签推送响应
type TagPushResp struct {
	types.BaseResp
	ID string `json:"id,omitempty"` // 推送唯一标识
}

// OnlineBroadcastResp 在线广播响应
type OnlineBroadcastResp struct {
	types.BaseResp
}

// PushRecallResp 推送撤回响应
type PushRecallResp struct {
	types.BaseResp
}

// PushQueryTaskResp 查询推送任务响应
type PushQueryTaskResp struct {
	types.BaseResp
	ID           string `json:"id,omitempty"`           // 任务 ID
	PushStatus   int    `json:"pushStatus,omitempty"`   // 推送状态
	InvalidCount int    `json:"invalidCount,omitempty"` // 无效数量
	TotalCount   int    `json:"totalCount,omitempty"`   // 总数
}

// PushQueryStatusResp 查询推送状态响应
type PushQueryStatusResp struct {
	types.BaseResp
	ID         string `json:"id,omitempty"`         // 任务 ID
	PushStatus int    `json:"pushStatus,omitempty"` // 推送状态
}

// PushDeleteResp 删除推送任务响应
type PushDeleteResp struct {
	types.BaseResp
}
