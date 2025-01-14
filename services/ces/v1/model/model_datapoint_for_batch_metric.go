package model

import (
	"encoding/json"

	"strings"
)

//
type DatapointForBatchMetric struct {
	// 指标值，该字段名称与请求参数中filter使用的查询值相同；字段名称可为：max/min/average/sum/variance。

	Max *float64 `json:"max,omitempty"`
	// 指标值，该字段名称与请求参数中filter使用的查询值相同；字段名称可为：max/min/average/sum/variance。

	Min *float64 `json:"min,omitempty"`
	// 指标值，该字段名称与请求参数中filter使用的查询值相同；字段名称可为：max/min/average/sum/variance。

	Average *float64 `json:"average,omitempty"`
	// 指标值，该字段名称与请求参数中filter使用的查询值相同；字段名称可为：max/min/average/sum/variance。

	Sum *float64 `json:"sum,omitempty"`
	// 指标值，该字段名称与请求参数中filter使用的查询值相同；字段名称可为：max/min/average/sum/variance。

	Variance *string `json:"variance,omitempty"`
	// 指标采集时间，UNIX时间戳，单位毫秒。

	Timestamp int64 `json:"timestamp"`
}

func (o DatapointForBatchMetric) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DatapointForBatchMetric struct{}"
	}

	return strings.Join([]string{"DatapointForBatchMetric", string(data)}, " ")
}
