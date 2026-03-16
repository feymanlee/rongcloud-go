package chatroomblock

import (
	"strconv"

	"github.com/feymanlee/rongcloud-go/internal/core"
)

// API 路径常量
const (
	pathAdd    = "/chatroom/block/add.json"
	pathRemove = "/chatroom/block/remove.json"
	pathQuery  = "/chatroom/block/query.json"
)

// API 聊天室封禁接口
type API interface {
	// Add 添加聊天室封禁用户
	Add(chatroomID, userID string, minute uint) (*AddResp, error)
	// Remove 移除聊天室封禁用户
	Remove(chatroomID, userID string) (*RemoveResp, error)
	// Query 查询聊天室封禁用户列表
	Query(chatroomID string) (*QueryResp, error)
}

type api struct {
	client core.Client
}

// NewAPI 创建聊天室封禁 API 实例
func NewAPI(client core.Client) API {
	return &api{client: client}
}

// Add 添加聊天室封禁用户
func (a *api) Add(chatroomID, userID string, minute uint) (*AddResp, error) {
	params := map[string]string{
		"chatroomId": chatroomID,
		"userId":     userID,
		"minute":     strconv.FormatUint(uint64(minute), 10),
	}
	resp := &AddResp{}
	if err := a.client.Post(pathAdd, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Remove 移除聊天室封禁用户
func (a *api) Remove(chatroomID, userID string) (*RemoveResp, error) {
	params := map[string]string{
		"chatroomId": chatroomID,
		"userId":     userID,
	}
	resp := &RemoveResp{}
	if err := a.client.Post(pathRemove, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Query 查询聊天室封禁用户列表
func (a *api) Query(chatroomID string) (*QueryResp, error) {
	params := map[string]string{
		"chatroomId": chatroomID,
	}
	resp := &QueryResp{}
	if err := a.client.Post(pathQuery, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
