package model

type SystemMessage struct {
	Message           `json:",inline" bson:",inline"`
	Text              string            `json:"text" bson:"text"`
	SystemMessageType SystemMessageType `json:"systemMessageType" bson:"systemMessageType"`
}

type PlainSystemMessage struct {
	SystemMessage `json:",inline" bson:",inline"`
}

type UploadProofCopySystemMessage struct {
	SystemMessage `json:",inline" bson:",inline"`
	ImagePath     string `json:"imagePath" bson:"imagePath"`
}

type UploadProductSystemMessage struct {
	SystemMessage    `json:",inline" bson:",inline"`
	FilePath         string `json:"filePath" bson:"filePath"`
	DisplayImagePath string `json:"displayImagePath" bson:"displayImagePath"`
}

type AcceptProductSystemMessage struct {
	SystemMessage `json:",inline" bson:",inline"`
	Rating        int     `json:"rating,omitempty" bson:"rating,omitempty"`
	Comment       *string `json:"comment,omitempty" bson:"comment,omitempty"`
}

type SystemMessageType string

const (
	SystemMessageTypePlain           SystemMessageType = "Plain"
	SystemMessageTypeUploadProofCopy SystemMessageType = "UploadProofCopy"
	SystemMessageTypeUploadProduct   SystemMessageType = "UploadProduct"
	SystemMessageTypeAcceptProduct   SystemMessageType = "AcceptProduct"
)
