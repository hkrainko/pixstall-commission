package update_commission

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
		case model.CommissionErrorNotFound:
			return http.StatusNotFound, ErrorResponse{
				Message: commError.Error(),
			}
		case model.CommissionErrorUnAuth:
			return http.StatusUnauthorized, ErrorResponse{
				Message: commError.Error(),
			}
		case model.CommissionErrorDecisionNotAllowed:
			return http.StatusMethodNotAllowed, ErrorResponse{
				Message: commError.Error(),
			}
		case model.CommissionErrorPriceInvalid,
			model.CommissionErrorDayNeedInvalid:
			return http.StatusBadRequest, ErrorResponse{
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