package model

import (
	"encoding/json"

	"strings"
)

type QueryResourcesReq struct {
	// 资源ID列表。 查询指定资源ID的资源（当only_main_resource=0时，查询指定资源及其附属资源）。最大支持50个ID同时作为条件查询，多个ID以英文逗号分隔。  说明： 资源ID是指开通资源以后，云服务针对该资源分配的标志，譬如云主机ECS的资源ID是server_id。

	ResourceIds *[]string `json:"resource_ids,omitempty"`
	// 订单号。 查询指定订单下的资源。

	OrderId *string `json:"order_id,omitempty"`
	// 是否只查询主资源，该参数对于请求参数是子资源ID的时候无效，如果resource_ids是子资源ID，只能查询自己。 0：查询主资源及附属资源。1：只查询主资源。 默认值为0。  说明： 主资源是指有关联的几个资源中，处于主导位置的资源。 对于ECS而言，虚拟机VM是主资源，磁盘EVS是辅资源。对于VPC而言，共享带宽的情况下，带宽为主资源，对应的从资源为弹性IP（可能包含多个IP）；独享带宽的情况下，弹性IP为主资源，对应的从资源为带宽。

	OnlyMainResource *int32 `json:"only_main_resource,omitempty"`
	// 资源状态。 查询指定状态的资源。多个状态以英文逗号分隔。 2：使用中4：已冻结5：已过期

	StatusList *[]int32 `json:"status_list,omitempty"`
	// 偏移量，从0开始。默认值为0。

	Offset *int32 `json:"offset,omitempty"`
	// 每次查询的条数。默认值为10。

	Limit *int32 `json:"limit,omitempty"`
}

func (o QueryResourcesReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QueryResourcesReq struct{}"
	}

	return strings.Join([]string{"QueryResourcesReq", string(data)}, " ")
}
