package dao

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pixstall-commission/app/message/repo/mongo/dao"
	"pixstall-commission/domain/commission/model"
	model2 "pixstall-commission/domain/message/model"
	"time"
)

type Commission struct {
	ObjectID         primitive.ObjectID `bson:"_id,omitempty"`
	model.Commission `bson:",inline"`
	Messages         []dao.Message `bson:"messages"`
}

func NewFromCommissionCreator(d model.CommissionCreator) Commission {
	now := time.Now()
	return Commission{
		ObjectID: primitive.ObjectID{},
		Commission: model.Commission{
			ID:                            "cm-" + uuid.NewString(),
			OpenCommissionID:              d.OpenCommissionID,
			ArtistID:                      d.ArtistID,
			RequesterID:                   d.RequesterID,
			Price:                         d.Price,
			DayNeed:                       d.DayNeed,
			Size:                          d.Size,
			Resolution:                    d.Resolution,
			ExportFormat:                  d.ExportFormat,
			Desc:                          d.Desc,
			PaymentMethod:                 d.PaymentMethod,
			IsR18:                         d.IsR18,
			BePrivate:                     d.BePrivate,
			Anonymous:                     d.Anonymous,
			RefImagePaths:                 d.RefImagePaths,
			CreateTime:                    now,
			CompleteTime:                  nil,
			LastUpdateTime:                now,
			ValidationHistory:             []model.CommissionValidation{},
			State:                         model.CommissionStatePendingValidation,
		},
		Messages: []dao.Message{},
	}
}

func (c Commission) ToDomainCommission() model.Commission {

	var dMessaging []model2.Messaging
	for _, value := range c.Messages {
		dMessaging = append(dMessaging, value.ToDomainMessaging(c.ArtistID, c.RequesterID))
	}

	return model.Commission{
		ID:                            c.ID,
		OpenCommissionID:              c.OpenCommissionID,
		ArtistID:                      c.ArtistID,
		ArtistName:                    c.ArtistName,
		ArtistProfilePath:             c.ArtistProfilePath,
		RequesterID:                   c.RequesterID,
		RequesterName:                 c.RequesterName,
		RequesterProfilePath:          c.RequesterProfilePath,
		Price:                         c.Price,
		DayNeed:                       c.DayNeed,
		Size:                          c.Size,
		Resolution:                    c.Resolution,
		ExportFormat:                  c.ExportFormat,
		Desc:                          c.Desc,
		PaymentMethod:                 c.PaymentMethod,
		IsR18:                         c.IsR18,
		BePrivate:                     c.BePrivate,
		Anonymous:                     c.Anonymous,
		RefImagePaths:                 c.RefImagePaths,
		Messages:                      dMessaging,
		DraftChangingRequestTime:      c.DraftChangingRequestTime,
		CompletionRevisionRequestTime: c.CompletionRevisionRequestTime,
		CreateTime:                    c.CreateTime,
		CompleteTime:                  c.CompleteTime,
		LastUpdateTime:                c.LastUpdateTime,
		ValidationHistory:             c.ValidationHistory,
		State:                         c.State,
	}
}
