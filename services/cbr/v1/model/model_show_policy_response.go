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
type ShowPolicyResponse struct {
	Policy *Policy `json:"policy,omitempty"`
}

func (o ShowPolicyResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowPolicyResponse", string(data)}, " ")
}