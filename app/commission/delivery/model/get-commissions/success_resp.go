package get_commissions

import "pixstall-commission/domain/commission/model"

type Response struct {
	RequesterId string             `json:"requesterId"`
	Commissions []model.Commission `json:"commissions"`
	Offset      int                `json:"offSet"`
	Count       int                `json:"count"`
	Total       int                `json:"total"`
}

func NewResponse(requesterId string, result model.GetCommissionsResult, offset int) *Response {
	return &Response{
		RequesterId: requesterId,
		Commissions: result.Commissions,
		Offset:      offset,
		Count:       len(result.Commissions),
		Total:       result.Total,
	}
}
