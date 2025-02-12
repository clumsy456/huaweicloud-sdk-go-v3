package model

import (
	"encoding/json"

	"strings"
)

// 服务配置信息
type MqsForwarding struct {
	// MQS服务的URL

	Url string `json:"url"`
	// 用于登录MQS的用户名

	UserName string `json:"user_name"`
	// 用于登录MQS的密码

	Password string `json:"password"`
	// 订阅的MQS主题

	Topic string `json:"topic"`
	// 是否加密传输

	EncryptTransport *bool `json:"encrypt_transport,omitempty"`
}

func (o MqsForwarding) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MqsForwarding struct{}"
	}

	return strings.Join([]string{"MqsForwarding", string(data)}, " ")
}
