package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListEnvironmentVariablesV2Response struct {
	// 满足条件的环境变量总数

	Total *int32 `json:"total,omitempty"`
	// 本次返回的列表长度

	Size *int32 `json:"size,omitempty"`
	// 本次返回的环境变量列表

	Variables      *[]EnvVariableResp `json:"variables,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListEnvironmentVariablesV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListEnvironmentVariablesV2Response struct{}"
	}

	return strings.Join([]string{"ListEnvironmentVariablesV2Response", string(data)}, " ")
}
