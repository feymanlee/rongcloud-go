package chatroompriority

import "github.com/feymanlee/rongcloud-go/internal/core"

// API 路径常量
const (
	pathAdd    = "/chatroom/message/priority/add.json"
	pathRemove = "/chatroom/message/priority/remove.json"
	pathQuery  = "/chatroom/message/priority/query.json"
)

// API 聊天室消息优先级接口
type API interface {
	// Add 添加聊天室消息优先级
	Add(objectName string) (*AddResp, error)
	// Remove 移除聊天室消息优先级
	Remove(objectName string) (*RemoveResp, error)
	// Query 查询聊天室消息优先级列表
	Query() (*QueryResp, error)
}

type api struct {
	client core.Client
}

// NewAPI 创建聊天室消息优先级 API 实例
func NewAPI(client core.Client) API {
	return &api{client: client}
}

// Add 添加聊天室消息优先级
func (a *api) Add(objectName string) (*AddResp, error) {
	params := map[string]string{
		"objectName": objectName,
	}
	resp := &AddResp{}
	if err := a.client.Post(pathAdd, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Remove 移除聊天室消息优先级
func (a *api) Remove(objectName string) (*RemoveResp, error) {
	params := map[string]string{
		"objectName": objectName,
	}
	resp := &RemoveResp{}
	if err := a.client.Post(pathRemove, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Query 查询聊天室消息优先级列表
func (a *api) Query() (*QueryResp, error) {
	params := map[string]string{}
	resp := &QueryResp{}
	if err := a.client.Post(pathQuery, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
