package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
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

func (m mongoMessageRepo) AddNewMessage(ctx context.Context, creator model.MessageCreator) (*model.Message, error) {
	panic("implement me")
}

func (m mongoMessageRepo) GetMessage(ctx context.Context, commId string, lastMsgTime time.Time, count int) (*model.Message, error) {
	panic("implement me")
}