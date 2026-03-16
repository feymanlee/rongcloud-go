# RongCloud IM Server SDK for Go

融云即时通讯服务端 Go SDK，覆盖 ~198 个 API，提供模块化、类型安全的接口。

## 安装

```bash
go get github.com/feymanlee/rongcloud-go
```

## 快速开始

```go
package main

import (
	"fmt"
	"log"

	rc "github.com/feymanlee/rongcloud-go"
)

func main() {
	client := rc.NewRC(&rc.Options{
		AppKey:    "your-app-key",
		AppSecret: "your-app-secret",
		Region:    rc.RegionBeijing,
	})

	// 注册用户，获取 Token
	resp, err := client.User().GetToken("user001", "张三", "https://example.com/avatar.png")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Token:", resp.Token)

	// 发送单聊消息
	_, err = client.Message().SendPrivate(&message.SendPrivateReq{
		FromUserId:  "user001",
		ToUserId:    "user002",
		ObjectName:  "RC:TxtMsg",
		Content:     `{"content":"hello"}`,
		PushContent: "你收到一条消息",
	})
	if err != nil {
		log.Fatal(err)
	}
}
```

## 区域配置

SDK 支持多区域部署，每个区域配置主备双域名，网络异常时自动切换：

| 区域 | 变量 | 主域名 |
|------|------|--------|
| 北京 | `RegionBeijing` | `api.rong-api.com` |
| 新加坡 | `RegionSingapore` | `api.sg-light-api.com` |
| 新加坡 B | `RegionSingaporeB` | `api.sg-b-light-api.com` |
| 北美 | `RegionNorthAmerica` | `api.us-light-api.com` |
| 沙特 | `RegionSAU` | `api.sau-light-api.com` |

## 配置选项

```go
client := rc.NewRC(&rc.Options{
	AppKey:    "your-app-key",    // 必填，融云 App Key
	AppSecret: "your-app-secret", // 必填，融云 App Secret
	Region:    rc.RegionBeijing,  // 必填，服务区域
	Timeout:   15 * time.Second,  // 可选，请求超时时间，默认 10s
})
```

## 模块概览

所有模块通过 `RC` 接口按需懒加载，线程安全。

### 用户管理

| 模块 | 访问方法 | API 数量 | 说明 |
|------|----------|---------|------|
| `user` | `User()` | 14 | 注册、Token、封禁、注销 |
| `user_tag` | `UserTag()` | 3 | 用户标签设置与查询 |
| `user_block` | `UserBlock()` | 8 | 黑名单、白名单、消息过滤 |
| `user_profile` | `UserProfile()` | 4 | 用户资料托管 (JSON) |

### 社交关系

| 模块 | 访问方法 | API 数量 | 说明 |
|------|----------|---------|------|
| `friend` | `Friend()` | 8 | 好友增删查、备注、黑名单 |

### 消息与会话

| 模块 | 访问方法 | API 数量 | 说明 |
|------|----------|---------|------|
| `message` | `Message()` | 15 | 单聊/群聊/聊天室/系统消息、撤回、历史 |
| `conversation` | `Conversation()` | 1 | 会话免打扰设置 |

### 推送

| 模块 | 访问方法 | API 数量 | 说明 |
|------|----------|---------|------|
| `push` | `Push()` | 14 | 推送、广播、标签推送、任务查询 |
| `push_period` | `PushPeriod()` | 3 | 推送免打扰时段 |
| `notification` | `Notification()` | 4 | 会话级/全局免打扰 |

### 群组

| 模块 | 访问方法 | API 数量 | 说明 |
|------|----------|---------|------|
| `group` | `Group()` | 30 | 基础群组 + 托管群组全部接口 |
| `group_mute` | `GroupMute()` | 9 | 群组禁言、全员禁言、白名单 |

### 超级群

| 模块 | 访问方法 | API 数量 | 说明 |
|------|----------|---------|------|
| `ultragroup` | `UltraGroup()` | 17 | 创建、消息、扩展、免打扰、频道 |
| `ultragroup_channel` | `UltraGroupChannel()` | 5 | 私有频道管理 |
| `ultragroup_usergroup` | `UltraGroupUserGroup()` | 10 | 用户组管理 (JSON) |
| `ultragroup_mute` | `UltraGroupMute()` | 8 | 超级群禁言与白名单 |

### 聊天室

| 模块 | 访问方法 | API 数量 | 说明 |
|------|----------|---------|------|
| `chatroom` | `Chatroom()` | 11 | 创建、销毁、查询、保活 |
| `chatroom_mute` | `ChatroomMute()` | 13 | 禁言、全员禁言、全局禁言、白名单 |
| `chatroom_block` | `ChatroomBlock()` | 3 | 聊天室封禁 |
| `chatroom_kv` | `ChatroomKV()` | 5 | KV 属性存储 |
| `chatroom_priority` | `ChatroomPriority()` | 3 | 消息优先级管理 |
| `chatroom_whitelist` | `ChatroomWhitelist()` | 6 | 消息类型/用户白名单 |

### 内容审核

| 模块 | 访问方法 | API 数量 | 说明 |
|------|----------|---------|------|
| `sensitive_word` | `SensitiveWord()` | 4 | 敏感词增删查 |

## API 调用示例

### 用户管理

```go
// 注册用户
resp, err := client.User().GetToken("uid001", "用户名", "https://example.com/avatar.png")

// 封禁用户 30 分钟
_, err = client.User().BlockAdd("uid001", 30)

// 查询用户在线状态
status, err := client.User().OnlineStatusCheck("uid001")
```

### 群组

```go
// 创建群组
_, err := client.Group().Create([]string{"uid001", "uid002"}, "group001", "我的群组")

// 群组禁言
err = client.GroupMute().MuteAdd("group001", "uid001", "60")
```

### 聊天室

```go
// 创建聊天室
err := client.Chatroom().Create(map[string]string{"room001": "聊天室名称"})

// 设置 KV 属性
err = client.ChatroomKV().Set("room001", "uid001", "key1", "value1", "0", "")
```

### 消息

```go
// 发送单聊消息
_, err := client.Message().SendPrivate(&message.SendPrivateReq{
	FromUserId:  "uid001",
	ToUserId:    "uid002",
	ObjectName:  "RC:TxtMsg",
	Content:     `{"content":"hello"}`,
})

// 撤回消息
_, err = client.Message().RecallPrivate(&message.RecallReq{
	FromUserId:  "uid001",
	TargetId:    "uid002",
	MessageUID:  "msg-uid",
	SentTime:    "1234567890",
})
```

### 推送

```go
// 全员广播
_, err := client.Push().BroadcastSend("system", "RC:TxtMsg", `{"content":"公告"}`, "系统公告", "")
```

## 错误处理

所有 API 方法在响应码非 200 时自动返回 `Error` 类型，可获取错误码与详情：

```go
resp, err := client.User().GetToken("uid", "name", "portrait")
if err != nil {
	if rcErr, ok := err.(rongcloud.Error); ok {
		fmt.Println("错误码:", rcErr.Code())
		fmt.Println("错误信息:", rcErr.Message())
	}
}
```

## 架构设计

- **双协议支持**：V1 API 使用 `application/x-www-form-urlencoded`，V2 API 使用 `application/json`
- **域名自动切换**：主域名请求失败时自动切换到备用域名，30 秒冷却间隔
- **签名认证**：每次请求自动生成 `SHA1(AppSecret + Nonce + Timestamp)` 签名
- **懒加载**：模块按需初始化，通过 `sync.Once` 保证线程安全
- **HTTP 客户端**：基于 [go-resty/resty](https://github.com/go-resty/resty) 实现

## 项目结构

```
rongcloud-go/
├── rongcloud.go                  # 入口：RC 接口 + NewRC()
├── internal/
│   ├── core/
│   │   ├── client.go             # HTTP 客户端、认证、域名切换
│   │   └── error.go              # 错误类型
│   ├── types/
│   │   └── types.go              # BaseResp 基础响应
│   └── enum/
│       └── code.go               # 状态码常量
├── user/                         # 用户管理
├── user_tag/                     # 用户标签
├── user_block/                   # 黑名单/白名单
├── user_profile/                 # 用户资料托管
├── friend/                       # 好友管理
├── message/                      # 消息管理
├── conversation/                 # 会话管理
├── push/                         # 推送管理
├── push_period/                  # 推送免打扰时段
├── notification/                 # 会话免打扰
├── group/                        # 群组管理
├── group_mute/                   # 群组禁言
├── ultragroup/                   # 超级群
├── ultragroup_channel/           # 超级群私有频道
├── ultragroup_usergroup/         # 超级群用户组
├── ultragroup_mute/              # 超级群禁言
├── chatroom/                     # 聊天室
├── chatroom_mute/                # 聊天室禁言
├── chatroom_block/               # 聊天室封禁
├── chatroom_kv/                  # 聊天室 KV 属性
├── chatroom_priority/            # 聊天室消息优先级
├── chatroom_whitelist/           # 聊天室白名单
└── sensitive_word/               # 敏感词
```

## License

MIT
