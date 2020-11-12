/*
 * CES
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowResourceGroupRequest struct {
	ContentType string  `json:"Content-Type"`
	GroupId     string  `json:"group_id"`
	Status      *string `json:"status,omitempty"`
	Namespace   *string `json:"namespace,omitempty"`
	Dname       *string `json:"dname,omitempty"`
	Start       *string `json:"start,omitempty"`
	Limit       *string `json:"limit,omitempty"`
}

func (o ShowResourceGroupRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowResourceGroupRequest", string(data)}, " ")
}