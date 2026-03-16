package chatroompriority

import "github.com/feymanlee/rongcloud-go/internal/types"

// ---------- Add ----------

// AddResp 添加聊天室消息优先级响应
type AddResp struct {
	types.BaseResp
}

// ---------- Remove ----------

// RemoveResp 移除聊天室消息优先级响应
type RemoveResp struct {
	types.BaseResp
}

// ---------- Query ----------

// QueryResp 查询聊天室消息优先级列表响应
type QueryResp struct {
	types.BaseResp
	ObjectNames []string `json:"objectNames"`
}
