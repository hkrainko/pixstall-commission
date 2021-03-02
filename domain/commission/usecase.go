package commission

import (
	"context"
	"pixstall-commission/domain/commission/model"
)

type UseCase interface {
	AddCommission(ctx context.Context, creator model.CommissionCreator) (*model.Commission, error)
	GetCommissions(ctx context.Context, filter model.CommissionFilter, sorter model.CommissionSorter) (*[]model.Commission, error)
	GetWorks(ctx context.Context, artistID string, filter model.CommissionFilter, sorter model.CommissionSorter) (*[]model.Commission, error)
	UpdateCommissions(ctx context.Context, updater model.CommissionUpdater) error
	OpenCommissionValidation(ctx context.Context, validation model.CommissionOpenCommissionValidation) error
	UsersValidation(ctx context.Context, validation model.CommissionUsersValidation) error
	SendMessage(ctx context.Context, commID string) error
	GetMessages(ctx context.Context, commID string, requesterID string) *model.Commission
}
