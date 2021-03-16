package model

type SystemMessage struct {
	Message           `json:",inline" bson:",inline"`
	Text              string            `json:"text" bson:"text"`
	SystemMessageType SystemMessageType `json:"systemMessageType" bson:"systemMessageType"`
}

type PlainSystemMessage struct {
	SystemMessage `json:",inline" bson:",inline"`
}

type ProofCopySystemMessage struct {
	SystemMessage `json:",inline" bson:",inline"`
	ImagePath     string `json:"imagePath" bson:"imagePath"`
}

type CompletionSystemMessage struct {
	SystemMessage `json:",inline" bson:",inline"`
	FilePath      string  `json:"filePath" bson:"filePath"`
}

type SystemMessageType string

const (
	SystemMessageTypePlain      SystemMessageType = "Plain"
	SystemMessageTypeProofCopy  SystemMessageType = "ProofCopy"
	SystemMessageTypeCompletion SystemMessageType = "Completion"
)
