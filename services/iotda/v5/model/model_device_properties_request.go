package model

import (
	"encoding/json"

	"strings"
)

type DevicePropertiesRequest struct {
	// 设备执行的属性，Json格式，里面是一个个健值对，如果serviceId不为空，每个健都是profile中属性的参数名（paraName）;如果serviceId为空则由用户自定义属性格式。设属性令示例：[{\"service_id\": \"Temperature\",\"properties\": {\"value\": 57}},{\"service_id\": \"Battery\",\"properties\": {\"level\": 80}}]，具体格式需要应用和设备约定。

	Services *interface{} `json:"services,omitempty"`
}

func (o DevicePropertiesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DevicePropertiesRequest struct{}"
	}

	return strings.Join([]string{"DevicePropertiesRequest", string(data)}, " ")
}
