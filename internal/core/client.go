package core

import (
	"crypto/sha1"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/feymanlee/rongcloud-go/internal/enum"
	"github.com/feymanlee/rongcloud-go/internal/types"
)

// Client HTTP 客户端接口
type Client interface {
	// Post 发送 form-encoded 请求 (V1 APIs)
	Post(path string, params map[string]string, resp any) error
	// PostJSON 发送 JSON 请求 (V2 APIs)
	PostJSON(path string, body any, resp any) error
}

// Region 区域配置
type Region struct {
	PrimaryDomain string
	BackupDomain  string
}

var (
	RegionBeijing      = Region{"https://api.rong-api.com", "https://api-b.rong-api.com"}
	RegionSingapore    = Region{"https://api.sg-light-api.com", "https://api-b.sg-light-api.com"}
	RegionSingaporeB   = Region{"https://api.sg-b-light-api.com", "https://api-b.sg-b-light-api.com"}
	RegionNorthAmerica = Region{"https://api.us-light-api.com", "https://api-b.us-light-api.com"}
	RegionSAU          = Region{"https://api.sau-light-api.com", "https://api-b.sau-light-api.com"}
)

const (
	defaultTimeout           = 10 * time.Second
	defaultChangeURIDuration = 30 * time.Second
)

// Options 客户端选项
type Options struct {
	AppKey    string
	AppSecret string
	Region    Region
	Timeout   time.Duration
}

type client struct {
	appKey    string
	appSecret string
	resty     *resty.Client

	primaryDomain string
	backupDomain  string
	currentURI    string

	uriLock           sync.Mutex
	lastChangeURITime time.Time
	changeURIDuration time.Duration
}

// NewClient 创建 HTTP 客户端
func NewClient(opt *Options) Client {
	timeout := opt.Timeout
	if timeout == 0 {
		timeout = defaultTimeout
	}

	c := &client{
		appKey:            opt.AppKey,
		appSecret:         opt.AppSecret,
		primaryDomain:     opt.Region.PrimaryDomain,
		backupDomain:      opt.Region.BackupDomain,
		currentURI:        opt.Region.PrimaryDomain,
		changeURIDuration: defaultChangeURIDuration,
	}

	c.resty = resty.New().
		SetTimeout(timeout)

	return c
}

// Post 发送 form-encoded 请求
func (c *client) Post(path string, params map[string]string, resp any) error {
	url := c.currentURI + path

	r := c.resty.R().
		SetFormData(params).
		SetResult(resp)

	c.fillHeader(r, "application/x-www-form-urlencoded")

	response, err := r.Post(url)
	if err != nil {
		c.changeURI()
		return err
	}

	if response.StatusCode() >= 500 {
		c.changeURI()
	}

	return c.checkResp(resp)
}

// PostJSON 发送 JSON 请求
func (c *client) PostJSON(path string, body any, resp any) error {
	url := c.currentURI + path

	r := c.resty.R().
		SetBody(body).
		SetResult(resp)

	c.fillHeader(r, "application/json")

	response, err := r.Post(url)
	if err != nil {
		c.changeURI()
		return err
	}

	if response.StatusCode() >= 500 {
		c.changeURI()
	}

	return c.checkResp(resp)
}

// fillHeader 填充认证头
func (c *client) fillHeader(r *resty.Request, contentType string) {
	nonce := strconv.Itoa(rand.Int())
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	h := sha1.New()
	_, _ = io.WriteString(h, c.appSecret+nonce+timestamp)
	signature := fmt.Sprintf("%x", h.Sum(nil))

	r.SetHeaders(map[string]string{
		"Content-Type": contentType,
		"App-Key":      c.appKey,
		"Nonce":        nonce,
		"Timestamp":    timestamp,
		"Signature":    signature,
	})
}

// checkResp 检查响应
func (c *client) checkResp(resp any) error {
	if r, ok := resp.(types.BaseRespInterface); ok {
		if r.GetCode() != enum.SuccessCode {
			return NewError(r.GetCode(), r.GetErrorMessage())
		}
	}
	return nil
}

// changeURI 切换域名
func (c *client) changeURI() {
	now := time.Now()
	c.uriLock.Lock()
	defer c.uriLock.Unlock()

	if now.Sub(c.lastChangeURITime) >= c.changeURIDuration {
		if c.currentURI == c.primaryDomain {
			c.currentURI = c.backupDomain
		} else {
			c.currentURI = c.primaryDomain
		}
		c.lastChangeURITime = now
	}
}
