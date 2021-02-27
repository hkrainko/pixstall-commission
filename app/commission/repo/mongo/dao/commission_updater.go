package dao

import (
	"go.mongodb.org/mongo-driver/bson"
	"pixstall-commission/domain/commission/model"
)

func NewUpdaterFromCommissionUpdater(d model.CommissionUpdater) bson.D {
	setter := bson.D{}

	if d.ArtistName != nil {
		setter = append(setter, bson.E{Key: "artistName", Value: d.ArtistName})
	}
	if d.ArtistProfilePath != nil {
		setter = append(setter, bson.E{Key: "artistProfilePath", Value: d.ArtistProfilePath})
	}
	if d.RequesterName != nil {
		setter = append(setter, bson.E{Key: "requesterName", Value: d.RequesterName})
	}
	if d.RequesterProfilePath != nil {
		setter = append(setter, bson.E{Key: "requesterProfilePath", Value: d.RequesterProfilePath})
	}
	if d.CompleteTime != nil {
		setter = append(setter, bson.E{Key: "completeTime", Value: d.CompleteTime})
	}
	if d.Validation != nil {
		setter = append(setter, bson.E{Key: "$push", Value: bson.M{"validationHistory": d.Validation}})
	}
	if d.State != nil {
		setter = append(setter, bson.E{Key: "state", Value: d.State})
	}
	return bson.D{{"$set", setter}}
}
