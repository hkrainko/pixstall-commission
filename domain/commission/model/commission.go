package model

import (
	"pixstall-commission/domain/message/model"
	"time"
)

type Commission struct {
	ID               string `json:"id" bson:"id"`
	OpenCommissionID string `json:"openCommissionId" bson:"openCommissionId"`

	ArtistID             string  `json:"artistId" bson:"artistId"`
	ArtistName           string  `json:"artistName" bson:"artistName"`
	ArtistProfilePath    *string `json:"artistProfilePath" bson:"artistProfilePath,omitempty"`
	RequesterID          string  `json:"requesterId" bson:"requesterId"`
	RequesterName        string  `json:"requesterName" bson:"requesterName"`
	RequesterProfilePath *string `json:"requesterProfilePath" bson:"requesterProfilePath,omitempty"`

	Price         Price    `json:"price" bson:"price"`
	DayNeed       int      `json:"dayNeed" bson:"dayNeed"`
	Size          *Size    `json:"size" bson:"size,omitempty"`
	Resolution    *float64 `json:"resolution" bson:"resolution,omitempty"`
	ExportFormat  *string  `json:"exportFormat" bson:"exportFormat,omitempty"`
	Desc          string   `json:"desc" bson:"desc"`
	PaymentMethod string   `json:"paymentMethod" bson:"paymentMethod"`
	IsR18         bool     `json:"isR18" bson:"isR18"`
	BePrivate     bool     `json:"bePrivate" bson:"bePrivate"`
	Anonymous     bool     `json:"anonymous" bson:"anonymous"`
	RefImagePaths []string `json:"refImagePaths" bson:"refImagePaths"`

	Messages []model.Message `json:"messages" bson:"messages"`

	CreateTime     time.Time       `json:"createTime" bson:"createTime"`
	CompleteTime   *time.Time      `json:"completeTime" bson:"completeTime,omitempty"`
	LastUpdateTime time.Time       `json:"lastUpdateTime" bson:"lastUpdateTime"`
	State          CommissionState `json:"state" bson:"state"`
}

type CommissionState string

const (
	CommissionStatePendingValidation     CommissionState = "PENDING_VALIDATION"
	CommissionStateInValid               CommissionState = "INVALID"
	CommissionStatePendingArtistApproval CommissionState = "PENDING_ARTIST_APPROVAL"
	CommissionStateRejectByArtist        CommissionState = "REJECT_BY_ARTIST"
	CommissionStateRejectByClient        CommissionState = "REJECT_BY_CLIENT"
)
