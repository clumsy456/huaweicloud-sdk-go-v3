/*
 * ELB
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type UpdateMemberRequestBody struct {
	Member *UpdateMemberReq `json:"member"`
}

func (o UpdateMemberRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateMemberRequestBody", string(data)}, " ")
}
