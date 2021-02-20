package model

import (
	"time"
)

type Commission struct {
	ID               string `json:"id"`
	OpenCommissionID string `json:"openCommissionId"`

	ArtistID             string  `json:"artistId"`
	ArtistName           string  `json:"artistName"`
	ArtistProfilePath    *string `json:"artistProfilePath"`
	RequesterID          string  `json:"requesterId"`
	RequesterName        string  `json:"requesterName"`
	RequesterProfilePath *string `json:"requesterProfilePath"`

	Price         Price    `json:"price"`
	DayNeed       int      `json:"dayNeed"`
	Size          *Size    `json:"size"`
	Resolution    *float64 `json:"resolution"`
	ExportFormat  *string  `json:"exportFormat"`
	Desc          string   `json:"desc"`
	PaymentMethod string   `json:"paymentMethod"`
	IsR18         bool     `json:"isR18"`
	BePrivate     bool     `json:"bePrivate"`
	Anonymous     bool     `json:"anonymous"`
	RefImagePaths []string `json:"refImagePaths"`

	CreateTime     time.Time       `json:"createTime"`
	CompleteTime   *time.Time      `json:"completeTime"`
	LastUpdateTime time.Time       `json:"lastUpdateTime"`
	State          CommissionState `json:"state"`
}

type CommissionState string

const (
	CommissionStatePending        CommissionState = "P"
	CommissionStateRejectByArtist CommissionState = "RJ_BY_ARTIST"
	CommissionStateRejectByClient CommissionState = "RJ_BY_CLIENT"
)
