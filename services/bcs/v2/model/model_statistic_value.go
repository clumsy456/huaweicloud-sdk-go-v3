package model

import (
	"encoding/json"

	"strings"
)

// 统计结果
type StatisticValue struct {
	// 统计方式。

	Statistic *string `json:"statistic,omitempty"`
	// 统计结果。

	Value *float64 `json:"value,omitempty"`
}

func (o StatisticValue) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StatisticValue struct{}"
	}

	return strings.Join([]string{"StatisticValue", string(data)}, " ")
}
