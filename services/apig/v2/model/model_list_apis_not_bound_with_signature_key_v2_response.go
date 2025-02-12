package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListApisNotBoundWithSignatureKeyV2Response struct {
	// 满足查询条件的API的总个数

	Total *int32 `json:"total,omitempty"`
	// 本次查询返回的API列表长度

	Size *int32 `json:"size,omitempty"`
	// 本次查询返回的API列表

	Apis           *[]SignUnbindingApiResp `json:"apis,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListApisNotBoundWithSignatureKeyV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListApisNotBoundWithSignatureKeyV2Response struct{}"
	}

	return strings.Join([]string{"ListApisNotBoundWithSignatureKeyV2Response", string(data)}, " ")
}
