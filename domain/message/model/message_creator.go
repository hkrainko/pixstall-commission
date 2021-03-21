package model

import (
	"pixstall-commission/domain/file/model"
)

type MessageCreator struct {
	CommissionID string           `json:"commissionId"`
	Form         string           `json:"from"`
	Text         string           `json:"text"`
	Image        *model.ImageFile `json:"image"`
	ImagePath    *string          `json:"imagePath"`
}
