# 融云回调处理模块

本模块提供融云 IM 回调服务的处理功能，包括签名验证、回调消息解析和 HTTP Handler 封装。

## 功能特性

- **签名验证**: 支持融云回调请求的 SHA1 签名验证
- **回调类型支持**:
  - 消息路由回调（全量消息路由）- `application/x-www-form-urlencoded`
  - 用户在线状态回调 - `application/json` (数组)
  - 消息审核结果回调 - `application/json`
  - 聊天室状态回调 - `application/json`
  - 聊天室 KV 属性回调 - `application/json`
  - 用户注销/激活回调 - `application/json`
- **自定义回调路径**: 支持配置自定义回调 URL 路径
- **自定义响应**: 支持在回调处理器中自定义 HTTP 响应
- **IP 白名单**: 内置融云回调服务器 IP 白名单
- **HTTP Handler**: 提供开箱即用的 HTTP 处理器

## 快速开始

### 基础签名验证

```go
package main

import (
    "net/http"
    "github.com/feymanlee/rongcloud-go/callback"
)

func handleCallback(w http.ResponseWriter, r *http.Request) {
    // 验证签名
    if !callback.VerifyRequest(r, "your-app-secret") {
        http.Error(w, "Invalid signature", http.StatusUnauthorized)
        return
    }

    // 处理回调...
    w.Write([]byte("OK"))
}
```

### 使用 Handler 包装器

```go
package main

import (
    "log"
    "net/http"
    "github.com/feymanlee/rongcloud-go/callback"
)

func main() {
    handler := callback.NewHandler(callback.HandlerConfig{
        AppSecret: "your-app-secret",
        OnMessageRoute: func(w callback.ResponseWriter, msg callback.MessageRouteCallback) error {
            log.Printf("收到消息: from=%s, to=%s, type=%s",
                msg.FromUserId, msg.ToUserId, msg.ObjectName)
            return nil // 返回 nil 自动响应 200 OK
        },
        OnUserOnlineStatus: func(w callback.ResponseWriter, statuses []callback.UserOnlineStatusCallback) error {
            for _, status := range statuses {
                log.Printf("用户状态: user=%s, status=%s, os=%s",
                    status.UserID, status.Status, status.OS)
            }
            return nil
        },
    })

    // 将同一个 Handler 注册到多个路径（与融云后台配置的路径对应）
    http.Handle("/message/sync", handler)
    http.Handle("/user/onlinestatus", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

**融云后台配置说明：**

上面的代码将 Handler 注册到了 `/message/sync` 和 `/user/onlinestatus` 路径。在融云控制台配置回调地址时，需要填写完整的 URL：

```
https://your-domain.com/message/sync
https://your-domain.com/user/onlinestatus
```

例如：
- 如果你的服务部署在 `https://api.example.com`
- 代码中注册的路径是 `/message/sync`
- 那么融云后台消息路由回调应该配置为：`https://api.example.com/message/sync`

**路径匹配说明：**

Handler 通过请求路径（`r.URL.Path`）来识别回调类型。默认路径如下：
- 消息路由: `/message/sync`
- 用户在线状态: `/user/onlinestatus`
- 审核结果: `/moderation/audit-result`
- 聊天室状态: `/chatroom/status`
- 聊天室 KV: `/chatroom/kv`
- 用户注销/激活: `/user/deactivation`

**配置方式：**

推荐做法：创建一个 Handler 实例，配置所有需要的回调处理器，然后将其注册到多个路径：

```go
// 创建一个 Handler，配置所有回调处理器
handler := callback.NewHandler(callback.HandlerConfig{
    AppSecret: "your-app-secret",
    OnMessageRoute: func(w callback.ResponseWriter, msg callback.MessageRouteCallback) error {
        // 处理消息
        return nil
    },
    OnUserOnlineStatus: func(w callback.ResponseWriter, statuses []callback.UserOnlineStatusCallback) error {
        // 处理用户状态
        return nil
    },
})

// 将同一个 Handler 注册到不同路径
http.Handle("/message/sync", handler)
http.Handle("/user/onlinestatus", handler)
```

融云后台配置：
- 消息路由回调地址：`https://your-domain.com/message/sync`
- 用户在线状态回调地址：`https://your-domain.com/user/onlinestatus`

### 自定义回调路径

融云控制台允许自定义回调地址，可在配置中指定：

```go
handler := callback.NewHandler(callback.HandlerConfig{
    AppSecret:            "your-app-secret",
    MessageRoutePath:     "/api/callback/message",      // 自定义消息路由路径
    UserOnlineStatusPath: "/api/callback/online",       // 自定义在线状态路径
    AuditResultPath:      "/api/callback/audit",        // 自定义审核结果路径
    OnMessageRoute: func(w callback.ResponseWriter, msg callback.MessageRouteCallback) error {
        // 处理消息
        return nil
    },
})
```

**默认路径：**
- 消息路由: `/message/sync`
- 用户在线状态: `/user/onlinestatus`
- 审核结果: `/moderation/audit-result`
- 聊天室状态: `/chatroom/status`
- 聊天室 KV: `/chatroom/kv`
- 用户注销/激活: `/user/deactivation`

### 自定义响应

回调处理器可以通过 `ResponseWriter` 自定义 HTTP 响应：

```go
handler := callback.NewHandler(callback.HandlerConfig{
    AppSecret: "your-app-secret",
    OnMessageRoute: func(w callback.ResponseWriter, msg callback.MessageRouteCallback) error {
        // 检查消息内容
        if msg.SensitiveType > 0 {
            // 拒绝敏感消息，返回 403
            w.WriteResponse(http.StatusForbidden, "Message contains sensitive content")
            return nil
        }

        // 返回 nil 将自动响应 200 OK
        return nil
    },
    OnAuditResult: func(w callback.ResponseWriter, result callback.AuditResultCallback) error {
        // 自定义 JSON 响应
        w.Header().Set("Content-Type", "application/json")
        w.WriteResponse(200, `{"status":"processed"}`)
        return nil
    },
})
```

**响应规则：**
- 调用 `w.WriteResponse(code, body)` → 使用自定义响应
- 返回 `nil` 且未调用 WriteResponse → 自动返回 `200 OK`
- 返回 `error` 且未调用 WriteResponse → 返回 `500 Internal Server Error`

## 回调类型详解

### 消息路由回调 (MessageRouteCallback)

**请求格式**: `application/x-www-form-urlencoded`
**发送方式**: 每条消息单独发送一个 POST 请求（非批量）

```go
type MessageRouteCallback struct {
    FromUserId     string   // 发送用户 ID
    ToUserId       string   // 目标 ID（targetId）
    ObjectName     string   // 消息类型，如 RC:TxtMsg
    Content        string   // 消息内容
    ChannelType    string   // 会话类型：PERSON/GROUP/TEMPGROUP/ULTRAGROUP
    MsgTimestamp   string   // 服务端时间（毫秒）
    MsgUID         string   // 消息唯一标识
    OriginalMsgUID string   // 原始消息 ID（超级群修改消息时使用）
    SensitiveType  int      // 敏感词类型：0=无，1=屏蔽，2=替换
    Source         string   // 发送源头：iOS/Android/HarmonyOS/Websocket 等
    BusChannel     string   // 超级群频道 ID
    GroupUserIds   []string // 群定向消息接收成员（可选）
    AIGenerated    bool     // 是否为 AI 生成消息
}
```

### 用户在线状态回调 (UserOnlineStatusCallback)

**请求格式**: `application/json`（数组格式）

```go
type UserOnlineStatusCallback struct {
    UserID    string // 用户 ID
    Status    string // 0=上线，1=离线，2=登出
    OS        string // iOS, Android, HarmonyOS, Websocket, PC, MiniProgram
    Time      int64  // 发生时间（毫秒）
    ClientIP  string // IP 地址及端口
    SessionID string // 连接唯一 ID
}
```

### 审核结果回调 (AuditResultCallback)

**请求格式**: `application/json`
**签名位置**: HTTP Header（而非 URL Query）

```go
type AuditResultCallback struct {
    Result          int    // 10000=通过，10001=不通过
    Content         string // 审核内容（JSON 字符串）
    MsgUID          string // 消息 ID
    ServiceProvider string // 审核服务商
    ResultDetail    string // 审核详情（JSON 字符串）
}
```

### 其他回调类型

```go
// 聊天室状态回调
type ChatroomStatusCallback struct {
    ChatroomId string // 聊天室 ID
    Status     int    // 0=创建，1=销毁
    Time       int64  // 发生时间（毫秒）
}

// 聊天室 KV 属性回调
type ChatroomKVCallback struct {
    ChatroomId string // 聊天室 ID
    UserId     string // 操作用户 ID
    Key        string // 属性名
    Value      string // 属性值
    Time       int64  // 发生时间（毫秒）
}

// 用户注销/激活回调
type UserDeactivationCallback struct {
    UserId string // 用户 ID
    Type   int    // 0=注销，1=激活
    Time   int64  // 发生时间（毫秒）
}
```

## API 参考

### Handler 配置

```go
type HandlerConfig struct {
    AppSecret            string                                              // 应用密钥
    MessageRoutePath     string                                              // 自定义消息路由路径
    UserOnlineStatusPath string                                              // 自定义在线状态路径
    AuditResultPath      string                                              // 自定义审核结果路径
    ChatroomStatusPath   string                                              // 自定义聊天室状态路径
    ChatroomKVPath       string                                              // 自定义聊天室 KV 路径
    UserDeactivationPath string                                              // 自定义用户注销路径
    OnMessageRoute       func(ResponseWriter, MessageRouteCallback) error    // 消息路由处理器
    OnUserOnlineStatus   func(ResponseWriter, []UserOnlineStatusCallback) error // 在线状态处理器
    OnAuditResult        func(ResponseWriter, AuditResultCallback) error     // 审核结果处理器
    OnChatroomStatus     func(ResponseWriter, ChatroomStatusCallback) error  // 聊天室状态处理器
    OnChatroomKV         func(ResponseWriter, ChatroomKVCallback) error      // 聊天室 KV 处理器
    OnUserDeactivation   func(ResponseWriter, UserDeactivationCallback) error // 用户注销处理器
}
```

### 签名验证

```go
// 验证回调签名（SHA1(AppSecret + Nonce + Timestamp)）
func VerifyCallback(appSecret, nonce, timestamp, signature string) bool

// 从 HTTP 请求中提取回调参数
func ExtractParams(r *http.Request) CallbackParams

// 便捷方法：从请求中提取参数并验证签名
func VerifyRequest(r *http.Request, appSecret string) bool
```

### IP 白名单

```go
// 检查 IP 是否在融云白名单中
func IsValidIP(ip string) bool

// 国内数据中心 IP 列表
var IPWhitelist []string

// 海外数据中心 IP 列表（新加坡、北美、沙特）
var IPWhitelistOverseas []string
```

### ResponseWriter

```go
type ResponseWriter interface {
    // WriteResponse 写入自定义 HTTP 响应
    WriteResponse(code int, body string)
    // Header 获取 HTTP Header 以便设置自定义头部
    Header() http.Header
}
```

## 注意事项

1. **签名参数位置**:
   - 大部分回调的签名参数在 URL Query 中
   - 审核结果回调的签名在 HTTP Header 中

2. **请求格式**:
   - 消息路由回调使用 `application/x-www-form-urlencoded`
   - 其他回调使用 `application/json`

3. **响应格式**:
   - 成功响应需返回 `OK`（或任何 2xx 响应）
   - 失败返回非 2xx 状态码
   - 返回 5xx 或网络错误时，融云会重试（最多 3 次）

4. **IP 白名单**:
   - 强烈建议配置服务器防火墙只允许融云 IP 访问回调接口
   - 国内数据中心和海外数据中心 IP 列表不同

5. **并发处理**:
   - Handler 是线程安全的
   - 回调函数可能被并发调用，需要自行保证线程安全

6. **超时处理**:
   - 回调处理应尽快完成（建议 < 5 秒）
   - 大规模消息可能导致临时暂停（1 分钟后恢复）

## 融云文档

- [回调服务概述](https://docs.rongcloud.cn/platform-chat-api/auth-callback)
- [全量消息路由](https://docs.rongcloud.cn/platform-chat-api/message/sync)
- [用户在线状态](https://docs.rongcloud.cn/platform-chat-api/user/onlinestatus)
- [消息审核结果](https://docs.rongcloud.cn/platform-chat-api/moderation/audit-result)
