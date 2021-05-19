package get_messages

import (
	model2 "pixstall-commission/domain/message/model"
)

type Response struct {
	CommissionID  string             `json:"commissionId"`
	LastMessageID *string            `json:"lastMessageId,omitempty"`
	Messages      []model2.Messaging `json:"messages"`
}

func NewResponse(commID string, lastMessageID *string, msgs []model2.Messaging) *Response {
	return &Response{
		CommissionID:  commID,
		LastMessageID: lastMessageID,
		Messages:      msgs,
	}
}
