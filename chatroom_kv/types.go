package chatroomkv

import "github.com/feymanlee/rongcloud-go/internal/types"

// ---------- Set ----------

// SetResp 设置聊天室 KV 属性响应
type SetResp struct {
	types.BaseResp
}

// ---------- Remove ----------

// RemoveResp 移除聊天室 KV 属性响应
type RemoveResp struct {
	types.BaseResp
}

// ---------- Query ----------

// QueryResp 查询聊天室 KV 属性响应
type QueryResp struct {
	types.BaseResp
	Key   string `json:"key"`
	Value string `json:"value"`
}

// ---------- BatchSet ----------

// BatchSetReq 批量设置聊天室 KV 属性请求体
type BatchSetReq struct {
	ChatroomID   string            `json:"chatroomId"`
	AutoDelete   int               `json:"autoDelete"`
	EntryOwnerID string            `json:"entryOwnerId"`
	EntryInfo    map[string]string `json:"entryInfo"`
}

// BatchSetResp 批量设置聊天室 KV 属性响应
type BatchSetResp struct {
	types.BaseResp
}

// ---------- QueryAll ----------

// EntryInfo KV 属性信息
type EntryInfo struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// QueryAllResp 查询聊天室全部 KV 属性响应
type QueryAllResp struct {
	types.BaseResp
	Keys []EntryInfo `json:"keys"`
}
