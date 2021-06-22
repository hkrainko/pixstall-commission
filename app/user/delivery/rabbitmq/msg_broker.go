package rabbitmq

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"pixstall-commission/domain/commission"
	"pixstall-commission/domain/user/model"
	"time"
)

type UserMessageBroker struct {
	commUseCase commission.UseCase
	ch          *amqp.Channel
}

func NewUserMessageBroker(commUseCase commission.UseCase, conn *amqp.Connection) UserMessageBroker {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel %v", err)
	}
	err = ch.Qos(5, 0, false)
	if err != nil {
		log.Fatalf("Failed to set QoS %v", err)
	}

	return UserMessageBroker{
		commUseCase: commUseCase,
		ch:          ch,
	}
}

func (u UserMessageBroker) StartUserEventQueue() {
	q, err := u.ch.QueueDeclare(
		"user-event-to-comm", // name
		true,                // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue %v", err)
	}
	err = u.ch.QueueBind(
		q.Name,
		"user.event.updated",
		"user",
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to bind queue %v", err)
	}
	msgs, err := u.ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer %v", err)
	}
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			d.Ack(false)

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			go func() {
				for {
					select {
					case <-ctx.Done():
						switch ctx.Err() {
						case context.DeadlineExceeded:
							log.Println("context.DeadlineExceeded")
						case context.Canceled:
							log.Println("context.Canceled")
						default:
							log.Println("default")
						}
						return // returning not to leak the goroutine
					}
				}
			}()

			switch d.RoutingKey {
			case "user.event.updated":
				err := u.consumeUserEventUpdated(ctx, d.Body)
				if err != nil {
					//TODO: error handling, store it ?
				}
				cancel()
			default:
				cancel()
			}
		}
	}()
	<-forever
}

func (u UserMessageBroker) StopAllQueues() {
	err := u.ch.Close()
	if err != nil {
		log.Printf("StopUserQueue err %v", err)
	}
	log.Printf("StopUserQueue success")
}

// Private handler
func (u UserMessageBroker) consumeUserEventUpdated(ctx context.Context, body []byte) error {
	req := model.UserUpdater{}
	err := json.Unmarshal(body, &req)
	if err != nil {
		return err
	}
	return u.commUseCase.UpdateCommissionByUserUpdatedEvent(ctx, req)
}