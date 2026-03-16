package userprofile

import "github.com/feymanlee/rongcloud-go/internal/types"

// ProfileItem 用户资料项
type ProfileItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// SetReq 设置用户资料请求
type SetReq struct {
	UserID   string        `json:"userId"`
	Profile  []ProfileItem `json:"profile,omitempty"`
	Expansion map[string]string `json:"expansion,omitempty"`
}

// GetReq 获取用户资料请求
type GetReq struct {
	UserID string   `json:"userId"`
	Keys   []string `json:"keys,omitempty"`
}

// GetResp 获取用户资料响应
type GetResp struct {
	types.BaseResp
	Data map[string]string `json:"data"`
}

// BatchGetReq 批量获取用户资料请求
type BatchGetReq struct {
	UserIDs []string `json:"userIds"`
	Keys    []string `json:"keys,omitempty"`
}

// BatchGetResp 批量获取用户资料响应
type BatchGetResp struct {
	types.BaseResp
	Data map[string]map[string]string `json:"data"`
}

// CleanExpansionReq 清除扩展信息请求
type CleanExpansionReq struct {
	UserID string   `json:"userId"`
	Keys   []string `json:"keys,omitempty"`
}
