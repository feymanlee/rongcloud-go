package ultragroupchannel

import "github.com/feymanlee/rongcloud-go/internal/types"

// CreateResp 创建私有频道响应
type CreateResp struct {
	types.BaseResp
}

// DismissResp 删除私有频道响应
type DismissResp struct {
	types.BaseResp
}

// MembersAddResp 添加私有频道成员响应
type MembersAddResp struct {
	types.BaseResp
}

// MembersRemoveResp 移除私有频道成员响应
type MembersRemoveResp struct {
	types.BaseResp
}

// MembersQueryResp 查询私有频道成员响应
type MembersQueryResp struct {
	types.BaseResp
	Users []string `json:"users"` // 成员用户 ID 列表
}
