package model

import "image"

type MessageCreator struct {
	CommissionID string       `json:"commissionId"`
	Text         string       `json:"text"`
	Image        *image.Image `json:"image"`
	ImagePath    *string      `json:"imagePath"`
}
