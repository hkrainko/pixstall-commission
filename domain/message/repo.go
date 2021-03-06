package message

import (
	"context"
	"pixstall-commission/domain/message/model"
	"time"
)

type Repo interface {
	AddNewMessage(ctx context.Context, creator model.MessageCreator) (*model.Message, error)
	GetMessage(ctx context.Context, commId string, lastMsgTime time.Time, count int) (*model.Message, error)
}