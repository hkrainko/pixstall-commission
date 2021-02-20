package usecase

import (
	"context"
	"pixstall-commission/domain/commission"
	"pixstall-commission/domain/commission/model"
	"pixstall-commission/domain/image"
)

type commissionUseCase struct {
	commRepo commission.Repo
	imageRepo image.Repo
}

func NewCommissionUseCase(commRepo commission.Repo, imageRepo image.Repo) commission.UseCase {
	return &commissionUseCase{
		commRepo: commRepo,
		imageRepo: imageRepo,
	}
}

func (c commissionUseCase) AddCommission(ctx context.Context, creator model.CommissionCreator) error {
	panic("implement me")
}

func (c commissionUseCase) GetCommissions(ctx context.Context, requesterID string) (*model.Commission, error) {
	panic("implement me")
}

func (c commissionUseCase) SendMessage(ctx context.Context, commID string) error {
	panic("implement me")
}

func (c commissionUseCase) GetMessages(ctx context.Context, commID string, requesterID string) *model.Commission {
	panic("implement me")
}