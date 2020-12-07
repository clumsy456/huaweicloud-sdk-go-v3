/*
 * BCS
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
type ListBlockchainChannelsResponse struct {
	// 通道信息列表
	Channels       *[]Channel `json:"channels,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListBlockchainChannelsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListBlockchainChannelsResponse", string(data)}, " ")
}