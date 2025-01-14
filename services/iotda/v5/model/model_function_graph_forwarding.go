package model

import (
	"encoding/json"

	"strings"
)

// 函数工作流转发配置信息
type FunctionGraphForwarding struct {
	// 函数的URN（Uniform Resource Name），唯一标识函数。

	FuncUrn string `json:"func_urn"`
	// 函数名称。

	FuncName string `json:"func_name"`
}

func (o FunctionGraphForwarding) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "FunctionGraphForwarding struct{}"
	}

	return strings.Join([]string{"FunctionGraphForwarding", string(data)}, " ")
}
