package model

import (
	"encoding/json"

	"strings"
)

type QueryCustomerOnDemandResourcesReq struct {
	// 客户账号ID。 您可以调用查询客户列表接口获取customer_id。

	CustomerId string `json:"customer_id"`
	// 云服务区编码，例如：“cn-north-1”。具体请参见地区和终端节点对应云服务的“区域”列的值。

	RegionCode *string `json:"region_code,omitempty"`
	// 云服务类型编码，例如ECS的云服务类型编码为“hws.service.type.ec2”。您可以调用查询云服务类型列表接口获取。

	ServiceTypeCode *string `json:"service_type_code,omitempty"`
	// 资源ID批量查询。 用于查询指定资源ID对应的资源。 最多支持同时传递50个ID的列表。

	ResourceIds *[]string `json:"resource_ids,omitempty"`
	// 生效时间的开始时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。 其中，HH范围是0～23，mm和ss范围是0～59。

	EffectiveTimeBegin *string `json:"effective_time_begin,omitempty"`
	// 生效时间的结束时间 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。 其中，HH范围是0～23，mm和ss范围是0～59。

	EffectiveTimeEnd *string `json:"effective_time_end,omitempty"`
	// 偏移量，从0开始。默认值为0。

	Offset *int32 `json:"offset,omitempty"`
	// 一次查询的条数，默认值为10。

	Limit *int32 `json:"limit,omitempty"`
	// 资源状态： 1：正常（已开通）3：冻结中4：变更中5：正在关闭6：已关闭

	Status *int32 `json:"status,omitempty"`
	// 精英服务商ID，如果需要查询精英服务商子客户的按需资源列表，必须携带该字段，否则只能查询自己的子客户按需资源。

	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
}

func (o QueryCustomerOnDemandResourcesReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QueryCustomerOnDemandResourcesReq struct{}"
	}

	return strings.Join([]string{"QueryCustomerOnDemandResourcesReq", string(data)}, " ")
}
