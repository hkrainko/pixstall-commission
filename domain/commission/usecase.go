package commission

import (
	"context"
	"pixstall-commission/domain/commission/model"
)

type UseCase interface {
	AddCommission(ctx context.Context, creator model.CommissionCreator) (*model.Commission, error)
	GetCommissions(ctx context.Context, requesterID string) (*[]model.Commission, error)
	UpdateCommissions(ctx context.Context, updater model.CommissionUpdater) error
	SendMessage(ctx context.Context, commID string) error
	GetMessages(ctx context.Context, commID string, requesterID string) *model.Commission
}
