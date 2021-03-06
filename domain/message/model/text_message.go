package model

type TextMessage struct {
	Message `json:",inline" bson:",inline"`
	From    string `json:"from" bson:"from"`
	Text    string `json:"text" bson:"text"`
}