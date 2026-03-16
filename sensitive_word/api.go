package sensitiveword

import (
	"github.com/feymanlee/rongcloud-go/internal/core"
	"github.com/feymanlee/rongcloud-go/internal/types"
)

const (
	pathAdd         = "/sensitiveword/add.json"
	pathBatchDelete = "/sensitiveword/batch/delete.json"
	pathList        = "/sensitiveword/list.json"
	pathDelete      = "/sensitiveword/delete.json"
)

// API 敏感词管理接口
type API interface {
	// Add 添加敏感词
	Add(word, replaceWord string) error
	// BatchDelete 批量删除敏感词
	BatchDelete(words string) error
	// List 查询敏感词列表
	List(wordType string) (*ListResp, error)
	// Delete 删除敏感词
	Delete(word string) error
}

type api struct {
	client core.Client
}

// NewAPI 创建敏感词管理接口实例
func NewAPI(client core.Client) API {
	return &api{client: client}
}

func (a *api) Add(word, replaceWord string) error {
	resp := &types.BaseResp{}
	return a.client.Post(pathAdd, map[string]string{
		"word":        word,
		"replaceWord": replaceWord,
	}, resp)
}

func (a *api) BatchDelete(words string) error {
	resp := &types.BaseResp{}
	return a.client.Post(pathBatchDelete, map[string]string{
		"words": words,
	}, resp)
}

func (a *api) List(wordType string) (*ListResp, error) {
	resp := &ListResp{}
	err := a.client.Post(pathList, map[string]string{
		"type": wordType,
	}, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *api) Delete(word string) error {
	resp := &types.BaseResp{}
	return a.client.Post(pathDelete, map[string]string{
		"word": word,
	}, resp)
}
