package callback

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestHandler_ServeHTTP_MessageRoute(t *testing.T) {
	appSecret := "test-secret"
	nonce := "123456"
	timestamp := "1234567890"
	validSig := sha1Sum(appSecret + nonce + timestamp)

	tests := []struct {
		name           string
		path           string
		body           string
		contentType    string
		signature      string
		handlerConfig  HandlerConfig
		expectedStatus int
		expectedBody   string
	}{
		{
			name:        "valid message route callback (form-urlencoded)",
			path:        "/message/sync",
			contentType: "application/x-www-form-urlencoded",
			signature:   validSig,
			body: url.Values{
				"fromUserId":   []string{"user1"},
				"targetId":     []string{"user2"},
				"objectName":   []string{"RC:TxtMsg"},
				"content":      []string{"hello"},
				"channelType":  []string{"PERSON"},
				"msgTimestamp": []string{"1234567890"},
				"msgUID":       []string{"msg-uid-123"},
			}.Encode(),
			handlerConfig: HandlerConfig{
				OnMessageRoute: func(w ResponseWriter, msg MessageRouteCallback) error {
					if msg.FromUserId != "user1" {
						t.Errorf("expected FromUserId=user1, got %s", msg.FromUserId)
					}
					// 不调用 WriteResponse，使用默认响应
					return nil
				},
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "OK",
		},
		{
			name:        "custom response",
			path:        "/message/sync",
			contentType: "application/x-www-form-urlencoded",
			signature:   validSig,
			body:        url.Values{"fromUserId": []string{"user1"}}.Encode(),
			handlerConfig: HandlerConfig{
				OnMessageRoute: func(w ResponseWriter, msg MessageRouteCallback) error {
					w.WriteResponse(201, "Created")
					return nil
				},
			},
			expectedStatus: 201,
			expectedBody:   "Created",
		},
		{
			name:        "custom path",
			path:        "/my/custom/path",
			contentType: "application/x-www-form-urlencoded",
			signature:   validSig,
			body:        url.Values{"fromUserId": []string{"user1"}}.Encode(),
			handlerConfig: HandlerConfig{
				MessageRoutePath: "/my/custom/path",
				OnMessageRoute: func(w ResponseWriter, msg MessageRouteCallback) error {
					return nil
				},
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "OK",
		},
		{
			name:        "invalid signature",
			path:        "/message/sync",
			contentType: "application/x-www-form-urlencoded",
			signature:   "invalid-sig",
			body:        "fromUserId=user1",
			handlerConfig: HandlerConfig{
			},
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:        "handler returns error",
			path:        "/message/sync",
			contentType: "application/x-www-form-urlencoded",
			signature:   validSig,
			body:        "fromUserId=user1",
			handlerConfig: HandlerConfig{
				OnMessageRoute: func(w ResponseWriter, msg MessageRouteCallback) error {
					return errors.New("handler error")
				},
			},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:        "handler returns error but already wrote response",
			path:        "/message/sync",
			contentType: "application/x-www-form-urlencoded",
			signature:   validSig,
			body:        "fromUserId=user1",
			handlerConfig: HandlerConfig{
				OnMessageRoute: func(w ResponseWriter, msg MessageRouteCallback) error {
					w.WriteResponse(400, "Bad Request")
					return errors.New("some error") // 这个错误会被忽略，因为已经写了响应
				},
			},
			expectedStatus: 400,
			expectedBody:   "Bad Request",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			urlPath := tt.path + "?nonce=" + nonce + "&timestamp=" + timestamp + "&signature=" + tt.signature
			req := httptest.NewRequest(http.MethodPost, urlPath, strings.NewReader(tt.body))
			req.Header.Set("Content-Type", tt.contentType)

			rec := httptest.NewRecorder()
			handler := NewHandler(appSecret, tt.handlerConfig)
			handler.ServeHTTP(rec, req)

			if rec.Code != tt.expectedStatus {
				t.Errorf("status = %d, want %d", rec.Code, tt.expectedStatus)
			}
			if tt.expectedBody != "" && rec.Body.String() != tt.expectedBody {
				t.Errorf("body = %q, want %q", rec.Body.String(), tt.expectedBody)
			}
		})
	}
}

func TestHandler_ServeHTTP_UserOnlineStatus(t *testing.T) {
	appSecret := "test-secret"
	nonce := "123456"
	timestamp := "1234567890"
	validSig := sha1Sum(appSecret + nonce + timestamp)

	var receivedCallbacks []UserOnlineStatusCallback

	config := HandlerConfig{
		OnUserOnlineStatus: func(w ResponseWriter, statuses []UserOnlineStatusCallback) error {
			receivedCallbacks = statuses
			return nil
		},
	}

	body := []UserOnlineStatusCallback{
		{
			UserID:    "user123",
			Status:    "1",
			OS:        "iOS",
			Time:      1234567890,
			ClientIP:  "192.168.1.1:8080",
			SessionID: "session-abc",
		},
	}
	bodyJSON, _ := json.Marshal(body)

	url := "/user/onlinestatus?nonce=" + nonce + "&timestamp=" + timestamp + "&signature=" + validSig
	req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(bodyJSON))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	handler := NewHandler(appSecret, config)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	if len(receivedCallbacks) != 1 {
		t.Fatalf("expected 1 callback, got %d", len(receivedCallbacks))
	}

	if receivedCallbacks[0].UserID != "user123" {
		t.Errorf("expected UserID=user123, got %s", receivedCallbacks[0].UserID)
	}
}

func TestHandler_ServeHTTP_AuditResult(t *testing.T) {
	appSecret := "test-secret"
	nonce := "123456"
	timestamp := "1234567890"
	validSig := sha1Sum(appSecret + nonce + timestamp)

	config := HandlerConfig{
		OnAuditResult: func(w ResponseWriter, result AuditResultCallback) error {
			// 自定义响应
			w.WriteResponse(200, `{"status":"received"}`)
			return nil
		},
	}

	body := AuditResultCallback{
		Result:          10000,
		Content:         `{"appKey":"test","fromUserId":"user1"}`,
		MsgUID:          "msg-123",
		ServiceProvider: "shumei",
		ResultDetail:    "{}",
	}
	bodyJSON, _ := json.Marshal(body)

	url := "/moderation/audit-result?nonce=" + nonce + "&timestamp=" + timestamp + "&signature=" + validSig
	req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(bodyJSON))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	handler := NewHandler(appSecret, config)
	handler.ServeHTTP(rec, req)

	if rec.Code != 200 {
		t.Errorf("status = %d, want 200", rec.Code)
	}

	if rec.Body.String() != `{"status":"received"}` {
		t.Errorf("body = %q, want %q", rec.Body.String(), `{"status":"received"}`)
	}
}

func TestHandler_ServeHTTP_NotFound(t *testing.T) {
	appSecret := "test-secret"
	nonce := "123456"
	timestamp := "1234567890"
	validSig := sha1Sum(appSecret + nonce + timestamp)

	config := HandlerConfig{
		// 没有配置任何处理器，使用默认路径
	}

	url := "/unknown/path?nonce=" + nonce + "&timestamp=" + timestamp + "&signature=" + validSig
	req := httptest.NewRequest(http.MethodPost, url, nil)

	rec := httptest.NewRecorder()
	handler := NewHandler(appSecret, config)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusNotFound)
	}
}

func TestHandler_DefaultPaths(t *testing.T) {
	// 测试默认路径是否正确设置
	config := HandlerConfig{}
	handler := NewHandler("test-secret", config)

	if handler.config.MessageRoutePath != "/message/sync" {
		t.Errorf("default MessageRoutePath = %s, want /message/sync", handler.config.MessageRoutePath)
	}
	if handler.config.UserOnlineStatusPath != "/user/onlinestatus" {
		t.Errorf("default UserOnlineStatusPath = %s, want /user/onlinestatus", handler.config.UserOnlineStatusPath)
	}
	if handler.config.AuditResultPath != "/moderation/audit-result" {
		t.Errorf("default AuditResultPath = %s, want /moderation/audit-result", handler.config.AuditResultPath)
	}
}
