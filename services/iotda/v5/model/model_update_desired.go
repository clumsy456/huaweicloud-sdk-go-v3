package model

import (
	"encoding/json"

	"strings"
)

// 修改设备影子预期数据结构体。
type UpdateDesired struct {
	// 设备的服务ID，在设备关联的产品模型中定义。

	ServiceId string `json:"service_id"`
	// 设备影子期望属性数据，Json格式，里面是一个个键值对，每个键都是产品模型中属性的参数名(property_name)，目前如样例所示只支持一层结构；如果想要删除整个desired可以填写空object(例如\"desired\":{})，如果想要删除某一个属性期望值可以将属性置位null(例如{\"temperature\":null})

	Desired *interface{} `json:"desired"`
	// 设备影子的版本，携带改参数时平台会校验值必须等于当前影子版本，初始从0开始。

	Version *int64 `json:"version,omitempty"`
}

func (o UpdateDesired) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDesired struct{}"
	}

	return strings.Join([]string{"UpdateDesired", string(data)}, " ")
}
