package model

import (
	"encoding/json"

	"strings"
)

type PeriodProductInfo struct {
	// ID标识，同一次询价中不能重复，用于标识返回询价结果和请求的映射关系。

	Id string `json:"id"`
	// 云服务类型编码，例如ECS的云服务类型编码为“hws.service.type.ec2”。您可以调用查询云服务类型列表接口获取。

	CloudServiceType string `json:"cloud_service_type"`
	// 资源类型编码，例如ECS的VM为“hws.resource.type.vm”。您可以调用查询资源类型列表接口获取。 ResourceType是CloudServiceType中的一种资源，CloudServiceType由多种ResourceType组合提供。

	ResourceType string `json:"resource_type"`
	// 云服务产品的资源规格。如果是VM的资源规格，则需要在规格后面添加“.win”或“.linux”，例如“s2.small.1.linux”。具体请参见对应云服务的相关介绍。

	ResourceSpec string `json:"resource_spec"`
	// 云服务区编码，例如：“cn-north-1”。具体请参见地区和终端节点对应云服务的“区域”列的值。

	Region string `json:"region"`
	// 可用区标识。

	AvailableZone *string `json:"available_zone,omitempty"`
	// 资源容量大小，例如购买的卷大小或带宽大小。 线性产品时该参数不能为空。

	ResourceSize *int32 `json:"resource_size,omitempty"`
	// 资源容量度量标识。 15：Mbps（购买带宽时使用）17：GB（购买云硬盘时使用）14：个 线性产品时该参数不能为空。

	SizeMeasureId *int32 `json:"size_measure_id,omitempty"`
	// 订购包年/包月产品的周期类型。 0：天2：月3：年4：小时

	PeriodType int32 `json:"period_type"`
	// 订购包年/包月产品的周期数。

	PeriodNum int32 `json:"period_num"`
	// 订购包年/包月产品的数量。

	SubscriptionNum int32 `json:"subscription_num"`
}

func (o PeriodProductInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PeriodProductInfo struct{}"
	}

	return strings.Join([]string{"PeriodProductInfo", string(data)}, " ")
}
