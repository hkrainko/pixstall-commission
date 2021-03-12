package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	dao2 "pixstall-commission/app/commission/repo/mongo/dao"
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

func (m mongoMessageRepo) AddNewMessage(ctx context.Context, userId *string, messaging model.Messaging) error {
	daoMessage := dao.NewFromMessaging(messaging)

	filter := m.getAddNewMessageFilter(messaging, userId)

	change := bson.M{"$push": bson.M{"messages": daoMessage}}
	_, err := m.collection.UpdateOne(ctx, filter, change)
	if err != nil {
		return dError.UnknownError
	}
	return nil
}

func (m mongoMessageRepo) GetMessages(ctx context.Context, userId string, commId string, offset int, count int) ([]model.Messaging, error) {
	filter := bson.M{"id": commId}
	opts := options.FindOneOptions{
		Projection: bson.M{
			"id":          1,
			"artistId":    1,
			"requesterId": 1,
			"messages":    bson.M{"$slice": []int{offset, count}},
			"state":       1,
		},
	}

	var comm dao2.Commission
	err := m.collection.FindOne(ctx, filter, &opts).Decode(&comm)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return nil, dError.NotFoundError
		default:
			return nil, dError.UnknownError
		}
	}
	var dMessaging []model.Messaging
	for _, daoMsg := range comm.Messages {
		dMessaging = append(dMessaging, daoMsg.ToDomainMessaging(comm.ArtistID, comm.RequesterID))
	}
	return dMessaging, nil
}

// Private
func (m mongoMessageRepo) getAddNewMessageFilter(messaging model.Messaging, userId *string) bson.M {
	switch messaging.(type) {
	case model.SystemMessage:
		return bson.M{"id": messaging.GetCommissionID()}
	default:
		return bson.M{
			"$and": []bson.M{
				{"id": messaging.GetCommissionID()},
				{"$or": []bson.M{
					{"artistId": userId},
					{"requesterId": userId},
				}},
			},
		}
	}
}
