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
type CreateMemberRequestBody struct {
	Member *CreateMemberReq `json:"member"`
}

func (o CreateMemberRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateMemberRequestBody", string(data)}, " ")
}
