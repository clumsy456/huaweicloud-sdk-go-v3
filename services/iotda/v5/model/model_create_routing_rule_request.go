package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateRoutingRuleRequest struct {
	// 实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。

	InstanceId *string `json:"Instance-Id,omitempty"`

	Body *AddRuleReq `json:"body,omitempty"`
}

func (o CreateRoutingRuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateRoutingRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateRoutingRuleRequest", string(data)}, " ")
}
