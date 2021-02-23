package msg_broker

import (
	"context"
	"pixstall-commission/domain/commission/model"
)

type Repo interface {
	SendCommissionCreatedMessage(ctx context.Context, commission model.Commission) error
}