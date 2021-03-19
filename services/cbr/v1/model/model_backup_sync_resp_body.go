package model

import (
	"encoding/json"

	"strings"
)

type BackupSyncRespBody struct {
	// 备份副本ID

	BackupId string `json:"backup_id"`
	// 同步任务ID

	OperationLogId string `json:"operation_log_id"`
}

func (o BackupSyncRespBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BackupSyncRespBody struct{}"
	}

	return strings.Join([]string{"BackupSyncRespBody", string(data)}, " ")
}
