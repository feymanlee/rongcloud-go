package notification

import (
	"github.com/feymanlee/rongcloud-go/internal/core"
	"github.com/feymanlee/rongcloud-go/internal/types"
)

const (
	pathSet       = "/conversation/notification/set.json"
	pathGet       = "/conversation/notification/get.json"
	pathGlobalSet = "/conversation/notification/global/set.json"
	pathGlobalGet = "/conversation/notification/global/get.json"
)

// API 会话免打扰接口
type API interface {
	// Set 设置会话免打扰
	Set(conversationType, requestId, targetId, isMuted string) error
	// Get 获取会话免打扰状态
	Get(conversationType, requestId, targetId string) (*GetResp, error)
	// GlobalSet 设置全局会话免打扰
	GlobalSet(requestId, conversationType, isMuted string) error
	// GlobalGet 获取全局会话免打扰状态
	GlobalGet(requestId, conversationType string) (*GlobalGetResp, error)
}

type api struct {
	client core.Client
}

// NewAPI 创建会话免打扰接口实例
func NewAPI(client core.Client) API {
	return &api{client: client}
}

func (a *api) Set(conversationType, requestId, targetId, isMuted string) error {
	resp := &types.BaseResp{}
	return a.client.Post(pathSet, map[string]string{
		"conversationType": conversationType,
		"requestId":        requestId,
		"targetId":         targetId,
		"isMuted":          isMuted,
	}, resp)
}

func (a *api) Get(conversationType, requestId, targetId string) (*GetResp, error) {
	resp := &GetResp{}
	err := a.client.Post(pathGet, map[string]string{
		"conversationType": conversationType,
		"requestId":        requestId,
		"targetId":         targetId,
	}, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *api) GlobalSet(requestId, conversationType, isMuted string) error {
	resp := &types.BaseResp{}
	return a.client.Post(pathGlobalSet, map[string]string{
		"requestId":        requestId,
		"conversationType": conversationType,
		"isMuted":          isMuted,
	}, resp)
}

func (a *api) GlobalGet(requestId, conversationType string) (*GlobalGetResp, error) {
	resp := &GlobalGetResp{}
	err := a.client.Post(pathGlobalGet, map[string]string{
		"requestId":        requestId,
		"conversationType": conversationType,
	}, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
