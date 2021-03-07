package get_commission

import "pixstall-commission/domain/commission/model"

type Response struct {
	Commission model.Commission `json:"commission"`
}

func NewResponse(dCommission model.Commission) *Response {
	return &Response{
		Commission: dCommission,
	}
}
