package model

import (
	"encoding/json"

	"strings"
)

// 创建标签请求结构体。
type BindTagsDto struct {
	// 要绑定标签的资源类型。 - device：设备。

	ResourceType string `json:"resource_type"`
	// 要绑定标签的资源id。例如，资源类型为device，那么对应的资源id就是device_id。

	ResourceId string `json:"resource_id"`
	// 要绑定到指定资源的标签列表，标签列表中各项标签键值之间不允许重复，一个资源最多可以绑定10个标签。

	Tags []TagV5Dto `json:"tags"`
}

func (o BindTagsDto) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BindTagsDto struct{}"
	}

	return strings.Join([]string{"BindTagsDto", string(data)}, " ")
}
