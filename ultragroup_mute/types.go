package ultragroupmute

import "github.com/feymanlee/rongcloud-go/internal/types"

// MuteAddResp 添加禁言成员响应
type MuteAddResp struct {
	types.BaseResp
}

// MuteRemoveResp 移除禁言成员响应
type MuteRemoveResp struct {
	types.BaseResp
}

// MutedUserInfo 禁言用户信息
type MutedUserInfo struct {
	Id   string `json:"id"`   // 用户 ID
	Time string `json:"time"` // 禁言时间
}

// MuteQueryResp 查询禁言成员响应
type MuteQueryResp struct {
	types.BaseResp
	Users []MutedUserInfo `json:"users"` // 禁言用户列表
}

// MuteAllAddResp 全员禁言响应
type MuteAllAddResp struct {
	types.BaseResp
}

// MuteAllRemoveResp 取消全员禁言响应
type MuteAllRemoveResp struct {
	types.BaseResp
}

// MuteAllQueryResp 查询全员禁言状态响应
type MuteAllQueryResp struct {
	types.BaseResp
	Status bool `json:"status"` // 全员禁言状态
}

// WhitelistAddResp 添加禁言白名单响应
type WhitelistAddResp struct {
	types.BaseResp
}

// WhitelistRemoveResp 移除禁言白名单响应
type WhitelistRemoveResp struct {
	types.BaseResp
}
