package msg

import (
	"pixstall-commission/domain/commission/model"
	"time"
)

type CompletedCommission struct {
	ID               string `json:"id"`
	OpenCommissionID string `json:"openCommissionId"`

	ArtistID             string  `json:"artistId"`
	ArtistName           string  `json:"artistName"`
	ArtistProfilePath    *string `json:"artistProfilePath,omitempty"`
	RequesterID          string  `json:"requesterId"`
	RequesterName        string  `json:"requesterName"`
	RequesterProfilePath *string `json:"requesterProfilePath,omitempty"`

	DayUsed   time.Duration `json:"dayUsed"`
	IsR18     bool          `json:"isR18"`
	Anonymous bool          `json:"anonymous"`

	DisplayImage       model.DisplayImage `json:"displayImage"`
	CompletionFilePath string             `json:"completionFilePath"`
	Rating             int                `json:"rating"`
	Comment            *string            `json:"comment,omitempty"`

	CreateTime    time.Time `json:"createTime" bson:"createTime"`
	CompletedTime time.Time `json:"completedTime" bson:"completedTime,omitempty"`
}

func NewCompletedCommission(comm model.Commission) CompletedCommission {

	dayUsed := comm.CompletedTime.Sub(comm.CreateTime)

	return CompletedCommission{
		ID:                   comm.ID,
		OpenCommissionID:     comm.OpenCommissionID,
		ArtistID:             comm.ArtistID,
		ArtistName:           comm.ArtistName,
		ArtistProfilePath:    comm.ArtistProfilePath,
		RequesterID:          comm.RequesterID,
		RequesterName:        comm.RequesterName,
		RequesterProfilePath: comm.RequesterProfilePath,
		DayUsed:              dayUsed,
		IsR18:                comm.IsR18,
		Anonymous:            comm.Anonymous,
		DisplayImage:         *comm.DisplayImage,
		CompletionFilePath:   *comm.CompletionFilePath,
		Rating:               *comm.Rating,
		Comment:              comm.Comment,
		CreateTime:           comm.CreateTime,
		CompletedTime:         *comm.CompletedTime,
	}

}
