// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"pixstall-commission/domain/message/model"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
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
