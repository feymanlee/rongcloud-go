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
	UserId           string `json:"userId"`           // 好友用户 ID
	Name             string `json:"name"`             // 好友名称
	RemarkName       string `json:"remarkName"`       // 好友备注名
	FriendExtProfile string `json:"friendExtProfile"` // 好友扩展资料
	Time             int64  `json:"time"`             // 建立好友关系时间

	// 兼容旧字段名；官方 get-friend 接口不返回这些字段。
	Message  string `json:"-"`
	Status   string `json:"-"`
	Remark   string `json:"-"`
	FriendId string `json:"-"`
}

// QueryResp 查询好友列表响应
type QueryResp struct {
	types.BaseResp
	PageToken  string       `json:"pageToken"`  // 下一页游标
	TotalCount int          `json:"totalCount"` // 好友总数
	FriendList []FriendInfo `json:"friendList"` // 好友列表

	// 兼容旧字段名。
	Friends []FriendInfo `json:"-"`
}

// CheckFriendResult 好友关系检查结果
type CheckFriendResult struct {
	UserId string `json:"userId"` // 目标用户 ID
	Result int    `json:"result"` // 检查结果
}

// CheckResp 检查好友关系响应
type CheckResp struct {
	types.BaseResp
	Results []CheckFriendResult `json:"results"`
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

func (r *QueryResp) syncCompatFields() {
	if r == nil {
		return
	}
	r.Friends = r.FriendList
	for i := range r.FriendList {
		r.FriendList[i].Remark = r.FriendList[i].RemarkName
		r.FriendList[i].FriendId = r.FriendList[i].UserId
	}
	r.Friends = r.FriendList
}
