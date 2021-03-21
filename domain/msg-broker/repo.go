package msg_broker

import (
	"context"
	"pixstall-commission/domain/commission/model"
	model2 "pixstall-commission/domain/message/model"
)

type Repo interface {
	SendCommissionCreatedMessage(ctx context.Context, commission model.Commission) error
	SendCommissionCompletedMessage(ctx context.Context, commission model.Commission) error
	SendCommissionMessageReceivedMessage(ctx context.Context, messaging model2.Messaging) error
}