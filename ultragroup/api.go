package ultragroup

import (
	"strconv"
	"strings"

	"github.com/feymanlee/rongcloud-go/internal/core"
)

const (
	pathCreate           = "/ultragroup/create.json"
	pathDismiss          = "/ultragroup/dis.json"
	pathJoin             = "/ultragroup/join.json"
	pathQuit             = "/ultragroup/quit.json"
	pathRefresh          = "/ultragroup/refresh.json"
	pathQueryMembers     = "/ultragroup/member/query.json"
	pathQueryUser        = "/user/ultragroup/query.json"
	pathHisMsgPublish    = "/message/ultragroup/publish.json"
	pathHisMsgRecall     = "/message/ultragroup/recall.json"
	pathExpansionSet     = "/ultragroup/expansion/set.json"
	pathExpansionRemove  = "/ultragroup/expansion/remove.json"
	pathExpansionQuery   = "/ultragroup/expansion/query.json"
	pathMsgModify        = "/ultragroup/msg/modify.json"
	pathNotDisturbSet    = "/ultragroup/notdisturb/set.json"
	pathNotDisturbGet    = "/ultragroup/notdisturb/get.json"
	pathChannelCreate    = "/ultragroup/channel/create.json"
	pathChannelDel       = "/ultragroup/channel/del.json"
)

// API 超级群相关接口
type API interface {
	// Create 创建超级群
	Create(userId, groupId, groupName string) (*CreateResp, error)
	// Dismiss 解散超级群
	Dismiss(groupId string) (*DismissResp, error)
	// Join 加入超级群
	Join(userId, groupId string) (*JoinResp, error)
	// Quit 退出超级群
	Quit(userId, groupId string) (*QuitResp, error)
	// Refresh 刷新超级群信息
	Refresh(groupId, groupName string) (*RefreshResp, error)
	// QueryMembers 查询超级群成员列表
	QueryMembers(groupId string, page, pageSize int) (*QueryMembersResp, error)
	// QueryUser 查询用户所属超级群
	QueryUser(userId string) (*QueryUserResp, error)
	// HisMsgPublish 发送超级群历史消息
	HisMsgPublish(fromUserId, toGroupId, objectName, content string) (*HisMsgPublishResp, error)
	// HisMsgRecall 撤回超级群历史消息
	HisMsgRecall(fromUserId, toGroupId, messageUID, busChannel string, sentTime string) (*HisMsgRecallResp, error)
	// ExpansionSet 设置超级群消息扩展信息
	ExpansionSet(groupId, msgUID, busChannel, extraKeyVal string) (*ExpansionSetResp, error)
	// ExpansionRemove 删除超级群消息扩展信息
	ExpansionRemove(groupId, msgUID, busChannel string, extraKeys []string) (*ExpansionRemoveResp, error)
	// ExpansionQuery 查询超级群消息扩展信息
	ExpansionQuery(groupId, msgUID, busChannel string) (*ExpansionQueryResp, error)
	// MsgModify 修改超级群消息
	MsgModify(fromUserId, groupId, msgUID, busChannel, content, objectName string, sentTime string) (*MsgModifyResp, error)
	// NotDisturbSet 设置超级群免打扰
	NotDisturbSet(groupId, busChannel string, unpushLevel int) (*NotDisturbSetResp, error)
	// NotDisturbGet 查询超级群免打扰
	NotDisturbGet(groupId string) (*NotDisturbGetResp, error)
	// ChannelCreate 创建超级群频道
	ChannelCreate(groupId, busChannel, channelType string) (*ChannelCreateResp, error)
	// ChannelDel 删除超级群频道
	ChannelDel(groupId, busChannel string) (*ChannelDelResp, error)
}

type api struct {
	client core.Client
}

// NewAPI 创建超级群 API 实例
func NewAPI(client core.Client) API {
	return &api{client: client}
}

// Create 创建超级群
func (a *api) Create(userId, groupId, groupName string) (*CreateResp, error) {
	resp := &CreateResp{}
	params := map[string]string{
		"userId":    userId,
		"groupId":   groupId,
		"groupName": groupName,
	}
	err := a.client.Post(pathCreate, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Dismiss 解散超级群
func (a *api) Dismiss(groupId string) (*DismissResp, error) {
	resp := &DismissResp{}
	params := map[string]string{
		"groupId": groupId,
	}
	err := a.client.Post(pathDismiss, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Join 加入超级群
func (a *api) Join(userId, groupId string) (*JoinResp, error) {
	resp := &JoinResp{}
	params := map[string]string{
		"userId":  userId,
		"groupId": groupId,
	}
	err := a.client.Post(pathJoin, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Quit 退出超级群
func (a *api) Quit(userId, groupId string) (*QuitResp, error) {
	resp := &QuitResp{}
	params := map[string]string{
		"userId":  userId,
		"groupId": groupId,
	}
	err := a.client.Post(pathQuit, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Refresh 刷新超级群信息
func (a *api) Refresh(groupId, groupName string) (*RefreshResp, error) {
	resp := &RefreshResp{}
	params := map[string]string{
		"groupId":   groupId,
		"groupName": groupName,
	}
	err := a.client.Post(pathRefresh, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// QueryMembers 查询超级群成员列表
func (a *api) QueryMembers(groupId string, page, pageSize int) (*QueryMembersResp, error) {
	resp := &QueryMembersResp{}
	params := map[string]string{
		"groupId":  groupId,
		"page":     strconv.Itoa(page),
		"pageSize": strconv.Itoa(pageSize),
	}
	err := a.client.Post(pathQueryMembers, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// QueryUser 查询用户所属超级群
func (a *api) QueryUser(userId string) (*QueryUserResp, error) {
	resp := &QueryUserResp{}
	params := map[string]string{
		"userId": userId,
	}
	err := a.client.Post(pathQueryUser, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// HisMsgPublish 发送超级群历史消息
func (a *api) HisMsgPublish(fromUserId, toGroupId, objectName, content string) (*HisMsgPublishResp, error) {
	resp := &HisMsgPublishResp{}
	params := map[string]string{
		"fromUserId": fromUserId,
		"toGroupId":  toGroupId,
		"objectName": objectName,
		"content":    content,
	}
	err := a.client.Post(pathHisMsgPublish, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// HisMsgRecall 撤回超级群历史消息
func (a *api) HisMsgRecall(fromUserId, toGroupId, messageUID, busChannel string, sentTime string) (*HisMsgRecallResp, error) {
	resp := &HisMsgRecallResp{}
	params := map[string]string{
		"fromUserId": fromUserId,
		"toGroupId":  toGroupId,
		"messageUID": messageUID,
		"busChannel": busChannel,
		"sentTime":   sentTime,
	}
	err := a.client.Post(pathHisMsgRecall, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ExpansionSet 设置超级群消息扩展信息
func (a *api) ExpansionSet(groupId, msgUID, busChannel, extraKeyVal string) (*ExpansionSetResp, error) {
	resp := &ExpansionSetResp{}
	params := map[string]string{
		"groupId":     groupId,
		"msgUID":      msgUID,
		"busChannel":  busChannel,
		"extraKeyVal": extraKeyVal,
	}
	err := a.client.Post(pathExpansionSet, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ExpansionRemove 删除超级群消息扩展信息
func (a *api) ExpansionRemove(groupId, msgUID, busChannel string, extraKeys []string) (*ExpansionRemoveResp, error) {
	resp := &ExpansionRemoveResp{}
	params := map[string]string{
		"groupId":    groupId,
		"msgUID":     msgUID,
		"busChannel": busChannel,
		"extraKey":   strings.Join(extraKeys, ","),
	}
	err := a.client.Post(pathExpansionRemove, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ExpansionQuery 查询超级群消息扩展信息
func (a *api) ExpansionQuery(groupId, msgUID, busChannel string) (*ExpansionQueryResp, error) {
	resp := &ExpansionQueryResp{}
	params := map[string]string{
		"groupId":    groupId,
		"msgUID":     msgUID,
		"busChannel": busChannel,
	}
	err := a.client.Post(pathExpansionQuery, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// MsgModify 修改超级群消息
func (a *api) MsgModify(fromUserId, groupId, msgUID, busChannel, content, objectName string, sentTime string) (*MsgModifyResp, error) {
	resp := &MsgModifyResp{}
	params := map[string]string{
		"fromUserId": fromUserId,
		"groupId":    groupId,
		"msgUID":     msgUID,
		"busChannel": busChannel,
		"content":    content,
		"objectName": objectName,
		"sentTime":   sentTime,
	}
	err := a.client.Post(pathMsgModify, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// NotDisturbSet 设置超级群免打扰
func (a *api) NotDisturbSet(groupId, busChannel string, unpushLevel int) (*NotDisturbSetResp, error) {
	resp := &NotDisturbSetResp{}
	params := map[string]string{
		"groupId":      groupId,
		"busChannel":   busChannel,
		"unpushLevel":  strconv.Itoa(unpushLevel),
	}
	err := a.client.Post(pathNotDisturbSet, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// NotDisturbGet 查询超级群免打扰
func (a *api) NotDisturbGet(groupId string) (*NotDisturbGetResp, error) {
	resp := &NotDisturbGetResp{}
	params := map[string]string{
		"groupId": groupId,
	}
	err := a.client.Post(pathNotDisturbGet, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ChannelCreate 创建超级群频道
func (a *api) ChannelCreate(groupId, busChannel, channelType string) (*ChannelCreateResp, error) {
	resp := &ChannelCreateResp{}
	params := map[string]string{
		"groupId":    groupId,
		"busChannel": busChannel,
		"type":       channelType,
	}
	err := a.client.Post(pathChannelCreate, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ChannelDel 删除超级群频道
func (a *api) ChannelDel(groupId, busChannel string) (*ChannelDelResp, error) {
	resp := &ChannelDelResp{}
	params := map[string]string{
		"groupId":    groupId,
		"busChannel": busChannel,
	}
	err := a.client.Post(pathChannelDel, params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
