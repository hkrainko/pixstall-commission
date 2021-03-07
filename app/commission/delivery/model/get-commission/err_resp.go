package get_commission

import (
	"net/http"
	error2 "pixstall-commission/domain/error"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(err error) (int, interface{}) {
	switch err {
	case error2.NotFoundError:
		return http.StatusNotFound, ErrorResponse{
			Message: err.Error(),
		}
	default:
		return http.StatusInternalServerError, ErrorResponse{
			Message: err.Error(),
		}
	}
}
