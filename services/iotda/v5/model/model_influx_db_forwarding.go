package model

import (
	"encoding/json"

	"strings"
)

// InfluxDB配置信息
type InfluxDbForwarding struct {
	Address *NetAddress `json:"address"`
	// 连接InfluxDB数据库的库名,不存在会自动创建

	DbName string `json:"db_name"`
	// 连接InfluxDB数据库的用户名

	Username string `json:"username"`
	// 连接InfluxDB数据库的密码

	Password string `json:"password"`
	// InfluxDB数据库的measurement,不存在会自动创建

	Measurement string `json:"measurement"`
	// InfluxDB数据库和流转数据的对应关系列表。

	ColumnMappings []ColumnMapping `json:"column_mappings"`
}

func (o InfluxDbForwarding) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "InfluxDbForwarding struct{}"
	}

	return strings.Join([]string{"InfluxDbForwarding", string(data)}, " ")
}
