package chatroomkv

import (
	"strconv"

	"github.com/feymanlee/rongcloud-go/internal/core"
)

// API 路径常量
const (
	pathSet      = "/chatroom/entry/set.json"
	pathRemove   = "/chatroom/entry/remove.json"
	pathQuery    = "/chatroom/entry/query.json"
	pathBatchSet = "/chatroom/entry/batch/set.json"
	pathQueryAll = "/chatroom/entry/query/all.json"
)

// API 聊天室 KV 属性接口
type API interface {
	// Set 设置聊天室 KV 属性
	Set(chatroomID, userID, key, value string, autoDelete int, objectName string) (*SetResp, error)
	// Remove 移除聊天室 KV 属性
	Remove(chatroomID, userID, key string) (*RemoveResp, error)
	// Query 查询聊天室 KV 属性
	Query(chatroomID, key string) (*QueryResp, error)
	// BatchSet 批量设置聊天室 KV 属性 (JSON)
	BatchSet(req *BatchSetReq) (*BatchSetResp, error)
	// QueryAll 查询聊天室全部 KV 属性
	QueryAll(chatroomID string) (*QueryAllResp, error)
}

type api struct {
	client core.Client
}

// NewAPI 创建聊天室 KV 属性 API 实例
func NewAPI(client core.Client) API {
	return &api{client: client}
}

// Set 设置聊天室 KV 属性
func (a *api) Set(chatroomID, userID, key, value string, autoDelete int, objectName string) (*SetResp, error) {
	params := map[string]string{
		"chatroomId": chatroomID,
		"userId":     userID,
		"key":        key,
		"value":      value,
		"autoDelete": strconv.Itoa(autoDelete),
	}
	if objectName != "" {
		params["objectName"] = objectName
	}
	resp := &SetResp{}
	if err := a.client.Post(pathSet, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Remove 移除聊天室 KV 属性
func (a *api) Remove(chatroomID, userID, key string) (*RemoveResp, error) {
	params := map[string]string{
		"chatroomId": chatroomID,
		"userId":     userID,
		"key":        key,
	}
	resp := &RemoveResp{}
	if err := a.client.Post(pathRemove, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Query 查询聊天室 KV 属性
func (a *api) Query(chatroomID, key string) (*QueryResp, error) {
	params := map[string]string{
		"chatroomId": chatroomID,
		"key":        key,
	}
	resp := &QueryResp{}
	if err := a.client.Post(pathQuery, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// BatchSet 批量设置聊天室 KV 属性 (JSON)
func (a *api) BatchSet(req *BatchSetReq) (*BatchSetResp, error) {
	resp := &BatchSetResp{}
	if err := a.client.PostJSON(pathBatchSet, req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// QueryAll 查询聊天室全部 KV 属性
func (a *api) QueryAll(chatroomID string) (*QueryAllResp, error) {
	params := map[string]string{
		"chatroomId": chatroomID,
	}
	resp := &QueryAllResp{}
	if err := a.client.Post(pathQueryAll, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
