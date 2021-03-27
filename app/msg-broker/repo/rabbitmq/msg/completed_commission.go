package msg

import (
	"pixstall-commission/domain/commission/model"
	model2 "pixstall-commission/domain/file/model"
	"time"
)

type CompletedCommission struct {
	ID               string `json:"id" bson:"id"`
	OpenCommissionID string `json:"openCommissionId" bson:"openCommissionId"`

	ArtistID             string  `json:"artistId" bson:"artistId"`
	ArtistName           string  `json:"artistName" bson:"artistName"`
	ArtistProfilePath    *string `json:"artistProfilePath" bson:"artistProfilePath,omitempty"`
	RequesterID          string  `json:"requesterId" bson:"requesterId"`
	RequesterName        string  `json:"requesterName" bson:"requesterName"`
	RequesterProfilePath *string `json:"requesterProfilePath" bson:"requesterProfilePath,omitempty"`

	Price                          model.Price  `json:"price" bson:"price"`
	DayNeed                        int          `json:"dayNeed" bson:"dayNeed"`
	Size                           *model2.Size `json:"size" bson:"size,omitempty"`
	Resolution                     *float64     `json:"resolution" bson:"resolution,omitempty"`
	ExportFormat                   *string      `json:"exportFormat" bson:"exportFormat,omitempty"`
	Desc                           string       `json:"desc" bson:"desc"`
	PaymentMethod                  string       `json:"paymentMethod" bson:"paymentMethod"`
	IsR18                          bool         `json:"isR18" bson:"isR18"`
	BePrivate                      bool         `json:"bePrivate" bson:"bePrivate"`
	Anonymous                      bool         `json:"anonymous" bson:"anonymous"`
	RefImagePaths                  []string     `json:"refImagePaths" bson:"refImagePaths"`
	TimesAllowedDraftToChange      *int         `json:"timesAllowedDraftToChange" bson:"timesAllowedDraftToChange"`
	TimesAllowedCompletionToChange *int         `json:"timesAllowedCompletionToChange" bson:"timesAllowedCompletionToChange"`
	DraftChangingRequestTime       int          `json:"draftChangingRequestTime" bson:"draftChangingRequestTime"`
	ProofCopyRevisionRequestTime   int          `json:"proofCopyRevisionRequestTime" bson:"proofCopyRevisionRequestTime"`

	ProofCopyImagePaths []string           `json:"proofCopyImagePaths" bson:"proofCopyImagePaths"`
	DisplayImage        model.DisplayImage `json:"displayImage" bson:"displayImage"`
	CompletionFilePath  *string            `json:"completionFilePath,omitempty" bson:"completionFilePath,omitempty"`
	Rating              *int               `json:"rating,omitempty" bson:"rating,omitempty"`
	Comment             *string            `json:"comment,omitempty" bson:"comment,omitempty"`

	CreateTime     time.Time             `json:"createTime" bson:"createTime"`
	CompleteTime   *time.Time            `json:"completeTime" bson:"completeTime,omitempty"`
	LastUpdateTime time.Time             `json:"lastUpdateTime" bson:"lastUpdateTime"`
	State          model.CommissionState `json:"state" bson:"state"`
}

func NewCompletedCommission(comm model.Commission) CompletedCommission {

	return CompletedCommission{
		ID:                             comm.ID,
		OpenCommissionID:               comm.OpenCommissionID,
		ArtistID:                       comm.ArtistID,
		ArtistName:                     comm.ArtistName,
		ArtistProfilePath:              comm.ArtistProfilePath,
		RequesterID:                    comm.RequesterID,
		RequesterName:                  comm.RequesterName,
		RequesterProfilePath:           comm.RequesterProfilePath,
		Price:                          comm.Price,
		DayNeed:                        comm.DayNeed,
		Size:                           comm.Size,
		Resolution:                     comm.Resolution,
		ExportFormat:                   comm.ExportFormat,
		Desc:                           comm.Desc,
		PaymentMethod:                  comm.PaymentMethod,
		IsR18:                          comm.IsR18,
		BePrivate:                      comm.BePrivate,
		Anonymous:                      comm.Anonymous,
		RefImagePaths:                  comm.RefImagePaths,
		TimesAllowedDraftToChange:      comm.TimesAllowedDraftToChange,
		TimesAllowedCompletionToChange: comm.TimesAllowedCompletionToChange,
		DraftChangingRequestTime:       comm.DraftChangingRequestTime,
		ProofCopyRevisionRequestTime:   comm.ProofCopyRevisionRequestTime,
		ProofCopyImagePaths:            comm.ProofCopyImagePaths,
		DisplayImage:                   *comm.DisplayImage,
		CompletionFilePath:             comm.CompletionFilePath,
		Rating:                         comm.Rating,
		Comment:                        comm.Comment,
		CreateTime:                     comm.CreateTime,
		CompleteTime:                   comm.CompleteTime,
		LastUpdateTime:                 comm.LastUpdateTime,
		State:                          comm.State,
	}

}
