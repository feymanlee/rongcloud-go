package groupmute

import "github.com/feymanlee/rongcloud-go/internal/types"

// MuteAddResp 群组禁言添加响应
type MuteAddResp struct {
	types.BaseResp
}

// MuteRemoveResp 群组禁言移除响应
type MuteRemoveResp struct {
	types.BaseResp
}

// MuteUser 群组禁言用户信息
type MuteUser struct {
	UserID string `json:"userId"` // 用户 ID
	Time   string `json:"time"`   // 禁言时间
}

// MuteQueryResp 查询群组禁言用户列表响应
type MuteQueryResp struct {
	types.BaseResp
	Users []MuteUser `json:"users"` // 禁言用户列表
}

// MuteAllAddResp 群组全员禁言添加响应
type MuteAllAddResp struct {
	types.BaseResp
}

// MuteAllRemoveResp 群组全员禁言移除响应
type MuteAllRemoveResp struct {
	types.BaseResp
}

// MuteAllGroup 全员禁言群组信息
type MuteAllGroup struct {
	GroupID string `json:"groupId"` // 群组 ID
	Stat    int    `json:"stat"`    // 禁言状态
}

// MuteAllQueryResp 查询全员禁言群组列表响应
type MuteAllQueryResp struct {
	types.BaseResp
	GroupInfo []MuteAllGroup `json:"groupinfo"` // 全员禁言群组列表
}

// WhitelistAddResp 群组禁言白名单添加响应
type WhitelistAddResp struct {
	types.BaseResp
}

// WhitelistRemoveResp 群组禁言白名单移除响应
type WhitelistRemoveResp struct {
	types.BaseResp
}

// WhitelistQueryResp 查询群组禁言白名单响应
type WhitelistQueryResp struct {
	types.BaseResp
	UserIDs []string `json:"userids"` // 白名单用户 ID 列表
}
