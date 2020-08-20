/*
 * Cbr
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowMembersDetailResponse struct {
	// 添加备份共享成员响应信息
	Members []Member `json:"members,omitempty"`
}

func (o ShowMembersDetailResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowMembersDetailResponse", string(data)}, " ")
}