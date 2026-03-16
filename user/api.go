package user

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/feymanlee/rongcloud-go/internal/core"
)

// API 路径常量
const (
	pathGetToken           = "/user/getToken.json"
	pathUpdate             = "/user/refresh.json"
	pathUserInfoGet        = "/user/info.json"
	pathTokenExpire        = "/user/token/expire.json"
	pathBlockAdd           = "/user/block.json"
	pathBlockRemove        = "/user/unblock.json"
	pathBlockQuery         = "/user/block/query.json"
	pathOnlineStatusCheck  = "/user/checkOnline.json"
	pathBan                = "/user/ban.json"
	pathBanQuery           = "/user/ban/query.json"
	pathUnBan              = "/user/unban.json"
	pathDeactivate         = "/user/deactivate.json"
	pathDeactivateQuery    = "/user/deactivate/query.json"
	pathReactivate         = "/user/reactivate.json"
)

// API 用户相关接口
type API interface {
	// GetToken 获取用户 Token
	GetToken(userID, name, portraitURI string) (*GetTokenResp, error)
	// Update 刷新用户信息
	Update(userID, name, portraitURI string) (*UpdateResp, error)
	// UserInfoGet 获取用户信息
	UserInfoGet(userID string) (*UserInfoGetResp, error)
	// TokenExpire 使用户 Token 过期
	TokenExpire(userID string, t int64) (*TokenExpireResp, error)
	// BlockAdd 封禁用户
	BlockAdd(userID string, minute uint64) (*BlockAddResp, error)
	// BlockRemove 解除封禁用户
	BlockRemove(userID string) (*BlockRemoveResp, error)
	// BlockQuery 查询被封禁用户列表
	BlockQuery() (*BlockQueryResp, error)
	// OnlineStatusCheck 查询用户在线状态
	OnlineStatusCheck(userID string) (*OnlineStatusCheckResp, error)
	// Ban 全局封禁用户
	Ban(userID string, minute uint64) (*BanResp, error)
	// BanQuery 查询全局封禁用户列表
	BanQuery() (*BanQueryResp, error)
	// UnBan 解除全局封禁用户
	UnBan(userID string) (*UnBanResp, error)
	// Deactivate 注销用户
	Deactivate(userIDs []string) (*DeactivateResp, error)
	// DeactivateQuery 查询注销用户列表
	DeactivateQuery(pageNo, pageSize int) (*DeactivateQueryResp, error)
	// Reactivate 重新激活用户
	Reactivate(userIDs []string) (*ReactivateResp, error)
}

type api struct {
	client core.Client
}

// NewAPI 创建用户 API 实例
func NewAPI(client core.Client) API {
	return &api{client: client}
}

// GetToken 获取用户 Token
func (a *api) GetToken(userID, name, portraitURI string) (*GetTokenResp, error) {
	params := map[string]string{
		"userId": userID,
		"name":   name,
	}
	if portraitURI != "" {
		params["portraitUri"] = portraitURI
	}
	resp := &GetTokenResp{}
	if err := a.client.Post(pathGetToken, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Update 刷新用户信息
func (a *api) Update(userID, name, portraitURI string) (*UpdateResp, error) {
	params := map[string]string{
		"userId":      userID,
		"name":        name,
		"portraitUri": portraitURI,
	}
	resp := &UpdateResp{}
	if err := a.client.Post(pathUpdate, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// UserInfoGet 获取用户信息
func (a *api) UserInfoGet(userID string) (*UserInfoGetResp, error) {
	params := map[string]string{
		"userId": userID,
	}
	resp := &UserInfoGetResp{}
	if err := a.client.Post(pathUserInfoGet, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// TokenExpire 使用户 Token 过期
func (a *api) TokenExpire(userID string, t int64) (*TokenExpireResp, error) {
	params := map[string]string{
		"userId": userID,
		"time":   fmt.Sprintf("%v", t),
	}
	resp := &TokenExpireResp{}
	if err := a.client.Post(pathTokenExpire, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// BlockAdd 封禁用户
func (a *api) BlockAdd(userID string, minute uint64) (*BlockAddResp, error) {
	params := map[string]string{
		"userId": userID,
		"minute": strconv.FormatUint(minute, 10),
	}
	resp := &BlockAddResp{}
	if err := a.client.Post(pathBlockAdd, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// BlockRemove 解除封禁用户
func (a *api) BlockRemove(userID string) (*BlockRemoveResp, error) {
	params := map[string]string{
		"userId": userID,
	}
	resp := &BlockRemoveResp{}
	if err := a.client.Post(pathBlockRemove, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// BlockQuery 查询被封禁用户列表
func (a *api) BlockQuery() (*BlockQueryResp, error) {
	params := map[string]string{}
	resp := &BlockQueryResp{}
	if err := a.client.Post(pathBlockQuery, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// OnlineStatusCheck 查询用户在线状态
func (a *api) OnlineStatusCheck(userID string) (*OnlineStatusCheckResp, error) {
	params := map[string]string{
		"userId": userID,
	}
	resp := &OnlineStatusCheckResp{}
	if err := a.client.Post(pathOnlineStatusCheck, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Ban 全局封禁用户
func (a *api) Ban(userID string, minute uint64) (*BanResp, error) {
	params := map[string]string{
		"userId": userID,
		"minute": strconv.FormatUint(minute, 10),
	}
	resp := &BanResp{}
	if err := a.client.Post(pathBan, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// BanQuery 查询全局封禁用户列表
func (a *api) BanQuery() (*BanQueryResp, error) {
	params := map[string]string{}
	resp := &BanQueryResp{}
	if err := a.client.Post(pathBanQuery, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// UnBan 解除全局封禁用户
func (a *api) UnBan(userID string) (*UnBanResp, error) {
	params := map[string]string{
		"userId": userID,
	}
	resp := &UnBanResp{}
	if err := a.client.Post(pathUnBan, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Deactivate 注销用户
func (a *api) Deactivate(userIDs []string) (*DeactivateResp, error) {
	params := map[string]string{
		"userId": strings.Join(userIDs, ","),
	}
	resp := &DeactivateResp{}
	if err := a.client.Post(pathDeactivate, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// DeactivateQuery 查询注销用户列表
func (a *api) DeactivateQuery(pageNo, pageSize int) (*DeactivateQueryResp, error) {
	params := map[string]string{
		"pageNo":   strconv.Itoa(pageNo),
		"pageSize": strconv.Itoa(pageSize),
	}
	resp := &DeactivateQueryResp{}
	if err := a.client.Post(pathDeactivateQuery, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Reactivate 重新激活用户
func (a *api) Reactivate(userIDs []string) (*ReactivateResp, error) {
	params := map[string]string{
		"userId": strings.Join(userIDs, ","),
	}
	resp := &ReactivateResp{}
	if err := a.client.Post(pathReactivate, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
