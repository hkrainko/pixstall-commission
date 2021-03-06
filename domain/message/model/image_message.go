package model

type ImageMessage struct {
	Message   `json:",inline" bson:",inline"`
	From      string  `json:"from" bson:"from"`
	Text      *string `json:"text" bson:"text"`
	ImagePath string  `json:"imagePath" bson:"imagePath"`
}