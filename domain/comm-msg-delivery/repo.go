package comm_msg_delivery

import (
	"context"
	"pixstall-commission/domain/message/model"
)

type Repo interface {
	DeliverCommissionMessage(ctx context.Context, messaging model.Messaging) error
}