package model

import (
	"pixstall-commission/domain/file/model"
	"time"
)

type CommissionUpdater struct {
	ID                             string     `json:"id"`
	ArtistName                     *string    `json:"artistName"`
	ArtistProfilePath              *string    `json:"artistProfilePath"`
	RequesterName                  *string    `json:"requesterName"`
	RequesterProfilePath           *string    `json:"requesterProfilePath"`
	TimesAllowedDraftToChange      *int       `json:"timesAllowedDraftToChange"`
	TimesAllowedCompletionToChange *int       `json:"timesAllowedCompletionToChange"`
	CompletionRevisionRequestTime  *int       `json:"completionRevisionRequestTime"`
	CompletedTime                  *time.Time `json:"completedTime"`
	Validation                     *CommissionValidation
	State                          *CommissionState `json:"state"`
	Rating                         *int             `json:"rating"`
	Comment                        *string          `json:"comment"`
	CompletionFile                 *model.File      `json:"completionFile"`
	CompletionFilePath             *string          `json:"completionFilePath"`
	DisplayImageFile               *model.ImageFile `json:"displayImageFile"`
	DisplayImage                   *DisplayImage    `json:"displayImage"`
	ProofCopyImage                 *model.ImageFile `json:"proofCopyImage"`
	ProofCopyImagePath             *string          `json:"proofCopyImagePath"`

	// TODO: allow requester to edit before artist approval
	Price         *Price
	DayNeed       *int
	Size          *model.Size
	Resolution    *float64
	ExportFormat  *string
	Desc          *string
	PaymentMethod *string
	BePrivate     *bool
	Anonymous     *bool
}
