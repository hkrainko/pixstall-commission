package model

type CommissionUsersValidation struct {
	CommID               string  `json:"commissionId"`
	IsValid              bool    `json:"isValid"`
	InvalidReason        *string `json:"invalidReason"`
	ArtistName           *string `json:"artistName"`
	ArtistProfilePath    *string `json:"artistProfilePath"`
	RequesterName        *string `json:"requesterName"`
	RequesterProfilePath *string `json:"requesterProfilePath"`
}
