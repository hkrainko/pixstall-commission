// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"pixstall-commission/app/commission/delivery/ws/msg"
	"pixstall-commission/domain/commission"
	"pixstall-commission/domain/message/model"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	commUseCase commission.UseCase
}

func NewHub(commUseCase commission.UseCase) *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		commUseCase: commUseCase,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			// TODO: write to usecase

			var dat map[string]interface{}
			err := json.Unmarshal(message, &dat)
			if err != nil {
				fmt.Printf("err: %v\n", err.Error())
				continue
			}
			msgType := dat["type"]
			switch msgType {
			case "chat":
				var wsCreator = msg.WSMessageCreator{}
				err := json.Unmarshal(message, &wsCreator)
				if err != nil {
					fmt.Printf("err2: %v\n", err.Error())
					continue
				}
				fmt.Printf("%v", wsCreator)
				ctx := context.Background()
				err = h.commUseCase.HandleInboundCommissionMessage(ctx, wsCreator.MessageCreator)
				if err != nil {
					// TODO: reply error to client
				}
			case "cmd":
				fmt.Printf("cmd")
			default:
				continue
			}
			b := []byte("reply" + string(message))
			for client := range h.clients {
				select {
				case client.send <- b:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

func (h *Hub) DeliverCommissionMessage(ctx context.Context, messaging model.Messaging) error {
	j, err := json.Marshal(messaging)
	if err != nil {
		return err
	}
	for client := range h.clients {
		if messaging.GetArtistID() == client.userId || messaging.GetRequesterID() == client.userId {
			fmt.Printf("deliver comm msg to %v, msg:%v", client.userId, messaging)
			client.send <- j
		}
	}
	return nil
}
