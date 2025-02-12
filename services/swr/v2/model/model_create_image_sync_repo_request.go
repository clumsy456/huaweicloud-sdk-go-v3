package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type CreateImageSyncRepoRequest struct {
	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json

	ContentType CreateImageSyncRepoRequestContentType `json:"Content-Type"`
	// 组织名称

	Namespace string `json:"namespace"`
	// 镜像仓库名称

	Repository string `json:"repository"`

	Body *CreateImageSyncRepoRequestBody `json:"body,omitempty"`
}

func (o CreateImageSyncRepoRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateImageSyncRepoRequest struct{}"
	}

	return strings.Join([]string{"CreateImageSyncRepoRequest", string(data)}, " ")
}

type CreateImageSyncRepoRequestContentType struct {
	value string
}

type CreateImageSyncRepoRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 CreateImageSyncRepoRequestContentType
	APPLICATION_JSON             CreateImageSyncRepoRequestContentType
}

func GetCreateImageSyncRepoRequestContentTypeEnum() CreateImageSyncRepoRequestContentTypeEnum {
	return CreateImageSyncRepoRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: CreateImageSyncRepoRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: CreateImageSyncRepoRequestContentType{
			value: "application/json",
		},
	}
}

func (c CreateImageSyncRepoRequestContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CreateImageSyncRepoRequestContentType) UnmarshalJSON(b []byte) error {
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
