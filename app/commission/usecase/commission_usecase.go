package usecase

import (
	"context"
	commMsgDelivery "pixstall-commission/domain/comm-msg-delivery"
	"pixstall-commission/domain/commission"
	"pixstall-commission/domain/commission/model"
	error2 "pixstall-commission/domain/error"
	model2 "pixstall-commission/domain/file/model"
	dImage "pixstall-commission/domain/image"
	"pixstall-commission/domain/message"
	model3 "pixstall-commission/domain/message/model"
	msgBroker "pixstall-commission/domain/msg-broker"
	"time"
)

type commissionUseCase struct {
	commRepo        commission.Repo
	imageRepo       dImage.Repo
	msgBrokerRepo   msgBroker.Repo
	msgRepo         message.Repo
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
		commRepo:        commRepo,
		imageRepo:       imageRepo,
		msgBrokerRepo:   msgBrokerRepo,
		msgRepo:         msgRepo,
		commMsgDeliRepo: commMsgDeliRepo,
	}
}

func (c commissionUseCase) AddCommission(ctx context.Context, creator model.CommissionCreator) (*model.Commission, error) {
	// Upload
	files := func(f []model2.ImageFile) []model2.File {
		var result []model2.File
		for _, v := range f {
			result = append(result, v.File)
		}
		return result
	}(creator.RefImages)

	var storedPaths []string
	for _, f := range files {
		path, err := c.imageRepo.SaveFile(ctx, f, model2.FileTypeCommissionRef, creator.RequesterID, []string{creator.ArtistID})
		if err != nil {
			return nil, err
		}
		storedPaths = append(storedPaths, *path)
	}
	creator.RefImagePaths = storedPaths
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

func (c commissionUseCase) GetCommissions(ctx context.Context, filter model.CommissionFilter, sorter model.CommissionSorter) (*model.GetCommissionsResult, error) {
	return c.commRepo.GetCommissions(ctx, filter, sorter)
}

func (c commissionUseCase) GetCommission(ctx context.Context, commId string) (*model.Commission, error) {
	return c.commRepo.GetCommission(ctx, commId)
}

func (c commissionUseCase) GetWorks(ctx context.Context, artistID string, filter model.CommissionFilter, sorter model.CommissionSorter) (*model.GetCommissionsResult, error) {
	filter.ArtistID = &artistID
	return c.commRepo.GetCommissions(ctx, filter, sorter)
}

func (c commissionUseCase) UpdateCommission(ctx context.Context, updater model.CommissionUpdater) error {
	return c.commRepo.UpdateCommission(ctx, updater)
}

func (c commissionUseCase) UpdateCommissionByUser(ctx context.Context, userId string, updater model.CommissionUpdater, decision model.CommissionDecision) error {
	comm, err := c.commRepo.GetCommission(ctx, updater.ID)
	if err != nil {
		return err
	}
	filteredUpdater, err := c.getFilteredUpdater(userId, *comm, updater, decision)
	if err != nil {
		return err
	}
	if filteredUpdater.ProofCopyImage != nil {
		path, err := c.imageRepo.SaveFile(ctx, filteredUpdater.ProofCopyImage.File, model2.FileTypeCommissionProofCopy, comm.ArtistID, []string{comm.RequesterID})
		if err != nil {
			return model.CommissionErrorUnknown
		}
		filteredUpdater.ProofCopyImagePath = path
	}
	if filteredUpdater.DisplayImageFile != nil {
		var acl []string
		if comm.BePrivate {
			acl = append(acl, comm.RequesterID)
		} else {
			acl = append(acl, "*")
		}
		path, err := c.imageRepo.SaveFile(ctx, filteredUpdater.DisplayImageFile.File, model2.FileTypeArtwork, comm.ArtistID, acl)
		if err != nil {
			return model.CommissionErrorUnknown
		}
		filteredUpdater.DisplayImage = &model.DisplayImage{
			Path:        *path,
			Volume:      filteredUpdater.DisplayImageFile.Volume,
			Size:        filteredUpdater.DisplayImageFile.Size,
			ContentType: filteredUpdater.DisplayImageFile.ContentType,
		}
	}
	if filteredUpdater.CompletionFile != nil {
		path, err := c.imageRepo.SaveFile(ctx, *filteredUpdater.CompletionFile, model2.FileTypeCompletion, comm.ArtistID, []string{comm.RequesterID})
		if err != nil {
			return model.CommissionErrorUnknown
		}
		filteredUpdater.CompletionFilePath = path
	}
	err = c.commRepo.UpdateCommission(ctx, *filteredUpdater)
	if err != nil {
		return err
	}

	if *filteredUpdater.State == model.CommissionStateCompleted {
		comm, err := c.commRepo.GetCommission(ctx, updater.ID)
		if err == nil && !comm.BePrivate && comm.DisplayImage != nil {
			err = c.msgBrokerRepo.SendCommissionCompletedMessage(ctx, *comm)
			if err != nil {
				// TODO: send the message later
			}
		}
	}
	msg, err := NewSystemMessage(decision, *comm, *filteredUpdater)
	if err == nil {
		err = c.msgRepo.AddNewMessage(ctx, nil, msg)
		if err == nil {
			_ = c.commMsgDeliRepo.DeliverCommissionMessage(ctx, msg)
		}
	}
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
		var acl []string
		if msgCreator.Form == comm.ArtistID {
			acl = append(acl, comm.RequesterID)
		} else {
			acl = append(acl, comm.ArtistID)
		}
		path, err := c.imageRepo.SaveFile(ctx, msgCreator.Image.File, model2.FileTypeMessage, msgCreator.Form, acl)
		if err == nil {
			msgCreator.ImagePath = path
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

func (c commissionUseCase) getFilteredUpdater(userID string, comm model.Commission, updater model.CommissionUpdater, decision model.CommissionDecision) (*model.CommissionUpdater, error) {
	err := c.isDecisionAllowToMadeByUser(userID, comm, updater, decision)
	if err != nil {
		return nil, err
	}
	result := model.CommissionUpdater{ID: updater.ID}
	switch decision {
	case model.CommissionDecisionRequesterModify:
		state := model.CommissionStatePendingRequesterModificationValidation
		//TODO, check price....
		result.State = &state
		return &result, nil
	case model.CommissionDecisionArtistAccept:
		now := time.Now()
		state := model.CommissionStateInProgress
		result.StartTime = &now
		result.State = &state
		return &result, nil
	case model.CommissionDecisionArtistDecline:
		state := model.CommissionStateDeclinedByArtist
		result.State = &state
		return &result, nil
	case model.CommissionDecisionRequesterCancel:
		state := model.CommissionStateCancelledByRequester
		result.State = &state
		return &result, nil
	case model.CommissionDecisionArtistUploadProofCopy:
		state := model.CommissionStatePendingRequesterAcceptance
		result.State = &state
		result.ProofCopyImage = updater.ProofCopyImage
		return &result, nil
	case model.CommissionDecisionRequesterAcceptProofCopy:
		state := model.CommissionStatePendingUploadProduct
		result.State = &state
		return &result, nil
	case model.CommissionDecisionRequesterRequestRevision:
		state := model.CommissionStateInProgress
		result.State = &state
		compRevReqTime := comm.ProofCopyRevisionRequestTime + 1
		result.CompletionRevisionRequestTime = &compRevReqTime
		return &result, nil
	case model.CommissionDecisionArtistUploadProduct:
		state := model.CommissionStatePendingRequesterAcceptProduct
		result.State = &state
		result.DisplayImageFile = updater.DisplayImageFile
		result.CompletionFile = updater.CompletionFile
		return &result, nil
	case model.CommissionDecisionRequesterAcceptProduct:
		now := time.Now()
		state := model.CommissionStateCompleted
		result.State = &state
		result.Rating = updater.Rating
		result.Comment = updater.Comment
		result.CompletedTime = &now
		return &result, nil
	}
	return nil, model.CommissionErrorDecisionNotAllowed
}

func (c commissionUseCase) isDecisionAllowToMadeByUser(userID string, comm model.Commission, updater model.CommissionUpdater, decision model.CommissionDecision) error {
	if userID != comm.ArtistID && userID != comm.RequesterID {
		return error2.UnAuthError
	}
	switch decision {
	case model.CommissionDecisionRequesterModify:
		if userID == comm.RequesterID && comm.State == model.CommissionStatePendingArtistApproval {
			return nil
		}
	case model.CommissionDecisionArtistAccept, model.CommissionDecisionArtistDecline:
		if userID == comm.ArtistID && comm.State == model.CommissionStatePendingArtistApproval {
			return nil
		}
	case model.CommissionDecisionRequesterCancel:
		if userID == comm.RequesterID && comm.State == model.CommissionStatePendingArtistApproval {
			return nil
		}
	case model.CommissionDecisionArtistUploadProofCopy:
		if userID == comm.ArtistID && comm.State == model.CommissionStateInProgress && updater.ProofCopyImage != nil {
			return nil
		}
	case model.CommissionDecisionRequesterAcceptProofCopy:
		if userID == comm.RequesterID && comm.State == model.CommissionStatePendingRequesterAcceptance {
			return nil
		}
	case model.CommissionDecisionRequesterRequestRevision:
		if userID == comm.RequesterID && comm.State == model.CommissionStatePendingRequesterAcceptance {
			if comm.TimesAllowedCompletionToChange != nil && comm.ProofCopyRevisionRequestTime >= *comm.TimesAllowedCompletionToChange {
				return model.CommissionErrorRevisionExceed
			}
			return nil
		}
	case model.CommissionDecisionArtistUploadProduct:
		if userID == comm.ArtistID && comm.State == model.CommissionStatePendingUploadProduct && updater.CompletionFile != nil {
			return nil
		}
	case model.CommissionDecisionRequesterAcceptProduct:
		if userID == comm.RequesterID && comm.State == model.CommissionStatePendingRequesterAcceptProduct && updater.Rating != nil {
			return nil
		}
	default:
		return model.CommissionErrorDecisionNotAllowed
	}
	return model.CommissionErrorDecisionNotAllowed
}
