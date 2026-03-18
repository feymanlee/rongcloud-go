package userprofile

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/feymanlee/rongcloud-go/internal/core"
	"github.com/feymanlee/rongcloud-go/internal/types"
)

const (
	pathSet             = "/user/profile/set.json"
	pathGet             = "/user/profile/get.json"
	pathBatchGet        = "/user/profile/batch/get.json"
	pathBatchQuery      = "/user/profile/batch/query.json"
	pathClean           = "/user/profile/clean.json"
	pathQuery           = "/user/profile/query.json"
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
	// Clean 清除用户托管资料
	Clean(req *CleanReq) error
	// Query 分页查询用户列表
	Query(req *QueryReq) (*QueryResp, error)
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

func (a *api) Clean(req *CleanReq) error {
	resp := &types.BaseResp{}
	params := map[string]string{}
	if len(req.UserIDs) > 0 {
		params["userId"] = strings.Join(req.UserIDs, ",")
	}
	return a.client.Post(pathClean, params, resp)
}

func (a *api) Query(req *QueryReq) (*QueryResp, error) {
	resp := &QueryResp{}
	params := map[string]string{}
	if req.Page > 0 {
		params["page"] = strconv.Itoa(req.Page)
	}
	if req.Size > 0 {
		params["size"] = strconv.Itoa(req.Size)
	}
	if req.Order != 0 {
		params["order"] = strconv.Itoa(req.Order)
	}
	err := a.client.Post(pathQuery, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *api) CleanExpansion(req *CleanExpansionReq) error {
	resp := &types.BaseResp{}
	return a.client.PostJSON(pathCleanExpansion, req, resp)
}
