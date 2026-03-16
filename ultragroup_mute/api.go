package ultragroupmute

import (
	"strings"

	"github.com/feymanlee/rongcloud-go/internal/core"
)

const (
	pathMuteAdd          = "/ultragroup/ban/add.json"
	pathMuteRemove       = "/ultragroup/ban/remove.json"
	pathMuteQuery        = "/ultragroup/ban/query.json"
	pathMuteAllAdd       = "/ultragroup/ban/all/add.json"
	pathMuteAllRemove    = "/ultragroup/ban/all/remove.json"
	pathMuteAllQuery     = "/ultragroup/ban/all/query.json"
	pathWhitelistAdd     = "/ultragroup/ban/whitelist/add.json"
	pathWhitelistRemove  = "/ultragroup/ban/whitelist/remove.json"
)

// API 超级群禁言相关接口
type API interface {
	// MuteAdd 添加禁言成员
	MuteAdd(groupId string, userIds []string, busChannel string) (*MuteAddResp, error)
	// MuteRemove 移除禁言成员
	MuteRemove(groupId string, userIds []string, busChannel string) (*MuteRemoveResp, error)
	// MuteQuery 查询禁言成员列表
	MuteQuery(groupId, busChannel string) (*MuteQueryResp, error)
	// MuteAllAdd 设置全员禁言
	MuteAllAdd(groupId, busChannel string) (*MuteAllAddResp, error)
	// MuteAllRemove 取消全员禁言
	MuteAllRemove(groupId, busChannel string) (*MuteAllRemoveResp, error)
	// MuteAllQuery 查询全员禁言状态
	MuteAllQuery(groupId, busChannel string) (*MuteAllQueryResp, error)
	// WhitelistAdd 添加禁言白名单成员
	WhitelistAdd(groupId string, userIds []string, busChannel string) (*WhitelistAddResp, error)
	// WhitelistRemove 移除禁言白名单成员
	WhitelistRemove(groupId string, userIds []string, busChannel string) (*WhitelistRemoveResp, error)
}

type api struct {
	client core.Client
}

// NewAPI 创建超级群禁言 API 实例
func NewAPI(client core.Client) API {
	return &api{client: client}
}

// MuteAdd 添加禁言成员
func (a *api) MuteAdd(groupId string, userIds []string, busChannel string) (*MuteAddResp, error) {
	resp := &MuteAddResp{}
	params := map[string]string{
		"groupId":    groupId,
		"userIds":    strings.Join(userIds, ","),
		"busChannel": busChannel,
	}
	err := a.client.Post(pathMuteAdd, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// MuteRemove 移除禁言成员
func (a *api) MuteRemove(groupId string, userIds []string, busChannel string) (*MuteRemoveResp, error) {
	resp := &MuteRemoveResp{}
	params := map[string]string{
		"groupId":    groupId,
		"userIds":    strings.Join(userIds, ","),
		"busChannel": busChannel,
	}
	err := a.client.Post(pathMuteRemove, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// MuteQuery 查询禁言成员列表
func (a *api) MuteQuery(groupId, busChannel string) (*MuteQueryResp, error) {
	resp := &MuteQueryResp{}
	params := map[string]string{
		"groupId":    groupId,
		"busChannel": busChannel,
	}
	err := a.client.Post(pathMuteQuery, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// MuteAllAdd 设置全员禁言
func (a *api) MuteAllAdd(groupId, busChannel string) (*MuteAllAddResp, error) {
	resp := &MuteAllAddResp{}
	params := map[string]string{
		"groupId":    groupId,
		"busChannel": busChannel,
	}
	err := a.client.Post(pathMuteAllAdd, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// MuteAllRemove 取消全员禁言
func (a *api) MuteAllRemove(groupId, busChannel string) (*MuteAllRemoveResp, error) {
	resp := &MuteAllRemoveResp{}
	params := map[string]string{
		"groupId":    groupId,
		"busChannel": busChannel,
	}
	err := a.client.Post(pathMuteAllRemove, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// MuteAllQuery 查询全员禁言状态
func (a *api) MuteAllQuery(groupId, busChannel string) (*MuteAllQueryResp, error) {
	resp := &MuteAllQueryResp{}
	params := map[string]string{
		"groupId":    groupId,
		"busChannel": busChannel,
	}
	err := a.client.Post(pathMuteAllQuery, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// WhitelistAdd 添加禁言白名单成员
func (a *api) WhitelistAdd(groupId string, userIds []string, busChannel string) (*WhitelistAddResp, error) {
	resp := &WhitelistAddResp{}
	params := map[string]string{
		"groupId":    groupId,
		"userIds":    strings.Join(userIds, ","),
		"busChannel": busChannel,
	}
	err := a.client.Post(pathWhitelistAdd, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// WhitelistRemove 移除禁言白名单成员
func (a *api) WhitelistRemove(groupId string, userIds []string, busChannel string) (*WhitelistRemoveResp, error) {
	resp := &WhitelistRemoveResp{}
	params := map[string]string{
		"groupId":    groupId,
		"userIds":    strings.Join(userIds, ","),
		"busChannel": busChannel,
	}
	err := a.client.Post(pathWhitelistRemove, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
