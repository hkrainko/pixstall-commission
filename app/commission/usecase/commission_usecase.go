package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"image"
	commMsgDelivery "pixstall-commission/domain/comm-msg-delivery"
	"pixstall-commission/domain/commission"
	"pixstall-commission/domain/commission/model"
	dImage "pixstall-commission/domain/image"
	model2 "pixstall-commission/domain/image/model"
	"pixstall-commission/domain/message"
	model3 "pixstall-commission/domain/message/model"
	msgBroker "pixstall-commission/domain/msg-broker"
)

type commissionUseCase struct {
	commRepo      commission.Repo
	imageRepo     dImage.Repo
	msgBrokerRepo msgBroker.Repo
	msgRepo       message.Repo
	commMsgDeliRepo commMsgDelivery.Repo
}

func NewCommissionUseCase(
	commRepo commission.Repo,
	imageRepo dImage.Repo,
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
	storedPaths, err := c.storeImages(ctx, "commissions", "rf-" + creator.RequesterID + "-" + uuid.NewString(), creator.RefImages)
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

func (c commissionUseCase) UpdateCommission(ctx context.Context, updater model.CommissionUpdater) error {
	return c.commRepo.UpdateCommission(ctx, updater)
}

func (c commissionUseCase) UpdateCommissionByUser(ctx context.Context, userId string, updater model.CommissionUpdater) error {
	comm, err := c.commRepo.GetCommission(ctx, updater.ID)
	if err != nil {
		return err
	}
	if updater.State != nil {
		err :=  c.isStateAllowToUpdateByUser(userId, *comm, *updater.State)
		if err != nil {
			return err
		}
	}
	err = c.commRepo.UpdateCommission(ctx, updater)
	if err != nil {
		return err
	}


	//err = c.msgRepo.AddNewMessage(ctx, nil, msg)
	//if err != nil {
	//	_ = c.commMsgDeliRepo.DeliverCommissionMessage(ctx, msg)
	//}
	//Ignore the error from sending message as we only care the state changed
	return nil
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

func (c commissionUseCase) GetMessages(ctx context.Context, userId string, commId string, offset int, count int) ([]model3.Messaging, error) {
	return c.msgRepo.GetMessages(ctx, userId, commId, offset, count)
}

func (c commissionUseCase) HandleInboundCommissionMessage(ctx context.Context, msgCreator model3.MessageCreator) error {
	comm, err := c.commRepo.GetCommission(ctx, msgCreator.CommissionID)
	if err != nil {
		return err
	}
	if !c.isStateAllowToSendMessage(comm.State) {
		return model.CommissionErrorNotAllowSendMessage
	}
	if msgCreator.Image != nil {
		storedPaths, err := c.storeImage(ctx, "messages", "img-msg-" + msgCreator.CommissionID + "-" + uuid.NewString(), *msgCreator.Image)
		if err == nil {
			msgCreator.ImagePath = storedPaths
		}
	}
	messaging := newMessagingFromUser(msgCreator, comm.ArtistID, comm.RequesterID)

	err = c.msgRepo.AddNewMessage(ctx, &msgCreator.Form, messaging)
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
func (c commissionUseCase) storeImages(ctx context.Context, path string, name string, images []image.Image) (*[]string, error) {
	if len(images) > 0 {
		pathImages := make([]model2.PathImage, 0, len(images))
		for _, refImage := range images {
			pathImages = append(pathImages, model2.PathImage{
				Path:  path + "/",
				Name:  name,
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

func (c commissionUseCase) storeImage(ctx context.Context, path string, name string, image image.Image) (*string, error) {
	pathImage := model2.PathImage{
		Path:  path + "/",
		Name:  name,
		Image: image,
	}
	savedPath, err := c.imageRepo.SaveImage(ctx, pathImage)
	if err != nil {
		return nil, err
	}
	return savedPath, nil
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

func (c commissionUseCase) isStateAllowToUpdateByUser(userId string, comm model.Commission, toState model.CommissionState) error {
	if userId != comm.ArtistID && userId != comm.RequesterID {
		return model.CommissionErrorUnAuth
	}
	switch comm.State {
	case model.CommissionStatePendingValidation,
	model.CommissionStateInvalidatedDueToOpenCommission,
	model.CommissionStateInvalidatedDueToUsers:
		return model.CommissionErrorStateNotAllowed
	case model.CommissionStatePendingArtistApproval:
		if userId == comm.ArtistID {
			if toState == model.CommissionStateInProgress || toState == model.CommissionStateRejectedByArtist {
				return nil
			} else {
				return model.CommissionErrorStateNotAllowed
			}
		} else {
			if toState == model.CommissionStateRejectedByRequester {
				return nil
			} else {
				return model.CommissionErrorStateNotAllowed
			}
		}
	case model.CommissionStateInProgress:
		if userId == comm.ArtistID {
			if toState == model.CommissionStatePendingRequesterAcceptance {
				return nil
			} else {
				return model.CommissionErrorStateNotAllowed
			}
		} else {
			if toState == model.CommissionStateRejectedByRequester {
				return nil
			} else {
				return model.CommissionErrorStateNotAllowed
			}
		}
	case model.CommissionStatePendingRequesterAcceptance:
		if userId == comm.RequesterID {
			switch toState {
			case model.CommissionStateCompleted:
				return nil
			case model.CommissionStateInProgress:
				if comm.TimesAllowedCompletionToChange == nil ||
					comm.CompletionChangingRequestTime < *comm.TimesAllowedCompletionToChange {
					return nil
				} else {
					return model.CommissionErrorStateNotAllowed
				}
			default:
				return model.CommissionErrorStateNotAllowed
			}
		} else {
			return model.CommissionErrorStateNotAllowed
		}
	default:
		return nil
	}
}