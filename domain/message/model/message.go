package model

import "time"

type Message struct {
	ID               string
	OpenCommissionID string
	CreateTime       time.Time    `json:"createTime"`
	LastUpdatedTime  time.Time    `json:"completeTime"`
	State            MessageState `json:"state"`
}

type TextMessage struct {
	Message
	text string
}

type SystemMessage struct {
	Message
	text string
}

type ImageMessage struct {
	Message
	imagePath string
}

type MessageState string

const (
	MessageStateNormal MessageState = "N"
)
