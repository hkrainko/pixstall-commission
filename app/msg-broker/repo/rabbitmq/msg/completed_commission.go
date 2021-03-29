package msg

import (
	"pixstall-commission/domain/commission/model"
	model2 "pixstall-commission/domain/file/model"
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

	DayUsed    time.Duration `json:"dayUsed"`
	Size       model2.Size   `json:"size"`
	Volume     float64       `json:"volume"`
	Resolution float64       `json:"resolution"`
	Format     string        `json:"format"`
	IsR18      bool          `json:"isR18"`
	Anonymous  bool          `json:"anonymous"`

	DisplayImagePath   string  `json:"displayImagePath"`
	CompletionFilePath string  `json:"completionFilePath"`
	Rating             int     `json:"rating"`
	Comment            *string `json:"comment,omitempty"`

	CreateTime   time.Time `json:"createTime" bson:"createTime"`
	CompleteTime time.Time `json:"completeTime" bson:"completeTime,omitempty"`
}

func NewCompletedCommission(comm model.Commission) CompletedCommission {

	return CompletedCommission{
		ID:                   comm.ID,
		OpenCommissionID:     comm.OpenCommissionID,
		ArtistID:             comm.ArtistID,
		ArtistName:           comm.ArtistName,
		ArtistProfilePath:    comm.ArtistProfilePath,
		RequesterID:          comm.RequesterID,
		RequesterName:        comm.RequesterName,
		RequesterProfilePath: comm.RequesterProfilePath,
		DayUsed:              0,
		Size:                 comm.Size,
		Volume:               0,
		Resolution:           0,
		Format:               "",
		IsR18:                comm.IsR18,
		Anonymous:            comm.Anonymous,
		DisplayImagePath:     comm.DisplayImage,
		CompletionFilePath:   *comm.CompletionFilePath,
		Rating:               0,
		Comment:              nil,
		CreateTime:           time.Time{},
		CompleteTime:         time.Time{},
	}

}
