package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListPublicZonesResponse struct {
	Links *PageLink `json:"links,omitempty"`

	Zones *string `json:"zones,omitempty"`

	Metadata       *Metedata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListPublicZonesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListPublicZonesResponse struct{}"
	}

	return strings.Join([]string{"ListPublicZonesResponse", string(data)}, " ")
}
