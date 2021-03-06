package model

import "image"

type MessageCreator struct {
	CommissionID string       `json:"commissionId"`
	Form         *string      `json:"from"` //system message if empty
	Text         string       `json:"text"`
	Image        *image.Image `json:"image"`
	ImagePath    *string      `json:"imagePath"`
}
