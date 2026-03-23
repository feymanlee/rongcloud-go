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
	AppSecret string // 应用密钥，用于验证签名

	// 自定义回调路径（可选，默认使用标准路径）
	MessageRoutePath     string // 消息路由回调路径，默认 "/message/sync"
	UserOnlineStatusPath string // 用户在线状态回调路径，默认 "/user/onlinestatus"
	AuditResultPath      string // 审核结果回调路径，默认 "/moderation/audit-result"
	ChatroomStatusPath   string // 聊天室状态回调路径，默认 "/chatroom/status"
	ChatroomKVPath       string // 聊天室 KV 回调路径，默认 "/chatroom/kv"
	UserDeactivationPath string // 用户注销/激活回调路径，默认 "/user/deactivation"

	// 回调处理器 - 可以通过 ResponseWriter 自定义响应
	// 返回 error 时，如果没有调用 WriteResponse，会返回 500 错误
	OnMessageRoute     func(w ResponseWriter, msg MessageRouteCallback) error
	OnUserOnlineStatus func(w ResponseWriter, statuses []UserOnlineStatusCallback) error
	OnAuditResult      func(w ResponseWriter, result AuditResultCallback) error
	OnChatroomStatus   func(w ResponseWriter, status ChatroomStatusCallback) error
	OnChatroomKV       func(w ResponseWriter, kv ChatroomKVCallback) error
	OnUserDeactivation func(w ResponseWriter, deactivation UserDeactivationCallback) error
}

// Handler 融云回调 HTTP 处理器
type Handler struct {
	config HandlerConfig
}

// NewHandler 创建回调处理器
func NewHandler(config HandlerConfig) *Handler {
	// 设置默认路径
	if config.MessageRoutePath == "" {
		config.MessageRoutePath = "/message/sync"
	}
	if config.UserOnlineStatusPath == "" {
		config.UserOnlineStatusPath = "/user/onlinestatus"
	}
	if config.AuditResultPath == "" {
		config.AuditResultPath = "/moderation/audit-result"
	}
	if config.ChatroomStatusPath == "" {
		config.ChatroomStatusPath = "/chatroom/status"
	}
	if config.ChatroomKVPath == "" {
		config.ChatroomKVPath = "/chatroom/kv"
	}
	if config.UserDeactivationPath == "" {
		config.UserDeactivationPath = "/user/deactivation"
	}
	return &Handler{config: config}
}

// ServeHTTP 实现 http.Handler 接口
// 根据请求路径分发到不同的回调处理器
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 对于消息路由回调，appKey 在请求体中，需要先解析表单
	if r.URL.Path == h.config.MessageRoutePath {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Invalid form data", http.StatusBadRequest)
			return
		}
	}

	// 验证签名
	if !VerifyRequest(r, h.config.AppSecret) {
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
