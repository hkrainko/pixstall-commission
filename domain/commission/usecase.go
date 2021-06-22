package commission

import (
	"context"
	"pixstall-commission/domain/commission/model"
	dMsgModel "pixstall-commission/domain/message/model"
	model2 "pixstall-commission/domain/user/model"
)

type UseCase interface {
	AddCommission(ctx context.Context, creator model.CommissionCreator) (*model.Commission, error)
	GetCommissions(ctx context.Context, filter model.CommissionFilter, sorter model.CommissionSorter) (*model.GetCommissionsResult, error)
	GetCommission(ctx context.Context, commId string) (*model.Commission, error)
	GetWorks(ctx context.Context, artistID string, filter model.CommissionFilter, sorter model.CommissionSorter) (*model.GetCommissionsResult, error)
	UpdateCommission(ctx context.Context, updater model.CommissionUpdater) error
	UpdateCommissionByUser(ctx context.Context, userId string, updater model.CommissionUpdater, decision model.CommissionDecision) error
	UpdateCommissionByUserUpdatedEvent(ctx context.Context, updater model2.UserUpdater) error
	OpenCommissionValidation(ctx context.Context, validation model.CommissionOpenCommissionValidation) error
	UsersValidation(ctx context.Context, validation model.CommissionUsersValidation) error
	GetMessages(ctx context.Context, userId string, filter dMsgModel.MessageFilter) ([]dMsgModel.Messaging, error)
	HandleInboundCommissionMessage(ctx context.Context, msgCreator dMsgModel.MessageCreator) error
	HandleOutBoundCommissionMessage(ctx context.Context, message dMsgModel.Messaging) error
}
