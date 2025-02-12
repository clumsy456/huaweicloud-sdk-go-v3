package model

import (
	"encoding/json"

	"strings"
)

// 发送给SMN消息结构
type ActionSmnForwarding struct {
	// SMN服务对应的region区域

	RegionName string `json:"region_name"`
	// SMN服务对应的projectId信息

	ProjectId string `json:"project_id"`
	// SMN服务对应的主题名称

	ThemeName string `json:"theme_name"`
	// SMN服务对应的topic的主题URN

	TopicUrn string `json:"topic_urn"`
	// 短信或邮件的内容。

	MessageContent string `json:"message_content"`
	// 短信或邮件的主题。

	MessageTitle string `json:"message_title"`
}

func (o ActionSmnForwarding) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ActionSmnForwarding struct{}"
	}

	return strings.Join([]string{"ActionSmnForwarding", string(data)}, " ")
}
