package model

type SystemMessage struct {
	Message `json:",inline" bson:",inline"`
	Text    string `json:"text" bson:"text"`
}
