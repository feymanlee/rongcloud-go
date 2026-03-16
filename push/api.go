package push

import (
	"encoding/json"

	"github.com/feymanlee/rongcloud-go/internal/core"
	"github.com/feymanlee/rongcloud-go/internal/types"
)

const (
	pathPushUser            = "/push.json"
	pathPushCustom          = "/push/custom.json"
	pathBroadcastSend       = "/message/broadcast.json"
	pathBroadcastRecall     = "/message/broadcast/recall.json"
	pathBroadcastPush       = "/push/broadcast.json"
	pathSystemSend          = "/message/system/publish.json"
	pathSystemSendTemplate  = "/message/system/publish_template.json"
	pathSystemBroadcast     = "/message/broadcast.json"
	pathTagPush             = "/push/tag.json"
	pathOnlineBroadcast     = "/message/online/broadcast.json"
	pathPushRecall          = "/push/recall.json"
	pathPushQueryTask       = "/push/query/task.json"
	pathPushQueryStatus     = "/push/query/status.json"
	pathPushDelete          = "/push/delete.json"
)

// API 推送相关接口
type API interface {
	// PushUser 向应用中指定用户发送推送通知
	PushUser(platform []PlatForm, audience Audience, message Message, notification Notification) (*PushUserResp, error)
	// PushCustom 向用户发送自定义推送通知
	PushCustom(platform []PlatForm, audience Audience, message Message, notification Notification) (*PushCustomResp, error)
	// BroadcastSend 发送广播消息
	BroadcastSend(fromUserID, objectName, content, pushContent, pushData string) error
	// BroadcastRecall 撤回广播消息
	BroadcastRecall(objectName, content string) error
	// BroadcastPush 广播推送通知
	BroadcastPush(platform []PlatForm, fromUserID string, audience Audience, message Message, notification Notification) (*BroadcastPushResp, error)
	// SystemSend 发送系统消息
	SystemSend(fromUserID, toUserID, objectName, content, pushContent, pushData string) error
	// SystemSendTemplate 发送系统模板消息
	SystemSendTemplate(req *SystemSendTemplateReq) error
	// SystemBroadcast 系统广播消息
	SystemBroadcast(fromUserID, objectName, content, pushContent, pushData string) error
	// TagPush 标签推送
	TagPush(req *TagPushReq) (*TagPushResp, error)
	// OnlineBroadcast 在线广播消息
	OnlineBroadcast(fromUserID, objectName, content, pushContent, pushData string) error
	// PushRecall 撤回推送
	PushRecall(objectName, content string) error
	// PushQueryTask 查询推送任务
	PushQueryTask(taskID string) (*PushQueryTaskResp, error)
	// PushQueryStatus 查询推送状态
	PushQueryStatus(taskID string) (*PushQueryStatusResp, error)
	// PushDelete 删除推送任务
	PushDelete(taskID string) error
}

type api struct {
	client core.Client
}

// NewAPI 创建推送 API 实例
func NewAPI(client core.Client) API {
	return &api{client: client}
}

func (a *api) PushUser(platform []PlatForm, audience Audience, message Message, notification Notification) (*PushUserResp, error) {
	platformJSON, err := json.Marshal(platform)
	if err != nil {
		return nil, err
	}
	audienceJSON, err := json.Marshal(audience)
	if err != nil {
		return nil, err
	}
	messageJSON, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}
	notificationJSON, err := json.Marshal(notification)
	if err != nil {
		return nil, err
	}
	params := map[string]string{
		"platform":     string(platformJSON),
		"audience":     string(audienceJSON),
		"message":      string(messageJSON),
		"notification": string(notificationJSON),
	}
	resp := &PushUserResp{}
	if err := a.client.Post(pathPushUser, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *api) PushCustom(platform []PlatForm, audience Audience, message Message, notification Notification) (*PushCustomResp, error) {
	platformJSON, err := json.Marshal(platform)
	if err != nil {
		return nil, err
	}
	audienceJSON, err := json.Marshal(audience)
	if err != nil {
		return nil, err
	}
	messageJSON, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}
	notificationJSON, err := json.Marshal(notification)
	if err != nil {
		return nil, err
	}
	params := map[string]string{
		"platform":     string(platformJSON),
		"audience":     string(audienceJSON),
		"message":      string(messageJSON),
		"notification": string(notificationJSON),
	}
	resp := &PushCustomResp{}
	if err := a.client.Post(pathPushCustom, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *api) BroadcastSend(fromUserID, objectName, content, pushContent, pushData string) error {
	params := map[string]string{
		"fromUserId":  fromUserID,
		"objectName":  objectName,
		"content":     content,
		"pushContent": pushContent,
		"pushData":    pushData,
	}
	resp := &BroadcastSendResp{}
	return a.client.Post(pathBroadcastSend, params, resp)
}

func (a *api) BroadcastRecall(objectName, content string) error {
	params := map[string]string{
		"objectName": objectName,
		"content":    content,
	}
	resp := &BroadcastRecallResp{}
	return a.client.Post(pathBroadcastRecall, params, resp)
}

func (a *api) BroadcastPush(platform []PlatForm, fromUserID string, audience Audience, message Message, notification Notification) (*BroadcastPushResp, error) {
	platformJSON, err := json.Marshal(platform)
	if err != nil {
		return nil, err
	}
	audienceJSON, err := json.Marshal(audience)
	if err != nil {
		return nil, err
	}
	messageJSON, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}
	notificationJSON, err := json.Marshal(notification)
	if err != nil {
		return nil, err
	}
	params := map[string]string{
		"platform":     string(platformJSON),
		"fromuserid":   fromUserID,
		"audience":     string(audienceJSON),
		"message":      string(messageJSON),
		"notification": string(notificationJSON),
	}
	resp := &BroadcastPushResp{}
	if err := a.client.Post(pathBroadcastPush, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *api) SystemSend(fromUserID, toUserID, objectName, content, pushContent, pushData string) error {
	params := map[string]string{
		"fromUserId":  fromUserID,
		"toUserId":    toUserID,
		"objectName":  objectName,
		"content":     content,
		"pushContent": pushContent,
		"pushData":    pushData,
	}
	resp := &SystemSendResp{}
	return a.client.Post(pathSystemSend, params, resp)
}

func (a *api) SystemSendTemplate(req *SystemSendTemplateReq) error {
	resp := &SystemSendTemplateResp{}
	return a.client.PostJSON(pathSystemSendTemplate, req, resp)
}

func (a *api) SystemBroadcast(fromUserID, objectName, content, pushContent, pushData string) error {
	params := map[string]string{
		"fromUserId":  fromUserID,
		"objectName":  objectName,
		"content":     content,
		"pushContent": pushContent,
		"pushData":    pushData,
	}
	resp := &types.BaseResp{}
	return a.client.Post(pathSystemBroadcast, params, resp)
}

func (a *api) TagPush(req *TagPushReq) (*TagPushResp, error) {
	resp := &TagPushResp{}
	if err := a.client.PostJSON(pathTagPush, req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *api) OnlineBroadcast(fromUserID, objectName, content, pushContent, pushData string) error {
	params := map[string]string{
		"fromUserId":  fromUserID,
		"objectName":  objectName,
		"content":     content,
		"pushContent": pushContent,
		"pushData":    pushData,
	}
	resp := &OnlineBroadcastResp{}
	return a.client.Post(pathOnlineBroadcast, params, resp)
}

func (a *api) PushRecall(objectName, content string) error {
	params := map[string]string{
		"objectName": objectName,
		"content":    content,
	}
	resp := &PushRecallResp{}
	return a.client.Post(pathPushRecall, params, resp)
}

func (a *api) PushQueryTask(taskID string) (*PushQueryTaskResp, error) {
	params := map[string]string{
		"taskId": taskID,
	}
	resp := &PushQueryTaskResp{}
	if err := a.client.Post(pathPushQueryTask, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *api) PushQueryStatus(taskID string) (*PushQueryStatusResp, error) {
	params := map[string]string{
		"taskId": taskID,
	}
	resp := &PushQueryStatusResp{}
	if err := a.client.Post(pathPushQueryStatus, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *api) PushDelete(taskID string) error {
	params := map[string]string{
		"taskId": taskID,
	}
	resp := &PushDeleteResp{}
	return a.client.Post(pathPushDelete, params, resp)
}
