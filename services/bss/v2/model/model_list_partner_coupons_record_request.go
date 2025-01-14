package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListPartnerCouponsRecordRequest struct {
	// 操作类型。 1：发放2：手动回收3：解绑自动回收4：过期回收5：退订回收6：退款充值撤销7：建立关联回收

	OperationTypes *[]string `json:"operation_types,omitempty"`
	// 额度ID。 请从“查询优惠券额度”接口的响应参数中获取。

	QuotaId *string `json:"quota_id,omitempty"`
	// 额度类型： 0：代金券额度1：现金券额度

	QuotaType *int32 `json:"quota_type,omitempty"`
	// 代金券ID列表。 请从“发放优惠券”接口的响应参数中获取。

	CouponIds *[]string `json:"coupon_ids,omitempty"`
	// 客户账号ID。您可以调用查询客户列表接口获取customer_id。

	CustomerId *string `json:"customer_id,omitempty"`
	// 操作时间（开始）。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。 其中，HH范围是0～23，mm和ss范围是0～59。 输入这个条件，会查询出操作时间大于这个时间的记录。

	OperationTimeBegin *string `json:"operation_time_begin,omitempty"`
	// 操作时间（结束）。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。 其中，HH范围是0～23，mm和ss范围是0～59。 输入这个条件，会查询出操作时间小于这个时间的记录。

	OperationTimeEnd *string `json:"operation_time_end,omitempty"`
	// 操作结果： 0：成功-1：失败（非0的记录）

	Result *string `json:"result,omitempty"`
	// 偏移量，从0开始。默认值为0。

	Offset *int32 `json:"offset,omitempty"`
	// 每页记录数。默认值为10。

	Limit *int32 `json:"limit,omitempty"`
	// 精英服务商ID。 华为云伙伴能力中心（一级经销商）查询精英服务商发放给子客户的优惠券发放记录时，需要携带该参数，否则只能查询发给自己子客户的优惠券发放记录。

	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
}

func (o ListPartnerCouponsRecordRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListPartnerCouponsRecordRequest struct{}"
	}

	return strings.Join([]string{"ListPartnerCouponsRecordRequest", string(data)}, " ")
}
