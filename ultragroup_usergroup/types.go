package ultragroupusergroup

import "github.com/feymanlee/rongcloud-go/internal/types"

// CreateReq 创建用户组请求
type CreateReq struct {
	GroupId     string `json:"groupId"`     // 超级群 ID
	UserGroupId string `json:"userGroupId"` // 用户组 ID
}

// CreateResp 创建用户组响应
type CreateResp struct {
	types.BaseResp
}

// DeleteReq 删除用户组请求
type DeleteReq struct {
	GroupId     string `json:"groupId"`     // 超级群 ID
	UserGroupId string `json:"userGroupId"` // 用户组 ID
}

// DeleteResp 删除用户组响应
type DeleteResp struct {
	types.BaseResp
}

// QueryReq 查询用户组请求
type QueryReq struct {
	GroupId string `json:"groupId"` // 超级群 ID
}

// UserGroupInfo 用户组信息
type UserGroupInfo struct {
	UserGroupId string `json:"userGroupId"` // 用户组 ID
}

// QueryResp 查询用户组响应
type QueryResp struct {
	types.BaseResp
	UserGroups []UserGroupInfo `json:"userGroups"` // 用户组列表
}

// MemberAddReq 添加用户组成员请求
type MemberAddReq struct {
	GroupId     string   `json:"groupId"`     // 超级群 ID
	UserGroupId string   `json:"userGroupId"` // 用户组 ID
	UserIds     []string `json:"userIds"`     // 用户 ID 列表
}

// MemberAddResp 添加用户组成员响应
type MemberAddResp struct {
	types.BaseResp
}

// MemberRemoveReq 移除用户组成员请求
type MemberRemoveReq struct {
	GroupId     string   `json:"groupId"`     // 超级群 ID
	UserGroupId string   `json:"userGroupId"` // 用户组 ID
	UserIds     []string `json:"userIds"`     // 用户 ID 列表
}

// MemberRemoveResp 移除用户组成员响应
type MemberRemoveResp struct {
	types.BaseResp
}

// MemberQueryReq 查询用户组成员请求
type MemberQueryReq struct {
	GroupId     string `json:"groupId"`     // 超级群 ID
	UserGroupId string `json:"userGroupId"` // 用户组 ID
}

// MemberQueryResp 查询用户组成员响应
type MemberQueryResp struct {
	types.BaseResp
	UserIds []string `json:"userIds"` // 成员用户 ID 列表
}

// ChannelBindReq 绑定频道请求
type ChannelBindReq struct {
	GroupId     string `json:"groupId"`     // 超级群 ID
	BusChannel  string `json:"busChannel"`  // 频道 ID
	UserGroupId string `json:"userGroupId"` // 用户组 ID
}

// ChannelBindResp 绑定频道响应
type ChannelBindResp struct {
	types.BaseResp
}

// ChannelUnbindReq 解绑频道请求
type ChannelUnbindReq struct {
	GroupId     string `json:"groupId"`     // 超级群 ID
	BusChannel  string `json:"busChannel"`  // 频道 ID
	UserGroupId string `json:"userGroupId"` // 用户组 ID
}

// ChannelUnbindResp 解绑频道响应
type ChannelUnbindResp struct {
	types.BaseResp
}

// ChannelQueryReq 查询频道绑定的用户组请求
type ChannelQueryReq struct {
	GroupId    string `json:"groupId"`    // 超级群 ID
	BusChannel string `json:"busChannel"` // 频道 ID
}

// ChannelQueryResp 查询频道绑定的用户组响应
type ChannelQueryResp struct {
	types.BaseResp
	UserGroups []UserGroupInfo `json:"userGroups"` // 用户组列表
}

// UserQueryReq 查询用户所属用户组请求
type UserQueryReq struct {
	GroupId string `json:"groupId"` // 超级群 ID
	UserId  string `json:"userId"`  // 用户 ID
}

// UserQueryResp 查询用户所属用户组响应
type UserQueryResp struct {
	types.BaseResp
	UserGroups []UserGroupInfo `json:"userGroups"` // 用户组列表
}
