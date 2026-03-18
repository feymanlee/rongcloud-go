package userprofile

import (
	"encoding/json"
	"strings"

	"github.com/feymanlee/rongcloud-go/internal/core"
	"github.com/feymanlee/rongcloud-go/internal/types"
)

const (
	pathSet             = "/user/profile/set.json"
	pathGet             = "/user/profile/get.json"
	pathBatchGet        = "/user/profile/batch/get.json"
	pathBatchQuery      = "/user/profile/batch/query.json"
	pathCleanExpansion  = "/user/profile/expansion/clean.json"
)

// API 用户资料托管接口
type API interface {
	// Set 设置用户资料
	Set(req *SetReq) error
	// Get 获取用户资料
	Get(req *GetReq) (*GetResp, error)
	// BatchGet 批量获取用户资料
	BatchGet(req *BatchGetReq) (*BatchGetResp, error)
	// BatchQuery 批量查询用户资料（返回完整 userProfile 和 userExtProfile）
	BatchQuery(req *BatchQueryReq) (*BatchQueryResp, error)
	// CleanExpansion 清除扩展信息
	CleanExpansion(req *CleanExpansionReq) error
}

type api struct {
	client core.Client
}

// NewAPI 创建用户资料托管接口实例
func NewAPI(client core.Client) API {
	return &api{client: client}
}

func (a *api) Set(req *SetReq) error {
	resp := &types.BaseResp{}
	params := map[string]string{
		"userId": req.UserID,
	}
	if req.UserProfile != nil {
		b, _ := json.Marshal(req.UserProfile)
		params["userProfile"] = string(b)
	}
	if len(req.UserExtProfile) > 0 {
		b, _ := json.Marshal(req.UserExtProfile)
		params["userExtProfile"] = string(b)
	}
	return a.client.Post(pathSet, params, resp)
}

func (a *api) Get(req *GetReq) (*GetResp, error) {
	resp := &GetResp{}
	err := a.client.PostJSON(pathGet, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *api) BatchGet(req *BatchGetReq) (*BatchGetResp, error) {
	resp := &BatchGetResp{}
	err := a.client.PostJSON(pathBatchGet, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *api) BatchQuery(req *BatchQueryReq) (*BatchQueryResp, error) {
	resp := &BatchQueryResp{}
	params := map[string]string{
		"userId": strings.Join(req.UserIDs, ","),
	}
	err := a.client.Post(pathBatchQuery, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *api) CleanExpansion(req *CleanExpansionReq) error {
	resp := &types.BaseResp{}
	return a.client.PostJSON(pathCleanExpansion, req, resp)
}
