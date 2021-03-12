package message

import (
	"context"
	"pixstall-commission/domain/message/model"
)

type Repo interface {
	AddNewMessage(ctx context.Context, userId *string, messaging model.Messaging) error
	GetMessages(ctx context.Context, userId string, commId string, offset int, count int) ([]model.Messaging, error)
}