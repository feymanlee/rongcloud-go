package core

import (
	"testing"
)

func TestNewError(t *testing.T) {
	e := NewError(1001, "test error")
	if e == nil {
		t.Fatal("NewError returned nil")
	}
	if e.Code() != 1001 {
		t.Errorf("expected code 1001, got %d", e.Code())
	}
	if e.Message() != "test error" {
		t.Errorf("expected message 'test error', got %q", e.Message())
	}
}

func TestErrorFormat(t *testing.T) {
	e := NewError(1001, "test error")
	expected := "rongcloud: code=1001, message=test error"
	if e.Error() != expected {
		t.Errorf("expected %q, got %q", expected, e.Error())
	}
}

func TestErrorCode(t *testing.T) {
	tests := []struct {
		code int
		msg  string
	}{
		{200, "success"},
		{1001, "app key error"},
		{0, ""},
		{-1, "invalid params"},
	}

	for _, tt := range tests {
		e := NewError(tt.code, tt.msg)
		if e.Code() != tt.code {
			t.Errorf("Code(): expected %d, got %d", tt.code, e.Code())
		}
		if e.Message() != tt.msg {
			t.Errorf("Message(): expected %q, got %q", tt.msg, e.Message())
		}
	}
}

func TestErrorImplementsErrorInterface(t *testing.T) {
	var _ error = NewError(1001, "test")
}

func TestErrorImplementsCoreErrorInterface(t *testing.T) {
	var _ Error = NewError(1001, "test")
}
