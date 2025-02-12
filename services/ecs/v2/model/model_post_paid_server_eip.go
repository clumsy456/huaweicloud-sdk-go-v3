package model

import (
	"encoding/json"

	"strings"
)

//
type PostPaidServerEip struct {
	// 弹性IP地址类型。

	Iptype string `json:"iptype"`

	Bandwidth *PostPaidServerEipBandwidth `json:"bandwidth"`

	Extendparam *PostPaidServerEipExtendParam `json:"extendparam,omitempty"`
}

func (o PostPaidServerEip) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PostPaidServerEip struct{}"
	}

	return strings.Join([]string{"PostPaidServerEip", string(data)}, " ")
}
