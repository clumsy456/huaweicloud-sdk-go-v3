package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreatePostalRequest struct {
	// 语言。 中文：zh_CN 缺省为zh_CN。

	XLanguage *string `json:"X-Language,omitempty"`

	Body *AddPostalReq `json:"body,omitempty"`
}

func (o CreatePostalRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreatePostalRequest struct{}"
	}

	return strings.Join([]string{"CreatePostalRequest", string(data)}, " ")
}
