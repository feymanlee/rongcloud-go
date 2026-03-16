package userblock

import (
	"github.com/feymanlee/rongcloud-go/internal/core"
	"github.com/feymanlee/rongcloud-go/internal/types"
)

const (
	pathBlacklistAdd    = "/user/blacklist/add.json"
	pathBlacklistRemove = "/user/blacklist/remove.json"
	pathBlacklistQuery  = "/user/blacklist/query.json"
	pathWhitelistAdd    = "/user/whitelist/add.json"
	pathWhitelistRemove = "/user/whitelist/remove.json"
	pathWhitelistQuery  = "/user/whitelist/query.json"
	pathMsgFilterAdd    = "/user/msgfilter/add.json"
	pathMsgFilterRemove = "/user/msgfilter/remove.json"
)

// API 用户黑名单/白名单接口
type API interface {
	// BlacklistAdd 添加黑名单
	BlacklistAdd(userId, blackUserId string) error
	// BlacklistRemove 移除黑名单
	BlacklistRemove(userId, blackUserId string) error
	// BlacklistQuery 查询黑名单
	BlacklistQuery(userId string) (*BlacklistResp, error)
	// WhitelistAdd 添加白名单
	WhitelistAdd(userId, whiteUserId string) error
	// WhitelistRemove 移除白名单
	WhitelistRemove(userId, whiteUserId string) error
	// WhitelistQuery 查询白名单
	WhitelistQuery(userId string) (*WhitelistResp, error)
	// MsgFilterAdd 添加消息过滤
	MsgFilterAdd(userId, filterType, targetIds string) error
	// MsgFilterRemove 移除消息过滤
	MsgFilterRemove(userId, filterType, targetIds string) error
}

type api struct {
	client core.Client
}

// NewAPI 创建用户黑名单/白名单接口实例
func NewAPI(client core.Client) API {
	return &api{client: client}
}

func (a *api) BlacklistAdd(userId, blackUserId string) error {
	resp := &types.BaseResp{}
	return a.client.Post(pathBlacklistAdd, map[string]string{
		"userId":      userId,
		"blackUserId": blackUserId,
	}, resp)
}

func (a *api) BlacklistRemove(userId, blackUserId string) error {
	resp := &types.BaseResp{}
	return a.client.Post(pathBlacklistRemove, map[string]string{
		"userId":      userId,
		"blackUserId": blackUserId,
	}, resp)
}

func (a *api) BlacklistQuery(userId string) (*BlacklistResp, error) {
	resp := &BlacklistResp{}
	err := a.client.Post(pathBlacklistQuery, map[string]string{
		"userId": userId,
	}, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *api) WhitelistAdd(userId, whiteUserId string) error {
	resp := &types.BaseResp{}
	return a.client.Post(pathWhitelistAdd, map[string]string{
		"userId":      userId,
		"whiteUserId": whiteUserId,
	}, resp)
}

func (a *api) WhitelistRemove(userId, whiteUserId string) error {
	resp := &types.BaseResp{}
	return a.client.Post(pathWhitelistRemove, map[string]string{
		"userId":      userId,
		"whiteUserId": whiteUserId,
	}, resp)
}

func (a *api) WhitelistQuery(userId string) (*WhitelistResp, error) {
	resp := &WhitelistResp{}
	err := a.client.Post(pathWhitelistQuery, map[string]string{
		"userId": userId,
	}, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *api) MsgFilterAdd(userId, filterType, targetIds string) error {
	resp := &types.BaseResp{}
	return a.client.Post(pathMsgFilterAdd, map[string]string{
		"userId":    userId,
		"type":      filterType,
		"targetIds": targetIds,
	}, resp)
}

func (a *api) MsgFilterRemove(userId, filterType, targetIds string) error {
	resp := &types.BaseResp{}
	return a.client.Post(pathMsgFilterRemove, map[string]string{
		"userId":    userId,
		"type":      filterType,
		"targetIds": targetIds,
	}, resp)
}
