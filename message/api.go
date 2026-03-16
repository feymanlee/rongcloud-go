package message

import (
	"strconv"
	"strings"

	"github.com/feymanlee/rongcloud-go/internal/core"
)

const (
	pathPrivatePublish         = "/message/private/publish.json"
	pathGroupPublish           = "/message/group/publish.json"
	pathChatroomPublish        = "/message/chatroom/publish.json"
	pathSystemPublish          = "/message/system/publish.json"
	pathBroadcast              = "/message/broadcast.json"
	pathRecall                 = "/message/recall.json"
	pathHistory                = "/message/history.json"
	pathHistoryDelete          = "/message/history/delete.json"
	pathPrivatePublishTemplate = "/message/private/publish_template.json"
	pathStatusMessageGroup     = "/statusmessage/group/publish.json"
	pathPrivateRecall          = "/message/private/recall.json"
	pathGroupRecall            = "/message/group/recall.json"
	pathExpansionSet           = "/message/expansion/set.json"
	pathExpansionQuery         = "/message/expansion/query.json"
)

// API 消息相关接口
type API interface {
	// SendPrivate 发送单聊消息
	SendPrivate(req *SendPrivateReq) (*SendResp, error)
	// SendGroup 发送群组消息
	SendGroup(req *SendGroupReq) (*SendResp, error)
	// SendChatroom 发送聊天室消息
	SendChatroom(req *SendChatroomReq) (*SendResp, error)
	// SendSystem 发送系统消息
	SendSystem(req *SendSystemReq) (*SendResp, error)
	// SendBroadcast 发送广播消息
	SendBroadcast(req *SendBroadcastReq) (*SendResp, error)
	// RecallPrivate 撤回单聊消息（通用撤回接口，conversationType=1）
	RecallPrivate(req *RecallReq) (*RecallResp, error)
	// RecallGroup 撤回群聊消息（通用撤回接口，conversationType=3）
	RecallGroup(req *RecallReq) (*RecallResp, error)
	// HistoryQuery 查询历史消息日志文件 URL
	HistoryQuery(date string) (*HistoryResp, error)
	// HistoryDelete 删除历史消息日志文件
	HistoryDelete(date string) (*HistoryDeleteResp, error)
	// SendPrivateTemplate 发送单聊模板消息（JSON 请求）
	SendPrivateTemplate(req *SendPrivateTemplateReq) (*SendResp, error)
	// SendStatusMessage 发送群组状态消息
	SendStatusMessage(req *SendStatusMessageReq) (*SendResp, error)
	// PrivateRecallMessage 单聊消息撤回（专用接口）
	PrivateRecallMessage(req *PrivateRecallReq) (*RecallResp, error)
	// GroupRecallMessage 群聊消息撤回（专用接口）
	GroupRecallMessage(req *GroupRecallReq) (*RecallResp, error)
	// ExpansionSet 设置消息扩展
	ExpansionSet(req *ExpansionSetReq) (*ExpansionSetResp, error)
	// ExpansionQuery 查询消息扩展
	ExpansionQuery(msgUID string, pageNo int) (*ExpansionQueryResp, error)
}

type api struct {
	client core.Client
}

// NewAPI 创建消息 API 实例
func NewAPI(client core.Client) API {
	return &api{client: client}
}

// SendPrivate 发送单聊消息
func (a *api) SendPrivate(req *SendPrivateReq) (*SendResp, error) {
	resp := &SendResp{}
	params := map[string]string{
		"fromUserId":  req.FromUserId,
		"toUserId":    strings.Join(req.ToUserId, ","),
		"objectName":  req.ObjectName,
		"content":     req.Content,
		"pushContent": req.PushContent,
		"pushData":    req.PushData,
	}
	err := a.client.Post(pathPrivatePublish, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SendGroup 发送群组消息
func (a *api) SendGroup(req *SendGroupReq) (*SendResp, error) {
	resp := &SendResp{}
	params := map[string]string{
		"fromUserId":  req.FromUserId,
		"toGroupId":   strings.Join(req.ToGroupId, ","),
		"objectName":  req.ObjectName,
		"content":     req.Content,
		"pushContent": req.PushContent,
		"pushData":    req.PushData,
	}
	err := a.client.Post(pathGroupPublish, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SendChatroom 发送聊天室消息
func (a *api) SendChatroom(req *SendChatroomReq) (*SendResp, error) {
	resp := &SendResp{}
	params := map[string]string{
		"fromUserId":   req.FromUserId,
		"toChatroomId": strings.Join(req.ToChatroomId, ","),
		"objectName":   req.ObjectName,
		"content":      req.Content,
	}
	err := a.client.Post(pathChatroomPublish, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SendSystem 发送系统消息
func (a *api) SendSystem(req *SendSystemReq) (*SendResp, error) {
	resp := &SendResp{}
	params := map[string]string{
		"fromUserId":  req.FromUserId,
		"toUserId":    strings.Join(req.ToUserId, ","),
		"objectName":  req.ObjectName,
		"content":     req.Content,
		"pushContent": req.PushContent,
		"pushData":    req.PushData,
	}
	err := a.client.Post(pathSystemPublish, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SendBroadcast 发送广播消息
func (a *api) SendBroadcast(req *SendBroadcastReq) (*SendResp, error) {
	resp := &SendResp{}
	params := map[string]string{
		"fromUserId": req.FromUserId,
		"objectName": req.ObjectName,
		"content":    req.Content,
	}
	err := a.client.Post(pathBroadcast, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// RecallPrivate 撤回单聊消息（通用撤回接口，conversationType=1）
func (a *api) RecallPrivate(req *RecallReq) (*RecallResp, error) {
	resp := &RecallResp{}
	params := map[string]string{
		"fromUserId":       req.FromUserId,
		"targetId":         req.TargetId,
		"messageUID":       req.MessageUID,
		"sentTime":         strconv.Itoa(req.SentTime),
		"conversationType": "1",
	}
	err := a.client.Post(pathRecall, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// RecallGroup 撤回群聊消息（通用撤回接口，conversationType=3）
func (a *api) RecallGroup(req *RecallReq) (*RecallResp, error) {
	resp := &RecallResp{}
	params := map[string]string{
		"fromUserId":       req.FromUserId,
		"targetId":         req.TargetId,
		"messageUID":       req.MessageUID,
		"sentTime":         strconv.Itoa(req.SentTime),
		"conversationType": "3",
	}
	err := a.client.Post(pathRecall, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// HistoryQuery 查询历史消息日志文件 URL
func (a *api) HistoryQuery(date string) (*HistoryResp, error) {
	resp := &HistoryResp{}
	params := map[string]string{
		"date": date,
	}
	err := a.client.Post(pathHistory, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// HistoryDelete 删除历史消息日志文件
func (a *api) HistoryDelete(date string) (*HistoryDeleteResp, error) {
	resp := &HistoryDeleteResp{}
	params := map[string]string{
		"date": date,
	}
	err := a.client.Post(pathHistoryDelete, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SendPrivateTemplate 发送单聊模板消息（JSON 请求）
func (a *api) SendPrivateTemplate(req *SendPrivateTemplateReq) (*SendResp, error) {
	resp := &SendResp{}
	err := a.client.PostJSON(pathPrivatePublishTemplate, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SendStatusMessage 发送群组状态消息
func (a *api) SendStatusMessage(req *SendStatusMessageReq) (*SendResp, error) {
	resp := &SendResp{}
	params := map[string]string{
		"fromUserId": req.FromUserId,
		"toGroupId":  strings.Join(req.ToGroupId, ","),
		"objectName": req.ObjectName,
		"content":    req.Content,
	}
	err := a.client.Post(pathStatusMessageGroup, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// PrivateRecallMessage 单聊消息撤回（专用接口）
func (a *api) PrivateRecallMessage(req *PrivateRecallReq) (*RecallResp, error) {
	resp := &RecallResp{}
	params := map[string]string{
		"fromUserId": req.FromUserId,
		"targetId":   req.TargetId,
		"messageUID": req.MessageUID,
		"sentTime":   strconv.Itoa(req.SentTime),
	}
	err := a.client.Post(pathPrivateRecall, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GroupRecallMessage 群聊消息撤回（专用接口）
func (a *api) GroupRecallMessage(req *GroupRecallReq) (*RecallResp, error) {
	resp := &RecallResp{}
	params := map[string]string{
		"fromUserId": req.FromUserId,
		"targetId":   req.TargetId,
		"messageUID": req.MessageUID,
		"sentTime":   strconv.Itoa(req.SentTime),
	}
	err := a.client.Post(pathGroupRecall, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ExpansionSet 设置消息扩展
func (a *api) ExpansionSet(req *ExpansionSetReq) (*ExpansionSetResp, error) {
	resp := &ExpansionSetResp{}
	params := map[string]string{
		"msgUID":           req.MsgUID,
		"userId":           req.UserId,
		"conversationType": req.ConversationType,
		"targetId":         req.TargetId,
		"extraKeyVal":      req.ExtraKeyVal,
		"isSyncSender":     strconv.Itoa(req.IsSyncSender),
	}
	err := a.client.Post(pathExpansionSet, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ExpansionQuery 查询消息扩展
func (a *api) ExpansionQuery(msgUID string, pageNo int) (*ExpansionQueryResp, error) {
	resp := &ExpansionQueryResp{}
	params := map[string]string{
		"msgUID": msgUID,
		"pageNo": strconv.Itoa(pageNo),
	}
	err := a.client.Post(pathExpansionQuery, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
