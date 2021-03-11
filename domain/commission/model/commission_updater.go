package model

import "time"

type CommissionUpdater struct {
	ID                             string     `json:"id"`
	ArtistName                     *string    `json:"artistName"`
	ArtistProfilePath              *string    `json:"artistProfilePath"`
	RequesterName                  *string    `json:"requesterName"`
	RequesterProfilePath           *string    `json:"requesterProfilePath"`
	TimesAllowedDraftToChange      *int       `json:"timesAllowedDraftToChange"`
	TimesAllowedCompletionToChange *int       `json:"timesAllowedCompletionToChange"`
	CompleteTime                   *time.Time `json:"completeTime"`
	Validation                     *CommissionValidation
	State                          *CommissionState `json:"state"`
}
