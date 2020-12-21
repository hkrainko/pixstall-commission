package model

import (
	"time"
)

type Commission struct {
	ID               string
	OpenCommissionID string
	ArtistID         string
	ClientID         string
	BidPrice         float64
	Size             Size
	RequestDate      time.Time
	CompleteTime     time.Time
	State            CommissionState
}

type CommissionState string

const (
	CommissionStatePending        CommissionState = "P"
	CommissionStateRejectByArtist CommissionState = "RJ_BY_ARTIST"
	CommissionStateRejectByClient CommissionState = "RJ_BY_CLIENT"
)

type Size struct {
	Width  float64
	Height float64
}
