package msg

import "pixstall-commission/domain/message/model"

type WSMessageCreator struct {
	Type                 WSMessageType `json:"type"`
	model.MessageCreator `json:"body"`
}

type WSMessageType string

const (
	WSMessageTypeChat WSMessageType = "chat"
	WSMessageTypeCMD  WSMessageType = "cmd"
)
