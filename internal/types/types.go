package types

// BaseResp 基础响应
type BaseResp struct {
	Code         int    `json:"code"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// BaseRespInterface 基础响应接口
type BaseRespInterface interface {
	GetCode() int
	GetErrorMessage() string
}

func (r *BaseResp) GetCode() int {
	return r.Code
}

func (r *BaseResp) GetErrorMessage() string {
	return r.ErrorMessage
}
