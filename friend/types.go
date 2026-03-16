package friend

import "github.com/feymanlee/rongcloud-go/internal/types"

// AddResp 添加好友响应
type AddResp struct {
	types.BaseResp
}

// RemoveResp 删除好友响应
type RemoveResp struct {
	types.BaseResp
}

// BatchRemoveResp 批量删除好友响应
type BatchRemoveResp struct {
	types.BaseResp
}

// SetRemarkResp 设置好友备注响应
type SetRemarkResp struct {
	types.BaseResp
}

// FriendInfo 好友信息
type FriendInfo struct {
	UserId   string `json:"userId"`   // 好友用户 ID
	Message  string `json:"message"`  // 好友消息
	Status   string `json:"status"`   // 好友状态
	Remark   string `json:"remark"`   // 好友备注
	FriendId string `json:"friendId"` // 好友 ID
}

// QueryResp 查询好友列表响应
type QueryResp struct {
	types.BaseResp
	Friends []FriendInfo `json:"friends"` // 好友列表
}

// DirectionFriendQueryResp 查询方向好友响应
type DirectionFriendQueryResp struct {
	types.BaseResp
	Friends []FriendInfo `json:"friends"` // 好友列表
}

// BlacklistResp 获取黑名单响应
type BlacklistResp struct {
	types.BaseResp
	Users []string `json:"users"` // 黑名单用户列表
}
