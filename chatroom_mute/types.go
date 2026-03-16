package chatroommute

import "github.com/feymanlee/rongcloud-go/internal/types"

// ---------- MuteAdd ----------

// MuteAddResp 添加聊天室禁言响应
type MuteAddResp struct {
	types.BaseResp
}

// ---------- MuteRemove ----------

// MuteRemoveResp 移除聊天室禁言响应
type MuteRemoveResp struct {
	types.BaseResp
}

// ---------- MuteQuery ----------

// MutedUser 被禁言用户信息
type MutedUser struct {
	UserID string `json:"userId"`
	Time   string `json:"time"`
}

// MuteQueryResp 查询聊天室禁言用户列表响应
type MuteQueryResp struct {
	types.BaseResp
	Users []MutedUser `json:"users"`
}

// ---------- MuteAllAdd ----------

// MuteAllAddResp 添加聊天室全体禁言响应
type MuteAllAddResp struct {
	types.BaseResp
}

// ---------- MuteAllRemove ----------

// MuteAllRemoveResp 移除聊天室全体禁言响应
type MuteAllRemoveResp struct {
	types.BaseResp
}

// ---------- MuteAllQuery ----------

// MuteAllQueryResp 查询聊天室全体禁言列表响应
type MuteAllQueryResp struct {
	types.BaseResp
	ChatRoomIDs []string `json:"chatroomids"`
}

// ---------- GlobalMuteAdd ----------

// GlobalMuteAddResp 添加全局聊天室禁言响应
type GlobalMuteAddResp struct {
	types.BaseResp
}

// ---------- GlobalMuteRemove ----------

// GlobalMuteRemoveResp 移除全局聊天室禁言响应
type GlobalMuteRemoveResp struct {
	types.BaseResp
}

// ---------- GlobalMuteQuery ----------

// GlobalMutedUser 全局禁言用户信息
type GlobalMutedUser struct {
	UserID string `json:"userId"`
	Time   string `json:"time"`
}

// GlobalMuteQueryResp 查询全局聊天室禁言用户列表响应
type GlobalMuteQueryResp struct {
	types.BaseResp
	Users []GlobalMutedUser `json:"users"`
}

// ---------- WhitelistAdd ----------

// WhitelistAddResp 添加聊天室禁言白名单响应
type WhitelistAddResp struct {
	types.BaseResp
}

// ---------- WhitelistRemove ----------

// WhitelistRemoveResp 移除聊天室禁言白名单响应
type WhitelistRemoveResp struct {
	types.BaseResp
}

// ---------- WhitelistQuery ----------

// WhitelistQueryResp 查询聊天室禁言白名单响应
type WhitelistQueryResp struct {
	types.BaseResp
	Users []string `json:"users"`
}

// ---------- UserWhitelistAdd ----------

// UserWhitelistAddResp 添加聊天室用户禁言白名单响应
type UserWhitelistAddResp struct {
	types.BaseResp
}
