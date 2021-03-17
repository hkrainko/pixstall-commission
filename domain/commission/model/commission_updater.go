package model

import (
	"image"
	"mime/multipart"
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
	CompleteTime                   *time.Time `json:"completeTime"`
	Validation                     *CommissionValidation
	State                          *CommissionState `json:"state"`
	Rating                         *int             `json:"rating"`
	Comment                        *string          `json:"comment"`
	CompletionFile                 *multipart.File  `json:"completionFile"`
	CompletionFilePath             *string          `json:"completionFilePath"`
	DisplayImage                   *image.Image     `json:"displayImage"`
	DisplayImagePath               *string          `json:"displayImagePath"`
	ProofCopyImage                 *image.Image     `json:"proofCopyImage"`
	ProofCopyImagePath             *string          `json:"proofCopyImagePath"`

	// TODO: allow requester to edit before artist approval
	Price         *Price
	DayNeed       *int
	Size          *Size
	Resolution    *float64
	ExportFormat  *string
	Desc          *string
	PaymentMethod *string
	BePrivate     *bool
	Anonymous     *bool
}
