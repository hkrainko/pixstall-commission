package model

import (
	"pixstall-commission/domain/file/model"
)

type CommissionCreator struct {
	OpenCommissionID string            `json:"openCommissionId"`
	ArtistID         string            `json:"artistID"`
	RequesterID      string            `json:"requesterID"`
	Price            Price             `json:"price"`
	DayNeed          int               `json:"dayNeed"`
	Size             *model.Size             `json:"size"`
	Resolution       *float64          `json:"resolution"`
	ExportFormat     *string           `json:"exportFormat"`
	Desc             string            `json:"desc"`
	PaymentMethod    string            `json:"paymentMethod"`
	IsR18            bool              `json:"isR18"`
	BePrivate        bool              `json:"bePrivate"`
	Anonymous        bool              `json:"anonymous"`
	RefImages        []model.ImageFile `json:"refImages"`
	RefImagePaths    []string          `json:"refImagePaths"`
}

type Price struct {
	Amount   float64  `json:"amount" bson:"amount"`
	Currency Currency `json:"currency" bson:"currency"`
}

type DayNeed struct {
	From int `json:"from" bson:"from"`
	To   int `json:"to" bson:"to"`
}

type Currency string

const (
	CurrencyHKD Currency = "HKD"
	CurrencyTWD Currency = "TWD"
	CurrencyUSE Currency = "USD"
)
