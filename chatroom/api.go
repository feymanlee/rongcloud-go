package chatroom

import (
	"strconv"

	"github.com/feymanlee/rongcloud-go/internal/core"
)

// API 路径常量
const (
	pathCreate          = "/chatroom/create.json"
	pathDestroy         = "/chatroom/destroy.json"
	pathQuery           = "/chatroom/query.json"
	pathQueryMembers    = "/chatroom/user/query.json"
	pathIsExist         = "/chatroom/user/exist.json"
	pathBlockAdd        = "/chatroom/user/block/add.json"
	pathBlockRemove     = "/chatroom/user/block/rollback.json"
	pathBlockQuery      = "/chatroom/user/block/list.json"
	pathKeepaliveAdd    = "/chatroom/keepalive/add.json"
	pathKeepaliveRemove = "/chatroom/keepalive/remove.json"
	pathKeepaliveQuery  = "/chatroom/keepalive/query.json"
)

// API 聊天室基础接口
type API interface {
	// Create 创建聊天室
	Create(chatrooms map[string]string) (*CreateResp, error)
	// Destroy 销毁聊天室
	Destroy(chatroomID string) (*DestroyResp, error)
	// Query 查询聊天室信息
	Query(chatroomID string) (*QueryResp, error)
	// QueryMembers 查询聊天室成员
	QueryMembers(chatroomID string, count, order int) (*QueryMembersResp, error)
	// IsExist 查询用户是否在聊天室
	IsExist(chatroomID, userID string) (*IsExistResp, error)
	// BlockAdd 添加聊天室封禁用户
	BlockAdd(chatroomID, userID string, minute uint) (*BlockAddResp, error)
	// BlockRemove 移除聊天室封禁用户
	BlockRemove(chatroomID, userID string) (*BlockRemoveResp, error)
	// BlockQuery 查询聊天室封禁用户列表
	BlockQuery(chatroomID string) (*BlockQueryResp, error)
	// KeepaliveAdd 添加保活聊天室
	KeepaliveAdd(chatroomID string) (*KeepaliveAddResp, error)
	// KeepaliveRemove 移除保活聊天室
	KeepaliveRemove(chatroomID string) (*KeepaliveRemoveResp, error)
	// KeepaliveQuery 查询保活聊天室列表
	KeepaliveQuery() (*KeepaliveQueryResp, error)
}

type api struct {
	client core.Client
}

// NewAPI 创建聊天室基础 API 实例
func NewAPI(client core.Client) API {
	return &api{client: client}
}

// Create 创建聊天室
func (a *api) Create(chatrooms map[string]string) (*CreateResp, error) {
	params := make(map[string]string)
	for id, name := range chatrooms {
		params["chatroom["+id+"]"] = name
	}
	resp := &CreateResp{}
	if err := a.client.Post(pathCreate, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Destroy 销毁聊天室
func (a *api) Destroy(chatroomID string) (*DestroyResp, error) {
	params := map[string]string{
		"chatroomId": chatroomID,
	}
	resp := &DestroyResp{}
	if err := a.client.Post(pathDestroy, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Query 查询聊天室信息
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

// QueryMembers 查询聊天室成员
func (a *api) QueryMembers(chatroomID string, count, order int) (*QueryMembersResp, error) {
	params := map[string]string{
		"chatroomId": chatroomID,
		"count":      strconv.Itoa(count),
		"order":      strconv.Itoa(order),
	}
	resp := &QueryMembersResp{}
	if err := a.client.Post(pathQueryMembers, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// IsExist 查询用户是否在聊天室
func (a *api) IsExist(chatroomID, userID string) (*IsExistResp, error) {
	params := map[string]string{
		"chatroomId": chatroomID,
		"userId":     userID,
	}
	resp := &IsExistResp{}
	if err := a.client.Post(pathIsExist, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// BlockAdd 添加聊天室封禁用户
func (a *api) BlockAdd(chatroomID, userID string, minute uint) (*BlockAddResp, error) {
	params := map[string]string{
		"chatroomId": chatroomID,
		"userId":     userID,
		"minute":     strconv.FormatUint(uint64(minute), 10),
	}
	resp := &BlockAddResp{}
	if err := a.client.Post(pathBlockAdd, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// BlockRemove 移除聊天室封禁用户
func (a *api) BlockRemove(chatroomID, userID string) (*BlockRemoveResp, error) {
	params := map[string]string{
		"chatroomId": chatroomID,
		"userId":     userID,
	}
	resp := &BlockRemoveResp{}
	if err := a.client.Post(pathBlockRemove, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// BlockQuery 查询聊天室封禁用户列表
func (a *api) BlockQuery(chatroomID string) (*BlockQueryResp, error) {
	params := map[string]string{
		"chatroomId": chatroomID,
	}
	resp := &BlockQueryResp{}
	if err := a.client.Post(pathBlockQuery, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// KeepaliveAdd 添加保活聊天室
func (a *api) KeepaliveAdd(chatroomID string) (*KeepaliveAddResp, error) {
	params := map[string]string{
		"chatroomId": chatroomID,
	}
	resp := &KeepaliveAddResp{}
	if err := a.client.Post(pathKeepaliveAdd, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// KeepaliveRemove 移除保活聊天室
func (a *api) KeepaliveRemove(chatroomID string) (*KeepaliveRemoveResp, error) {
	params := map[string]string{
		"chatroomId": chatroomID,
	}
	resp := &KeepaliveRemoveResp{}
	if err := a.client.Post(pathKeepaliveRemove, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// KeepaliveQuery 查询保活聊天室列表
func (a *api) KeepaliveQuery() (*KeepaliveQueryResp, error) {
	params := map[string]string{}
	resp := &KeepaliveQueryResp{}
	if err := a.client.Post(pathKeepaliveQuery, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
