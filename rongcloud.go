package rongcloud

import (
	"sync"

	"github.com/feymanlee/rongcloud-go/callback"
	"github.com/feymanlee/rongcloud-go/chatroom"
	"github.com/feymanlee/rongcloud-go/chatroom_block"
	"github.com/feymanlee/rongcloud-go/chatroom_kv"
	"github.com/feymanlee/rongcloud-go/chatroom_mute"
	"github.com/feymanlee/rongcloud-go/chatroom_priority"
	"github.com/feymanlee/rongcloud-go/chatroom_whitelist"
	"github.com/feymanlee/rongcloud-go/conversation"
	"github.com/feymanlee/rongcloud-go/friend"
	"github.com/feymanlee/rongcloud-go/group"
	"github.com/feymanlee/rongcloud-go/group_mute"
	"github.com/feymanlee/rongcloud-go/internal/core"
	"github.com/feymanlee/rongcloud-go/message"
	"github.com/feymanlee/rongcloud-go/notification"
	"github.com/feymanlee/rongcloud-go/push"
	"github.com/feymanlee/rongcloud-go/push_period"
	"github.com/feymanlee/rongcloud-go/sensitive_word"
	"github.com/feymanlee/rongcloud-go/ultragroup"
	"github.com/feymanlee/rongcloud-go/ultragroup_channel"
	"github.com/feymanlee/rongcloud-go/ultragroup_mute"
	"github.com/feymanlee/rongcloud-go/ultragroup_usergroup"
	"github.com/feymanlee/rongcloud-go/user"
	"github.com/feymanlee/rongcloud-go/user_block"
	"github.com/feymanlee/rongcloud-go/user_profile"
	"github.com/feymanlee/rongcloud-go/user_tag"
)

// Error 错误类型别名
type Error = core.Error

// Region 区域配置别名
type Region = core.Region

var (
	RegionBeijing      = core.RegionBeijing
	RegionSingapore    = core.RegionSingapore
	RegionSingaporeB   = core.RegionSingaporeB
	RegionNorthAmerica = core.RegionNorthAmerica
	RegionSAU          = core.RegionSAU
)

// RC 融云 IM 服务端 SDK 接口
type RC interface {
	// User 用户管理
	User() user.API
	// UserTag 用户标签
	UserTag() usertag.API
	// UserBlock 用户黑名单/白名单
	UserBlock() userblock.API
	// UserProfile 用户资料托管
	UserProfile() userprofile.API
	// Friend 好友管理
	Friend() friend.API
	// Message 消息管理
	Message() message.API
	// Conversation 会话管理
	Conversation() conversation.API
	// Push 推送管理
	Push() push.API
	// Group 群组管理
	Group() group.API
	// GroupMute 群组禁言
	GroupMute() groupmute.API
	// UltraGroup 超级群
	UltraGroup() ultragroup.API
	// UltraGroupChannel 超级群私有频道
	UltraGroupChannel() ultragroupchannel.API
	// UltraGroupUserGroup 超级群用户组
	UltraGroupUserGroup() ultragroupusergroup.API
	// UltraGroupMute 超级群禁言
	UltraGroupMute() ultragroupmute.API
	// Chatroom 聊天室
	Chatroom() chatroom.API
	// ChatroomMute 聊天室禁言
	ChatroomMute() chatroommute.API
	// ChatroomBlock 聊天室封禁
	ChatroomBlock() chatroomblock.API
	// ChatroomKV 聊天室KV属性
	ChatroomKV() chatroomkv.API
	// ChatroomPriority 聊天室消息优先级
	ChatroomPriority() chatroompriority.API
	// ChatroomWhitelist 聊天室白名单
	ChatroomWhitelist() chatroomwhitelist.API
	// SensitiveWord 敏感词
	SensitiveWord() sensitiveword.API
	// Notification 会话免打扰
	Notification() notification.API
	// PushPeriod 推送免打扰时段
	PushPeriod() pushperiod.API
	// Callback 回调服务
	Callback() callback.API
}

// Options 创建 RC 实例的选项
type Options = core.Options

// NewRC 创建融云 IM 服务端 SDK 实例
func NewRC(opt *Options) RC {
	return &rc{
		client:    core.NewClient(opt),
		appSecret: opt.AppSecret,
	}
}

type rc struct {
	client    core.Client
	appSecret string

	user struct {
		once     sync.Once
		instance user.API
	}
	userTag struct {
		once     sync.Once
		instance usertag.API
	}
	userBlock struct {
		once     sync.Once
		instance userblock.API
	}
	userProfile struct {
		once     sync.Once
		instance userprofile.API
	}
	friend struct {
		once     sync.Once
		instance friend.API
	}
	message struct {
		once     sync.Once
		instance message.API
	}
	conversation struct {
		once     sync.Once
		instance conversation.API
	}
	push struct {
		once     sync.Once
		instance push.API
	}
	group struct {
		once     sync.Once
		instance group.API
	}
	groupMute struct {
		once     sync.Once
		instance groupmute.API
	}
	ultraGroup struct {
		once     sync.Once
		instance ultragroup.API
	}
	ultraGroupChannel struct {
		once     sync.Once
		instance ultragroupchannel.API
	}
	ultraGroupUserGroup struct {
		once     sync.Once
		instance ultragroupusergroup.API
	}
	ultraGroupMute struct {
		once     sync.Once
		instance ultragroupmute.API
	}
	chatroom struct {
		once     sync.Once
		instance chatroom.API
	}
	chatroomMute struct {
		once     sync.Once
		instance chatroommute.API
	}
	chatroomBlock struct {
		once     sync.Once
		instance chatroomblock.API
	}
	chatroomKV struct {
		once     sync.Once
		instance chatroomkv.API
	}
	chatroomPriority struct {
		once     sync.Once
		instance chatroompriority.API
	}
	chatroomWhitelist struct {
		once     sync.Once
		instance chatroomwhitelist.API
	}
	sensitiveWord struct {
		once     sync.Once
		instance sensitiveword.API
	}
	notification struct {
		once     sync.Once
		instance notification.API
	}
	pushPeriod struct {
		once     sync.Once
		instance pushperiod.API
	}
	cb struct {
		once     sync.Once
		instance callback.API
	}
}

func (r *rc) User() user.API {
	r.user.once.Do(func() {
		r.user.instance = user.NewAPI(r.client)
	})
	return r.user.instance
}

func (r *rc) UserTag() usertag.API {
	r.userTag.once.Do(func() {
		r.userTag.instance = usertag.NewAPI(r.client)
	})
	return r.userTag.instance
}

func (r *rc) UserBlock() userblock.API {
	r.userBlock.once.Do(func() {
		r.userBlock.instance = userblock.NewAPI(r.client)
	})
	return r.userBlock.instance
}

func (r *rc) UserProfile() userprofile.API {
	r.userProfile.once.Do(func() {
		r.userProfile.instance = userprofile.NewAPI(r.client)
	})
	return r.userProfile.instance
}

func (r *rc) Friend() friend.API {
	r.friend.once.Do(func() {
		r.friend.instance = friend.NewAPI(r.client)
	})
	return r.friend.instance
}

func (r *rc) Message() message.API {
	r.message.once.Do(func() {
		r.message.instance = message.NewAPI(r.client)
	})
	return r.message.instance
}

func (r *rc) Conversation() conversation.API {
	r.conversation.once.Do(func() {
		r.conversation.instance = conversation.NewAPI(r.client)
	})
	return r.conversation.instance
}

func (r *rc) Push() push.API {
	r.push.once.Do(func() {
		r.push.instance = push.NewAPI(r.client)
	})
	return r.push.instance
}

func (r *rc) Group() group.API {
	r.group.once.Do(func() {
		r.group.instance = group.NewAPI(r.client)
	})
	return r.group.instance
}

func (r *rc) GroupMute() groupmute.API {
	r.groupMute.once.Do(func() {
		r.groupMute.instance = groupmute.NewAPI(r.client)
	})
	return r.groupMute.instance
}

func (r *rc) UltraGroup() ultragroup.API {
	r.ultraGroup.once.Do(func() {
		r.ultraGroup.instance = ultragroup.NewAPI(r.client)
	})
	return r.ultraGroup.instance
}

func (r *rc) UltraGroupChannel() ultragroupchannel.API {
	r.ultraGroupChannel.once.Do(func() {
		r.ultraGroupChannel.instance = ultragroupchannel.NewAPI(r.client)
	})
	return r.ultraGroupChannel.instance
}

func (r *rc) UltraGroupUserGroup() ultragroupusergroup.API {
	r.ultraGroupUserGroup.once.Do(func() {
		r.ultraGroupUserGroup.instance = ultragroupusergroup.NewAPI(r.client)
	})
	return r.ultraGroupUserGroup.instance
}

func (r *rc) UltraGroupMute() ultragroupmute.API {
	r.ultraGroupMute.once.Do(func() {
		r.ultraGroupMute.instance = ultragroupmute.NewAPI(r.client)
	})
	return r.ultraGroupMute.instance
}

func (r *rc) Chatroom() chatroom.API {
	r.chatroom.once.Do(func() {
		r.chatroom.instance = chatroom.NewAPI(r.client)
	})
	return r.chatroom.instance
}

func (r *rc) ChatroomMute() chatroommute.API {
	r.chatroomMute.once.Do(func() {
		r.chatroomMute.instance = chatroommute.NewAPI(r.client)
	})
	return r.chatroomMute.instance
}

func (r *rc) ChatroomBlock() chatroomblock.API {
	r.chatroomBlock.once.Do(func() {
		r.chatroomBlock.instance = chatroomblock.NewAPI(r.client)
	})
	return r.chatroomBlock.instance
}

func (r *rc) ChatroomKV() chatroomkv.API {
	r.chatroomKV.once.Do(func() {
		r.chatroomKV.instance = chatroomkv.NewAPI(r.client)
	})
	return r.chatroomKV.instance
}

func (r *rc) ChatroomPriority() chatroompriority.API {
	r.chatroomPriority.once.Do(func() {
		r.chatroomPriority.instance = chatroompriority.NewAPI(r.client)
	})
	return r.chatroomPriority.instance
}

func (r *rc) ChatroomWhitelist() chatroomwhitelist.API {
	r.chatroomWhitelist.once.Do(func() {
		r.chatroomWhitelist.instance = chatroomwhitelist.NewAPI(r.client)
	})
	return r.chatroomWhitelist.instance
}

func (r *rc) SensitiveWord() sensitiveword.API {
	r.sensitiveWord.once.Do(func() {
		r.sensitiveWord.instance = sensitiveword.NewAPI(r.client)
	})
	return r.sensitiveWord.instance
}

func (r *rc) Notification() notification.API {
	r.notification.once.Do(func() {
		r.notification.instance = notification.NewAPI(r.client)
	})
	return r.notification.instance
}

func (r *rc) PushPeriod() pushperiod.API {
	r.pushPeriod.once.Do(func() {
		r.pushPeriod.instance = pushperiod.NewAPI(r.client)
	})
	return r.pushPeriod.instance
}

func (r *rc) Callback() callback.API {
	r.cb.once.Do(func() {
		r.cb.instance = callback.NewAPI(r.client, r.appSecret)
	})
	return r.cb.instance
}
