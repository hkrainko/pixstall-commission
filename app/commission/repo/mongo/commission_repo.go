package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	model2 "pixstall-commission/app/commission/repo/mongo/model"
	"pixstall-commission/domain/commission"
	"pixstall-commission/domain/commission/model"
)

type mongoCommissionRepo struct {
	db         *mongo.Database
	collection *mongo.Collection
}

const (
	CommissionCollection = "Commissions"
	commIDPrefix = "CM-"
)

func NewMongoCommissionRepo(db *mongo.Database) commission.Repo {
	return &mongoCommissionRepo{
		db:         db,
		collection: db.Collection(CommissionCollection),
	}
}

func (m mongoCommissionRepo) AddCommission(ctx context.Context, creator model.CommissionCreator) (*string, error) {
	newComm := model2.NewFromCommissionCreator(creator)
	result, err := m.collection.InsertOne(ctx, newComm)
	if err != nil {
		fmt.Printf("AddCommission error %v\n", err)
		return nil, err
	}
	fmt.Printf("AddCommission %v", result.InsertedID)
	return &newComm.ID, nil
}

func (m mongoCommissionRepo) UpdateCommission(ctx context.Context, updater model.CommissionUpdater) error {
	panic("implement me")
}