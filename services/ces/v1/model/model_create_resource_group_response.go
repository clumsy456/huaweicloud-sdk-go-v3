package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateResourceGroupResponse struct {
	// 创建的资源分组ID，如：rg1606377637506DmVOENVyL。

	GroupId        *string `json:"group_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateResourceGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateResourceGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateResourceGroupResponse", string(data)}, " ")
}
