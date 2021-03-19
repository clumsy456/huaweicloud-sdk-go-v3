package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListWhitelistsRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	PageReverse *bool `json:"page_reverse,omitempty"`

	Id *string `json:"id,omitempty"`

	EnableWhitelist *bool `json:"enable_whitelist,omitempty"`

	ListenerId *string `json:"listener_id,omitempty"`

	Whitelist *string `json:"whitelist,omitempty"`
}

func (o ListWhitelistsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListWhitelistsRequest struct{}"
	}

	return strings.Join([]string{"ListWhitelistsRequest", string(data)}, " ")
}
