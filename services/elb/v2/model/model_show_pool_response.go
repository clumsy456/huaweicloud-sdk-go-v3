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

// Response Object
type ShowPoolResponse struct {
	Pool *PoolResp `json:"pool,omitempty"`
}

func (o ShowPoolResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowPoolResponse", string(data)}, " ")
}
