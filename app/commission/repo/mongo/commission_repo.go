package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
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
	panic("implement me")
}

func (m mongoCommissionRepo) UpdateCommission(ctx context.Context, updater model.CommissionUpdater) error {
	panic("implement me")
}