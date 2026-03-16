package usertag

import (
	"encoding/json"

	"github.com/feymanlee/rongcloud-go/internal/core"
	"github.com/feymanlee/rongcloud-go/internal/types"
)

const (
	pathTagSet      = "/user/tag/set.json"
	pathTagBatchSet = "/user/tag/batch/set.json"
	pathTagGet      = "/user/tag/get.json"
)

// API 用户标签接口
type API interface {
	// TagSet 设置用户标签
	TagSet(req *SetReq) error
	// TagBatchSet 批量设置用户标签
	TagBatchSet(req *BatchSetReq) error
	// TagGet 获取用户标签
	TagGet(req *GetReq) (*GetResp, error)
}

type api struct {
	client core.Client
}

// NewAPI 创建用户标签接口实例
func NewAPI(client core.Client) API {
	return &api{client: client}
}

// TagSet 设置用户标签
func (a *api) TagSet(req *SetReq) error {
	tagsJSON, _ := json.Marshal(req.Tags)
	resp := &types.BaseResp{}
	return a.client.Post(pathTagSet, map[string]string{
		"userId": req.UserID,
		"tags":   string(tagsJSON),
	}, resp)
}

// TagBatchSet 批量设置用户标签
func (a *api) TagBatchSet(req *BatchSetReq) error {
	userIdsJSON, _ := json.Marshal(req.UserIDs)
	tagsJSON, _ := json.Marshal(req.Tags)
	resp := &types.BaseResp{}
	return a.client.Post(pathTagBatchSet, map[string]string{
		"userIds": string(userIdsJSON),
		"tags":    string(tagsJSON),
	}, resp)
}

// TagGet 获取用户标签
func (a *api) TagGet(req *GetReq) (*GetResp, error) {
	userIdsJSON, _ := json.Marshal(req.UserIDs)
	resp := &GetResp{}
	err := a.client.Post(pathTagGet, map[string]string{
		"userIds": string(userIdsJSON),
	}, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
