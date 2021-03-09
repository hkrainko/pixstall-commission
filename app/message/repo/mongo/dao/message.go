package dao

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pixstall-commission/domain/message/model"
	"time"
)

type Message struct {
	ObjectID        primitive.ObjectID `bson:"_id,omitempty"`
	ID              string             `bson:"id"`
	ArtistID        string             `bson:"artistId"`
	RequesterID     string             `bson:"requesterId"`
	CommissionID    string             `bson:"commissionId"`
	CreateTime      time.Time          `bson:"createTime"`
	LastUpdatedTime time.Time          `bson:"completeTime"`
	State           model.MessageState `bson:"state"`
	MessageType     model.MessageType  `bson:"messageType"`

	From      *string `bson:"from,omitempty"`
	Text      *string `bson:"text,omitempty"`
	ImagePath *string `bson:"imagePath,omitempty"`
}

func NewFromMessaging(d model.Messaging) *Message {

	msg := Message{
		ObjectID:        primitive.ObjectID{},
		ID:              d.GetID(),
		ArtistID:        d.GetArtistID(),
		RequesterID:     d.GetRequesterID(),
		CommissionID:    d.GetCommissionID(),
		CreateTime:      d.GetCreateTime(),
		LastUpdatedTime: d.GetLastUpdatedTime(),
		State:           d.GetState(),
		MessageType:     d.GetMessageType(),
		From:            nil,
		Text:            nil,
		ImagePath:       nil,
	}
	switch v := d.(type) {
	case model.TextMessage:
		msg.From = &v.From
		msg.Text = &v.Text
	case model.ImageMessage:
		msg.From = &v.From
		msg.ImagePath = &v.ImagePath
	case model.SystemMessage:
		msg.Text = &v.Text
	}
	return &msg
}

func (m *Message) ToDomainMessaging() model.Messaging {

	msg := model.Message{
		ID:              m.ID,
		ArtistID:        m.ArtistID,
		RequesterID:     m.RequesterID,
		CommissionID:    m.CommissionID,
		CreateTime:      m.CreateTime,
		LastUpdatedTime: m.LastUpdatedTime,
		State:           m.State,
		MessageType:     m.MessageType,
	}

	switch m.MessageType {
	case model.MessageTypeText:
		return &model.TextMessage{
			Message: msg,
			From:    *m.From,
			Text:    *m.Text,
		}
	case model.MessageTypeImage:
		return &model.ImageMessage{
			Message:   msg,
			From:      *m.From,
			Text:      m.Text,
			ImagePath: *m.ImagePath,
		}
	case model.MessageTypeSystem:
		return &model.SystemMessage{
			Message: msg,
			Text:    *m.Text,
		}
	default:
		return msg
	}
}
