package sensitiveword

import "github.com/feymanlee/rongcloud-go/internal/types"

// Word 敏感词
type Word struct {
	Type        string `json:"type"`
	Word        string `json:"word"`
	ReplaceWord string `json:"replaceWord,omitempty"`
}

// ListResp 敏感词列表响应
type ListResp struct {
	types.BaseResp
	Words []Word `json:"words"`
}
