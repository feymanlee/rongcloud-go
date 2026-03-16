package groupmute

import (
	"github.com/feymanlee/rongcloud-go/internal/core"
)

const (
	pathMuteAdd         = "/group/ban/add.json"
	pathMuteRemove      = "/group/ban/remove.json"
	pathMuteQuery       = "/group/ban/query.json"
	pathMuteAllAdd      = "/group/ban/all/add.json"
	pathMuteAllRemove   = "/group/ban/all/remove.json"
	pathMuteAllQuery    = "/group/ban/all/query.json"
	pathWhitelistAdd    = "/group/ban/whitelist/add.json"
	pathWhitelistRemove = "/group/ban/whitelist/remove.json"
	pathWhitelistQuery  = "/group/ban/whitelist/query.json"
)

// API 群组禁言相关接口
type API interface {
	// MuteAdd 添加群组禁言成员
	MuteAdd(groupID, userID, minute string) (*MuteAddResp, error)
	// MuteRemove 移除群组禁言成员
	MuteRemove(groupID, userID string) (*MuteRemoveResp, error)
	// MuteQuery 查询群组禁言成员列表
	MuteQuery(groupID string) (*MuteQueryResp, error)
	// MuteAllAdd 添加群组全员禁言
	MuteAllAdd(groupID string) (*MuteAllAddResp, error)
	// MuteAllRemove 移除群组全员禁言
	MuteAllRemove(groupID string) (*MuteAllRemoveResp, error)
	// MuteAllQuery 查询全员禁言群组列表
	MuteAllQuery() (*MuteAllQueryResp, error)
	// WhitelistAdd 添加群组禁言白名单
	WhitelistAdd(groupID, userID string) (*WhitelistAddResp, error)
	// WhitelistRemove 移除群组禁言白名单
	WhitelistRemove(groupID, userID string) (*WhitelistRemoveResp, error)
	// WhitelistQuery 查询群组禁言白名单
	WhitelistQuery(groupID string) (*WhitelistQueryResp, error)
}

type api struct {
	client core.Client
}

// NewAPI 创建群组禁言 API 实例
func NewAPI(client core.Client) API {
	return &api{client: client}
}

// MuteAdd 添加群组禁言成员
func (a *api) MuteAdd(groupID, userID, minute string) (*MuteAddResp, error) {
	resp := &MuteAddResp{}
	params := map[string]string{
		"groupId": groupID,
		"userId":  userID,
		"minute":  minute,
	}
	if err := a.client.Post(pathMuteAdd, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// MuteRemove 移除群组禁言成员
func (a *api) MuteRemove(groupID, userID string) (*MuteRemoveResp, error) {
	resp := &MuteRemoveResp{}
	params := map[string]string{
		"groupId": groupID,
		"userId":  userID,
	}
	if err := a.client.Post(pathMuteRemove, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// MuteQuery 查询群组禁言成员列表
func (a *api) MuteQuery(groupID string) (*MuteQueryResp, error) {
	resp := &MuteQueryResp{}
	params := map[string]string{
		"groupId": groupID,
	}
	if err := a.client.Post(pathMuteQuery, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// MuteAllAdd 添加群组全员禁言
func (a *api) MuteAllAdd(groupID string) (*MuteAllAddResp, error) {
	resp := &MuteAllAddResp{}
	params := map[string]string{
		"groupId": groupID,
	}
	if err := a.client.Post(pathMuteAllAdd, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// MuteAllRemove 移除群组全员禁言
func (a *api) MuteAllRemove(groupID string) (*MuteAllRemoveResp, error) {
	resp := &MuteAllRemoveResp{}
	params := map[string]string{
		"groupId": groupID,
	}
	if err := a.client.Post(pathMuteAllRemove, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// MuteAllQuery 查询全员禁言群组列表
func (a *api) MuteAllQuery() (*MuteAllQueryResp, error) {
	resp := &MuteAllQueryResp{}
	params := map[string]string{}
	if err := a.client.Post(pathMuteAllQuery, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// WhitelistAdd 添加群组禁言白名单
func (a *api) WhitelistAdd(groupID, userID string) (*WhitelistAddResp, error) {
	resp := &WhitelistAddResp{}
	params := map[string]string{
		"groupId": groupID,
		"userId":  userID,
	}
	if err := a.client.Post(pathWhitelistAdd, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// WhitelistRemove 移除群组禁言白名单
func (a *api) WhitelistRemove(groupID, userID string) (*WhitelistRemoveResp, error) {
	resp := &WhitelistRemoveResp{}
	params := map[string]string{
		"groupId": groupID,
		"userId":  userID,
	}
	if err := a.client.Post(pathWhitelistRemove, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// WhitelistQuery 查询群组禁言白名单
func (a *api) WhitelistQuery(groupID string) (*WhitelistQueryResp, error) {
	resp := &WhitelistQueryResp{}
	params := map[string]string{
		"groupId": groupID,
	}
	if err := a.client.Post(pathWhitelistQuery, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
