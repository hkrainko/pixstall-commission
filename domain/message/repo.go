package message

import (
	"context"
	"pixstall-commission/domain/message/model"
)

type Repo interface {
	CreateNewMessage(ctx context.Context, creator model.MessageCreator) (model.Message, error)
}