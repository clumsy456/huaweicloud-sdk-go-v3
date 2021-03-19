package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type GlanceDeleteImageRequest struct {
	ImageId string `json:"image_id"`

	Body *GlanceDeleteImageRequestBody `json:"body,omitempty"`
}

func (o GlanceDeleteImageRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GlanceDeleteImageRequest struct{}"
	}

	return strings.Join([]string{"GlanceDeleteImageRequest", string(data)}, " ")
}
