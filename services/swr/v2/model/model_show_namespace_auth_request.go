package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowNamespaceAuthRequest struct {
	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json

	ContentType ShowNamespaceAuthRequestContentType `json:"Content-Type"`
	// 组织名称

	Namespace string `json:"namespace"`
}

func (o ShowNamespaceAuthRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowNamespaceAuthRequest struct{}"
	}

	return strings.Join([]string{"ShowNamespaceAuthRequest", string(data)}, " ")
}

type ShowNamespaceAuthRequestContentType struct {
	value string
}

type ShowNamespaceAuthRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 ShowNamespaceAuthRequestContentType
	APPLICATION_JSON             ShowNamespaceAuthRequestContentType
}

func GetShowNamespaceAuthRequestContentTypeEnum() ShowNamespaceAuthRequestContentTypeEnum {
	return ShowNamespaceAuthRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: ShowNamespaceAuthRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: ShowNamespaceAuthRequestContentType{
			value: "application/json",
		},
	}
}

func (c ShowNamespaceAuthRequestContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ShowNamespaceAuthRequestContentType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
