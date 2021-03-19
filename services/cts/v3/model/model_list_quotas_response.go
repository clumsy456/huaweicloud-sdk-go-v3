package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListQuotasResponse struct {
	// 本次查询追踪器列表返回的追踪器数组。

	Resources      *[]Quota `json:"resources,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ListQuotasResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListQuotasResponse struct{}"
	}

	return strings.Join([]string{"ListQuotasResponse", string(data)}, " ")
}
