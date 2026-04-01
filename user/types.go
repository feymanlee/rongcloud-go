package user

import "github.com/feymanlee/rongcloud-go/internal/types"

// ---------- GetToken ----------

// GetTokenResp 获取 Token 响应
type GetTokenResp struct {
	types.BaseResp
	Token  string `json:"token"`
	UserID string `json:"userId"`
}

// ---------- Update (refresh) ----------

// UpdateResp 刷新用户信息响应
type UpdateResp struct {
	types.BaseResp
}

// ---------- UserInfoGet ----------

// UserInfoGetResp 获取用户信息响应
type UserInfoGetResp struct {
	types.BaseResp
	UserName     string `json:"userName"`
	UserPortrait string `json:"userPortrait"`
	CreateTime   string `json:"createTime"`
}

// ---------- TokenExpire ----------

// TokenExpireResp Token 过期响应
type TokenExpireResp struct {
	types.BaseResp
}

// ---------- Block ----------

// BlockAddResp 封禁用户响应
type BlockAddResp struct {
	types.BaseResp
}

// BlockRemoveResp 解除封禁用户响应
type BlockRemoveResp struct {
	types.BaseResp
}

// BlockUser 封禁用户信息
type BlockUser struct {
	UserID       string `json:"userId"`
	BlockEndTime string `json:"blockEndTime,omitempty"`
}

// BlockQueryResp 查询封禁用户列表响应
type BlockQueryResp struct {
	types.BaseResp
	Users []BlockUser `json:"users"`
}

// ---------- OnlineStatusCheck ----------

// OnlineStatusCheckResp 查询在线状态响应
type OnlineStatusCheckResp struct {
	types.BaseResp
	Status string `json:"status"`
}

// ---------- Ban (全局封禁) ----------

// BanResp 全局封禁响应
type BanResp struct {
	types.BaseResp
}

// BanUser 全局封禁用户信息
type BanUser struct {
	UserID  string `json:"userId"`
	Status  string `json:"status,omitempty"`
}

// BanQueryResp 查询全局封禁用户列表响应
type BanQueryResp struct {
	types.BaseResp
	Users []BanUser `json:"users"`
}

// UnBanResp 解除全局封禁响应
type UnBanResp struct {
	types.BaseResp
}

// ---------- Deactivate (注销) ----------

// DeactivateResp 注销用户响应
type DeactivateResp struct {
	types.BaseResp
	OperateID string `json:"operateId"`
}

// DeactivateQueryResp 查询注销用户列表响应
type DeactivateQueryResp struct {
	types.BaseResp
	Users []string `json:"users"`
}

// ---------- Delete (删除用户) ----------

// DeleteResp 删除用户响应
type DeleteResp struct {
	types.BaseResp
}

// ---------- Reactivate (重新激活) ----------

// ReactivateResp 重新激活用户响应
type ReactivateResp struct {
	types.BaseResp
	OperateID string `json:"operateId"`
}
