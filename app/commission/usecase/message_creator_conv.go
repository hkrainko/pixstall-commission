package usecase

import (
	"github.com/google/uuid"
	"pixstall-commission/domain/message/model"
	"time"
)

func newMessagingFromMessageCreator(d model.MessageCreator, artistID string, requesterID string) model.Messaging {
	if d.Form == nil {
		return getSystemMessage(d, artistID, requesterID)
	}
	if d.ImagePath != nil {
		return getImageMessage(d, artistID, requesterID)
	}
	return getTextMessage(d, artistID, requesterID)
}

// Private
func getSystemMessage(d model.MessageCreator, artistID string, requesterID string) model.SystemMessage {
	return model.SystemMessage{
		Message: model.Message{
			ID:              "msg-" + uuid.NewString(),
			ArtistID:        "",
			RequesterID:     "",
			CommissionID:    d.CommissionID,
			CreateTime:      time.Now(),
			LastUpdatedTime: time.Now(),
			State:           model.MessageStateSent,
			MessageType:     model.MessageTypeSystem,
		},
		Text: &d.Text,
	}
}

func getImageMessage(d model.MessageCreator, artistID string, requesterID string) model.ImageMessage {
	return model.ImageMessage{
		Message: model.Message{
			ID:              "msg-" + uuid.NewString(),
			ArtistID:        "",
			RequesterID:     "",
			CommissionID:    d.CommissionID,
			CreateTime:      time.Now(),
			LastUpdatedTime: time.Now(),
			State:           model.MessageStateSent,
			MessageType:     model.MessageTypeImage,
		},
		From:      *d.Form,
		Text:      &d.Text,
		ImagePath: *d.ImagePath,
	}
}

func getTextMessage(d model.MessageCreator, artistID string, requesterID string) model.TextMessage {
	return model.TextMessage{
		Message: model.Message{
			ID:              "msg-" + uuid.NewString(),
			ArtistID:        "",
			RequesterID:     "",
			CommissionID:    d.CommissionID,
			CreateTime:      time.Now(),
			LastUpdatedTime: time.Now(),
			State:           model.MessageStateSent,
			MessageType:     model.MessageTypeText,
		},
		From: *d.Form,
		Text: d.Text,
	}
}