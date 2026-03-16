package core

import "fmt"

// Error 错误接口
type Error interface {
	error
	Code() int
	Message() string
}

type apiError struct {
	code    int
	message string
}

// NewError 创建错误
func NewError(code int, message string) Error {
	return &apiError{code: code, message: message}
}

func (e *apiError) Error() string {
	return fmt.Sprintf("rongcloud: code=%d, message=%s", e.code, e.message)
}

func (e *apiError) Code() int {
	return e.code
}

func (e *apiError) Message() string {
	return e.message
}
