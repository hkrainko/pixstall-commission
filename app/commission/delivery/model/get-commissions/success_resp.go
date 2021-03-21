package get_commissions

import "pixstall-commission/domain/commission/model"

type Response struct {
	RequesterId string             `json:"requesterId"`
	Commissions []model.Commission `json:"commissions"`
	Offset      int                `json:"offSet"`
	Count       int                `json:"count"`
}

func NewResponse(requesterId string, dCommissions []model.Commission, offset int, count int) *Response {
	return &Response{
		RequesterId: requesterId,
		Commissions: dCommissions,
		Offset: offset,
		Count: count,
	}
}
