package callback

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// ResponseWriter 回调响应写入接口
type ResponseWriter interface {
	// WriteResponse 写入响应，code 为 HTTP 状态码，body 为响应内容
	// 如果未调用此方法，默认返回 200 OK
	WriteResponse(code int, body string)
	// Header 获取 HTTP Header 以便设置自定义头部
	Header() http.Header
}

// responseWriter 内部实现
type responseWriter struct {
	http.ResponseWriter
	written bool
	code    int
	body    string
}

func (w *responseWriter) WriteResponse(code int, body string) {
	if w.written {
		return
	}
	w.written = true
	w.code = code
	w.body = body
	w.WriteHeader(code)
	w.Write([]byte(body))
}

func (w *responseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

// HandlerConfig 回调处理器配置
type HandlerConfig struct {
	// 自定义回调路径（可选，默认使用标准路径）
	MessageRoutePath     string // 消息路由回调路径，默认 DefaultMessageRoutePath
	UserOnlineStatusPath string // 用户在线状态回调路径，默认 DefaultUserOnlineStatusPath
	AuditResultPath      string // 审核结果回调路径，默认 DefaultAuditResultPath
	ChatroomStatusPath   string // 聊天室状态回调路径，默认 DefaultChatroomStatusPath
	ChatroomKVPath       string // 聊天室 KV 回调路径，默认 DefaultChatroomKVPath
	UserDeactivationPath string // 用户注销/激活回调路径，默认 DefaultUserDeactivationPath
	MessageOperationPath string // 消息操作状态同步回调路径，默认 DefaultMessageOperationPath
	MessageCallbackPath  string // 消息回调服务路径，默认 DefaultMessageCallbackPath
	BotMessagePath       string // 机器人消息回调路径，默认 DefaultBotMessagePath

	// 回调处理器 - 可以通过 ResponseWriter 自定义响应
	// 返回 error 时，如果没有调用 WriteResponse，会返回 500 错误
	OnMessageRoute     func(w ResponseWriter, msg MessageRouteCallback) error
	OnUserOnlineStatus func(w ResponseWriter, statuses []UserOnlineStatusCallback) error
	OnAuditResult      func(w ResponseWriter, result AuditResultCallback) error
	OnChatroomStatus   func(w ResponseWriter, status ChatroomStatusCallback) error
	OnChatroomKV       func(w ResponseWriter, kv ChatroomKVCallback) error
	OnUserDeactivation func(w ResponseWriter, deactivation UserDeactivationCallback) error
	OnMessageOperation func(w ResponseWriter, operation MessageOperationCallback) error
	OnMessageCallback  func(w ResponseWriter, msg MessageCallback) error
	OnBotMessage       func(w ResponseWriter, msg BotMessageCallback) error
}

// Handler 融云回调 HTTP 处理器
type Handler struct {
	appSecret string
	config    HandlerConfig
}

// NewHandler 创建回调处理器
func NewHandler(appSecret string, config HandlerConfig) *Handler {
	// 设置默认路径
	if config.MessageRoutePath == "" {
		config.MessageRoutePath = DefaultMessageRoutePath
	}
	if config.UserOnlineStatusPath == "" {
		config.UserOnlineStatusPath = DefaultUserOnlineStatusPath
	}
	if config.AuditResultPath == "" {
		config.AuditResultPath = DefaultAuditResultPath
	}
	if config.ChatroomStatusPath == "" {
		config.ChatroomStatusPath = DefaultChatroomStatusPath
	}
	if config.ChatroomKVPath == "" {
		config.ChatroomKVPath = DefaultChatroomKVPath
	}
	if config.UserDeactivationPath == "" {
		config.UserDeactivationPath = DefaultUserDeactivationPath
	}
	if config.MessageOperationPath == "" {
		config.MessageOperationPath = DefaultMessageOperationPath
	}
	if config.MessageCallbackPath == "" {
		config.MessageCallbackPath = DefaultMessageCallbackPath
	}
	if config.BotMessagePath == "" {
		config.BotMessagePath = DefaultBotMessagePath
	}
	return &Handler{appSecret: appSecret, config: config}
}

// ServeHTTP 实现 http.Handler 接口
// 根据请求路径分发到不同的回调处理器
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 对于消息路由回调和消息回调服务，appKey 在请求体中，需要先解析表单
	if r.URL.Path == h.config.MessageRoutePath || r.URL.Path == h.config.MessageCallbackPath {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Invalid form data", http.StatusBadRequest)
			return
		}
	}

	// 验证签名
	if !VerifyRequest(r, h.appSecret) {
		http.Error(w, "Invalid signature", http.StatusUnauthorized)
		return
	}

	rw := &responseWriter{ResponseWriter: w}

	// 根据路径分发
	switch r.URL.Path {
	case h.config.MessageRoutePath:
		h.handleMessageRoute(rw, r)
	case h.config.UserOnlineStatusPath:
		h.handleUserOnlineStatus(rw, r)
	case h.config.AuditResultPath:
		h.handleAuditResult(rw, r)
	case h.config.ChatroomStatusPath:
		h.handleChatroomStatus(rw, r)
	case h.config.ChatroomKVPath:
		h.handleChatroomKV(rw, r)
	case h.config.UserDeactivationPath:
		h.handleUserDeactivation(rw, r)
	case h.config.MessageOperationPath:
		h.handleMessageOperation(rw, r)
	case h.config.MessageCallbackPath:
		h.handleMessageCallback(rw, r)
	case h.config.BotMessagePath:
		h.handleBotMessage(rw, r)
	default:
		http.Error(w, "Not found", http.StatusNotFound)
	}
}

// writeDefaultResponse 写入默认响应
func writeDefaultResponse(rw *responseWriter, err error) {
	if rw.written {
		return
	}
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteResponse(http.StatusOK, "OK")
}

// handleMessageRoute 处理消息路由回调
// 注意：消息路由回调使用 application/x-www-form-urlencoded 格式
func (h *Handler) handleMessageRoute(rw *responseWriter, r *http.Request) {
	if h.config.OnMessageRoute == nil {
		http.Error(rw, "Handler not configured", http.StatusInternalServerError)
		return
	}

	// 注意：ServeHTTP 中已经调用过 ParseForm 来提取 appKey 验证签名
	// 这里再次调用是安全的（ParseForm 会检查是否已经解析过）
	if err := r.ParseForm(); err != nil {
		http.Error(rw, "Invalid form data", http.StatusBadRequest)
		return
	}

	// 单个消息路由回调（form-urlencoded 格式）
	callback := MessageRouteCallback{
		FromUserId:     r.FormValue("fromUserId"),
		ToUserId:       r.FormValue("targetId"),
		ObjectName:     r.FormValue("objectName"),
		Content:        r.FormValue("content"),
		ChannelType:    r.FormValue("channelType"),
		MsgTimestamp:   r.FormValue("msgTimestamp"),
		MsgUID:         r.FormValue("msgUID"),
		OriginalMsgUID: r.FormValue("originalMsgUID"),
		Source:         r.FormValue("source"),
		BusChannel:     r.FormValue("busChannel"),
	}

	// 解析可选字段
	if v := r.FormValue("sensitiveType"); v != "" {
		callback.SensitiveType, _ = strconv.Atoi(v)
	}
	if v := r.FormValue("aiGenerated"); v == "true" {
		callback.AIGenerated = true
	}
	// groupUserIds 可能是逗号分隔的字符串
	if v := r.FormValue("groupUserIds"); v != "" {
		callback.GroupUserIds = []string{v}
	}

	err := h.config.OnMessageRoute(rw, callback)
	writeDefaultResponse(rw, err)
}

// handleUserOnlineStatus 处理用户在线状态回调
// 注意：用户在线状态回调使用 application/json 格式（数组）
func (h *Handler) handleUserOnlineStatus(rw *responseWriter, r *http.Request) {
	if h.config.OnUserOnlineStatus == nil {
		http.Error(rw, "Handler not configured", http.StatusInternalServerError)
		return
	}

	var callbacks []UserOnlineStatusCallback
	if err := json.NewDecoder(r.Body).Decode(&callbacks); err != nil {
		http.Error(rw, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err := h.config.OnUserOnlineStatus(rw, callbacks)
	writeDefaultResponse(rw, err)
}

// handleAuditResult 处理审核结果回调
// 注意：审核结果回调使用 application/json 格式
func (h *Handler) handleAuditResult(rw *responseWriter, r *http.Request) {
	if h.config.OnAuditResult == nil {
		http.Error(rw, "Handler not configured", http.StatusInternalServerError)
		return
	}

	var callback AuditResultCallback
	if err := json.NewDecoder(r.Body).Decode(&callback); err != nil {
		http.Error(rw, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err := h.config.OnAuditResult(rw, callback)
	writeDefaultResponse(rw, err)
}

// handleChatroomStatus 处理聊天室状态回调
func (h *Handler) handleChatroomStatus(rw *responseWriter, r *http.Request) {
	if h.config.OnChatroomStatus == nil {
		http.Error(rw, "Handler not configured", http.StatusInternalServerError)
		return
	}

	var callback ChatroomStatusCallback
	if err := json.NewDecoder(r.Body).Decode(&callback); err != nil {
		http.Error(rw, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err := h.config.OnChatroomStatus(rw, callback)
	writeDefaultResponse(rw, err)
}

// handleChatroomKV 处理聊天室 KV 属性回调
func (h *Handler) handleChatroomKV(rw *responseWriter, r *http.Request) {
	if h.config.OnChatroomKV == nil {
		http.Error(rw, "Handler not configured", http.StatusInternalServerError)
		return
	}

	var callback ChatroomKVCallback
	if err := json.NewDecoder(r.Body).Decode(&callback); err != nil {
		http.Error(rw, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err := h.config.OnChatroomKV(rw, callback)
	writeDefaultResponse(rw, err)
}

// handleUserDeactivation 处理用户注销/激活回调
func (h *Handler) handleUserDeactivation(rw *responseWriter, r *http.Request) {
	if h.config.OnUserDeactivation == nil {
		http.Error(rw, "Handler not configured", http.StatusInternalServerError)
		return
	}

	var callback UserDeactivationCallback
	if err := json.NewDecoder(r.Body).Decode(&callback); err != nil {
		http.Error(rw, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err := h.config.OnUserDeactivation(rw, callback)
	writeDefaultResponse(rw, err)
}

// handleMessageOperation 处理消息操作状态同步回调（消息撤回/删除）
func (h *Handler) handleMessageOperation(rw *responseWriter, r *http.Request) {
	if h.config.OnMessageOperation == nil {
		http.Error(rw, "Handler not configured", http.StatusInternalServerError)
		return
	}

	var callback MessageOperationCallback
	if err := json.NewDecoder(r.Body).Decode(&callback); err != nil {
		http.Error(rw, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err := h.config.OnMessageOperation(rw, callback)
	writeDefaultResponse(rw, err)
}

// handleMessageCallback 处理消息回调服务
// 注意：此回调使用 application/x-www-form-urlencoded 格式，且 appKey 在请求体中
func (h *Handler) handleMessageCallback(rw *responseWriter, r *http.Request) {
	if h.config.OnMessageCallback == nil {
		http.Error(rw, "Handler not configured", http.StatusInternalServerError)
		return
	}

	// ServeHTTP 中已经调用过 ParseForm 来提取 appKey 验证签名
	if err := r.ParseForm(); err != nil {
		http.Error(rw, "Invalid form data", http.StatusBadRequest)
		return
	}

	callback := MessageCallback{
		AppKey:         r.FormValue("appKey"),
		FromUserId:     r.FormValue("fromUserId"),
		TargetId:       r.FormValue("targetId"),
		ToUserIds:      r.FormValue("toUserIds"),
		MsgType:        r.FormValue("msgType"),
		Content:        r.FormValue("content"),
		PushContent:    r.FormValue("pushContent"),
		PushExt:        r.FormValue("pushExt"),
		ExtraContent:   r.FormValue("extraContent"),
		ChannelType:    r.FormValue("channelType"),
		MsgTimeStamp:   r.FormValue("msgTimeStamp"),
		MessageId:      r.FormValue("messageId"),
		OriginalMsgUID: r.FormValue("originalMsgUID"),
		OS:             r.FormValue("os"),
		BusChannel:     r.FormValue("busChannel"),
		ClientIp:       r.FormValue("clientIp"),
	}

	// 解析可选字段
	if v := r.FormValue("disablePush"); v == "true" {
		callback.DisablePush = true
	}
	if v := r.FormValue("expansion"); v == "true" {
		callback.Expansion = true
	}
	if v := r.FormValue("aiGenerated"); v == "true" {
		callback.AiGenerated = true
	}

	err := h.config.OnMessageCallback(rw, callback)
	writeDefaultResponse(rw, err)
}

// handleBotMessage 处理机器人消息回调
func (h *Handler) handleBotMessage(rw *responseWriter, r *http.Request) {
	if h.config.OnBotMessage == nil {
		http.Error(rw, "Handler not configured", http.StatusInternalServerError)
		return
	}

	var callback BotMessageCallback
	if err := json.NewDecoder(r.Body).Decode(&callback); err != nil {
		http.Error(rw, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err := h.config.OnBotMessage(rw, callback)
	writeDefaultResponse(rw, err)
}
