package chatroommute

import (
	"strconv"

	"github.com/feymanlee/rongcloud-go/internal/core"
)

// API 路径常量
const (
	pathMuteAdd          = "/chatroom/ban/add.json"
	pathMuteRemove       = "/chatroom/ban/remove.json"
	pathMuteQuery        = "/chatroom/ban/query.json"
	pathMuteAllAdd       = "/chatroom/ban/all/add.json"
	pathMuteAllRemove    = "/chatroom/ban/all/remove.json"
	pathMuteAllQuery     = "/chatroom/ban/all/query.json"
	pathGlobalMuteAdd    = "/chatroom/ban/global/add.json"
	pathGlobalMuteRemove = "/chatroom/ban/global/remove.json"
	pathGlobalMuteQuery  = "/chatroom/ban/global/query.json"
	pathWhitelistAdd     = "/chatroom/ban/whitelist/add.json"
	pathWhitelistRemove  = "/chatroom/ban/whitelist/remove.json"
	pathWhitelistQuery   = "/chatroom/ban/whitelist/query.json"
	pathUserWhitelistAdd = "/chatroom/user/ban/whitelist/add.json"
)

// API 聊天室禁言接口
type API interface {
	// MuteAdd 添加聊天室禁言用户
	MuteAdd(chatroomID, userID string, minute uint) (*MuteAddResp, error)
	// MuteRemove 移除聊天室禁言用户
	MuteRemove(chatroomID, userID string) (*MuteRemoveResp, error)
	// MuteQuery 查询聊天室禁言用户列表
	MuteQuery(chatroomID string) (*MuteQueryResp, error)
	// MuteAllAdd 添加聊天室全体禁言
	MuteAllAdd(chatroomID string) (*MuteAllAddResp, error)
	// MuteAllRemove 移除聊天室全体禁言
	MuteAllRemove(chatroomID string) (*MuteAllRemoveResp, error)
	// MuteAllQuery 查询聊天室全体禁言列表
	MuteAllQuery() (*MuteAllQueryResp, error)
	// GlobalMuteAdd 添加全局聊天室禁言用户
	GlobalMuteAdd(userID string, minute uint) (*GlobalMuteAddResp, error)
	// GlobalMuteRemove 移除全局聊天室禁言用户
	GlobalMuteRemove(userID string) (*GlobalMuteRemoveResp, error)
	// GlobalMuteQuery 查询全局聊天室禁言用户列表
	GlobalMuteQuery() (*GlobalMuteQueryResp, error)
	// WhitelistAdd 添加聊天室禁言白名单用户
	WhitelistAdd(chatroomID, userID string) (*WhitelistAddResp, error)
	// WhitelistRemove 移除聊天室禁言白名单用户
	WhitelistRemove(chatroomID, userID string) (*WhitelistRemoveResp, error)
	// WhitelistQuery 查询聊天室禁言白名单用户列表
	WhitelistQuery(chatroomID string) (*WhitelistQueryResp, error)
	// UserWhitelistAdd 添加聊天室用户禁言白名单
	UserWhitelistAdd(chatroomID, userID string) (*UserWhitelistAddResp, error)
}

type api struct {
	client core.Client
}

// NewAPI 创建聊天室禁言 API 实例
func NewAPI(client core.Client) API {
	return &api{client: client}
}

// MuteAdd 添加聊天室禁言用户
func (a *api) MuteAdd(chatroomID, userID string, minute uint) (*MuteAddResp, error) {
	params := map[string]string{
		"chatroomId": chatroomID,
		"userId":     userID,
		"minute":     strconv.FormatUint(uint64(minute), 10),
	}
	resp := &MuteAddResp{}
	if err := a.client.Post(pathMuteAdd, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// MuteRemove 移除聊天室禁言用户
func (a *api) MuteRemove(chatroomID, userID string) (*MuteRemoveResp, error) {
	params := map[string]string{
		"chatroomId": chatroomID,
		"userId":     userID,
	}
	resp := &MuteRemoveResp{}
	if err := a.client.Post(pathMuteRemove, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// MuteQuery 查询聊天室禁言用户列表
func (a *api) MuteQuery(chatroomID string) (*MuteQueryResp, error) {
	params := map[string]string{
		"chatroomId": chatroomID,
	}
	resp := &MuteQueryResp{}
	if err := a.client.Post(pathMuteQuery, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// MuteAllAdd 添加聊天室全体禁言
func (a *api) MuteAllAdd(chatroomID string) (*MuteAllAddResp, error) {
	params := map[string]string{
		"chatroomId": chatroomID,
	}
	resp := &MuteAllAddResp{}
	if err := a.client.Post(pathMuteAllAdd, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// MuteAllRemove 移除聊天室全体禁言
func (a *api) MuteAllRemove(chatroomID string) (*MuteAllRemoveResp, error) {
	params := map[string]string{
		"chatroomId": chatroomID,
	}
	resp := &MuteAllRemoveResp{}
	if err := a.client.Post(pathMuteAllRemove, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// MuteAllQuery 查询聊天室全体禁言列表
func (a *api) MuteAllQuery() (*MuteAllQueryResp, error) {
	params := map[string]string{}
	resp := &MuteAllQueryResp{}
	if err := a.client.Post(pathMuteAllQuery, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// GlobalMuteAdd 添加全局聊天室禁言用户
func (a *api) GlobalMuteAdd(userID string, minute uint) (*GlobalMuteAddResp, error) {
	params := map[string]string{
		"userId": userID,
		"minute": strconv.FormatUint(uint64(minute), 10),
	}
	resp := &GlobalMuteAddResp{}
	if err := a.client.Post(pathGlobalMuteAdd, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// GlobalMuteRemove 移除全局聊天室禁言用户
func (a *api) GlobalMuteRemove(userID string) (*GlobalMuteRemoveResp, error) {
	params := map[string]string{
		"userId": userID,
	}
	resp := &GlobalMuteRemoveResp{}
	if err := a.client.Post(pathGlobalMuteRemove, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// GlobalMuteQuery 查询全局聊天室禁言用户列表
func (a *api) GlobalMuteQuery() (*GlobalMuteQueryResp, error) {
	params := map[string]string{}
	resp := &GlobalMuteQueryResp{}
	if err := a.client.Post(pathGlobalMuteQuery, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// WhitelistAdd 添加聊天室禁言白名单用户
func (a *api) WhitelistAdd(chatroomID, userID string) (*WhitelistAddResp, error) {
	params := map[string]string{
		"chatroomId": chatroomID,
		"userId":     userID,
	}
	resp := &WhitelistAddResp{}
	if err := a.client.Post(pathWhitelistAdd, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// WhitelistRemove 移除聊天室禁言白名单用户
func (a *api) WhitelistRemove(chatroomID, userID string) (*WhitelistRemoveResp, error) {
	params := map[string]string{
		"chatroomId": chatroomID,
		"userId":     userID,
	}
	resp := &WhitelistRemoveResp{}
	if err := a.client.Post(pathWhitelistRemove, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// WhitelistQuery 查询聊天室禁言白名单用户列表
func (a *api) WhitelistQuery(chatroomID string) (*WhitelistQueryResp, error) {
	params := map[string]string{
		"chatroomId": chatroomID,
	}
	resp := &WhitelistQueryResp{}
	if err := a.client.Post(pathWhitelistQuery, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// UserWhitelistAdd 添加聊天室用户禁言白名单
func (a *api) UserWhitelistAdd(chatroomID, userID string) (*UserWhitelistAddResp, error) {
	params := map[string]string{
		"chatroomId": chatroomID,
		"userId":     userID,
	}
	resp := &UserWhitelistAddResp{}
	if err := a.client.Post(pathUserWhitelistAdd, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
