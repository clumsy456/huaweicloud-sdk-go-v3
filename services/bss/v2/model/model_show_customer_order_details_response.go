package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowCustomerOrderDetailsResponse struct {
	// 订单项个数。

	TotalCount *int32 `json:"total_count,omitempty"`

	OrderInfo *CustomerOrderV2 `json:"order_info,omitempty"`
	// 订单对应的订单项。 具体请参见表4。

	OrderLineItems *[]OrderLineItemEntityV2 `json:"order_line_items,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowCustomerOrderDetailsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowCustomerOrderDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowCustomerOrderDetailsResponse", string(data)}, " ")
}
