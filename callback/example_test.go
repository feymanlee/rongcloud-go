package callback_test

import (
	"fmt"
	"log"
	"net/http"

	"github.com/feymanlee/rongcloud-go/callback"
)

// ExampleVerifyCallback 演示如何验证回调签名
func ExampleVerifyCallback() {
	appSecret := "your-app-secret"
	nonce := "123456"
	timestamp := "1234567890"
	signature := "a1b2c3d4e5f6" // 融云提供的签名

	if callback.VerifyCallback(appSecret, nonce, timestamp, signature) {
		fmt.Println("签名验证通过")
	} else {
		fmt.Println("签名验证失败")
	}
}

// ExampleHandler 演示如何使用 Handler 处理回调
func ExampleHandler() {
	handler := callback.NewHandler("your-app-secret", callback.HandlerConfig{
		OnMessageRoute: func(w callback.ResponseWriter, msg callback.MessageRouteCallback) error {
			log.Printf("收到消息: from=%s, to=%s, type=%s\n",
				msg.FromUserId, msg.ToUserId, msg.ObjectName)
			// 返回 nil 会自动写入 200 OK
			// 如需自定义响应：w.WriteResponse(200, "自定义响应")
			return nil
		},
		OnUserOnlineStatus: func(w callback.ResponseWriter, statuses []callback.UserOnlineStatusCallback) error {
			for _, status := range statuses {
				log.Printf("用户状态变更: user=%s, status=%s, os=%s\n",
					status.UserID, status.Status, status.OS)
			}
			return nil
		},
		OnAuditResult: func(w callback.ResponseWriter, result callback.AuditResultCallback) error {
			if result.Result == 10000 {
				log.Printf("消息审核通过: msgUID=%s\n", result.MsgUID)
			} else {
				log.Printf("消息审核未通过: msgUID=%s\n", result.MsgUID)
			}
			return nil
		},
	})

	http.Handle("/rongcloud/callback", handler)
}

// ExampleIsValidIP 演示如何验证 IP 白名单
func ExampleIsValidIP() {
	clientIP := "39.105.128.42"

	if callback.IsValidIP(clientIP) {
		fmt.Println("IP 在融云白名单中")
	} else {
		fmt.Println("IP 不在白名单中，拒绝访问")
	}
}

// ExampleExtractParams 演示如何从请求中提取参数
func ExampleExtractParams() {
	// 模拟 HTTP 请求
	req, _ := http.NewRequest(http.MethodPost,
		"/callback?appKey=app-key&nonce=123&timestamp=456&signature=abc",
		nil)

	params := callback.ExtractParams(req)
	fmt.Printf("AppKey: %s, Nonce: %s\n", params.AppKey, params.Nonce)
}
