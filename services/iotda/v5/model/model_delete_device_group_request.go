package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteDeviceGroupRequest struct {
	// 实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。

	InstanceId *string `json:"Instance-Id,omitempty"`
	// 设备组ID，用于唯一标识一个设备组，在创建设备组时由物联网平台分配。

	GroupId string `json:"group_id"`
}

func (o DeleteDeviceGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteDeviceGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteDeviceGroupRequest", string(data)}, " ")
}
