package userprofile

import "github.com/feymanlee/rongcloud-go/internal/types"

// UserProfile 用户基本资料
type UserProfile struct {
	UniqueId    string `json:"uniqueId,omitempty"`    // 应用标识，最长 32 字符
	Name        string `json:"name,omitempty"`        // 用户昵称，最长 64 字符
	PortraitUri string `json:"portraitUri,omitempty"` // 头像 URL，最长 1024 字符
	Email       string `json:"email,omitempty"`       // 邮箱，最长 128 字符
	Birthday    string `json:"birthday,omitempty"`    // 生日，最长 32 字符
	Gender      *int   `json:"gender,omitempty"`      // 性别：0 未知，1 男，2 女
	Location    string `json:"location,omitempty"`    // 地理位置，最长 32 字符
	Role        *int   `json:"role,omitempty"`        // 用户角色，0-100
	Level       *int   `json:"level,omitempty"`       // 用户等级，0-100
}

// SetReq 设置用户资料请求
type SetReq struct {
	UserID         string            // 用户 ID
	UserProfile    *UserProfile      // 基本资料（JSON 编码后作为 userProfile 参数）
	UserExtProfile map[string]string // 自定义扩展属性（key 以 ext_ 开头，JSON 编码后作为 userExtProfile 参数）
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
