package callback

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"strings"
)

// VerifyCallback 验证回调签名
// 签名计算方式：SHA1(AppSecret + Nonce + Timestamp)
func VerifyCallback(appSecret, nonce, timestamp, signature string) bool {
	expected := sha1Sum(appSecret + nonce + timestamp)
	return strings.EqualFold(expected, signature)
}

// sha1Sum 计算 SHA1 哈希并返回十六进制字符串
func sha1Sum(data string) string {
	h := sha1.New()
	h.Write([]byte(data))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// ExtractParams 从 HTTP 请求中提取回调参数
// 优先从 Query 参数获取，如果不在 Query 中则尝试 Header（用于审核结果回调）
func ExtractParams(r *http.Request) CallbackParams {
	params := CallbackParams{}

	// 优先从 Query 参数获取
	params.AppKey = r.URL.Query().Get("appKey")
	params.Nonce = r.URL.Query().Get("nonce")
	params.Timestamp = r.URL.Query().Get("timestamp")
	params.Signature = r.URL.Query().Get("signature")

	// 如果不在 Query 中，尝试从 Header 获取（审核结果回调）
	if params.Signature == "" {
		params.Signature = r.Header.Get("Signature")
	}
	if params.Nonce == "" {
		params.Nonce = r.Header.Get("Nonce")
	}
	if params.Timestamp == "" {
		params.Timestamp = r.Header.Get("Timestamp")
	}
	if params.AppKey == "" {
		params.AppKey = r.Header.Get("App-Key")
	}

	return params
}

// VerifyRequest 便捷方法：从请求中提取参数并验证签名
func VerifyRequest(r *http.Request, appSecret string) bool {
	params := ExtractParams(r)
	if params.Signature == "" || params.Nonce == "" || params.Timestamp == "" {
		return false
	}
	return VerifyCallback(appSecret, params.Nonce, params.Timestamp, params.Signature)
}

// IPWhitelist 融云回调服务器 IP 白名单（国内数据中心）
var IPWhitelist = []string{
	"39.105.128.42",
	"39.105.147.30",
	"123.56.88.42",
	"182.92.215.38",
	"182.92.84.148",
	"39.106.150.151",
	"39.107.75.101",
	"101.201.34.95",
	"39.106.2.63",
	"101.200.62.251",
	"47.93.57.144",
}

// IPWhitelistOverseas 融云回调服务器 IP 白名单（海外数据中心）
var IPWhitelistOverseas = []string{
	// 新加坡
	"52.221.93.74",
	"8.219.168.45",
	"8.219.93.148",
	"8.219.215.35",
	"8.219.43.97",
	"47.245.124.194",
	"8.222.167.17",
	"47.236.149.1",
	"8.219.217.193",
	"8.222.202.67",
	"43.156.138.254",
	"43.163.81.196",
	"43.156.239.53",
	// 北美
	"52.41.206.152",
	// 沙特
	"8.213.17.80",
	"8.213.16.171",
	"8.213.28.96",
	"101.46.58.181",
	"80.238.230.175",
	"101.46.53.241",
}

// IsValidIP 检查 IP 是否在融云回调白名单中
func IsValidIP(ip string) bool {
	for _, validIP := range IPWhitelist {
		if validIP == ip {
			return true
		}
	}
	for _, validIP := range IPWhitelistOverseas {
		if validIP == ip {
			return true
		}
	}
	return false
}
