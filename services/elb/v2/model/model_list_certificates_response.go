/*
 * ELB
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListCertificatesResponse struct {
	Certificates *CertificateResp `json:"certificates,omitempty"`
	// 证书的个数
	InstanceNum *int32 `json:"instance_num,omitempty"`
}

func (o ListCertificatesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListCertificatesResponse", string(data)}, " ")
}
