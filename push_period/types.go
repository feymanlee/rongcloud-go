package pushperiod

import "github.com/feymanlee/rongcloud-go/internal/types"

// GetResp 获取推送免打扰时段响应
type GetResp struct {
	types.BaseResp
	Data struct {
		StartTime string `json:"startTime"`
		Period    int    `json:"period"`
		Level     int    `json:"unPushLevel"`
	} `json:"data"`
}
