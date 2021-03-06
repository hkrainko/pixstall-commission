package commission

import (
	"context"
	"pixstall-commission/domain/commission/model"
	dMsgModel "pixstall-commission/domain/message/model"
)

type UseCase interface {
	AddCommission(ctx context.Context, creator model.CommissionCreator) (*model.Commission, error)
	GetCommissions(ctx context.Context, filter model.CommissionFilter, sorter model.CommissionSorter) (*[]model.Commission, error)
	GetWorks(ctx context.Context, artistID string, filter model.CommissionFilter, sorter model.CommissionSorter) (*[]model.Commission, error)
	UpdateCommissions(ctx context.Context, updater model.CommissionUpdater) error
	OpenCommissionValidation(ctx context.Context, validation model.CommissionOpenCommissionValidation) error
	UsersValidation(ctx context.Context, validation model.CommissionUsersValidation) error
	HandleInboundCommissionMessage(ctx context.Context, msgCreator dMsgModel.MessageCreator) error
	HandleOutBoundCommissionMessage(ctx context.Context, message dMsgModel.Messaging) error
}
