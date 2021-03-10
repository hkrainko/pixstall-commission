package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pixstall-commission/app/message/repo/mongo/dao"
	dError "pixstall-commission/domain/error"
	"pixstall-commission/domain/message"
	"pixstall-commission/domain/message/model"
)

type mongoMessageRepo struct {
	db         *mongo.Database
	collection *mongo.Collection
}

const (
	CommissionCollection = "Commissions"
)

func NewMongoMessageRepo(db *mongo.Database) message.Repo {
	return &mongoMessageRepo{
		db:         db,
		collection: db.Collection(CommissionCollection),
	}
}

func (m mongoMessageRepo) AddNewMessage(ctx context.Context, messaging model.Messaging) error {
	daoMessage := dao.NewFromMessaging(messaging)

	filter := bson.M{"id": messaging.GetCommissionID()}
	change := bson.M{"$push": bson.M{"messages": daoMessage}}

	_, err := m.collection.UpdateOne(ctx, filter, change)
	if err != nil {
		return dError.UnknownError
	}
	return nil
}

func (m mongoMessageRepo) GetMessages(ctx context.Context, commId string, offset int, count int) ([]model.Messaging, error) {
	filter := bson.M{"id": commId}
	opts := options.FindOptions{
		AllowDiskUse:        nil,
		AllowPartialResults: nil,
		BatchSize:           nil,
		Collation:           nil,
		Comment:             nil,
		CursorType:          nil,
		Hint:                nil,
		Limit:               nil,
		Max:                 nil,
		MaxAwaitTime:        nil,
		MaxTime:             nil,
		Min:                 nil,
		NoCursorTimeout:     nil,
		OplogReplay:         nil,
		Projection:          bson.M{
			"$slice": []int{offset, count},
		},
		ReturnKey:           nil,
		ShowRecordID:        nil,
		Skip:                nil,
		Snapshot:            nil,
		Sort:                nil,
	}

	cursor, err := m.collection.Find(ctx, filter, &opts)
	if err != nil {
		return nil, err
	}
	var results []dao.Message
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	var dMessaging []model.Messaging
	for _, result := range results {
		dMessaging = append(dMessaging, result.ToDomainMessaging("", ""))
	}
	return dMessaging, nil
}