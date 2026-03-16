package chatroomblock

import "github.com/feymanlee/rongcloud-go/internal/types"

// ---------- Add ----------

// AddResp 添加聊天室封禁响应
type AddResp struct {
	types.BaseResp
}

// ---------- Remove ----------

// RemoveResp 移除聊天室封禁响应
type RemoveResp struct {
	types.BaseResp
}

// ---------- Query ----------

// BlockedUser 被封禁用户信息
type BlockedUser struct {
	UserID string `json:"userId"`
	Time   string `json:"time"`
}

// QueryResp 查询聊天室封禁用户列表响应
type QueryResp struct {
	types.BaseResp
	Users []BlockedUser `json:"users"`
}
