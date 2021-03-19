package model

import (
	"encoding/json"

	"strings"
)

type BatchTagActionDelRequestBody struct {
	// 操作标识（区分大小写）：删除时为“delete”。

	Action string `json:"action"`
	// 标签列表。单个实例总标签数上限10个。

	Tags []TagDelWithKeyValue `json:"tags"`
}

func (o BatchTagActionDelRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchTagActionDelRequestBody struct{}"
	}

	return strings.Join([]string{"BatchTagActionDelRequestBody", string(data)}, " ")
}
