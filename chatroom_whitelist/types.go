package chatroomwhitelist

import "github.com/feymanlee/rongcloud-go/internal/types"

// ---------- MsgAdd ----------

// MsgAddResp 添加聊天室消息白名单响应
type MsgAddResp struct {
	types.BaseResp
}

// ---------- MsgRemove ----------

// MsgRemoveResp 移除聊天室消息白名单响应
type MsgRemoveResp struct {
	types.BaseResp
}

// ---------- MsgQuery ----------

// MsgQueryResp 查询聊天室消息白名单响应
type MsgQueryResp struct {
	types.BaseResp
	WhitelistMsgType []string `json:"whitlistMsgType"`
}

// ---------- UserAdd ----------

// UserAddResp 添加聊天室用户白名单响应
type UserAddResp struct {
	types.BaseResp
}

// ---------- UserRemove ----------

// UserRemoveResp 移除聊天室用户白名单响应
type UserRemoveResp struct {
	types.BaseResp
}

// ---------- UserQuery ----------

// UserQueryResp 查询聊天室用户白名单响应
type UserQueryResp struct {
	types.BaseResp
	Users []string `json:"users"`
}
