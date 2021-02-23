package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	model "pixstall-commission/app/commission/repo/mongo/model"
	"pixstall-commission/domain/commission"
	"pixstall-commission/domain/commission/model"
)

type mongoCommissionRepo struct {
	db         *mongo.Database
	collection *mongo.Collection
}

const (
	CommissionCollection = "Commissions"
	commIDPrefix         = "CM-"
)

func NewMongoCommissionRepo(db *mongo.Database) commission.Repo {
	return &mongoCommissionRepo{
		db:         db,
		collection: db.Collection(CommissionCollection),
	}
}

func (m mongoCommissionRepo) AddCommission(ctx context.Context, creator model.CommissionCreator) (*model.Commission, error) {
	newComm := model.NewFromCommissionCreator(creator)
	result, err := m.collection.InsertOne(ctx, newComm)
	if err != nil {
		fmt.Printf("AddCommission error %v\n", err)
		return nil, err
	}
	fmt.Printf("AddCommission %v", result.InsertedID)
	dComm := newComm.ToDomainCommission()
	return &dComm, nil
}

func (m mongoCommissionRepo) GetCommission(ctx context.Context, commId string) (*model.Commission, error) {
	mongoComm := model.Commission{}
	err :=  m.collection.FindOne(ctx, bson.M{"": ""}).Decode(&mongoComm)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return nil, model.CommissionErrorNotFound
		default:
			return nil, model.CommissionErrorUnknown
		}
	}
	dComm := mongoComm.ToDomainCommission()
	return &dComm, nil
}

func (m mongoCommissionRepo) UpdateCommission(ctx context.Context, updater model.CommissionUpdater) error {
	panic("implement me")
}
