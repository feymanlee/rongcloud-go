package ultragroupchannel

import (
	"strconv"
	"strings"

	"github.com/feymanlee/rongcloud-go/internal/core"
)

const (
	pathCreate        = "/ultragroup/channel/private/create.json"
	pathDismiss       = "/ultragroup/channel/private/del.json"
	pathMembersAdd    = "/ultragroup/channel/private/users/add.json"
	pathMembersRemove = "/ultragroup/channel/private/users/del.json"
	pathMembersQuery  = "/ultragroup/channel/private/users/query.json"
)

// API 超级群私有频道相关接口
type API interface {
	// Create 创建私有频道
	Create(groupId, busChannel string) (*CreateResp, error)
	// Dismiss 删除私有频道
	Dismiss(groupId, busChannel string) (*DismissResp, error)
	// MembersAdd 添加私有频道成员
	MembersAdd(groupId, busChannel string, userIds []string) (*MembersAddResp, error)
	// MembersRemove 移除私有频道成员
	MembersRemove(groupId, busChannel string, userIds []string) (*MembersRemoveResp, error)
	// MembersQuery 查询私有频道成员
	MembersQuery(groupId, busChannel string, page, pageSize int) (*MembersQueryResp, error)
}

type api struct {
	client core.Client
}

// NewAPI 创建超级群私有频道 API 实例
func NewAPI(client core.Client) API {
	return &api{client: client}
}

// Create 创建私有频道
func (a *api) Create(groupId, busChannel string) (*CreateResp, error) {
	resp := &CreateResp{}
	params := map[string]string{
		"groupId":    groupId,
		"busChannel": busChannel,
	}
	err := a.client.Post(pathCreate, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Dismiss 删除私有频道
func (a *api) Dismiss(groupId, busChannel string) (*DismissResp, error) {
	resp := &DismissResp{}
	params := map[string]string{
		"groupId":    groupId,
		"busChannel": busChannel,
	}
	err := a.client.Post(pathDismiss, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// MembersAdd 添加私有频道成员
func (a *api) MembersAdd(groupId, busChannel string, userIds []string) (*MembersAddResp, error) {
	resp := &MembersAddResp{}
	params := map[string]string{
		"groupId":    groupId,
		"busChannel": busChannel,
		"userIds":    strings.Join(userIds, ","),
	}
	err := a.client.Post(pathMembersAdd, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// MembersRemove 移除私有频道成员
func (a *api) MembersRemove(groupId, busChannel string, userIds []string) (*MembersRemoveResp, error) {
	resp := &MembersRemoveResp{}
	params := map[string]string{
		"groupId":    groupId,
		"busChannel": busChannel,
		"userIds":    strings.Join(userIds, ","),
	}
	err := a.client.Post(pathMembersRemove, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// MembersQuery 查询私有频道成员
func (a *api) MembersQuery(groupId, busChannel string, page, pageSize int) (*MembersQueryResp, error) {
	resp := &MembersQueryResp{}
	params := map[string]string{
		"groupId":    groupId,
		"busChannel": busChannel,
		"page":       strconv.Itoa(page),
		"pageSize":   strconv.Itoa(pageSize),
	}
	err := a.client.Post(pathMembersQuery, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
