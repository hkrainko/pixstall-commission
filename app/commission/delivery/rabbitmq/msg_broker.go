package rabbitmq

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	create_commission "pixstall-commission/app/commission/delivery/model/create-commission"
	"pixstall-commission/domain/commission"
	"time"
)

type CommissionMessageBroker struct {
	commUseCase commission.UseCase
	ch          *amqp.Channel
}

func NewCommissionMessageBroker(commUseCase commission.UseCase, conn *amqp.Connection) CommissionMessageBroker {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel %v", err)
	}
	err = ch.Qos(5, 0, false)
	if err != nil {
		log.Fatalf("Failed to set QoS %v", err)
	}

	return CommissionMessageBroker{
		commUseCase: commUseCase,
		ch:          ch,
	}
}

func (c CommissionMessageBroker) StartQueue() {
	//TODO
	q, err := c.ch.QueueDeclare(
		"commission_artist", // name
		true,                             // durable
		false,                            // delete when unused
		false,                            // exclusive
		false,                            // no-wait
		nil,                              // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue %v", err)
	}
	err = c.ch.QueueBind(
		q.Name,
		"commission.new",
		"commission",
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to bind queue %v", err)
	}

	msgs, err := c.ch.Consume(
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
			case "commission.new":
				err := c.CreateCommission(ctx, d.Body)
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

func (c CommissionMessageBroker) StopQueue() {
	err := c.ch.Close()
	if err != nil {
		log.Printf("StopCommissionQueue err %v", err)
	}
	log.Printf("StopCommissionQueue success")
}

func (c CommissionMessageBroker) CreateCommission(ctx context.Context, body []byte) error {
	req := create_commission.Request{}
	err := json.Unmarshal(body, &req)
	if err != nil {
		return err
	}
	return nil
}
