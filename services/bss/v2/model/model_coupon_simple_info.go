package model

import (
	"encoding/json"

	"strings"
)

type CouponSimpleInfo struct {
	// 批量发放优惠券成功的客户ID。

	Id string `json:"id"`
	// 发放成功的券ID。

	CouponId string `json:"coupon_id"`
}

func (o CouponSimpleInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CouponSimpleInfo struct{}"
	}

	return strings.Join([]string{"CouponSimpleInfo", string(data)}, " ")
}
