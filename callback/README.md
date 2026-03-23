# 融云回调处理模块

本模块提供融云 IM 回调服务的处理功能，包括签名验证、回调消息解析和 HTTP Handler 封装。

## 功能特性

- **签名验证**: 支持融云回调请求的 SHA1 签名验证
- **回调类型支持**:
  - 消息路由回调（全量消息路由）- `application/x-www-form-urlencoded`
  - 消息回调服务（自定义条件消息抄送）- `application/x-www-form-urlencoded`
  - 用户在线状态回调 - `application/json` (数组)
  - 消息审核结果回调 - `application/json`
  - 聊天室状态回调 - `application/json`
  - 聊天室 KV 属性回调 - `application/json`
  - 用户注销/激活回调 - `application/json`
  - 消息操作状态同步回调（撤回/删除）- `application/json`
  - 机器人消息回调 - `application/json`
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

### 通过 RC 使用回调服务（推荐）

```go
package main

import (
    "log"
    "net/http"
    "github.com/feymanlee/rongcloud-go"
)

func main() {
    // 创建 RC 实例
    rc := rongcloud.NewRC(&rongcloud.Options{
        AppKey:    "your-app-key",
        AppSecret: "your-app-secret",
        Region:    rongcloud.RegionBeijing,
    })

    // 获取 callback API
    callbackAPI := rc.Callback()

    // 设置回调处理器
    callbackAPI.SetHandlerConfig(callback.HandlerConfig{
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

    // 获取 Handler 并注册到 HTTP 服务器
    http.Handle("/message/sync", callbackAPI.Handler())
    http.Handle("/user/onlinestatus", callbackAPI.Handler())
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### 直接使用 Handler（无需 RC）

```go
package main

import (
    "log"
    "net/http"
    "github.com/feymanlee/rongcloud-go/callback"
)

func main() {
    handler := callback.NewHandler("your-app-secret", callback.HandlerConfig{
        OnMessageRoute: func(w callback.ResponseWriter, msg callback.MessageRouteCallback) error {
            log.Printf("收到消息: from=%s, to=%s, type=%s",
                msg.FromUserId, msg.ToUserId, msg.ObjectName)
            return nil
        },
        OnUserOnlineStatus: func(w callback.ResponseWriter, statuses []callback.UserOnlineStatusCallback) error {
            for _, status := range statuses {
                log.Printf("用户状态: user=%s, status=%s, os=%s",
                    status.UserID, status.Status, status.OS)
            }
            return nil
        },
    })

    // 将 Handler 注册到 HTTP 服务器
    http.Handle("/message/sync", handler)
    http.Handle("/user/onlinestatus", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## 路径常量

```go
const (
    DefaultMessageRoutePath     = "/message/sync"
    DefaultUserOnlineStatusPath = "/user/onlinestatus"
    DefaultAuditResultPath      = "/moderation/audit-result"
    DefaultChatroomStatusPath   = "/chatroom/status"
    DefaultChatroomKVPath       = "/chatroom/kv"
    DefaultUserDeactivationPath = "/user/deactivation"
    DefaultMessageOperationPath = "/message/operation"
    DefaultMessageCallbackPath  = "/message/callback"
    DefaultBotMessagePath       = "/bot/message"
)
```

## 融云后台配置说明

代码将 Handler 注册到了默认路径。在融云控制台配置回调地址时，需要填写完整的 URL：

```
https://your-domain.com/message/sync
https://your-domain.com/user/onlinestatus
```

例如：
- 如果你的服务部署在 `https://api.example.com`
- 代码中注册的路径是 `/message/sync`
- 那么融云后台消息路由回调应该配置为：`https://api.example.com/message/sync`

## 路径匹配说明

Handler 通过请求路径（`r.URL.Path`）来识别回调类型。默认路径如下：
- 消息路由: `/message/sync`
- 用户在线状态: `/user/onlinestatus`
- 审核结果: `/moderation/audit-result`
- 聊天室状态: `/chatroom/status`
- 聊天室 KV: `/chatroom/kv`
- 用户注销/激活: `/user/deactivation`
- 消息操作状态: `/message/operation`
- 消息回调服务: `/message/callback`
- 机器人消息: `/bot/message`

## 自定义回调路径

融云控制台允许自定义回调地址，可在配置中指定：

```go
callbackAPI.SetHandlerConfig(callback.HandlerConfig{
    MessageRoutePath:     "/api/callback/message",      // 自定义消息路由路径
    UserOnlineStatusPath: "/api/callback/online",       // 自定义在线状态路径
    AuditResultPath:      "/api/callback/audit",        // 自定义审核结果路径
    MessageOperationPath: "/api/callback/operation",    // 自定义消息操作路径
    MessageCallbackPath:  "/api/callback/msgcb",        // 自定义消息回调服务路径
    BotMessagePath:       "/api/callback/bot",          // 自定义机器人消息路径
    OnMessageRoute: func(w callback.ResponseWriter, msg callback.MessageRouteCallback) error {
        // 处理消息
        return nil
    },
})
```

## 自定义响应

回调处理器可以通过 `ResponseWriter` 自定义 HTTP 响应：

```go
callbackAPI.SetHandlerConfig(callback.HandlerConfig{
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

// 消息操作状态同步回调（消息撤回/删除）
type MessageOperationCallback struct {
    EventType        int                           // 事件类型：1=消息撤回，2=消息删除
    FromUserId       string                        // 操作人用户 ID
    OptTime          int64                         // 操作时间戳（毫秒）
    Source           string                        // 操作来源
    ConversationInfo MessageOperationConversation  // 会话信息
    OriginalMsgInfo  MessageOperationOriginalMsg   // 原始消息信息
    RecallMsgInfo    *MessageOperationRecallInfo   // 撤回消息特有信息
    DeleteMsgInfo    *MessageOperationDeleteInfo   // 删除消息特有信息
}

// 消息回调服务
// 注意：此回调使用 application/x-www-form-urlencoded 格式，且 appKey 在请求体中
type MessageCallbackService struct {
    AppKey         string // 应用 App Key
    FromUserId     string // 发送用户 ID
    TargetId       string // 目标会话 ID
    ToUserIds      string // 群成员 ID 列表（逗号分隔）
    MsgType        string // 消息类型标识
    Content        string // JSON 结构的消息内容
    PushContent    string // 推送通知栏显示内容
    DisablePush    bool   // 是否为静默消息
    PushExt        string // 推送通知配置
    Expansion      bool   // 是否为可扩展消息
    ExtraContent   string // 消息扩展内容（JSON 字符串）
    ChannelType    string // 会话类型
    MsgTimeStamp   string // 服务器时间戳（毫秒）
    MessageId      string // 消息唯一标识
    OriginalMsgUID string // 原始消息 ID（超级群有效）
    OS             string // 消息来源
    BusChannel     string // 超级群频道 ID
    ClientIp       string // 用户 IP 地址及端口
    AiGenerated    bool   // 是否为 AI 生成消息
}

// 机器人消息回调
type BotMessageCallback struct {
    Type      string                 // 回调事件类型
    Timestamp int64                  // 触发时间（Unix 毫秒）
    Bot       BotInfo                // 机器人信息
    Data      map[string]interface{} // 事件特定数据
}

type BotInfo struct {
    UserId     string                 // 机器人用户 ID
    Name       string                 // 机器人名称
    ProfileUrl string                 // 机器人头像 URL
    Type       string                 // 机器人类型
    Metadata   map[string]interface{} // 机器人元数据
}
```

### 机器人消息事件类型

```go
const (
    // 消息事件
    BotEventMessagePrivate              = "message:private"                // 私聊消息
    BotEventMessageGroupMentioned       = "message:group_mentioned"        // 群组@消息
    BotEventMessagePrivateRecall        = "message:private_recall"         // 私聊消息撤回
    BotEventMessageGroupMentionedRecall = "message:group_mentioned_recall" // 群组@消息撤回
    BotEventMessagePrivateRead          = "message:private_read"           // 私聊已读回执
    BotEventMessageGroupRead            = "message:group_read"             // 群组已读回执
    // 群组事件
    BotEventGroupBotJoin  = "group:bot_join"  // 机器人被加入群组
    BotEventGroupBotLeft  = "group:bot_left"  // 机器人被移出群组
    BotEventGroupDismiss  = "group:dismiss"   // 群组解散
    BotEventGroupUserJoin = "group:user_join" // 其他用户加入群组
    BotEventGroupUserLeft = "group:user_left" // 其他用户离开群组
)
```

## API 参考

### callback.API 接口

```go
type API interface {
    // HandlerConfig 获取 Handler 配置，用于设置回调处理器
    HandlerConfig() HandlerConfig
    // SetHandlerConfig 设置 Handler 配置
    SetHandlerConfig(config HandlerConfig)
    // Handler 获取 HTTP Handler 实例
    // 注意：需要先在 HandlerConfig 中设置回调处理器
    Handler() *Handler
}
```

### HandlerConfig

```go
type HandlerConfig struct {
    // 自定义回调路径（可选，默认使用标准路径）
    MessageRoutePath     string // 消息路由回调路径，默认 DefaultMessageRoutePath
    UserOnlineStatusPath string // 用户在线状态回调路径，默认 DefaultUserOnlineStatusPath
    AuditResultPath      string // 审核结果回调路径，默认 DefaultAuditResultPath
    ChatroomStatusPath   string // 聊天室状态回调路径，默认 DefaultChatroomStatusPath
    ChatroomKVPath       string // 聊天室 KV 回调路径，默认 DefaultChatroomKVPath
    UserDeactivationPath string // 用户注销/激活回调路径，默认 DefaultUserDeactivationPath
    MessageOperationPath string // 消息操作状态同步回调路径，默认 DefaultMessageOperationPath
    MessageCallbackPath  string // 消息回调服务路径，默认 DefaultMessageCallbackPath
    BotMessagePath       string // 机器人消息回调路径，默认 DefaultBotMessagePath

    // 回调处理器 - 可以通过 ResponseWriter 自定义响应
    OnMessageRoute       func(ResponseWriter, MessageRouteCallback) error
    OnUserOnlineStatus   func(ResponseWriter, []UserOnlineStatusCallback) error
    OnAuditResult        func(ResponseWriter, AuditResultCallback) error
    OnChatroomStatus     func(ResponseWriter, ChatroomStatusCallback) error
    OnChatroomKV         func(ResponseWriter, ChatroomKVCallback) error
    OnUserDeactivation   func(ResponseWriter, UserDeactivationCallback) error
    OnMessageOperation   func(ResponseWriter, MessageOperationCallback) error
    OnMessageCallback    func(ResponseWriter, MessageCallbackService) error
    OnBotMessage         func(ResponseWriter, BotMessageCallback) error
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

// 从请求中提取 appKey，支持 Query、Header 和请求体
func ExtractAppKey(r *http.Request) string
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
   - 消息回调服务的 AppKey 在请求体（form-urlencoded）中，本库会自动处理

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
