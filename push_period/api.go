package pushperiod

import (
	"github.com/feymanlee/rongcloud-go/internal/core"
	"github.com/feymanlee/rongcloud-go/internal/types"
)

const (
	pathSet    = "/user/blockPushPeriod/set.json"
	pathGet    = "/user/blockPushPeriod/get.json"
	pathDelete = "/user/blockPushPeriod/delete.json"
)

// API 推送免打扰时段接口
type API interface {
	// Set 设置推送免打扰时段
	Set(userId, startTime, period, level string) error
	// Get 获取推送免打扰时段
	Get(userId string) (*GetResp, error)
	// Delete 删除推送免打扰时段
	Delete(userId string) error
}

type api struct {
	client core.Client
}

// NewAPI 创建推送免打扰时段接口实例
func NewAPI(client core.Client) API {
	return &api{client: client}
}

func (a *api) Set(userId, startTime, period, level string) error {
	resp := &types.BaseResp{}
	return a.client.Post(pathSet, map[string]string{
		"userId":    userId,
		"startTime": startTime,
		"period":    period,
		"level":     level,
	}, resp)
}

func (a *api) Get(userId string) (*GetResp, error) {
	resp := &GetResp{}
	err := a.client.Post(pathGet, map[string]string{
		"userId": userId,
	}, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *api) Delete(userId string) error {
	resp := &types.BaseResp{}
	return a.client.Post(pathDelete, map[string]string{
		"userId": userId,
	}, resp)
}
