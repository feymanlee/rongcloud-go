package embeddedconsole

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/feymanlee/rongcloud-go/internal/core"
)

const (
	defaultEndpoint      = "https://embed-console.rongcloud.cn/embed/access/token"
	defaultTimeout       = 10 * time.Second
	successCode          = 10000
	minUsefulLifeSeconds = 1
	maxUsefulLifeSeconds = 31536000
)

// API 控制台嵌入相关接口
type API interface {
	// GetAccessToken 获取控制台嵌入 access token
	GetAccessToken(pageCode string, usefulLife int64) (*GetAccessTokenResp, error)
}

type api struct {
	accessKey string
	endpoint  string
	client    *resty.Client
}

// NewAPI 创建控制台嵌入 API 实例
func NewAPI(accessKey string, timeout time.Duration) API {
	if timeout == 0 {
		timeout = defaultTimeout
	}

	return &api{
		accessKey: accessKey,
		endpoint:  defaultEndpoint,
		client:    resty.New().SetTimeout(timeout),
	}
}

// GetAccessToken 获取控制台嵌入 access token
func (a *api) GetAccessToken(pageCode string, usefulLife int64) (*GetAccessTokenResp, error) {
	if a.accessKey == "" {
		return nil, errors.New("rongcloud: accessKey is required")
	}
	if pageCode == "" {
		return nil, errors.New("rongcloud: pageCode is required")
	}
	if usefulLife < minUsefulLifeSeconds || usefulLife > maxUsefulLifeSeconds {
		return nil, fmt.Errorf("rongcloud: usefulLife must be between %d and %d", minUsefulLifeSeconds, maxUsefulLifeSeconds)
	}

	req := &GetAccessTokenReq{
		AccessKey:  a.accessKey,
		PageCode:   pageCode,
		UsefulLife: usefulLife,
	}
	resp := &GetAccessTokenResp{}

	httpResp, err := a.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(req).
		Post(a.endpoint)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(httpResp.Body(), resp); err != nil {
		if httpResp.StatusCode() < 200 || httpResp.StatusCode() >= 300 {
			return nil, fmt.Errorf("rongcloud: unexpected http status %d", httpResp.StatusCode())
		}
		return nil, err
	}

	if httpResp.StatusCode() < 200 || httpResp.StatusCode() >= 300 {
		if resp.Code != 0 {
			return nil, core.NewError(resp.Code, "")
		}
		return nil, fmt.Errorf("rongcloud: unexpected http status %d", httpResp.StatusCode())
	}

	if resp.Code != successCode {
		return nil, core.NewError(resp.Code, "")
	}

	return resp, nil
}
