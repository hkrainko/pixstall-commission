package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	commMsgDelivery "pixstall-commission/domain/comm-msg-delivery"
	"pixstall-commission/domain/commission"
	"pixstall-commission/domain/commission/model"
	"pixstall-commission/domain/image"
	model2 "pixstall-commission/domain/image/model"
	"pixstall-commission/domain/message"
	model3 "pixstall-commission/domain/message/model"
	msgBroker "pixstall-commission/domain/msg-broker"
)

type commissionUseCase struct {
	commRepo      commission.Repo
	imageRepo     image.Repo
	msgBrokerRepo msgBroker.Repo
	msgRepo       message.Repo
	commMsgDeliRepo commMsgDelivery.Repo
}

func NewCommissionUseCase(
	commRepo commission.Repo,
	imageRepo image.Repo,
	msgBrokerRepo msgBroker.Repo,
	msgRepo message.Repo,
	commMsgDeliRepo commMsgDelivery.Repo,
	) commission.UseCase {
	return &commissionUseCase{
		commRepo:      commRepo,
		imageRepo:     imageRepo,
		msgBrokerRepo: msgBrokerRepo,
		msgRepo:       msgRepo,
		commMsgDeliRepo: commMsgDeliRepo,
	}
}

func (c commissionUseCase) AddCommission(ctx context.Context, creator model.CommissionCreator) (*model.Commission, error) {
	// Upload
	storedPaths, err := c.storeRefImage(ctx, creator)
	if err == nil {
		creator.RefImagePaths = *storedPaths
	}
	newComm, err := c.commRepo.AddCommission(ctx, creator)
	if err != nil {
		return nil, err
	}
	err = c.msgBrokerRepo.SendCommissionCreatedMessage(ctx, *newComm)
	if err != nil {
		print(err)
	}
	return newComm, nil
}

func (c commissionUseCase) GetCommissions(ctx context.Context, filter model.CommissionFilter, sorter model.CommissionSorter) (*[]model.Commission, error) {
	return c.commRepo.GetCommissions(ctx, filter, sorter)
}

func (c commissionUseCase) GetCommission(ctx context.Context, commId string) (*model.Commission, error) {
	return c.commRepo.GetCommission(ctx, commId)
}

func (c commissionUseCase) GetWorks(ctx context.Context, artistID string, filter model.CommissionFilter, sorter model.CommissionSorter) (*[]model.Commission, error) {
	filter.ArtistID = &artistID
	return c.commRepo.GetCommissions(ctx, filter, sorter)
}

func (c commissionUseCase) UpdateCommissions(ctx context.Context, updater model.CommissionUpdater) error {
	return c.commRepo.UpdateCommission(ctx, updater)
}

func (c commissionUseCase) OpenCommissionValidation(ctx context.Context, validation model.CommissionOpenCommissionValidation) error {
	comm, err := c.commRepo.GetCommission(ctx, validation.ID)
	if err != nil {
		return err
	}
	if comm.State != model.CommissionStatePendingValidation {
		return nil
	}
	updater := model.CommissionUpdater{ID: validation.ID}
	if validation.IsValid {
		updater = c.getCommValidationOpenCommUpdater(comm.ValidationHistory, updater, validation)
	} else {
		state := model.CommissionStateInvalidatedDueToOpenCommission
		updater.State = &state
	}
	return c.commRepo.UpdateCommission(ctx, updater)
}

func (c commissionUseCase) UsersValidation(ctx context.Context, validation model.CommissionUsersValidation) error {
	comm, err := c.commRepo.GetCommission(ctx, validation.CommID)
	if err != nil {
		return err
	}
	if comm.State != model.CommissionStatePendingValidation {
		return nil
	}
	updater := model.CommissionUpdater{ID: validation.CommID}
	if validation.IsValid {
		updater = c.getCommValidationUsersUpdater(comm.ValidationHistory, updater, validation)
	} else {
		state := model.CommissionStateInvalidatedDueToUsers
		updater.State = &state
	}
	return c.commRepo.UpdateCommission(ctx, updater)
}

func (c commissionUseCase) HandleInboundCommissionMessage(ctx context.Context, msgCreator model3.MessageCreator) error {

	comm, err := c.commRepo.GetCommission(ctx, msgCreator.CommissionID)
	if err != nil {
		return err
	}
	if !c.isStateAllowToSendMessage(comm.State) {
		return model.CommissionErrorNotAllowSendMessage
	}
	messaging := newMessagingFromMessageCreator(msgCreator, comm.ArtistID, comm.RequesterID)

	err = c.msgRepo.AddNewMessage(ctx, messaging)
	if err != nil {
		return err
	}
	_ = c.msgBrokerRepo.SendCommissionMessageReceivedMessage(ctx, messaging)
	return nil
}

func (c commissionUseCase) HandleOutBoundCommissionMessage(ctx context.Context, message model3.Messaging) error {
	return c.commMsgDeliRepo.DeliverCommissionMessage(ctx, message)
}

// Private
func (c commissionUseCase) storeRefImage(ctx context.Context, creator model.CommissionCreator) (*[]string, error) {
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
			return &paths, nil
		}
	}
	return nil, errors.New("storeRefImage failed")
}

func (c commissionUseCase) getCommValidationOpenCommUpdater(history []model.CommissionValidation, updater model.CommissionUpdater, validation model.CommissionOpenCommissionValidation) model.CommissionUpdater {
	v := model.CommissionValidationOpenCommission
	if isCommValid := c.isCommValidationCompletable(
		history, v,
	); isCommValid {
		newState := model.CommissionStatePendingArtistApproval
		updater.State = &newState
	}
	updater.Validation = &v
	updater.TimesAllowedDraftToChange = validation.TimesAllowedDraftToChange
	updater.TimesAllowedCompletionToChange = validation.TimesAllowedCompletionToChange
	return updater
}

func (c commissionUseCase) getCommValidationUsersUpdater(history []model.CommissionValidation, updater model.CommissionUpdater, validation model.CommissionUsersValidation) model.CommissionUpdater {
	v := model.CommissionValidationUsers
	if isCommValid := c.isCommValidationCompletable(
		history, v,
	); isCommValid {
		newState := model.CommissionStatePendingArtistApproval
		updater.State = &newState
	}
	updater.ArtistName = validation.ArtistName
	updater.ArtistProfilePath = validation.ArtistProfilePath
	updater.RequesterName = validation.RequesterName
	updater.RequesterProfilePath = validation.RequesterProfilePath
	updater.Validation = &v
	return updater
}

func (c commissionUseCase) isCommValidationCompletable(validationHistory []model.CommissionValidation, newValidation model.CommissionValidation) bool {
	switch newValidation {
	case model.CommissionValidationOpenCommission:
		for _, hV := range validationHistory {
			if hV == model.CommissionValidationUsers {
				return true
			}
		}
	case model.CommissionValidationUsers:
		for _, hV := range validationHistory {
			if hV == model.CommissionValidationOpenCommission {
				return true
			}
		}
	}
	return false
}

func (c commissionUseCase) isStateAllowToSendMessage(state model.CommissionState) bool {
	switch state {
	case model.CommissionStatePendingArtistApproval,
	model.CommissionStateInProgress,
	model.CommissionStatePendingRequesterAcceptance:
		return true
	default:
		return false
	}
}
