package model

import (
	"encoding/json"

	"strings"
)

// 单机转主备时必填。
type Single2Ha struct {
	// 实例节点可用区码（AZ）。

	AzCodeNewNode string `json:"az_code_new_node"`
	// 仅在支持SQL Server数据库实例进行单机转主备时可选，指定时会验证密码有效性。

	Password *string `json:"password,omitempty"`
	// Dec用户专属存储ID，每个az配置的专属存储不同，实例节点所在专属存储ID，仅支持DEC用户创建时使用。

	DsspoolId *string `json:"dsspool_id,omitempty"`
	// 仅包周期实例进行单机转主备时可指定，表示是否自动从客户的账户中支付。 - true，为自动支付。 - false，为手动支付，默认该方式。

	IsAutoPay *bool `json:"is_auto_pay,omitempty"`
}

func (o Single2Ha) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Single2Ha struct{}"
	}

	return strings.Join([]string{"Single2Ha", string(data)}, " ")
}
