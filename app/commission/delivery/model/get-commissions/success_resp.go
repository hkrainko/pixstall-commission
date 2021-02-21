package get_commissions

import "pixstall-commission/domain/commission/model"

type Response struct {
	RequesterId string             `json:"requesterId"`
	Commissions []model.Commission `json:"commissions"`
}

func NewResponse(requesterId string, dCommissions []model.Commission) *Response {
	return &Response{
		RequesterId: requesterId,
		Commissions: dCommissions,
	}
}