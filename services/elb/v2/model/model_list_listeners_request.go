package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListListenersRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	PageReverse *bool `json:"page_reverse,omitempty"`

	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	DefaultPoolId *string `json:"default_pool_id,omitempty"`

	DefaultTlsContainerRef *string `json:"default_tls_container_ref,omitempty"`

	ClientCaTlsContainerRef *string `json:"client_ca_tls_container_ref,omitempty"`

	Protocol *string `json:"protocol,omitempty"`

	ProtocolPort *int32 `json:"protocol_port,omitempty"`

	TlsCiphersPolicy *string `json:"tls_ciphers_policy,omitempty"`

	MemberTimeout *int32 `json:"member_timeout,omitempty"`

	ClientTimeout *int32 `json:"client_timeout,omitempty"`

	KeepaliveTimeout *int32 `json:"keepalive_timeout,omitempty"`

	TlsContainerId *string `json:"tls_container_id,omitempty"`
}

func (o ListListenersRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListListenersRequest struct{}"
	}

	return strings.Join([]string{"ListListenersRequest", string(data)}, " ")
}
