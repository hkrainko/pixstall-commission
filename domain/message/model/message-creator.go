package model

import "image"

type MessageCreator struct {
	OpenCommissionID string      `json:"openCommissionId"`
	From             string      `json:"from"`
	MessageType      MessageType `json:"messageType"`
}

type TextMessageCreator struct {
	Text string `json:"text"`
}

type ImageMessageCreator struct {
	Text      *string      `json:"text"`
	Image     *image.Image `json:"image"`
	ImagePath *string      `json:"imagePath"`
}

type SystemMessageCreator struct {
	Text string `json:"text"`
}
