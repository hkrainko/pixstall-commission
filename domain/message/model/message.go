package model

import "time"

type Message struct {
	ID               string
	OpenCommissionID string       `json:""`
	CreateTime       time.Time    `json:"createTime"`
	LastUpdatedTime  time.Time    `json:"completeTime"`
	State            MessageState `json:"state"`
	MessageType      MessageType  `json:"messageType"`
}

type TextMessage struct {
	Message
	From string `json:"from"`
	To   string `json:"to"`
	Text string `json:"text"`
}

type ImageMessage struct {
	Message
	From      string  `json:"from"`
	To        string  `json:"to"`
	Text      *string `json:"text"`
	ImagePath string  `json:"imagePath"`
}

type SystemMessage struct {
	Message
	Text      *string `json:"text"`
}

type MessageState string

const (
	MessageStateNormal MessageState = "N"
)

type MessageType string

const (
	MessageTypeText   MessageType = "Text"
	MessageTypeImage  MessageType = "Image"
	MessageTypeSystem MessageType = "System"
)
