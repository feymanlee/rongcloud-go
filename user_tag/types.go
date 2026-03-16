package usertag

import "github.com/feymanlee/rongcloud-go/internal/types"

// SetReq 设置用户标签请求
type SetReq struct {
	UserID string   `json:"userId"`
	Tags   []string `json:"tags"`
}

// BatchSetReq 批量设置用户标签请求
type BatchSetReq struct {
	UserIDs []string `json:"userIds"`
	Tags    []string `json:"tags"`
}

// GetReq 获取用户标签请求
type GetReq struct {
	UserIDs []string `json:"userIds"`
}

// GetResp 获取用户标签响应
type GetResp struct {
	types.BaseResp
	Result map[string][]string `json:"result"`
}
