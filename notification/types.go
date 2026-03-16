package notification

import "github.com/feymanlee/rongcloud-go/internal/types"

// GetResp 获取会话免打扰响应
type GetResp struct {
	types.BaseResp
	IsMuted int `json:"isMuted"`
}

// GlobalGetResp 获取全局会话免打扰响应
type GlobalGetResp struct {
	types.BaseResp
	IsMuted int `json:"isMuted"`
}
