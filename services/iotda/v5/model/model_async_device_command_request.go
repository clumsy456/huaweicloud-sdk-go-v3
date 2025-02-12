package model

import (
	"encoding/json"

	"strings"
)

type AsyncDeviceCommandRequest struct {
	// 设备命令所属的设备服务ID，在设备关联的产品模型中定义。如设备需要编解码插件来解析命令，此参数为必填项。

	ServiceId *string `json:"service_id,omitempty"`
	// 设备命令名称，在设备关联的产品模型中定义。如设备需要编解码插件来解析命令，此参数为必填项。

	CommandName *string `json:"command_name,omitempty"`
	// 设备执行的命令，Json格式，里面是一个个健值对，如果service_id不为空，每个健都是profile中命令的参数名（paraName）;如果service_id为空则由用户自定义命令格式。设备命令示例：{\"value\":\"1\"}，具体格式需要应用和设备约定， 最大32K。

	Paras *interface{} `json:"paras"`
	// 物联网平台缓存命令的时长， 单位秒, 平台最多缓存20条消息（即最多缓存20条PENDING状态的命令） 该参数在send_strategy字段为delay时有效，默认缓存24小时，最大缓存2天；

	ExpireTime *int32 `json:"expire_time,omitempty"`
	// 下发策略， immediately表示立即下发，此时expire_time无效；delay表示缓存下发，等数据上报或者设备上线之后下发，默认缓存下发。delay策略下，expire_time为0或空时，命令会默认缓存24小时。

	SendStrategy string `json:"send_strategy"`
}

func (o AsyncDeviceCommandRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AsyncDeviceCommandRequest struct{}"
	}

	return strings.Join([]string{"AsyncDeviceCommandRequest", string(data)}, " ")
}
