package model

import (
	"encoding/json"

	"strings"
)

// 监控数据信息
type MetricItemResultApi struct {
	Metric *MetricDemision `json:"metric,omitempty"`
	// 监控数据信息

	DataPoints *[]MetricDataPoints `json:"dataPoints,omitempty"`
}

func (o MetricItemResultApi) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MetricItemResultApi struct{}"
	}

	return strings.Join([]string{"MetricItemResultApi", string(data)}, " ")
}
