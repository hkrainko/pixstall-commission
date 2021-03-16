package rabbitmq

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/streadway/amqp"
	"log"
	"pixstall-commission/app/commission/delivery/rabbitmq/msg"
	"pixstall-commission/domain/commission"
	"pixstall-commission/domain/message/model"
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

func (c CommissionMessageBroker) StartUpdateCommissionQueue() {
	//TODO
	q, err := c.ch.QueueDeclare(
		"commission-update", // name
		true,                // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue %v", err)
	}
	err = c.ch.QueueBind(
		q.Name,
		"commission.cmd.update",
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
			case "commission.cmd.update":
				err := c.updateCommission(ctx, d.Body)
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

func (c CommissionMessageBroker) StartCommissionValidatedQueue() {
	//TODO
	q, err := c.ch.QueueDeclare(
		"commission-validated", // name
		true,                   // durable
		false,                  // delete when unused
		false,                  // exclusive
		false,                  // no-wait
		nil,                    // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue %v", err)
	}
	err = c.ch.QueueBind(
		q.Name,
		"commission.event.validation.#",
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
			case "commission.event.validation.open-comm":
				err := c.commOpenCommValidated(ctx, d.Body)
				if err != nil {
					//TODO: error handling, store it ?
				}
				cancel()
			case "commission.event.validation.users":
				err := c.commUsersValidated(ctx, d.Body)
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

func (c CommissionMessageBroker) StartCommissionMessageDeliverQueue() {
	//TODO
	q, err := c.ch.QueueDeclare(
		"",    // name
		false, // durable
		true,  // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue %v", err)
	}
	err = c.ch.QueueBind(
		q.Name,
		"comm-msg.event.#",
		"comm-msg",
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
			case "comm-msg.event.received":
				err := c.commMsgDeliver(ctx, d.Body)
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

func (c CommissionMessageBroker) StopAllQueues() {
	err := c.ch.Close()
	if err != nil {
		log.Printf("StopCommissionQueue err %v", err)
	}
	log.Printf("StopCommissionQueue success")
}

// Private handler
func (c CommissionMessageBroker) updateCommission(ctx context.Context, body []byte) error {
	req := msg.CommissionUpdater{}
	err := json.Unmarshal(body, &req)
	if err != nil {
		return err
	}
	return c.commUseCase.UpdateCommission(ctx, req.Updater)
}

func (c CommissionMessageBroker) commOpenCommValidated(ctx context.Context, body []byte) error {
	req := msg.CommissionOpenCommissionValidation{}
	err := json.Unmarshal(body, &req)
	if err != nil {
		return err
	}
	return c.commUseCase.OpenCommissionValidation(ctx, req.CommissionOpenCommissionValidation)
}

func (c CommissionMessageBroker) commUsersValidated(ctx context.Context, body []byte) error {
	req := msg.CommissionUsersValidation{}
	err := json.Unmarshal(body, &req)
	if err != nil {
		return err
	}
	return c.commUseCase.UsersValidation(ctx, req.CommissionUsersValidation)
}

func (c CommissionMessageBroker) commMsgDeliver(ctx context.Context, body []byte) error {
	messaging, err := c.getMessage(body)
	if err != nil {
		return err
	}
	return c.commUseCase.HandleOutBoundCommissionMessage(ctx, messaging)
}

// Private utilities

func (c CommissionMessageBroker) getMessage(body []byte) (model.Messaging, error) {
	req := msg.CommissionMessage{}
	err := json.Unmarshal(body, &req)
	if err != nil {
		return nil, err
	}
	switch req.Message.MessageType {
	case model.MessageTypeText:
		var result model.TextMessage
		err := json.Unmarshal(body, &result)
		if err != nil {
			return nil, err
		}
		return &result, nil
	case model.MessageTypeImage:
		var result model.ImageMessage
		err := json.Unmarshal(body, &result)
		if err != nil {
			return nil, err
		}
		return &result, nil
	case model.MessageTypeSystem:
		var result model.SystemMessage
		err := json.Unmarshal(body, &result)
		if err != nil {
			return nil, err
		}
		return c.getSystemMessage(body, result.SystemMessageType)
	default:
		break
	}
	return nil, errors.New("unknown message type")
}

func (c CommissionMessageBroker) getSystemMessage(body []byte, systemMessageType model.SystemMessageType) (model.Messaging, error) {
	switch systemMessageType {
	case model.SystemMessageTypePlain:
		var result model.PlainSystemMessage
		err := json.Unmarshal(body, &result)
		if err != nil {
			return nil, err
		}
		return &result, nil
	case model.SystemMessageTypeProofCopy:
		var result model.ProofCopySystemMessage
		err := json.Unmarshal(body, &result)
		if err != nil {
			return nil, err
		}
		return &result, nil
	case model.SystemMessageTypeCompletion:
		var result model.CompletionSystemMessage
		err := json.Unmarshal(body, &result)
		if err != nil {
			return nil, err
		}
		return &result, nil
	default:
		return nil, errors.New("unknown message type")
	}
}
