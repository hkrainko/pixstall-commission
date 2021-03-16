package usecase

import (
	"fmt"
	"github.com/google/uuid"
	model2 "pixstall-commission/domain/commission/model"
	"pixstall-commission/domain/message/model"
	"time"
)

func newMessagingFromUser(d model.MessageCreator, artistID string, requesterID string) model.Messaging {
	if d.ImagePath != nil {
		return getImageMessage(d, artistID, requesterID)
	}
	return getTextMessage(d, artistID, requesterID)
}

// Private
func getImageMessage(d model.MessageCreator, artistID string, requesterID string) model.ImageMessage {
	return model.ImageMessage{
		Message: model.Message{
			ID:              "msg-" + uuid.NewString(),
			ArtistID:        artistID,
			RequesterID:     requesterID,
			CommissionID:    d.CommissionID,
			CreateTime:      time.Now(),
			LastUpdatedTime: time.Now(),
			State:           model.MessageStateSent,
			MessageType:     model.MessageTypeImage,
		},
		From:      d.Form,
		Text:      &d.Text,
		ImagePath: *d.ImagePath,
	}
}

func getTextMessage(d model.MessageCreator, artistID string, requesterID string) model.TextMessage {
	return model.TextMessage{
		Message: model.Message{
			ID:              "msg-" + uuid.NewString(),
			ArtistID:        artistID,
			RequesterID:     requesterID,
			CommissionID:    d.CommissionID,
			CreateTime:      time.Now(),
			LastUpdatedTime: time.Now(),
			State:           model.MessageStateSent,
			MessageType:     model.MessageTypeText,
		},
		From: d.Form,
		Text: d.Text,
	}
}

func NewPlainSystemMessage(commId string, text string, artistID string, requesterID string) model.PlainSystemMessage {
	return model.PlainSystemMessage{
		SystemMessage: model.SystemMessage{
			Message: model.Message{
				ID:              "msg-" + uuid.NewString(),
				ArtistID:        artistID,
				RequesterID:     requesterID,
				CommissionID:    commId,
				CreateTime:      time.Now(),
				LastUpdatedTime: time.Now(),
				State:           model.MessageStateSent,
				MessageType:     model.MessageTypeSystem,
			},
			Text:              text,
			SystemMessageType: model.SystemMessageTypePlain,
		},
	}
}

func NewProofCopySystemMessage(commId string, text string, artistID string, requesterID string, imagePath string) model.ProofCopySystemMessage {
	return model.ProofCopySystemMessage{
		SystemMessage: model.SystemMessage{
			Message: model.Message{
				ID:              "msg-" + uuid.NewString(),
				ArtistID:        artistID,
				RequesterID:     requesterID,
				CommissionID:    commId,
				CreateTime:      time.Now(),
				LastUpdatedTime: time.Now(),
				State:           model.MessageStateSent,
				MessageType:     model.MessageTypeSystem,
			},
			Text:              text,
			SystemMessageType: model.SystemMessageTypePlain,
		},
		ImagePath: imagePath,
	}
}

func NewCompletionSystemMessage(commId string, text string, artistID string, requesterID string, filePath string) model.CompletionSystemMessage {
	return model.CompletionSystemMessage{
		SystemMessage: model.SystemMessage{
			Message: model.Message{
				ID:              "msg-" + uuid.NewString(),
				ArtistID:        artistID,
				RequesterID:     requesterID,
				CommissionID:    commId,
				CreateTime:      time.Now(),
				LastUpdatedTime: time.Now(),
				State:           model.MessageStateSent,
				MessageType:     model.MessageTypeSystem,
			},
			Text:              text,
			SystemMessageType: model.SystemMessageTypePlain,
		},
		FilePath: filePath,
	}
}

func NewPlainSystemMessageForStateChange(comm model2.Commission, toState model2.CommissionState, filePath *string) model.Messaging {
	switch toState {
	case model2.CommissionStatePendingValidation:
		return NewPlainSystemMessage(comm.ID, "系統己收到委托。", comm.ArtistID, comm.RequesterID)
	case model2.CommissionStateInvalidatedDueToOpenCommission,
	model2.CommissionStateInvalidatedDueToUsers:
		return NewPlainSystemMessage(comm.ID, "系統審查失敗。", comm.ArtistID, comm.RequesterID)
	case model2.CommissionStatePendingArtistApproval:
		return NewPlainSystemMessage(
			comm.ID,
			"系統審查成功。",
			comm.ArtistID,
			comm.RequesterID)
	case model2.CommissionStateInProgress:
		return NewPlainSystemMessage(
			comm.ID,
			fmt.Sprintf("繪師 %v 接受委托。", comm.ArtistID),
			comm.ArtistID,
			comm.RequesterID)
	case model2.CommissionStatePendingRequesterAcceptance:
		return NewProofCopySystemMessage(
			comm.ID,
			fmt.Sprintf("繪師 @%v 上傳完稿。", comm.ArtistID),
			comm.ArtistID,
			comm.RequesterID,
			*filePath)
	case model2.CommissionStatePendingUploadProduct:
		return NewPlainSystemMessage(
			comm.ID,
			fmt.Sprintf("委托人 %v 確認完稿。", comm.RequesterID),
			comm.ArtistID,
			comm.RequesterID)
	case model2.CommissionStatePendingUploadProductDueToRevisionExceed:
		return NewPlainSystemMessage(
			comm.ID,
			"到達完稿可修改上限。",
			comm.ArtistID,
			comm.RequesterID)
	case model2.CommissionStateCompleted:
		return NewCompletionSystemMessage(
			comm.ID,
			"繪師已上傳完成品並完成委托。",
			comm.ArtistID,
			comm.RequesterID,
			*filePath,
		)
	default:
		return nil
	}
}
