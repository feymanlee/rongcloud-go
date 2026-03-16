package conversation

import (
	"strconv"

	"github.com/feymanlee/rongcloud-go/internal/core"
)

const (
	pathNotificationSet = "/conversation/notification/set.json"
)

// API 会话相关接口
type API interface {
	// Mute 设置会话免打扰
	Mute(conversationType int, requestId, targetId string, isMuted int) (*MuteResp, error)
}

type api struct {
	client core.Client
}

// NewAPI 创建会话 API 实例
func NewAPI(client core.Client) API {
	return &api{client: client}
}

// Mute 设置会话免打扰
func (a *api) Mute(conversationType int, requestId, targetId string, isMuted int) (*MuteResp, error) {
	resp := &MuteResp{}
	params := map[string]string{
		"conversationType": strconv.Itoa(conversationType),
		"requestId":        requestId,
		"targetId":         targetId,
		"isMuted":          strconv.Itoa(isMuted),
	}
	err := a.client.Post(pathNotificationSet, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
