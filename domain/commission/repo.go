package commission

import (
	"context"
	"pixstall-commission/domain/commission/model"
)

type Repo interface {
	AddCommission(ctx context.Context, creator model.CommissionCreator) (*string, error)
	UpdateCommission(ctx context.Context, updater model.CommissionUpdater) error
}