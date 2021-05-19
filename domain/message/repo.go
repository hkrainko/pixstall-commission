package message

import (
	"context"
	"pixstall-commission/domain/message/model"
)

type Repo interface {
	AddNewMessage(ctx context.Context, userId *string, messaging model.Messaging) error
	GetMessages(ctx context.Context, userId string, filter model.MessageFilter) ([]model.Messaging, error)
}