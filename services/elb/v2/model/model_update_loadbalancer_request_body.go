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
type UpdateLoadbalancerRequestBody struct {
	Loadbalancer *UpdateLoadbalancerReq `json:"loadbalancer"`
}

func (o UpdateLoadbalancerRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateLoadbalancerRequestBody", string(data)}, " ")
}
