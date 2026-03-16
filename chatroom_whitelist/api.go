package chatroomwhitelist

import "github.com/feymanlee/rongcloud-go/internal/core"

// API 路径常量
const (
	pathMsgAdd     = "/chatroom/whitelist/add.json"
	pathMsgRemove  = "/chatroom/whitelist/delete.json"
	pathMsgQuery   = "/chatroom/whitelist/query.json"
	pathUserAdd    = "/chatroom/user/whitelist/add.json"
	pathUserRemove = "/chatroom/user/whitelist/remove.json"
	pathUserQuery  = "/chatroom/user/whitelist/query.json"
)

// API 聊天室白名单接口
type API interface {
	// MsgAdd 添加聊天室消息白名单
	MsgAdd(objectNames string) (*MsgAddResp, error)
	// MsgRemove 移除聊天室消息白名单
	MsgRemove(objectNames string) (*MsgRemoveResp, error)
	// MsgQuery 查询聊天室消息白名单
	MsgQuery() (*MsgQueryResp, error)
	// UserAdd 添加聊天室用户白名单
	UserAdd(chatroomID, userID string) (*UserAddResp, error)
	// UserRemove 移除聊天室用户白名单
	UserRemove(chatroomID, userID string) (*UserRemoveResp, error)
	// UserQuery 查询聊天室用户白名单
	UserQuery(chatroomID string) (*UserQueryResp, error)
}

type api struct {
	client core.Client
}

// NewAPI 创建聊天室白名单 API 实例
func NewAPI(client core.Client) API {
	return &api{client: client}
}

// MsgAdd 添加聊天室消息白名单
func (a *api) MsgAdd(objectNames string) (*MsgAddResp, error) {
	params := map[string]string{
		"objectnames": objectNames,
	}
	resp := &MsgAddResp{}
	if err := a.client.Post(pathMsgAdd, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// MsgRemove 移除聊天室消息白名单
func (a *api) MsgRemove(objectNames string) (*MsgRemoveResp, error) {
	params := map[string]string{
		"objectnames": objectNames,
	}
	resp := &MsgRemoveResp{}
	if err := a.client.Post(pathMsgRemove, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// MsgQuery 查询聊天室消息白名单
func (a *api) MsgQuery() (*MsgQueryResp, error) {
	params := map[string]string{}
	resp := &MsgQueryResp{}
	if err := a.client.Post(pathMsgQuery, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// UserAdd 添加聊天室用户白名单
func (a *api) UserAdd(chatroomID, userID string) (*UserAddResp, error) {
	params := map[string]string{
		"chatroomId": chatroomID,
		"userId":     userID,
	}
	resp := &UserAddResp{}
	if err := a.client.Post(pathUserAdd, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// UserRemove 移除聊天室用户白名单
func (a *api) UserRemove(chatroomID, userID string) (*UserRemoveResp, error) {
	params := map[string]string{
		"chatroomId": chatroomID,
		"userId":     userID,
	}
	resp := &UserRemoveResp{}
	if err := a.client.Post(pathUserRemove, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// UserQuery 查询聊天室用户白名单
func (a *api) UserQuery(chatroomID string) (*UserQueryResp, error) {
	params := map[string]string{
		"chatroomId": chatroomID,
	}
	resp := &UserQueryResp{}
	if err := a.client.Post(pathUserQuery, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
