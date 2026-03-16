package rongcloud

import (
	"testing"
	"time"
)

func TestNewRC(t *testing.T) {
	rc := NewRC(&Options{
		AppKey:    "test-key",
		AppSecret: "test-secret",
		Region:    RegionBeijing,
		Timeout:   5 * time.Second,
	})
	if rc == nil {
		t.Fatal("NewRC returned nil")
	}
}

func TestModuleAccessors(t *testing.T) {
	rc := NewRC(&Options{
		AppKey:    "test-key",
		AppSecret: "test-secret",
		Region:    RegionBeijing,
	})

	// 验证所有模块访问器返回非 nil
	tests := []struct {
		name string
		fn   func() any
	}{
		{"User", func() any { return rc.User() }},
		{"UserTag", func() any { return rc.UserTag() }},
		{"UserBlock", func() any { return rc.UserBlock() }},
		{"UserProfile", func() any { return rc.UserProfile() }},
		{"Friend", func() any { return rc.Friend() }},
		{"Message", func() any { return rc.Message() }},
		{"Conversation", func() any { return rc.Conversation() }},
		{"Push", func() any { return rc.Push() }},
		{"Group", func() any { return rc.Group() }},
		{"GroupMute", func() any { return rc.GroupMute() }},
		{"UltraGroup", func() any { return rc.UltraGroup() }},
		{"UltraGroupChannel", func() any { return rc.UltraGroupChannel() }},
		{"UltraGroupUserGroup", func() any { return rc.UltraGroupUserGroup() }},
		{"UltraGroupMute", func() any { return rc.UltraGroupMute() }},
		{"Chatroom", func() any { return rc.Chatroom() }},
		{"ChatroomMute", func() any { return rc.ChatroomMute() }},
		{"ChatroomBlock", func() any { return rc.ChatroomBlock() }},
		{"ChatroomKV", func() any { return rc.ChatroomKV() }},
		{"ChatroomPriority", func() any { return rc.ChatroomPriority() }},
		{"ChatroomWhitelist", func() any { return rc.ChatroomWhitelist() }},
		{"SensitiveWord", func() any { return rc.SensitiveWord() }},
		{"Notification", func() any { return rc.Notification() }},
		{"PushPeriod", func() any { return rc.PushPeriod() }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.fn() == nil {
				t.Errorf("%s() returned nil", tt.name)
			}
		})
	}
}

func TestModuleAccessorsSingleton(t *testing.T) {
	rc := NewRC(&Options{
		AppKey:    "test-key",
		AppSecret: "test-secret",
		Region:    RegionBeijing,
	})

	// 验证多次调用返回同一实例（sync.Once）
	u1 := rc.User()
	u2 := rc.User()
	if u1 != u2 {
		t.Error("User() should return the same instance on multiple calls")
	}

	g1 := rc.Group()
	g2 := rc.Group()
	if g1 != g2 {
		t.Error("Group() should return the same instance on multiple calls")
	}
}

func TestRegionAliases(t *testing.T) {
	tests := []struct {
		name   string
		region Region
		domain string
	}{
		{"Beijing", RegionBeijing, "https://api.rong-api.com"},
		{"Singapore", RegionSingapore, "https://api.sg-light-api.com"},
		{"SingaporeB", RegionSingaporeB, "https://api.sg-b-light-api.com"},
		{"NorthAmerica", RegionNorthAmerica, "https://api.us-light-api.com"},
		{"SAU", RegionSAU, "https://api.sau-light-api.com"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.region.PrimaryDomain != tt.domain {
				t.Errorf("got %s, want %s", tt.region.PrimaryDomain, tt.domain)
			}
		})
	}
}
