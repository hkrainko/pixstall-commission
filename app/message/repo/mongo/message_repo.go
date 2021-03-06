package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"pixstall-commission/app/message/repo/mongo/dao"
	dError "pixstall-commission/domain/error"
	"pixstall-commission/domain/message"
	"pixstall-commission/domain/message/model"
	"time"
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

	filter := bson.M{"id": messaging.GetID()}
	change := bson.M{"$push": bson.M{"messages": daoMessage}}

	_, err := m.collection.UpdateOne(ctx, filter, change)
	if err != nil {
		return dError.UnknownError
	}
	return nil
}

func (m mongoMessageRepo) GetMessages(ctx context.Context, commId string, lastMsgTime time.Time, count int) (*model.Message, error) {
	panic("implement me")
}