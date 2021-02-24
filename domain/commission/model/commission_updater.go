package model

import "time"

type CommissionUpdater struct {
	ID                   string           `json:"id"`
	ArtistName           *string          `json:"artistName"`
	ArtistProfilePath    *string          `json:"artistProfilePath"`
	RequesterName        *string          `json:"requesterName"`
	RequesterProfilePath *string          `json:"requesterProfilePath"`
	CompleteTime         *time.Time       `json:"completeTime"`
	State                *CommissionState `json:"state"`
}
