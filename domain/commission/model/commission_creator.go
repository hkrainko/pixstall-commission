package model

import (
	"image"
)

type CommissionCreator struct {
	OpenCommissionID string        `json:"openCommissionId"`
	ArtistID         string        `json:"artistID"`
	RequesterID      string        `json:"requesterID"`
	Price            Price         `json:"price"`
	DayNeed          int           `json:"dayNeed"`
	Size             *Size         `json:"size"`
	Resolution       *float64      `json:"resolution"`
	ExportFormat     *string       `json:"exportFormat"`
	Desc             string        `json:"desc"`
	PaymentMethod    string        `json:"paymentMethod"`
	IsR18            bool          `json:"isR18"`
	BePrivate        bool          `json:"bePrivate"`
	Anonymous        bool          `json:"anonymous"`
	RefImages        []image.Image `json:"refImages"`
	RefImagePaths    []string      `json:"refImagePaths"`
}

type Price struct {
	Amount   float64  `json:"amount" bson:"amount"`
	Currency Currency `json:"currency" bson:"currency"`
}

type DayNeed struct {
	From int `json:"from" bson:"from"`
	To   int `json:"to" bson:"to"`
}

type Size struct {
	Width  float64 `json:"width" bson:"width"`
	Height float64 `json:"height" bson:"height"`
	Unit   string  `json:"unit" bson:"unit"`
}

type Currency string

const (
	CurrencyHKD Currency = "HKD"
	CurrencyTWD Currency = "TWD"
	CurrencyUSE Currency = "USD"
)
