package model

import "time"

type Messaging interface {
	GetID() string
	GetArtistID() string
	GetRequesterID() string
	GetOpenCommissionID() string
	GetCreateTime() time.Time
	GetLastUpdatedTime() time.Time
	GetState() MessageState
	GetMessageType() MessageType
}

type Message struct {
	ID               string       `json:"id" bson:"id"`
	ArtistID         string       `json:"artistId" bson:"artistId"`
	RequesterID      string       `json:"requesterId" bson:"requesterId"`
	OpenCommissionID string       `json:"openCommissionId" bson:"openCommissionId"`
	CreateTime       time.Time    `json:"createTime" bson:"createTime"`
	LastUpdatedTime  time.Time    `json:"completeTime" bson:"completeTime"`
	State            MessageState `json:"state" bson:"state"`
	MessageType      MessageType  `json:"messageType" bson:"messageType"`
}

func (m Message) GetID() string {
	return m.ID
}

func (m Message) GetArtistID() string {
	return m.ArtistID
}

func (m Message) GetRequesterID() string {
	return m.RequesterID
}

func (m Message) GetOpenCommissionID() string {
	return m.OpenCommissionID
}

func (m Message) GetCreateTime() time.Time {
	return m.CreateTime
}

func (m Message) GetLastUpdatedTime() time.Time {
	return m.LastUpdatedTime
}

func (m Message) GetState() MessageState {
	return m.State
}

func (m Message) GetMessageType() MessageType {
	return m.MessageType
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
