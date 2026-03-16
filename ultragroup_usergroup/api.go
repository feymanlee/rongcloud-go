package ultragroupusergroup

import (
	"github.com/feymanlee/rongcloud-go/internal/core"
)

const (
	pathCreate         = "/ultragroup/usergroup/create.json"
	pathDelete         = "/ultragroup/usergroup/del.json"
	pathQuery          = "/ultragroup/usergroup/query.json"
	pathMemberAdd      = "/ultragroup/usergroup/member/add.json"
	pathMemberRemove   = "/ultragroup/usergroup/member/del.json"
	pathMemberQuery    = "/ultragroup/usergroup/member/query.json"
	pathChannelBind    = "/ultragroup/usergroup/channel/bindv2.json"
	pathChannelUnbind  = "/ultragroup/usergroup/channel/unbind.json"
	pathChannelQuery   = "/ultragroup/usergroup/channel/query.json"
	pathUserQuery      = "/ultragroup/usergroup/user/query.json"
)

// API 超级群用户组相关接口
type API interface {
	// Create 创建用户组
	Create(req *CreateReq) (*CreateResp, error)
	// Delete 删除用户组
	Delete(req *DeleteReq) (*DeleteResp, error)
	// Query 查询用户组列表
	Query(req *QueryReq) (*QueryResp, error)
	// MemberAdd 添加用户组成员
	MemberAdd(req *MemberAddReq) (*MemberAddResp, error)
	// MemberRemove 移除用户组成员
	MemberRemove(req *MemberRemoveReq) (*MemberRemoveResp, error)
	// MemberQuery 查询用户组成员
	MemberQuery(req *MemberQueryReq) (*MemberQueryResp, error)
	// ChannelBind 绑定频道到用户组
	ChannelBind(req *ChannelBindReq) (*ChannelBindResp, error)
	// ChannelUnbind 解绑频道与用户组
	ChannelUnbind(req *ChannelUnbindReq) (*ChannelUnbindResp, error)
	// ChannelQuery 查询频道绑定的用户组
	ChannelQuery(req *ChannelQueryReq) (*ChannelQueryResp, error)
	// UserQuery 查询用户所属的用户组
	UserQuery(req *UserQueryReq) (*UserQueryResp, error)
}

type api struct {
	client core.Client
}

// NewAPI 创建超级群用户组 API 实例
func NewAPI(client core.Client) API {
	return &api{client: client}
}

// Create 创建用户组
func (a *api) Create(req *CreateReq) (*CreateResp, error) {
	resp := &CreateResp{}
	err := a.client.PostJSON(pathCreate, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Delete 删除用户组
func (a *api) Delete(req *DeleteReq) (*DeleteResp, error) {
	resp := &DeleteResp{}
	err := a.client.PostJSON(pathDelete, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Query 查询用户组列表
func (a *api) Query(req *QueryReq) (*QueryResp, error) {
	resp := &QueryResp{}
	err := a.client.PostJSON(pathQuery, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// MemberAdd 添加用户组成员
func (a *api) MemberAdd(req *MemberAddReq) (*MemberAddResp, error) {
	resp := &MemberAddResp{}
	err := a.client.PostJSON(pathMemberAdd, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// MemberRemove 移除用户组成员
func (a *api) MemberRemove(req *MemberRemoveReq) (*MemberRemoveResp, error) {
	resp := &MemberRemoveResp{}
	err := a.client.PostJSON(pathMemberRemove, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// MemberQuery 查询用户组成员
func (a *api) MemberQuery(req *MemberQueryReq) (*MemberQueryResp, error) {
	resp := &MemberQueryResp{}
	err := a.client.PostJSON(pathMemberQuery, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ChannelBind 绑定频道到用户组
func (a *api) ChannelBind(req *ChannelBindReq) (*ChannelBindResp, error) {
	resp := &ChannelBindResp{}
	err := a.client.PostJSON(pathChannelBind, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ChannelUnbind 解绑频道与用户组
func (a *api) ChannelUnbind(req *ChannelUnbindReq) (*ChannelUnbindResp, error) {
	resp := &ChannelUnbindResp{}
	err := a.client.PostJSON(pathChannelUnbind, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ChannelQuery 查询频道绑定的用户组
func (a *api) ChannelQuery(req *ChannelQueryReq) (*ChannelQueryResp, error) {
	resp := &ChannelQueryResp{}
	err := a.client.PostJSON(pathChannelQuery, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UserQuery 查询用户所属的用户组
func (a *api) UserQuery(req *UserQueryReq) (*UserQueryResp, error) {
	resp := &UserQueryResp{}
	err := a.client.PostJSON(pathUserQuery, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
