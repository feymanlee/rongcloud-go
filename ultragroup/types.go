package ultragroup

import "github.com/feymanlee/rongcloud-go/internal/types"

// CreateResp 创建超级群响应
type CreateResp struct {
	types.BaseResp
}

// DismissResp 解散超级群响应
type DismissResp struct {
	types.BaseResp
}

// JoinResp 加入超级群响应
type JoinResp struct {
	types.BaseResp
}

// QuitResp 退出超级群响应
type QuitResp struct {
	types.BaseResp
}

// RefreshResp 刷新超级群信息响应
type RefreshResp struct {
	types.BaseResp
}

// MemberInfo 超级群成员信息
type MemberInfo struct {
	Id string `json:"id"` // 用户 ID
}

// QueryMembersResp 查询超级群成员响应
type QueryMembersResp struct {
	types.BaseResp
	Members []MemberInfo `json:"members"` // 成员列表
}

// GroupInfo 超级群信息
type GroupInfo struct {
	GroupId   string `json:"groupId"`   // 超级群 ID
	GroupName string `json:"groupName"` // 超级群名称
}

// QueryUserResp 查询用户所属超级群响应
type QueryUserResp struct {
	types.BaseResp
	Groups []GroupInfo `json:"groups"` // 超级群列表
}

// HisMsgPublishResp 发送超级群历史消息响应
type HisMsgPublishResp struct {
	types.BaseResp
}

// HisMsgRecallResp 撤回超级群历史消息响应
type HisMsgRecallResp struct {
	types.BaseResp
}

// ExpansionSetResp 设置扩展信息响应
type ExpansionSetResp struct {
	types.BaseResp
}

// ExpansionRemoveResp 删除扩展信息响应
type ExpansionRemoveResp struct {
	types.BaseResp
}

// ExpansionInfo 扩展信息
type ExpansionInfo struct {
	Key   string `json:"key"`   // 扩展 key
	Value string `json:"value"` // 扩展 value
}

// ExpansionQueryResp 查询扩展信息响应
type ExpansionQueryResp struct {
	types.BaseResp
	Keys []ExpansionInfo `json:"keys"` // 扩展信息列表
}

// MsgModifyResp 修改消息响应
type MsgModifyResp struct {
	types.BaseResp
}

// NotDisturbSetResp 设置免打扰响应
type NotDisturbSetResp struct {
	types.BaseResp
}

// NotDisturbGetResp 查询免打扰响应
type NotDisturbGetResp struct {
	types.BaseResp
	UnPushLevel int `json:"unpushLevel"` // 免打扰级别
}

// ChannelCreateResp 创建频道响应
type ChannelCreateResp struct {
	types.BaseResp
}

// ChannelDelResp 删除频道响应
type ChannelDelResp struct {
	types.BaseResp
}
