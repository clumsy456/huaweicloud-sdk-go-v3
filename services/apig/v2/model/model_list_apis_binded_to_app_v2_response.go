package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListApisBindedToAppV2Response struct {
	// 符合条件的API总数

	Total *int32 `json:"total,omitempty"`
	// 本次返回的列表长度

	Size *int32 `json:"size,omitempty"`
	// 本次返回的API列表

	Auths          *[]AppAuthBindedApiResp `json:"auths,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListApisBindedToAppV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListApisBindedToAppV2Response struct{}"
	}

	return strings.Join([]string{"ListApisBindedToAppV2Response", string(data)}, " ")
}
