package ws

import (
	"context"
	"pixstall-commission/app/commission/delivery/ws"
	commMsgDelivery "pixstall-commission/domain/comm-msg-delivery"
	"pixstall-commission/domain/message/model"
)

type wsCommMsgDeliveryRepo struct {
	hub *ws.Hub
}

func NewWSCommMsgDeliveryRepo(hub *ws.Hub) commMsgDelivery.Repo {
	return &wsCommMsgDeliveryRepo{
		hub: hub,
	}
}

func (w wsCommMsgDeliveryRepo) DeliverCommissionMessage(ctx context.Context, messaging model.Messaging) error {
	return w.hub.DeliverCommissionMessage(ctx, messaging)
}