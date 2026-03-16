package testutil

import (
	"encoding/json"
	"sync"
)

// Call 记录一次 API 调用
type Call struct {
	Method string            // "Post" or "PostJSON"
	Path   string            // 请求路径
	Params map[string]string // Post 的 form 参数
	Body   any               // PostJSON 的请求体
}

// MockClient 模拟 core.Client，记录调用并返回预设响应
type MockClient struct {
	mu    sync.Mutex
	Calls []Call

	// PostFunc 自定义 Post 行为，为 nil 时默认返回 {"code":200}
	PostFunc func(path string, params map[string]string, resp any) error
	// PostJSONFunc 自定义 PostJSON 行为，为 nil 时默认返回 {"code":200}
	PostJSONFunc func(path string, body any, resp any) error
}

// NewMockClient 创建 MockClient
func NewMockClient() *MockClient {
	return &MockClient{}
}

// Post 实现 core.Client 接口
func (m *MockClient) Post(path string, params map[string]string, resp any) error {
	m.mu.Lock()
	m.Calls = append(m.Calls, Call{Method: "Post", Path: path, Params: params})
	m.mu.Unlock()

	if m.PostFunc != nil {
		return m.PostFunc(path, params, resp)
	}
	return defaultResp(resp)
}

// PostJSON 实现 core.Client 接口
func (m *MockClient) PostJSON(path string, body any, resp any) error {
	m.mu.Lock()
	m.Calls = append(m.Calls, Call{Method: "PostJSON", Path: path, Body: body})
	m.mu.Unlock()

	if m.PostJSONFunc != nil {
		return m.PostJSONFunc(path, body, resp)
	}
	return defaultResp(resp)
}

// LastCall 返回最近一次调用
func (m *MockClient) LastCall() Call {
	m.mu.Lock()
	defer m.mu.Unlock()
	if len(m.Calls) == 0 {
		return Call{}
	}
	return m.Calls[len(m.Calls)-1]
}

// Reset 清除所有调用记录
func (m *MockClient) Reset() {
	m.mu.Lock()
	m.Calls = nil
	m.mu.Unlock()
}

func defaultResp(resp any) error {
	data := []byte(`{"code":200}`)
	return json.Unmarshal(data, resp)
}
