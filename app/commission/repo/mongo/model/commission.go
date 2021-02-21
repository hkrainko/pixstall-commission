package model

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pixstall-commission/domain/commission/model"
	"time"
)

type Commission struct {
	ObjectID         primitive.ObjectID `bson:"_id,omitempty"`
	model.Commission `bson:",inline"`
}

func NewFromCommissionCreator(d model.CommissionCreator) Commission {
	return Commission{
		ObjectID: primitive.ObjectID{},
		Commission: model.Commission{
			ID:               "CM-" + "-" + uuid.NewString(),
			OpenCommissionID: d.OpenCommissionID,
			ArtistID:         d.ArtistID,
			RequesterID:      d.RequesterID,
			Price:            d.Price,
			DayNeed:          d.DayNeed,
			Size: &model.Size{
				Width:  0,
				Height: 0,
				Unit:   "",
			},
			Resolution:     nil,
			ExportFormat:   nil,
			Desc:           "",
			PaymentMethod:  "",
			IsR18:          false,
			BePrivate:      false,
			Anonymous:      false,
			RefImagePaths:  nil,
			CreateTime:     time.Time{},
			CompleteTime:   &time.Time{},
			LastUpdateTime: time.Time{},
			State:          "",
		},
	}
}
