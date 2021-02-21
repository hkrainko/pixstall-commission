package add_commission

type Response struct {
	CommID string `json:"commissionId"`
}

func NewResponse(commID string) *Response {
	return &Response{
		CommID: commID,
	}
}
