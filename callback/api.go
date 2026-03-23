package callback

import "github.com/feymanlee/rongcloud-go/internal/core"

const (
	// DefaultMessageRoutePath 消息路由回调默认路径
	DefaultMessageRoutePath = "/message/sync"
	// DefaultUserOnlineStatusPath 用户在线状态回调默认路径
	DefaultUserOnlineStatusPath = "/user/onlinestatus"
	// DefaultAuditResultPath 审核结果回调默认路径
	DefaultAuditResultPath = "/moderation/audit-result"
	// DefaultChatroomStatusPath 聊天室状态回调默认路径
	DefaultChatroomStatusPath = "/chatroom/status"
	// DefaultChatroomKVPath 聊天室 KV 回调默认路径
	DefaultChatroomKVPath = "/chatroom/kv"
	// DefaultUserDeactivationPath 用户注销/激活回调默认路径
	DefaultUserDeactivationPath = "/user/deactivation"
)

// API 回调处理接口
type API interface {
	// HandlerConfig 获取 Handler 配置，用于设置回调处理器
	HandlerConfig() HandlerConfig
	// SetHandlerConfig 设置 Handler 配置
	SetHandlerConfig(config HandlerConfig)
	// Handler 获取 HTTP Handler 实例
	// 注意：需要先在 HandlerConfig 中设置回调处理器
	Handler() *Handler
}

type api struct {
	client      core.Client
	appSecret   string
	handlerConfig HandlerConfig
}

// NewAPI 创建回调 API 实例
func NewAPI(client core.Client, appSecret string) API {
	return &api{
		client:    client,
		appSecret: appSecret,
	}
}

func (a *api) HandlerConfig() HandlerConfig {
	return a.handlerConfig
}

func (a *api) SetHandlerConfig(config HandlerConfig) {
	a.handlerConfig = config
}

func (a *api) Handler() *Handler {
	return NewHandler(a.appSecret, a.handlerConfig)
}
