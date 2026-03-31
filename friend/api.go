package friend

import (
	"errors"
	"strconv"
	"strings"

	"github.com/feymanlee/rongcloud-go/internal/core"
)

const (
	pathAdd                  = "/friend/add.json"
	pathDelete               = "/friend/delete.json"
	pathGet                  = "/friend/get.json"
	pathCheck                = "/friend/check.json"
	pathSetProfile           = "/friend/profile/set.json"
	pathDirectionFriendQuery = "/friend/direction_friend/query.json"
	pathBlacklistQuery       = "/friend/blacklist/query.json"
)

// API 好友相关接口
type API interface {
	// Add 添加好友
	Add(userId, targetId string, optType int, extra string) (*AddResp, error)
	// Remove 删除好友
	Remove(userId, targetId string) (*RemoveResp, error)
	// BatchRemove 批量删除好友
	BatchRemove(userId string, targetIds []string) (*BatchRemoveResp, error)
	// Query 查询好友列表（默认参数）
	Query(userId string) (*QueryResp, error)
	// QueryWithPage 按 pageToken 分页查询好友列表
	QueryWithPage(userId, pageToken string, size, order int) (*QueryResp, error)
	// Check 检查好友关系
	Check(userId string, targetIds []string) (*CheckResp, error)
	// QueryByFriendId 根据好友 ID 查询好友关系
	QueryByFriendId(userId, friendId string) (*CheckResp, error)
	// SetProfile 设置好友资料
	SetProfile(userId, targetId, remarkName, friendExtProfile string) (*SetProfileResp, error)
	// DirectionFriendQuery 查询方向好友
	DirectionFriendQuery(userId string) (*DirectionFriendQueryResp, error)
	// GetBlacklist 获取黑名单
	GetBlacklist(userId string) (*BlacklistResp, error)
}

type api struct {
	client core.Client
}

// NewAPI 创建好友 API 实例
func NewAPI(client core.Client) API {
	return &api{client: client}
}

// Add 添加好友
func (a *api) Add(userId, targetId string, optType int, extra string) (*AddResp, error) {
	resp := &AddResp{}
	params := map[string]string{
		"userId":   userId,
		"targetId": targetId,
	}
	if optType > 0 {
		params["optType"] = strconv.Itoa(optType)
	}
	if extra != "" {
		params["extra"] = extra
	}
	err := a.client.Post(pathAdd, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Remove 删除好友
func (a *api) Remove(userId, targetId string) (*RemoveResp, error) {
	if targetId == "" {
		return nil, errors.New("rongcloud: targetId is required")
	}

	resp := &RemoveResp{}
	params := map[string]string{
		"userId":    userId,
		"targetIds": targetId,
	}
	err := a.client.Post(pathDelete, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// BatchRemove 批量删除好友
func (a *api) BatchRemove(userId string, targetIds []string) (*BatchRemoveResp, error) {
	if len(targetIds) == 0 {
		return nil, errors.New("rongcloud: targetIds is required")
	}
	if len(targetIds) > 100 {
		return nil, errors.New("rongcloud: targetIds must not exceed 100")
	}

	resp := &BatchRemoveResp{}
	params := map[string]string{
		"userId":    userId,
		"targetIds": strings.Join(targetIds, ","),
	}
	err := a.client.Post(pathDelete, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Query 查询好友列表
func (a *api) Query(userId string) (*QueryResp, error) {
	return a.QueryWithPage(userId, "", 0, -1)
}

// QueryWithPage 按 pageToken 分页查询好友列表
func (a *api) QueryWithPage(userId, pageToken string, size, order int) (*QueryResp, error) {
	resp := &QueryResp{}
	params := map[string]string{
		"userId": userId,
	}
	if pageToken != "" {
		params["pageToken"] = pageToken
	}
	if size > 0 {
		params["size"] = strconv.Itoa(size)
	}
	if order >= 0 {
		params["order"] = strconv.Itoa(order)
	}
	err := a.client.Post(pathGet, params, resp)
	if err != nil {
		return nil, err
	}
	resp.syncCompatFields()
	return resp, nil
}

// Check 检查好友关系
func (a *api) Check(userId string, targetIds []string) (*CheckResp, error) {
	if len(targetIds) == 0 {
		return nil, errors.New("rongcloud: targetIds is required")
	}
	if len(targetIds) > 100 {
		return nil, errors.New("rongcloud: targetIds must not exceed 100")
	}

	resp := &CheckResp{}
	params := map[string]string{
		"userId":    userId,
		"targetIds": strings.Join(targetIds, ","),
	}
	err := a.client.Post(pathCheck, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// QueryByFriendId 根据好友 ID 查询好友关系
func (a *api) QueryByFriendId(userId, friendId string) (*CheckResp, error) {
	return a.Check(userId, []string{friendId})
}

// SetProfile 设置好友资料
func (a *api) SetProfile(userId, targetId, remarkName, friendExtProfile string) (*SetProfileResp, error) {
	resp := &SetProfileResp{}
	params := map[string]string{
		"userId":   userId,
		"targetId": targetId,
	}
	if remarkName != "" {
		params["remarkName"] = remarkName
	}
	if friendExtProfile != "" {
		params["friendExtProfile"] = friendExtProfile
	}
	err := a.client.Post(pathSetProfile, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DirectionFriendQuery 查询方向好友
func (a *api) DirectionFriendQuery(userId string) (*DirectionFriendQueryResp, error) {
	resp := &DirectionFriendQueryResp{}
	params := map[string]string{
		"userId": userId,
	}
	err := a.client.Post(pathDirectionFriendQuery, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetBlacklist 获取黑名单
func (a *api) GetBlacklist(userId string) (*BlacklistResp, error) {
	resp := &BlacklistResp{}
	params := map[string]string{
		"userId": userId,
	}
	err := a.client.Post(pathBlacklistQuery, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
