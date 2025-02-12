package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateNodeResponse struct {
	// API类型，固定值“Node”，该值不可修改。

	Kind *string `json:"kind,omitempty"`
	// API版本，固定值“v3”，该值不可修改。

	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *NodeMetadata `json:"metadata,omitempty"`

	Spec *V3NodeSpec `json:"spec,omitempty"`

	Status         *V3NodeStatus `json:"status,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o UpdateNodeResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateNodeResponse struct{}"
	}

	return strings.Join([]string{"UpdateNodeResponse", string(data)}, " ")
}
