package dao

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pixstall-commission/domain/message/model"
)

type Message struct {
	ObjectID        primitive.ObjectID `bson:"_id,omitempty"`
	model.Messaging `bson:",inline"`
}

func NewFromMessaging(d model.Messaging) *Message {
	return &Message{
		ObjectID:  primitive.ObjectID{},
		Messaging: d,
	}
}
