package dao

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
	now := time.Now()
	return Commission{
		ObjectID: primitive.ObjectID{},
		Commission: model.Commission{
			ID:                "CM-" + uuid.NewString(),
			OpenCommissionID:  d.OpenCommissionID,
			ArtistID:          d.ArtistID,
			RequesterID:       d.RequesterID,
			Price:             d.Price,
			DayNeed:           d.DayNeed,
			Size:              d.Size,
			Resolution:        d.Resolution,
			ExportFormat:      d.ExportFormat,
			Desc:              d.Desc,
			PaymentMethod:     d.PaymentMethod,
			IsR18:             d.IsR18,
			BePrivate:         d.BePrivate,
			Anonymous:         d.Anonymous,
			RefImagePaths:     d.RefImagePaths,
			CreateTime:        now,
			CompleteTime:      nil,
			LastUpdateTime:    now,
			ValidationHistory: []model.CommissionValidation{},
			State:             model.CommissionStatePendingValidation,
		},
	}
}

func (c Commission) ToDomainCommission() model.Commission {
	return model.Commission{
		ID:                   c.ID,
		OpenCommissionID:     c.OpenCommissionID,
		ArtistID:             c.ArtistID,
		ArtistName:           c.ArtistName,
		ArtistProfilePath:    c.ArtistProfilePath,
		RequesterID:          c.RequesterID,
		RequesterName:        c.RequesterName,
		RequesterProfilePath: c.RequesterProfilePath,
		Price:                c.Price,
		DayNeed:              c.DayNeed,
		Size:                 c.Size,
		Resolution:           c.Resolution,
		ExportFormat:         c.ExportFormat,
		Desc:                 c.Desc,
		PaymentMethod:        c.PaymentMethod,
		IsR18:                c.IsR18,
		BePrivate:            c.BePrivate,
		Anonymous:            c.Anonymous,
		RefImagePaths:        c.RefImagePaths,
		CreateTime:           c.CreateTime,
		CompleteTime:         c.CompleteTime,
		LastUpdateTime:       c.LastUpdateTime,
		ValidationHistory:    c.ValidationHistory,
		State:                c.State,
	}
}
