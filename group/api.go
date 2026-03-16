package group

import (
	"strconv"
	"strings"

	"github.com/feymanlee/rongcloud-go/internal/core"
)

const (
	// Basic Group
	pathCreate       = "/group/create.json"
	pathJoin         = "/group/join.json"
	pathQuit         = "/group/quit.json"
	pathDismiss      = "/group/dismiss.json"
	pathRefresh      = "/group/refresh.json"
	pathQueryUser    = "/user/group/query.json"
	pathQueryMembers = "/group/user/query.json"

	// Entrust Group
	pathEntrustCreate            = "/entrust/group/create.json"
	pathEntrustUpdateProfile     = "/entrust/group/profile/update.json"
	pathEntrustQueryProfiles     = "/entrust/group/profile/query.json"
	pathEntrustSetName           = "/entrust/group/name/set.json"
	pathEntrustOwnerTransfer     = "/entrust/group/owner/transfer.json"
	pathEntrustJoin              = "/entrust/group/join.json"
	pathEntrustQuit              = "/entrust/group/quit.json"
	pathEntrustKickOut           = "/entrust/group/kickout.json"
	pathEntrustMembers           = "/entrust/group/members/query.json"
	pathEntrustPagingMembers     = "/entrust/group/members/paging/query.json"
	pathEntrustDismiss           = "/entrust/group/dismiss.json"
	pathEntrustSetManagers       = "/entrust/group/managers/set.json"
	pathEntrustRemoveManagers    = "/entrust/group/managers/remove.json"
	pathEntrustQueryManagers     = "/entrust/group/managers/query.json"
	pathEntrustSetMemberInfo     = "/entrust/group/member/info/set.json"
	pathEntrustSetRemarkName     = "/entrust/group/remarkname/set.json"
	pathEntrustQueryRemarkName   = "/entrust/group/remarkname/query.json"
	pathEntrustImport            = "/entrust/group/import.json"
	pathEntrustQueryJoinedGroups = "/entrust/group/joined/query.json"
	pathEntrustFollow            = "/entrust/group/follow.json"
	pathEntrustUnfollow          = "/entrust/group/unfollow.json"
	pathEntrustQueryFollowed     = "/entrust/group/followed/query.json"
	pathEntrustPagingQuery       = "/entrust/group/paging/query.json"
)

// API 群组相关接口
type API interface {
	// Create 创建群组
	Create(userIDs []string, groupID, groupName string) (*CreateResp, error)
	// Join 加入群组
	Join(userIDs []string, groupID, groupName string) (*JoinResp, error)
	// Quit 退出群组
	Quit(userIDs []string, groupID string) (*QuitResp, error)
	// Dismiss 解散群组
	Dismiss(userID, groupID string) (*DismissResp, error)
	// Refresh 刷新群组信息
	Refresh(groupID, groupName string) (*RefreshResp, error)
	// QueryUser 查询用户所加入的群组
	QueryUser(userID string) (*QueryUserResp, error)
	// QueryMembers 查询群组成员
	QueryMembers(groupID string) (*QueryMembersResp, error)

	// EntrustCreate 委托创建群组
	EntrustCreate(req EntrustCreateReq) (*EntrustCreateResp, error)
	// EntrustUpdateProfile 委托更新群组信息
	EntrustUpdateProfile(req EntrustUpdateProfileReq) (*EntrustUpdateProfileResp, error)
	// EntrustQueryProfiles 委托查询群组资料
	EntrustQueryProfiles(groupIDs string) (*EntrustQueryProfilesResp, error)
	// EntrustSetName 委托设置群组名称
	EntrustSetName(groupID, name string) (*EntrustSetNameResp, error)
	// EntrustOwnerTransfer 委托转让群主
	EntrustOwnerTransfer(req EntrustOwnerTransferReq) (*EntrustOwnerTransferResp, error)
	// EntrustJoin 委托加入群组
	EntrustJoin(groupID, userIDs string) (*EntrustJoinResp, error)
	// EntrustQuit 委托退出群组
	EntrustQuit(req EntrustQuitReq) (*EntrustQuitResp, error)
	// EntrustKickOut 委托踢出群组成员
	EntrustKickOut(req EntrustKickOutReq) (*EntrustKickOutResp, error)
	// EntrustMembers 委托查询群组成员
	EntrustMembers(groupID string) (*EntrustMembersResp, error)
	// EntrustPagingMembers 委托分页查询群组成员
	EntrustPagingMembers(req EntrustPagingMembersReq) (*EntrustPagingMembersResp, error)
	// EntrustDismiss 委托解散群组
	EntrustDismiss(groupID string) (*EntrustDismissResp, error)
	// EntrustSetManagers 委托设置管理员
	EntrustSetManagers(groupID, userIDs string) (*EntrustSetManagersResp, error)
	// EntrustRemoveManagers 委托移除管理员
	EntrustRemoveManagers(groupID, userIDs string) (*EntrustRemoveManagersResp, error)
	// EntrustQueryManagers 委托查询管理员
	EntrustQueryManagers(groupID string) (*EntrustQueryManagersResp, error)
	// EntrustSetMemberInfo 委托设置成员信息
	EntrustSetMemberInfo(req EntrustSetMemberInfoReq) (*EntrustSetMemberInfoResp, error)
	// EntrustSetRemarkName 委托设置群备注名
	EntrustSetRemarkName(req EntrustSetRemarkNameReq) (*EntrustSetRemarkNameResp, error)
	// EntrustQueryRemarkName 委托查询群备注名
	EntrustQueryRemarkName(groupID, userID string) (*EntrustQueryRemarkNameResp, error)
	// EntrustImport 委托导入群组
	EntrustImport(req EntrustImportReq) (*EntrustImportResp, error)
	// EntrustQueryJoinedGroups 委托查询用户已加入的群组
	EntrustQueryJoinedGroups(req EntrustQueryJoinedGroupsReq) (*EntrustQueryJoinedGroupsResp, error)
	// EntrustFollow 委托关注群成员
	EntrustFollow(groupID, userID, followUserIDs string) (*EntrustFollowResp, error)
	// EntrustUnfollow 委托取消关注群成员
	EntrustUnfollow(groupID, userID, followUserIDs string) (*EntrustUnfollowResp, error)
	// EntrustQueryFollowed 委托查询已关注的群成员
	EntrustQueryFollowed(groupID, userID string) (*EntrustQueryFollowedResp, error)
	// EntrustPagingQuery 委托分页查询群组
	EntrustPagingQuery(req EntrustPagingQueryReq) (*EntrustPagingQueryResp, error)
}

type api struct {
	client core.Client
}

// NewAPI 创建群组 API 实例
func NewAPI(client core.Client) API {
	return &api{client: client}
}

// ---------- Basic Group ----------

// Create 创建群组
func (a *api) Create(userIDs []string, groupID, groupName string) (*CreateResp, error) {
	resp := &CreateResp{}
	params := map[string]string{
		"userId":    strings.Join(userIDs, ","),
		"groupId":   groupID,
		"groupName": groupName,
	}
	if err := a.client.Post(pathCreate, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Join 加入群组
func (a *api) Join(userIDs []string, groupID, groupName string) (*JoinResp, error) {
	resp := &JoinResp{}
	params := map[string]string{
		"userId":    strings.Join(userIDs, ","),
		"groupId":   groupID,
		"groupName": groupName,
	}
	if err := a.client.Post(pathJoin, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Quit 退出群组
func (a *api) Quit(userIDs []string, groupID string) (*QuitResp, error) {
	resp := &QuitResp{}
	params := map[string]string{
		"userId":  strings.Join(userIDs, ","),
		"groupId": groupID,
	}
	if err := a.client.Post(pathQuit, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Dismiss 解散群组
func (a *api) Dismiss(userID, groupID string) (*DismissResp, error) {
	resp := &DismissResp{}
	params := map[string]string{
		"userId":  userID,
		"groupId": groupID,
	}
	if err := a.client.Post(pathDismiss, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Refresh 刷新群组信息
func (a *api) Refresh(groupID, groupName string) (*RefreshResp, error) {
	resp := &RefreshResp{}
	params := map[string]string{
		"groupId":   groupID,
		"groupName": groupName,
	}
	if err := a.client.Post(pathRefresh, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// QueryUser 查询用户所加入的群组
func (a *api) QueryUser(userID string) (*QueryUserResp, error) {
	resp := &QueryUserResp{}
	params := map[string]string{
		"userId": userID,
	}
	if err := a.client.Post(pathQueryUser, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// QueryMembers 查询群组成员
func (a *api) QueryMembers(groupID string) (*QueryMembersResp, error) {
	resp := &QueryMembersResp{}
	params := map[string]string{
		"groupId": groupID,
	}
	if err := a.client.Post(pathQueryMembers, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// ---------- Entrust Group ----------

// EntrustCreate 委托创建群组
func (a *api) EntrustCreate(req EntrustCreateReq) (*EntrustCreateResp, error) {
	resp := &EntrustCreateResp{}
	params := map[string]string{
		"groupId": req.GroupID,
		"name":    req.Name,
		"owner":   req.Owner,
	}
	if req.UserIDs != "" {
		params["userIds"] = req.UserIDs
	}
	if req.GroupProfile != "" {
		params["groupProfile"] = req.GroupProfile
	}
	if req.GroupExtProfile != "" {
		params["groupExtProfile"] = req.GroupExtProfile
	}
	if req.Permissions != "" {
		params["permissions"] = req.Permissions
	}
	if err := a.client.Post(pathEntrustCreate, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// EntrustUpdateProfile 委托更新群组信息
func (a *api) EntrustUpdateProfile(req EntrustUpdateProfileReq) (*EntrustUpdateProfileResp, error) {
	resp := &EntrustUpdateProfileResp{}
	params := map[string]string{
		"groupId": req.GroupID,
	}
	if req.GroupProfile != "" {
		params["groupProfile"] = req.GroupProfile
	}
	if req.GroupExtProfile != "" {
		params["groupExtProfile"] = req.GroupExtProfile
	}
	if req.Permissions != "" {
		params["permissions"] = req.Permissions
	}
	if err := a.client.Post(pathEntrustUpdateProfile, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// EntrustQueryProfiles 委托查询群组资料
func (a *api) EntrustQueryProfiles(groupIDs string) (*EntrustQueryProfilesResp, error) {
	resp := &EntrustQueryProfilesResp{}
	params := map[string]string{
		"groupIds": groupIDs,
	}
	if err := a.client.Post(pathEntrustQueryProfiles, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// EntrustSetName 委托设置群组名称
func (a *api) EntrustSetName(groupID, name string) (*EntrustSetNameResp, error) {
	resp := &EntrustSetNameResp{}
	params := map[string]string{
		"groupId": groupID,
		"name":    name,
	}
	if err := a.client.Post(pathEntrustSetName, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// EntrustOwnerTransfer 委托转让群主
func (a *api) EntrustOwnerTransfer(req EntrustOwnerTransferReq) (*EntrustOwnerTransferResp, error) {
	resp := &EntrustOwnerTransferResp{}
	params := map[string]string{
		"groupId":  req.GroupID,
		"newOwner": req.NewOwner,
	}
	if req.IsDelBan != nil {
		params["isDelBan"] = strconv.Itoa(*req.IsDelBan)
	}
	if req.IsDelWhite != nil {
		params["isDelWhite"] = strconv.Itoa(*req.IsDelWhite)
	}
	if req.IsDelFollowed != nil {
		params["isDelFollowed"] = strconv.Itoa(*req.IsDelFollowed)
	}
	if req.IsQuit != nil {
		params["isQuit"] = strconv.Itoa(*req.IsQuit)
	}
	if err := a.client.Post(pathEntrustOwnerTransfer, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// EntrustJoin 委托加入群组
func (a *api) EntrustJoin(groupID, userIDs string) (*EntrustJoinResp, error) {
	resp := &EntrustJoinResp{}
	params := map[string]string{
		"groupId": groupID,
		"userIds": userIDs,
	}
	if err := a.client.Post(pathEntrustJoin, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// EntrustQuit 委托退出群组
func (a *api) EntrustQuit(req EntrustQuitReq) (*EntrustQuitResp, error) {
	resp := &EntrustQuitResp{}
	params := map[string]string{
		"groupId": req.GroupID,
		"userIds": req.UserIDs,
	}
	if req.IsDelBan != nil {
		params["isDelBan"] = strconv.Itoa(*req.IsDelBan)
	}
	if req.IsDelWhite != nil {
		params["isDelWhite"] = strconv.Itoa(*req.IsDelWhite)
	}
	if req.IsDelFollowed != nil {
		params["isDelFollowed"] = strconv.Itoa(*req.IsDelFollowed)
	}
	if err := a.client.Post(pathEntrustQuit, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// EntrustKickOut 委托踢出群组成员
func (a *api) EntrustKickOut(req EntrustKickOutReq) (*EntrustKickOutResp, error) {
	resp := &EntrustKickOutResp{}
	params := map[string]string{
		"groupId": req.GroupID,
		"userIds": req.UserIDs,
	}
	if req.IsDelBan != nil {
		params["isDelBan"] = strconv.Itoa(*req.IsDelBan)
	}
	if req.IsDelWhite != nil {
		params["isDelWhite"] = strconv.Itoa(*req.IsDelWhite)
	}
	if req.IsDelFollowed != nil {
		params["isDelFollowed"] = strconv.Itoa(*req.IsDelFollowed)
	}
	if err := a.client.Post(pathEntrustKickOut, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// EntrustMembers 委托查询群组成员
func (a *api) EntrustMembers(groupID string) (*EntrustMembersResp, error) {
	resp := &EntrustMembersResp{}
	params := map[string]string{
		"groupId": groupID,
	}
	if err := a.client.Post(pathEntrustMembers, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// EntrustPagingMembers 委托分页查询群组成员
func (a *api) EntrustPagingMembers(req EntrustPagingMembersReq) (*EntrustPagingMembersResp, error) {
	resp := &EntrustPagingMembersResp{}
	params := map[string]string{
		"groupId": req.GroupID,
	}
	if req.Type != 0 {
		params["type"] = strconv.Itoa(req.Type)
	}
	if req.PageToken != "" {
		params["pageToken"] = req.PageToken
	}
	if req.Size > 0 {
		params["size"] = strconv.Itoa(req.Size)
	}
	if req.Order != 0 {
		params["order"] = strconv.Itoa(req.Order)
	}
	if err := a.client.Post(pathEntrustPagingMembers, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// EntrustDismiss 委托解散群组
func (a *api) EntrustDismiss(groupID string) (*EntrustDismissResp, error) {
	resp := &EntrustDismissResp{}
	params := map[string]string{
		"groupId": groupID,
	}
	if err := a.client.Post(pathEntrustDismiss, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// EntrustSetManagers 委托设置管理员
func (a *api) EntrustSetManagers(groupID, userIDs string) (*EntrustSetManagersResp, error) {
	resp := &EntrustSetManagersResp{}
	params := map[string]string{
		"groupId": groupID,
		"userIds": userIDs,
	}
	if err := a.client.Post(pathEntrustSetManagers, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// EntrustRemoveManagers 委托移除管理员
func (a *api) EntrustRemoveManagers(groupID, userIDs string) (*EntrustRemoveManagersResp, error) {
	resp := &EntrustRemoveManagersResp{}
	params := map[string]string{
		"groupId": groupID,
		"userIds": userIDs,
	}
	if err := a.client.Post(pathEntrustRemoveManagers, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// EntrustQueryManagers 委托查询管理员
func (a *api) EntrustQueryManagers(groupID string) (*EntrustQueryManagersResp, error) {
	resp := &EntrustQueryManagersResp{}
	params := map[string]string{
		"groupId": groupID,
	}
	if err := a.client.Post(pathEntrustQueryManagers, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// EntrustSetMemberInfo 委托设置成员信息
func (a *api) EntrustSetMemberInfo(req EntrustSetMemberInfoReq) (*EntrustSetMemberInfoResp, error) {
	resp := &EntrustSetMemberInfoResp{}
	params := map[string]string{
		"groupId": req.GroupID,
		"userId":  req.UserID,
	}
	if req.Nickname != "" {
		params["nickname"] = req.Nickname
	}
	if req.Extra != "" {
		params["extra"] = req.Extra
	}
	if err := a.client.Post(pathEntrustSetMemberInfo, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// EntrustSetRemarkName 委托设置群备注名
func (a *api) EntrustSetRemarkName(req EntrustSetRemarkNameReq) (*EntrustSetRemarkNameResp, error) {
	resp := &EntrustSetRemarkNameResp{}
	params := map[string]string{
		"userId":     req.UserID,
		"groupId":    req.GroupID,
		"remarkName": req.RemarkName,
	}
	if err := a.client.Post(pathEntrustSetRemarkName, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// EntrustQueryRemarkName 委托查询群备注名
func (a *api) EntrustQueryRemarkName(groupID, userID string) (*EntrustQueryRemarkNameResp, error) {
	resp := &EntrustQueryRemarkNameResp{}
	params := map[string]string{
		"userId":  userID,
		"groupId": groupID,
	}
	if err := a.client.Post(pathEntrustQueryRemarkName, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// EntrustImport 委托导入群组
func (a *api) EntrustImport(req EntrustImportReq) (*EntrustImportResp, error) {
	resp := &EntrustImportResp{}
	params := map[string]string{
		"groupId": req.GroupID,
		"name":    req.Name,
		"owner":   req.Owner,
	}
	if req.GroupProfile != "" {
		params["groupProfile"] = req.GroupProfile
	}
	if req.GroupExtProfile != "" {
		params["groupExtProfile"] = req.GroupExtProfile
	}
	if req.Permissions != "" {
		params["permissions"] = req.Permissions
	}
	if err := a.client.Post(pathEntrustImport, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// EntrustQueryJoinedGroups 委托查询用户已加入的群组
func (a *api) EntrustQueryJoinedGroups(req EntrustQueryJoinedGroupsReq) (*EntrustQueryJoinedGroupsResp, error) {
	resp := &EntrustQueryJoinedGroupsResp{}
	params := map[string]string{
		"userId": req.UserID,
	}
	if req.Role != 0 {
		params["role"] = strconv.Itoa(req.Role)
	}
	if req.PageToken != "" {
		params["pageToken"] = req.PageToken
	}
	if req.Size > 0 {
		params["size"] = strconv.Itoa(req.Size)
	}
	if req.Order != nil {
		params["order"] = strconv.Itoa(*req.Order)
	}
	if err := a.client.Post(pathEntrustQueryJoinedGroups, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// EntrustFollow 委托关注群成员
func (a *api) EntrustFollow(groupID, userID, followUserIDs string) (*EntrustFollowResp, error) {
	resp := &EntrustFollowResp{}
	params := map[string]string{
		"groupId":       groupID,
		"userId":        userID,
		"followUserIds": followUserIDs,
	}
	if err := a.client.Post(pathEntrustFollow, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// EntrustUnfollow 委托取消关注群成员
func (a *api) EntrustUnfollow(groupID, userID, followUserIDs string) (*EntrustUnfollowResp, error) {
	resp := &EntrustUnfollowResp{}
	params := map[string]string{
		"groupId":       groupID,
		"userId":        userID,
		"followUserIds": followUserIDs,
	}
	if err := a.client.Post(pathEntrustUnfollow, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// EntrustQueryFollowed 委托查询已关注的群成员
func (a *api) EntrustQueryFollowed(groupID, userID string) (*EntrustQueryFollowedResp, error) {
	resp := &EntrustQueryFollowedResp{}
	params := map[string]string{
		"groupId": groupID,
		"userId":  userID,
	}
	if err := a.client.Post(pathEntrustQueryFollowed, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// EntrustPagingQuery 委托分页查询群组
func (a *api) EntrustPagingQuery(req EntrustPagingQueryReq) (*EntrustPagingQueryResp, error) {
	resp := &EntrustPagingQueryResp{}
	params := map[string]string{}
	if req.PageToken != "" {
		params["pageToken"] = req.PageToken
	}
	if req.Size > 0 {
		params["size"] = strconv.Itoa(req.Size)
	}
	if req.Order != 0 {
		params["order"] = strconv.Itoa(req.Order)
	}
	if err := a.client.Post(pathEntrustPagingQuery, params, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

