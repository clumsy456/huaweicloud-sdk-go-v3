package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ShowNamespace struct {
	// id

	Id int32 `json:"id"`
	// 组织名称

	Name string `json:"name"`
	// IAM用户名

	CreatorName string `json:"creator_name"`
	// 用户权限。7表示管理权限，3表示编辑权限，1表示读取权限。

	Auth ShowNamespaceAuth `json:"auth"`
}

func (o ShowNamespace) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowNamespace struct{}"
	}

	return strings.Join([]string{"ShowNamespace", string(data)}, " ")
}

type ShowNamespaceAuth struct {
	value int32
}

type ShowNamespaceAuthEnum struct {
	E_7 ShowNamespaceAuth
	E_3 ShowNamespaceAuth
	E_1 ShowNamespaceAuth
}

func GetShowNamespaceAuthEnum() ShowNamespaceAuthEnum {
	return ShowNamespaceAuthEnum{
		E_7: ShowNamespaceAuth{
			value: 7,
		}, E_3: ShowNamespaceAuth{
			value: 3,
		}, E_1: ShowNamespaceAuth{
			value: 1,
		},
	}
}

func (c ShowNamespaceAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ShowNamespaceAuth) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
