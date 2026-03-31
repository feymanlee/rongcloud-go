package embeddedconsole

// 控制台嵌入 page_code 常量
const (
	// PageCodeUserManage 用户管理页面。
	PageCodeUserManage = "user_manage"
	// PageCodeGroupManage 群组管理页面。
	PageCodeGroupManage = "group_manage"
	// PageCodeBlockedUserManage 封禁用户管理页面。
	PageCodeBlockedUserManage = "blocked_user_manage"
	// PageCodeChatroomManage 聊天室管理页面。
	PageCodeChatroomManage = "chatroom_manage"
	// PageCodeOperationLog 操作日志页面。
	PageCodeOperationLog = "operation_log"
	// PageCodeRecordingManage 录制文件管理页面。
	PageCodeRecordingManage = "recording_manage"
	// PageCodeCustomerServiceManage 客服管理页面。
	PageCodeCustomerServiceManage = "customer_service_manage"
	// PageCodeSensitiveWordSettings 敏感词设置页面。
	PageCodeSensitiveWordSettings = "sensitive_word_settings"
)

// GetAccessTokenReq 获取嵌入控制台 access token 请求
type GetAccessTokenReq struct {
	AccessKey  string `json:"access_key"`
	PageCode   string `json:"page_code"`
	UsefulLife int64  `json:"useful_life"`
}

// GetAccessTokenResp 获取嵌入控制台 access token 响应
type GetAccessTokenResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}
