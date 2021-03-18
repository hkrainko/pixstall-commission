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

func NewSystemMessage(decision model2.CommissionDecision, comm model2.Commission, updater model2.CommissionUpdater) (model.Messaging, error) {

	msg := model.Message{
		ID:              "msg-" + uuid.NewString(),
		ArtistID:        comm.ArtistID,
		RequesterID:     comm.RequesterID,
		CommissionID:    comm.ID,
		CreateTime:      time.Now(),
		LastUpdatedTime: time.Now(),
		State:           model.MessageStateSent,
		MessageType:     model.MessageTypeSystem,
	}

	switch decision {
	case model2.CommissionDecisionRequesterModify,
		model2.CommissionDecisionArtistAccept,
		model2.CommissionDecisionArtistDecline,
		model2.CommissionDecisionRequesterCancel,
		model2.CommissionDecisionRequesterAcceptProofCopy,
		model2.CommissionDecisionRequesterRequestRevision:
		return model.PlainSystemMessage{
			SystemMessage: model.SystemMessage{
				Message:           msg,
				Text:              decisionMessage(decision, comm),
				SystemMessageType: model.SystemMessageTypePlain,
			},
		}, nil
	case model2.CommissionDecisionArtistUploadProofCopy:
		return model.UploadProofCopySystemMessage{
			SystemMessage: model.SystemMessage{
				Message: msg,
				Text:              decisionMessage(decision, comm),
				SystemMessageType: model.SystemMessageTypeUploadProofCopy,
			},
			ImagePath: *updater.ProofCopyImagePath,
		}, nil
	case model2.CommissionDecisionArtistUploadProduct:
		return model.UploadProductSystemMessage{
			SystemMessage: model.SystemMessage{
				Message: msg,
				Text:              decisionMessage(decision, comm),
				SystemMessageType: model.SystemMessageTypeUploadProduct,
			},
			FilePath: *updater.CompletionFilePath,
		}, nil
	case model2.CommissionDecisionRequesterAcceptProduct:
		return model.AcceptProductSystemMessage{
			SystemMessage: model.SystemMessage{
				Message: msg,
				Text:              decisionMessage(decision, comm),
				SystemMessageType: model.SystemMessageTypeAcceptProduct,
			},
			Rating:  *updater.Rating,
			Comment: updater.Comment,
		}, nil
	default:
		return nil, model2.CommissionErrorUnknown
	}
}

func decisionMessage(decision model2.CommissionDecision, comm model2.Commission) string {
	switch decision {
	case model2.CommissionDecisionRequesterModify:
		return fmt.Sprintf("委托人 %v 嘗試更改委托內容。", comm.RequesterID)
	case model2.CommissionDecisionArtistAccept:
		return fmt.Sprintf("繪師 %v 接受委托。", comm.ArtistID)
	case model2.CommissionDecisionArtistDecline:
		return fmt.Sprintf("繪師 %v 拒絕委托。", comm.ArtistID)
	case model2.CommissionDecisionRequesterCancel:
		return fmt.Sprintf("委托人 %v 取消委托。", comm.RequesterID)
	case model2.CommissionDecisionArtistUploadProofCopy:
		return fmt.Sprintf("繪師 %v 上傳委托完稿。", comm.ArtistID)
	case model2.CommissionDecisionRequesterAcceptProofCopy:
		return fmt.Sprintf("委托人 %v 接受委托完稿。", comm.RequesterID)
	case model2.CommissionDecisionRequesterRequestRevision:
		return fmt.Sprintf("委托人 %v 對完稿提出修訂。", comm.RequesterID)
	case model2.CommissionDecisionArtistUploadProduct:
		return fmt.Sprintf("繪師 %v 上傳完成品。", comm.ArtistID)
	case model2.CommissionDecisionRequesterAcceptProduct:
		return fmt.Sprintf("委托人 %v 接受完成品。", comm.RequesterID)
	default:
		return ""
	}
}
