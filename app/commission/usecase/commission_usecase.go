package usecase

import (
	"context"
	"github.com/google/uuid"
	"pixstall-commission/domain/commission"
	"pixstall-commission/domain/commission/model"
	"pixstall-commission/domain/image"
	model2 "pixstall-commission/domain/image/model"
	msgBroker "pixstall-commission/domain/msg-broker"
)

type commissionUseCase struct {
	commRepo commission.Repo
	imageRepo image.Repo
	msgBrokerRepo msgBroker.Repo
}

func NewCommissionUseCase(commRepo commission.Repo, imageRepo image.Repo, msgBrokerRepo msgBroker.Repo) commission.UseCase {
	return &commissionUseCase{
		commRepo: commRepo,
		imageRepo: imageRepo,
		msgBrokerRepo: msgBrokerRepo,
	}
}

func (c commissionUseCase) AddCommission(ctx context.Context, creator model.CommissionCreator) (*model.Commission, error) {
	// Upload
	if len(creator.RefImages) > 0 {
		pathImages := make([]model2.PathImage, 0, len(creator.RefImages))
		for _, refImage := range creator.RefImages {
			pathImages = append(pathImages, model2.PathImage{
				Path:  "commissions/",
				Name:  "RF-" + creator.RequesterID + "-" + uuid.NewString(),
				Image: refImage,
			})
		}
		paths, err := c.imageRepo.SaveImages(ctx, pathImages)
		if err == nil {
			creator.RefImagePaths = paths
		}
	}
	newComm, err := c.commRepo.AddCommission(ctx, creator)
	if err != nil {
		return nil, err
	}
	return newComm, nil
}

func (c commissionUseCase) GetCommissions(ctx context.Context, requesterID string) (*[]model.Commission, error) {
	panic("implement me")
}

func (c commissionUseCase) UpdateCommissions(ctx context.Context, updater model.CommissionUpdater) error {

	dComm, err := c.commRepo.GetCommission(ctx, updater.ID)
	if err != nil {
		return err
	}
	if updater.State != nil {
		switch *updater.State {
		case model.CommissionStatePendingArtistApproval, model.CommissionStateInValid:
			//Don't change to approved is state != PendingValidation
			if dComm.State != model.CommissionStatePendingValidation {
				updater.State = nil
			}
		default:
			break
		}
	}

	return c.commRepo.UpdateCommission(ctx, updater)
}

func (c commissionUseCase) SendMessage(ctx context.Context, commID string) error {
	panic("implement me")
}

func (c commissionUseCase) GetMessages(ctx context.Context, commID string, requesterID string) *model.Commission {
	panic("implement me")
}