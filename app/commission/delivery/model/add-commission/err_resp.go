package add_commission

import (
	"net/http"
	"pixstall-commission/domain/commission/model"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(err error) (int, interface{}) {
	if commError, isError := err.(model.CommissionError); isError {
		switch commError {
		case model.CommissionErrorUnAuth:
			return http.StatusUnauthorized, ErrorResponse{
				Message: commError.Error(),
			}
		default:
			return http.StatusInternalServerError, ErrorResponse{
				Message: err.Error(),
			}
		}
	} else {
		return http.StatusInternalServerError, ErrorResponse{
			Message: err.Error(),
		}
	}
}