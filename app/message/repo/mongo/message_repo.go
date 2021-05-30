package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func (m mongoMessageRepo) GetMessages(ctx context.Context, userId string, filter model.MessageFilter) ([]model.Messaging, error) {
	var pipeline []bson.M
	pipeline = append(pipeline, bson.M{
		"$match": bson.M{
			"$and": bson.A{
				bson.M{"id": filter.CommissionID},
				bson.M{"$or": bson.A{
					bson.M{"artistId": userId},
					bson.M{"requesterId": userId},
				}},
			},
		},
	})
	if filter.LastMessageID != nil {
		pipeline = append(pipeline, bson.M{
			"$project": bson.M{
				"id": 1,
				"artistId": 1,
				"requesterId": 1,
				"messages": bson.M{
					"$slice": bson.A {
						"$messages",
						bson.M {
							"$subtract": bson.A{
								bson.M{
									"$indexOfArray": bson.A{
										"$messages",
										bson.M{"id": *filter.LastMessageID},
									},// TODO: bug here
								},
								filter.Count * -1,
							},
						},
						filter.Count,
					},
				},
				"state": 1,
			}})
	} else {
		pipeline = append(pipeline, bson.M{
			"$project": bson.M{
				"id": 1,
				"artistId": 1,
				"requesterId": 1,
				"messages": bson.M{
					"$slice": bson.A {
						"$messages",
						filter.Count * -1,
					},
				},
				"state": 1,
			}})
	}

	cursor, err := m.collection.Aggregate(ctx, pipeline)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return nil, dError.NotFoundError
		default:
			return nil, dError.UnknownError
		}
	}
	defer cursor.Close(ctx)
	var comm dao2.Commission
	for cursor.Next(ctx) {
		if err := cursor.Decode(&comm); err != nil {
			return nil, err
		}
	}

	dMessagings := make([]model.Messaging, 0)
	for _, daoMsg := range comm.Messages {
		dMessagings = append(dMessagings, daoMsg.ToDomainMessaging(comm.ArtistID, comm.RequesterID))
	}
	return dMessagings, nil
}

// Private
func (m mongoMessageRepo) getAddNewMessageFilter(messaging model.Messaging, userId *string) bson.M {
	switch messaging.GetMessageType() {
	case model.MessageTypeSystem:
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
