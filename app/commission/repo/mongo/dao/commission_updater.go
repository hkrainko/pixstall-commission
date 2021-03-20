package dao

import (
	"go.mongodb.org/mongo-driver/bson"
	"pixstall-commission/domain/commission/model"
)

func NewUpdaterFromCommissionUpdater(d model.CommissionUpdater) bson.D {
	updater := bson.D{}

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
	if d.CompletionRevisionRequestTime != nil {
		setter = append(setter, bson.E{Key: "completionRevisionRequestTime", Value: d.CompletionRevisionRequestTime})
	}
	if d.CompleteTime != nil {
		setter = append(setter, bson.E{Key: "completeTime", Value: d.CompleteTime})
	}
	if d.State != nil {
		setter = append(setter, bson.E{Key: "state", Value: d.State})
	}
	if d.DisplayImagePath != nil {
		setter = append(setter, bson.E{Key: "displayImagePath", Value: d.DisplayImagePath})
	}
	if d.CompletionFilePath != nil {
		setter = append(setter, bson.E{Key: "completionFilePath", Value: d.CompletionFilePath})
	}
	if d.Rating != nil {
		setter = append(setter, bson.E{Key: "rating", Value: d.Rating})
	}
	if d.Comment != nil {
		setter = append(setter, bson.E{Key: "comment", Value: d.Comment})
	}

	putter := bson.D{}
	if d.Validation != nil {
		putter = append(putter, bson.E{Key: "validationHistory", Value: d.Validation})
	}
	if d.ProofCopyImagePath != nil {
		putter = append(putter, bson.E{Key: "proofCopyImagePaths", Value: d.ProofCopyImagePath})
	}

	if len(setter) > 0 {
		updater = append(updater, bson.E{Key: "$set", Value: setter})
	}
	if len(putter) > 0 {
		updater = append(updater, bson.E{Key: "$push", Value: putter})
	}

	return updater
}
