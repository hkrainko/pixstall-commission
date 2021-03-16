package model

import "image"

type MessageCreator struct {
	CommissionID string       `json:"commissionId"`
	Form         string      `json:"from"`
	Text         string       `json:"text"`
	Image        *image.Image `json:"image"`
	ImagePath    *string      `json:"imagePath"`
}
