package userblock

import "github.com/feymanlee/rongcloud-go/internal/types"

// BlacklistResp 黑名单查询响应
type BlacklistResp struct {
	types.BaseResp
	Users []string `json:"users"`
}

// WhitelistResp 白名单查询响应
type WhitelistResp struct {
	types.BaseResp
	Users []string `json:"users"`
}
