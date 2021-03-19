package model

import (
	"encoding/json"

	"strings"
)

// lb状态树的策略状态信息
type LoadBalancerStatusPolicy struct {
	// 匹配动作。 支持REDIRECT_TO_POOL和REDIRECT_TO_LISTENER。

	Action *string `json:"action,omitempty"`
	// 策略ID。

	Id *string `json:"id,omitempty"`
	// provisioning的状态。 可以为：ACTIVE、PENDING_CREATE 或者ERROR。默认为ACTIVE。

	ProvisioningStatus *string `json:"provisioning_status,omitempty"`
	// 策略名称。

	Name *string `json:"name,omitempty"`
	// 规则。

	Rules *[]LoadBalancerStatusL7Rule `json:"rules,omitempty"`
}

func (o LoadBalancerStatusPolicy) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "LoadBalancerStatusPolicy struct{}"
	}

	return strings.Join([]string{"LoadBalancerStatusPolicy", string(data)}, " ")
}
