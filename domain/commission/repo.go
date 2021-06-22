package commission

import (
	"context"
	"pixstall-commission/domain/commission/model"
	model2 "pixstall-commission/domain/user/model"
)

type Repo interface {
	AddCommission(ctx context.Context, creator model.CommissionCreator) (*model.Commission, error)
	GetCommission(ctx context.Context, commId string) (*model.Commission, error)
	GetCommissions(ctx context.Context, filter model.CommissionFilter, sorter model.CommissionSorter) (*model.GetCommissionsResult, error)
	UpdateCommission(ctx context.Context, updater model.CommissionUpdater) error
	UpdateCommissionUser(ctx context.Context, updater model2.UserUpdater) error
}
