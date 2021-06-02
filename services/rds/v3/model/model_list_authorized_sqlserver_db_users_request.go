package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListAuthorizedSqlserverDbUsersRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID。

	InstanceId string `json:"instance_id"`
	// 数据库名。

	DbName string `json:"db-name"`
	// 分页页码，从1开始。

	Page int32 `json:"page"`
	// 每页数据条数。取值范围[1, 100]。

	Limit int32 `json:"limit"`
}

func (o ListAuthorizedSqlserverDbUsersRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAuthorizedSqlserverDbUsersRequest struct{}"
	}

	return strings.Join([]string{"ListAuthorizedSqlserverDbUsersRequest", string(data)}, " ")
}