package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	model2 "pixstall-commission/app/msg-broker/repo/rabbitmq/msg"
	"pixstall-commission/domain/commission/model"
	dMsgModel "pixstall-commission/domain/message/model"
	msg_broker "pixstall-commission/domain/msg-broker"
)

type rabbitmqMsgBrokerRepo struct {
	ch *amqp.Channel
}

func NewRabbitMQMsgBrokerRepo(conn *amqp.Connection) msg_broker.Repo {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel %v", err)
	}
	err = ch.Qos(5, 0, false)
	if err != nil {
		log.Fatalf("Failed to set QoS %v", err)
	}
	return rabbitmqMsgBrokerRepo{
		ch: ch,
	}
}

func (r rabbitmqMsgBrokerRepo) SendCommissionCreatedMessage(ctx context.Context, commission model.Commission) error {
	cComm := model2.CreatedCommission{Commission: commission}
	b, err := json.Marshal(cComm)
	if err != nil {
		return err
	}
	err = r.ch.Publish(
		"commission",
		"commission.event.created",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        b,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (r rabbitmqMsgBrokerRepo) SendCommissionCompletedMessage(ctx context.Context, commission model.Commission) error {
	cComm := model2.NewCompletedCommission(commission)
	b, err := json.Marshal(cComm)
	if err != nil {
		return err
	}
	err = r.ch.Publish(
		"commission",
		"commission.event.completed",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        b,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (r rabbitmqMsgBrokerRepo) SendCommissionMessageReceivedMessage(ctx context.Context, messaging dMsgModel.Messaging) error {
	b, err := json.Marshal(messaging)
	fmt.Printf("sent msg:%v", string(b))
	if err != nil {
		return err
	}
	err = r.ch.Publish(
		"comm-msg",
		"comm-msg.event.received",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        b,
		},
	)
	if err != nil {
		return err
	}
	return nil
}