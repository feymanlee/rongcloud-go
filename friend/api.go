package friend

import (
	"strings"

	"github.com/feymanlee/rongcloud-go/internal/core"
)

const (
	pathAdd                  = "/friend/add.json"
	pathRemove               = "/friend/remove.json"
	pathBatchRemove          = "/friend/batch/remove.json"
	pathQuery                = "/friend/query.json"
	pathSetRemark            = "/friend/set_remark.json"
	pathDirectionFriendQuery = "/friend/direction_friend/query.json"
	pathBlacklistQuery       = "/friend/blacklist/query.json"
)

// API 好友相关接口
type API interface {
	// Add 添加好友
	Add(userId, friendId, message, status string) (*AddResp, error)
	// Remove 删除好友
	Remove(userId, friendId string) (*RemoveResp, error)
	// BatchRemove 批量删除好友
	BatchRemove(userId string, friendIds []string) (*BatchRemoveResp, error)
	// Query 查询好友列表
	Query(userId string) (*QueryResp, error)
	// QueryByFriendId 根据好友 ID 查询好友信息
	QueryByFriendId(userId, friendId string) (*QueryResp, error)
	// SetRemark 设置好友备注
	SetRemark(userId, friendId, remark string) (*SetRemarkResp, error)
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
func (a *api) Add(userId, friendId, message, status string) (*AddResp, error) {
	resp := &AddResp{}
	params := map[string]string{
		"userId":   userId,
		"friendId": friendId,
		"message":  message,
		"status":   status,
	}
	err := a.client.Post(pathAdd, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Remove 删除好友
func (a *api) Remove(userId, friendId string) (*RemoveResp, error) {
	resp := &RemoveResp{}
	params := map[string]string{
		"userId":   userId,
		"friendId": friendId,
	}
	err := a.client.Post(pathRemove, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// BatchRemove 批量删除好友
func (a *api) BatchRemove(userId string, friendIds []string) (*BatchRemoveResp, error) {
	resp := &BatchRemoveResp{}
	params := map[string]string{
		"userId":    userId,
		"friendIds": strings.Join(friendIds, ","),
	}
	err := a.client.Post(pathBatchRemove, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Query 查询好友列表
func (a *api) Query(userId string) (*QueryResp, error) {
	resp := &QueryResp{}
	params := map[string]string{
		"userId": userId,
	}
	err := a.client.Post(pathQuery, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// QueryByFriendId 根据好友 ID 查询好友信息
func (a *api) QueryByFriendId(userId, friendId string) (*QueryResp, error) {
	resp := &QueryResp{}
	params := map[string]string{
		"userId":   userId,
		"friendId": friendId,
	}
	err := a.client.Post(pathQuery, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SetRemark 设置好友备注
func (a *api) SetRemark(userId, friendId, remark string) (*SetRemarkResp, error) {
	resp := &SetRemarkResp{}
	params := map[string]string{
		"userId":   userId,
		"friendId": friendId,
		"remark":   remark,
	}
	err := a.client.Post(pathSetRemark, params, resp)
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
