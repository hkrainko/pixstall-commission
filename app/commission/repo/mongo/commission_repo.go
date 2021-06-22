package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pixstall-commission/app/commission/repo/mongo/dao"
	"pixstall-commission/domain/commission"
	dModel "pixstall-commission/domain/commission/model"
	error2 "pixstall-commission/domain/error"
	model2 "pixstall-commission/domain/user/model"
)

type mongoCommissionRepo struct {
	db         *mongo.Database
	collection *mongo.Collection
}

const (
	CommissionCollection = "Commissions"
)

func NewMongoCommissionRepo(db *mongo.Database) commission.Repo {
	return &mongoCommissionRepo{
		db:         db,
		collection: db.Collection(CommissionCollection),
	}
}

func (m mongoCommissionRepo) AddCommission(ctx context.Context, creator dModel.CommissionCreator) (*dModel.Commission, error) {
	newComm := dao.NewFromCommissionCreator(creator)
	result, err := m.collection.InsertOne(ctx, newComm)
	if err != nil {
		fmt.Printf("AddCommission error %v\n", err)
		return nil, err
	}
	fmt.Printf("AddCommission %v", result.InsertedID)
	dComm := newComm.ToDomainCommission()
	return &dComm, nil
}

func (m mongoCommissionRepo) GetCommission(ctx context.Context, commId string) (*dModel.Commission, error) {
	mongoComm := dao.Commission{}
	opts := options.FindOneOptions{
		Projection: bson.M{
			"messages": bson.M{"$slice": -20},
		},
	}

	err := m.collection.FindOne(ctx, bson.M{"id": commId}, &opts).Decode(&mongoComm)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return nil, error2.NotFoundError
		default:
			return nil, error2.UnknownError
		}
	}
	dComm := mongoComm.ToDomainCommission()
	return &dComm, nil
}

func (m mongoCommissionRepo) GetCommissions(ctx context.Context, filter dModel.CommissionFilter, sorter dModel.CommissionSorter) (*dModel.GetCommissionsResult, error) {
	//daoFilter := dao.NewFilterFromDomainCommissionFilter(filter) // TODO: Change to aggregate filter

	var pipeline []bson.M
	if filter.ArtistID != nil {
		//For querying received commissions
		pipeline = append(pipeline, bson.M{"$match": bson.M{"artistId": filter.ArtistID}})
	} else if filter.RequesterID != nil {
		// For querying submitted commissions
		pipeline = append(pipeline, bson.M{"$match": bson.M{"requesterId": filter.RequesterID}})
	} else {
		return nil, error2.UnknownError
	}
	pipeline = append(pipeline, bson.M{
		"$facet": bson.M{
			"total": []bson.M{{
				"$count": "total",
			}},
			"commissions": bson.A{
				bson.D{{"$skip", filter.Offset}},
				bson.D{{"$limit", filter.Count}},
			},
		},
	})
	pipeline = append(pipeline, bson.M{
		"$addFields": bson.M{
			"total": bson.M{"$arrayElemAt": bson.A{"$total.total", 0}},
		},
	})

	cursor, err := m.collection.Aggregate(ctx, pipeline)
	defer cursor.Close(ctx)
	if err != nil {
		return nil, dModel.CommissionErrorUnknown
	}
	var getCommsResult *dModel.GetCommissionsResult
	for cursor.Next(ctx) {
		var r dao.GetCommissionsResult
		if err := cursor.Decode(&r); err != nil {
			return nil, err
		}
		getCommsResult = r.ToDomainGetCommissionsResult()
	}
	return getCommsResult, nil
}

func (m mongoCommissionRepo) UpdateCommission(ctx context.Context, commUpdater dModel.CommissionUpdater) error {
	filter := bson.M{
		"id": commUpdater.ID,
	}
	updater := dao.NewUpdaterFromCommissionUpdater(commUpdater)
	_, err := m.collection.UpdateOne(ctx, filter, updater)
	if err != nil {
		return dModel.CommissionErrorUnknown
	}
	return nil
}

func (m mongoCommissionRepo) UpdateCommissionUser(ctx context.Context, updater model2.UserUpdater) error {
	filter := bson.M{
		"$or": bson.A{bson.M{"artistId": updater.UserID}, bson.M{"requesterId": updater.UserID}},
	}
	update := bson.M{
		"$cond": getUpdateCommissionUserCondition(updater),
	}
	result, err := m.collection.UpdateMany(ctx, filter, update)
	if err != nil {
		fmt.Printf("%v\n", err)
		return err
	}
	fmt.Printf("UpdateCommissionUser count:%v\n", result.MatchedCount)
	return nil
}

func getUpdateCommissionUserCondition(updater model2.UserUpdater) bson.M {
	result := bson.M{}
	result["if"] = bson.M{"artistId": updater.UserID}

	if updater.UserName != nil {
		result["then"] = bson.M{"$set": bson.M{"artistName": updater.UserName}}
		result["else"] = bson.M{"$set": bson.M{"requesterName": updater.UserName}}
	}
	if updater.ProfilePath != nil {
		result["then"] = bson.M{"$set": bson.M{"artistProfilePath": updater.ProfilePath}}
		result["else"] = bson.M{"$set": bson.M{"requesterProfilePath": updater.ProfilePath}}
	}
	return result
}