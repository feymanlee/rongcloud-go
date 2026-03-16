package group

import "github.com/feymanlee/rongcloud-go/internal/types"

// ---------- Basic Group ----------

// CreateResp 创建群组响应
type CreateResp struct {
	types.BaseResp
}

// JoinResp 加入群组响应
type JoinResp struct {
	types.BaseResp
}

// QuitResp 退出群组响应
type QuitResp struct {
	types.BaseResp
}

// DismissResp 解散群组响应
type DismissResp struct {
	types.BaseResp
}

// RefreshResp 刷新群组信息响应
type RefreshResp struct {
	types.BaseResp
}

// GroupInfo 用户加入的群组信息
type GroupInfo struct {
	ID   string `json:"id"`   // 群组 ID
	Name string `json:"name"` // 群组名称
}

// QueryUserResp 查询用户所加入的群组响应
type QueryUserResp struct {
	types.BaseResp
	Groups []GroupInfo `json:"groups"` // 群组列表
}

// GroupUser 群组成员信息
type GroupUser struct {
	ID     string `json:"id"`     // 用户 ID
	UserID string `json:"userId"` // 用户 ID
	Time   string `json:"time"`   // 加入时间
}

// QueryMembersResp 查询群组成员响应
type QueryMembersResp struct {
	types.BaseResp
	ID    string      `json:"id"`    // 群组 ID
	Users []GroupUser `json:"users"` // 成员列表
}

// ---------- Entrust Group ----------

// EntrustCreateReq 委托创建群组请求
type EntrustCreateReq struct {
	GroupID         string `json:"groupId"`                   // 群组 ID
	Name            string `json:"name"`                      // 群组名称
	Owner           string `json:"owner"`                     // 群主用户 ID
	UserIDs         string `json:"userIds,omitempty"`         // 成员用户 ID 列表，逗号分隔
	GroupProfile    string `json:"groupProfile,omitempty"`    // 群组基本信息
	GroupExtProfile string `json:"groupExtProfile,omitempty"` // 群组扩展信息
	Permissions     string `json:"permissions,omitempty"`     // 权限设置
}

// EntrustCreateResp 委托创建群组响应
type EntrustCreateResp struct {
	types.BaseResp
}

// EntrustUpdateProfileReq 委托更新群组信息请求
type EntrustUpdateProfileReq struct {
	GroupID         string `json:"groupId"`                   // 群组 ID
	GroupProfile    string `json:"groupProfile,omitempty"`    // 群组基本信息
	GroupExtProfile string `json:"groupExtProfile,omitempty"` // 群组扩展信息
	Permissions     string `json:"permissions,omitempty"`     // 权限设置
}

// EntrustUpdateProfileResp 委托更新群组信息响应
type EntrustUpdateProfileResp struct {
	types.BaseResp
}

// EntrustGroupProfile 委托群组资料
type EntrustGroupProfile struct {
	GroupID         string `json:"groupId"`         // 群组 ID
	GroupName       string `json:"groupName"`       // 群组名称
	GroupProfile    string `json:"groupProfile"`    // 群组基本信息
	GroupExtProfile string `json:"groupExtProfile"` // 群组扩展信息
	Permissions     string `json:"permissions"`     // 权限设置
	CreateTime      int64  `json:"createTime"`      // 创建时间
	MemberCount     int    `json:"memberCount"`     // 成员数量
	Owner           string `json:"owner"`           // 群主用户 ID
}

// EntrustQueryProfilesResp 委托查询群组资料响应
type EntrustQueryProfilesResp struct {
	types.BaseResp
	Profiles []EntrustGroupProfile `json:"profiles"` // 群组资料列表
}

// EntrustSetNameResp 委托设置群组名称响应
type EntrustSetNameResp struct {
	types.BaseResp
}

// EntrustOwnerTransferReq 委托转让群主请求
type EntrustOwnerTransferReq struct {
	GroupID       string `json:"groupId"`                 // 群组 ID
	NewOwner      string `json:"newOwner"`                // 新群主用户 ID
	IsDelBan      *int   `json:"isDelBan,omitempty"`      // 是否删除禁言
	IsDelWhite    *int   `json:"isDelWhite,omitempty"`    // 是否删除白名单
	IsDelFollowed *int   `json:"isDelFollowed,omitempty"` // 是否删除关注
	IsQuit        *int   `json:"isQuit,omitempty"`        // 是否退出群组
}

// EntrustOwnerTransferResp 委托转让群主响应
type EntrustOwnerTransferResp struct {
	types.BaseResp
}

// EntrustJoinResp 委托加入群组响应
type EntrustJoinResp struct {
	types.BaseResp
	GroupID         string `json:"groupId"`         // 群组 ID
	Name            string `json:"name"`            // 群组名称
	Owner           string `json:"owner"`           // 群主用户 ID
	GroupProfile    string `json:"groupProfile"`    // 群组基本信息
	GroupExtProfile string `json:"groupExtProfile"` // 群组扩展信息
	Permissions     string `json:"permissions"`     // 权限设置
	RemarkName      string `json:"remarkName"`      // 备注名称
	CreateTime      int64  `json:"createTime"`      // 创建时间
	MemberCount     int    `json:"memberCount"`     // 成员数量
}

// EntrustQuitReq 委托退出群组请求
type EntrustQuitReq struct {
	GroupID       string `json:"groupId"`                 // 群组 ID
	UserIDs       string `json:"userIds"`                 // 用户 ID 列表，逗号分隔
	IsDelBan      *int   `json:"isDelBan,omitempty"`      // 是否删除禁言
	IsDelWhite    *int   `json:"isDelWhite,omitempty"`    // 是否删除白名单
	IsDelFollowed *int   `json:"isDelFollowed,omitempty"` // 是否删除关注
}

// EntrustQuitResp 委托退出群组响应
type EntrustQuitResp struct {
	types.BaseResp
}

// EntrustKickOutReq 委托踢出群组请求
type EntrustKickOutReq struct {
	GroupID       string `json:"groupId"`                 // 群组 ID
	UserIDs       string `json:"userIds"`                 // 用户 ID 列表，逗号分隔
	IsDelBan      *int   `json:"isDelBan,omitempty"`      // 是否删除禁言
	IsDelWhite    *int   `json:"isDelWhite,omitempty"`    // 是否删除白名单
	IsDelFollowed *int   `json:"isDelFollowed,omitempty"` // 是否删除关注
}

// EntrustKickOutResp 委托踢出群组响应
type EntrustKickOutResp struct {
	types.BaseResp
}

// EntrustMember 委托群组成员
type EntrustMember struct {
	UserID   string `json:"userId"`   // 用户 ID
	Nickname string `json:"nickname"` // 昵称
	Role     int    `json:"role"`     // 角色
	Time     int64  `json:"time"`     // 加入时间
	Extra    string `json:"extra"`    // 扩展信息
}

// EntrustMembersResp 委托查询群组成员响应
type EntrustMembersResp struct {
	types.BaseResp
	Members []EntrustMember `json:"members"` // 成员列表
}

// EntrustPagingMembersReq 委托分页查询群组成员请求
type EntrustPagingMembersReq struct {
	GroupID   string `json:"groupId"`             // 群组 ID
	Type      int    `json:"type,omitempty"`      // 查询类型
	PageToken string `json:"pageToken,omitempty"` // 分页 token
	Size      int    `json:"size,omitempty"`      // 每页大小
	Order     int    `json:"order,omitempty"`     // 排序方式
}

// EntrustPagingMembersResp 委托分页查询群组成员响应
type EntrustPagingMembersResp struct {
	types.BaseResp
	PageToken string          `json:"pageToken"` // 分页 token
	Members   []EntrustMember `json:"members"`   // 成员列表
}

// EntrustDismissResp 委托解散群组响应
type EntrustDismissResp struct {
	types.BaseResp
}

// EntrustSetManagersResp 委托设置管理员响应
type EntrustSetManagersResp struct {
	types.BaseResp
	UserIDs      []string `json:"userIds"`      // 未能设置的用户 ID（非群成员）
	ManagerCount int      `json:"managerCount"` // 管理员数量
}

// EntrustRemoveManagersResp 委托移除管理员响应
type EntrustRemoveManagersResp struct {
	types.BaseResp
}

// EntrustQueryManagersResp 委托查询管理员响应
type EntrustQueryManagersResp struct {
	types.BaseResp
	Members []EntrustMember `json:"members"` // 管理员列表
}

// EntrustSetMemberInfoReq 委托设置成员信息请求
type EntrustSetMemberInfoReq struct {
	GroupID  string `json:"groupId"`            // 群组 ID
	UserID   string `json:"userId"`             // 用户 ID
	Nickname string `json:"nickname,omitempty"` // 昵称
	Extra    string `json:"extra,omitempty"`    // 扩展信息
}

// EntrustSetMemberInfoResp 委托设置成员信息响应
type EntrustSetMemberInfoResp struct {
	types.BaseResp
}

// EntrustSetRemarkNameReq 委托设置群备注名请求
type EntrustSetRemarkNameReq struct {
	UserID     string `json:"userId"`     // 用户 ID
	GroupID    string `json:"groupId"`    // 群组 ID
	RemarkName string `json:"remarkName"` // 备注名称
}

// EntrustSetRemarkNameResp 委托设置群备注名响应
type EntrustSetRemarkNameResp struct {
	types.BaseResp
}

// EntrustQueryRemarkNameResp 委托查询群备注名响应
type EntrustQueryRemarkNameResp struct {
	types.BaseResp
	RemarkName string `json:"remarkName"` // 备注名称
}

// EntrustImportReq 委托导入群组请求
type EntrustImportReq struct {
	GroupID         string `json:"groupId"`                   // 群组 ID
	Name            string `json:"name"`                      // 群组名称
	Owner           string `json:"owner"`                     // 群主用户 ID
	GroupProfile    string `json:"groupProfile,omitempty"`    // 群组基本信息
	GroupExtProfile string `json:"groupExtProfile,omitempty"` // 群组扩展信息
	Permissions     string `json:"permissions,omitempty"`     // 权限设置
}

// EntrustImportResp 委托导入群组响应
type EntrustImportResp struct {
	types.BaseResp
}

// EntrustJoinedGroup 委托已加入群组信息
type EntrustJoinedGroup struct {
	GroupID         string `json:"groupId"`         // 群组 ID
	Name            string `json:"name"`            // 群组名称
	RemarkName      string `json:"remarkName"`      // 备注名称
	GroupProfile    string `json:"groupProfile"`    // 群组基本信息
	CreateTime      int64  `json:"createTime"`      // 创建时间
	Permissions     string `json:"permissions"`     // 权限设置
	GroupExtProfile string `json:"groupExtProfile"` // 群组扩展信息
	JoinTime        int64  `json:"joinTime"`        // 加入时间
	Role            int    `json:"role"`             // 角色
	Count           int    `json:"count"`            // 成员数量
	Owner           string `json:"owner"`            // 群主用户 ID
}

// EntrustQueryJoinedGroupsReq 委托查询已加入群组请求
type EntrustQueryJoinedGroupsReq struct {
	UserID    string `json:"userId"`              // 用户 ID
	Role      int    `json:"role,omitempty"`      // 角色
	PageToken string `json:"pageToken,omitempty"` // 分页 token
	Size      int    `json:"size,omitempty"`      // 每页大小
	Order     *int   `json:"order,omitempty"`     // 排序方式
}

// EntrustQueryJoinedGroupsResp 委托查询已加入群组响应
type EntrustQueryJoinedGroupsResp struct {
	types.BaseResp
	TotalCount int                  `json:"totalCount"` // 总数
	PageToken  string               `json:"pageToken"`  // 分页 token
	Groups     []EntrustJoinedGroup `json:"groups"`     // 群组列表
}

// EntrustFollowResp 委托关注群成员响应
type EntrustFollowResp struct {
	types.BaseResp
}

// EntrustUnfollowResp 委托取消关注群成员响应
type EntrustUnfollowResp struct {
	types.BaseResp
}

// FollowedMember 已关注的群成员
type FollowedMember struct {
	UserID    string `json:"userId"`    // 用户 ID
	Timestamp int64  `json:"timestamp"` // 关注时间
}

// EntrustQueryFollowedResp 委托查询已关注群成员响应
type EntrustQueryFollowedResp struct {
	types.BaseResp
	UserID  string           `json:"userId"`  // 用户 ID
	GroupID string           `json:"groupId"` // 群组 ID
	Members []FollowedMember `json:"members"` // 已关注成员列表
}

// EntrustGroupInfo 委托群组信息
type EntrustGroupInfo struct {
	GroupID     string `json:"groupId"`     // 群组 ID
	Owner       string `json:"owner"`       // 群主用户 ID
	Creator     string `json:"creator"`     // 创建者 ID
	Name        string `json:"name"`        // 群组名称
	PortraitURL string `json:"portraitUrl"` // 群组头像 URL
	Time        int64  `json:"time"`        // 创建时间
}

// EntrustPagingQueryReq 委托分页查询群组请求
type EntrustPagingQueryReq struct {
	PageToken string `json:"pageToken,omitempty"` // 分页 token
	Size      int    `json:"size,omitempty"`      // 每页大小
	Order     int    `json:"order,omitempty"`     // 排序方式
}

// EntrustPagingQueryResp 委托分页查询群组响应
type EntrustPagingQueryResp struct {
	types.BaseResp
	PageToken string             `json:"pageToken"` // 分页 token
	Groups    []EntrustGroupInfo `json:"groups"`    // 群组列表
}
