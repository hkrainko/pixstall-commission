package model

import "time"

type CommissionFilter struct {
	ArtistID       *string
	RequesterID    *string
	Count          *int64
	Offset         *int64
	PriceFrom      *Price
	PriceTo        *Price
	DayNeedFrom    *DayNeed
	DayNeedTo      *DayNeed
	CreateTimeFrom *time.Time
	CreateTimeTo   *time.Time
	State          *CommissionState
}
