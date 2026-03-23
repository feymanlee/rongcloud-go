package callback

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVerifyCallback(t *testing.T) {
	appSecret := "test-secret"
	nonce := "123456"
	timestamp := "1234567890"

	// 计算正确的签名
	expectedSig := sha1Sum(appSecret + nonce + timestamp)

	tests := []struct {
		name      string
		appSecret string
		nonce     string
		timestamp string
		signature string
		want      bool
	}{
		{
			name:      "valid signature",
			appSecret: appSecret,
			nonce:     nonce,
			timestamp: timestamp,
			signature: expectedSig,
			want:      true,
		},
		{
			name:      "invalid signature",
			appSecret: appSecret,
			nonce:     nonce,
			timestamp: timestamp,
			signature: "invalid-signature",
			want:      false,
		},
		{
			name:      "wrong secret",
			appSecret: "wrong-secret",
			nonce:     nonce,
			timestamp: timestamp,
			signature: expectedSig,
			want:      false,
		},
		{
			name:      "empty signature",
			appSecret: appSecret,
			nonce:     nonce,
			timestamp: timestamp,
			signature: "",
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := VerifyCallback(tt.appSecret, tt.nonce, tt.timestamp, tt.signature)
			if got != tt.want {
				t.Errorf("VerifyCallback() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractParams(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() *http.Request
		expected CallbackParams
	}{
		{
			name: "params from query",
			setup: func() *http.Request {
				req := httptest.NewRequest(http.MethodPost, "/callback?appKey=app-key&nonce=nonce-123&timestamp=123456&signature=abc123", nil)
				return req
			},
			expected: CallbackParams{
				AppKey:    "app-key",
				Nonce:     "nonce-123",
				Timestamp: "123456",
				Signature: "abc123",
			},
		},
		{
			name: "params from header",
			setup: func() *http.Request {
				req := httptest.NewRequest(http.MethodPost, "/callback", nil)
				req.Header.Set("App-Key", "header-app-key")
				req.Header.Set("Nonce", "header-nonce")
				req.Header.Set("Timestamp", "789012")
				req.Header.Set("Signature", "header-sig")
				return req
			},
			expected: CallbackParams{
				AppKey:    "header-app-key",
				Nonce:     "header-nonce",
				Timestamp: "789012",
				Signature: "header-sig",
			},
		},
		{
			name: "query takes precedence over header",
			setup: func() *http.Request {
				req := httptest.NewRequest(http.MethodPost, "/callback?appKey=query-key&signature=query-sig", nil)
				req.Header.Set("App-Key", "header-key")
				req.Header.Set("Signature", "header-sig")
				return req
			},
			expected: CallbackParams{
				AppKey:    "query-key",
				Signature: "query-sig",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := tt.setup()
			got := ExtractParams(req)
			if got != tt.expected {
				t.Errorf("ExtractParams() = %+v, want %+v", got, tt.expected)
			}
		})
	}
}

func TestVerifyRequest(t *testing.T) {
	appSecret := "test-secret"
	nonce := "123456"
	timestamp := "1234567890"
	validSig := sha1Sum(appSecret + nonce + timestamp)

	tests := []struct {
		name      string
		setup     func() *http.Request
		appSecret string
		want      bool
	}{
		{
			name: "valid request",
			setup: func() *http.Request {
				url := fmt.Sprintf("/callback?nonce=%s&timestamp=%s&signature=%s", nonce, timestamp, validSig)
				return httptest.NewRequest(http.MethodPost, url, nil)
			},
			appSecret: appSecret,
			want:      true,
		},
		{
			name: "missing signature",
			setup: func() *http.Request {
				return httptest.NewRequest(http.MethodPost, "/callback?nonce=123&timestamp=456", nil)
			},
			appSecret: appSecret,
			want:      false,
		},
		{
			name: "invalid signature",
			setup: func() *http.Request {
				return httptest.NewRequest(http.MethodPost, "/callback?nonce=123&timestamp=456&signature=invalid", nil)
			},
			appSecret: appSecret,
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := tt.setup()
			got := VerifyRequest(req, tt.appSecret)
			if got != tt.want {
				t.Errorf("VerifyRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidIP(t *testing.T) {
	tests := []struct {
		name string
		ip   string
		want bool
	}{
		{
			name: "valid domestic IP",
			ip:   "39.105.128.42",
			want: true,
		},
		{
			name: "valid overseas IP",
			ip:   "52.221.93.74",
			want: true,
		},
		{
			name: "invalid IP",
			ip:   "192.168.1.1",
			want: false,
		},
		{
			name: "empty IP",
			ip:   "",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidIP(tt.ip)
			if got != tt.want {
				t.Errorf("IsValidIP(%q) = %v, want %v", tt.ip, got, tt.want)
			}
		})
	}
}
