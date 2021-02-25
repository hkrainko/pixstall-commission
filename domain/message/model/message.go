package model

import "time"

type Message struct {
	ID               string       `json:"id" bson:"id"`
	OpenCommissionID string       `json:"openCommissionId" bson:"openCommissionId"`
	CreateTime       time.Time    `json:"createTime" bson:"createTime"`
	LastUpdatedTime  time.Time    `json:"completeTime" bson:"completeTime"`
	State            MessageState `json:"state" bson:"state"`
	MessageType      MessageType  `json:"messageType" bson:"messageType"`
}

type TextMessage struct {
	Message
	From string `json:"from" bson:"from"`
	To   string `json:"to" bson:"to"`
	Text string `json:"text" bson:"text"`
}

type ImageMessage struct {
	Message
	From      string  `json:"from" bson:"from"`
	To        string  `json:"to" bson:"to"`
	Text      *string `json:"text" bson:"text"`
	ImagePath string  `json:"imagePath" bson:"imagePath"`
}

type SystemMessage struct {
	Message
	Text *string `json:"text" bson:"text"`
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
