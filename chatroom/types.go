package chatroom

import "github.com/feymanlee/rongcloud-go/internal/types"

// ---------- Create ----------

// CreateResp 创建聊天室响应
type CreateResp struct {
	types.BaseResp
}

// ---------- Destroy ----------

// DestroyResp 销毁聊天室响应
type DestroyResp struct {
	types.BaseResp
}

// ---------- Query ----------

// ChatRoomInfo 聊天室信息
type ChatRoomInfo struct {
	ChatRoomID string `json:"chrmId"`
	Name       string `json:"name"`
	Time       string `json:"time"`
}

// QueryResp 查询聊天室响应
type QueryResp struct {
	types.BaseResp
	ChatRooms []ChatRoomInfo `json:"chatRooms"`
}

// ---------- QueryMembers ----------

// ChatRoomUser 聊天室用户信息
type ChatRoomUser struct {
	ID       string `json:"id"`
	UserID   string `json:"userId"`
	Time     string `json:"time"`
	IsInChrm int    `json:"isInChrm"`
}

// QueryMembersResp 查询聊天室成员响应
type QueryMembersResp struct {
	types.BaseResp
	Total int            `json:"total"`
	Users []ChatRoomUser `json:"users"`
}

// ---------- IsExist ----------

// IsExistResp 查询用户是否在聊天室响应
type IsExistResp struct {
	types.BaseResp
	Result []ChatRoomUser `json:"result"`
}

// ---------- BlockAdd ----------

// BlockAddResp 添加聊天室封禁用户响应
type BlockAddResp struct {
	types.BaseResp
}

// ---------- BlockRemove ----------

// BlockRemoveResp 移除聊天室封禁用户响应
type BlockRemoveResp struct {
	types.BaseResp
}

// ---------- BlockQuery ----------

// BlockQueryResp 查询聊天室封禁用户列表响应
type BlockQueryResp struct {
	types.BaseResp
	Users []ChatRoomUser `json:"users"`
}

// ---------- KeepaliveAdd ----------

// KeepaliveAddResp 添加保活聊天室响应
type KeepaliveAddResp struct {
	types.BaseResp
}

// ---------- KeepaliveRemove ----------

// KeepaliveRemoveResp 移除保活聊天室响应
type KeepaliveRemoveResp struct {
	types.BaseResp
}

// ---------- KeepaliveQuery ----------

// KeepaliveQueryResp 查询保活聊天室列表响应
type KeepaliveQueryResp struct {
	types.BaseResp
	ChatRoomIDs []string `json:"chatroomids"`
}
