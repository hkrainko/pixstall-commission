package get_messages

import (
	model2 "pixstall-commission/domain/message/model"
)

type Response struct {
	CommissionID string             `json:"commissionId"`
	Messages     []model2.Messaging `json:"messages"`
}

func NewResponse(commID string, msgs []model2.Messaging) *Response {
	return &Response{
		CommissionID: commID,
		Messages:     msgs,
	}
}
