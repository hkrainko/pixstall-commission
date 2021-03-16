package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"pixstall-commission/app/comm-msg-delivery/delivery/ws/msg"
	"pixstall-commission/domain/commission"
)

type CommissionMessageController struct {
	hub *Hub
	broadcast chan UserMessage
	register chan *Client
	unregister chan *Client
	commUseCase commission.UseCase
}

func NewCommissionMessageController(commUseCase commission.UseCase, hub *Hub) CommissionMessageController {
	return CommissionMessageController{
		hub: hub,
		broadcast:  make(chan UserMessage),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		commUseCase: commUseCase,
	}
}

func (c *CommissionMessageController) HandleConnection(ctx *gin.Context) {
	tokenUserID := ctx.GetString("userId")
	if tokenUserID == "" {
		return
	}
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{
		broadcast: c.broadcast,
		unregister: c.unregister,
		conn:   conn,
		send:   make(chan []byte, 256),
		userId: tokenUserID,
	}
	c.register <- client
	go client.writePump()
	go client.readPump()
}

func (c *CommissionMessageController) Run() {
	for {
		select {
		case client := <-c.register:
			c.hub.clients[client] = true
		case client := <-c.unregister:
			if _, ok := c.hub.clients[client]; ok {
				delete(c.hub.clients, client)
				close(client.send)
			}
		case message := <-c.broadcast:
			// TODO: write to usecase
			fmt.Printf("in msg userId:%v msg:%v\n", message.UserID, string(message.Byte))

			var dat = msg.WSMessage{}
			err := json.Unmarshal(message.Byte, &dat)
			if err != nil {
				fmt.Printf("err: %v\n", err.Error())
				continue
			}
			msgType := dat.Type

			//Not implemented as message will be sent by http
			switch msgType {
			case "chat":
				fmt.Printf("cmd")
				//var wsCreator = msg.WSMessageCreator{}
				//err := json.Unmarshal(message.Byte, &wsCreator)
				//if err != nil {
				//	fmt.Printf("err2: %v\n", err.Error())
				//	continue
				//}
				//wsCreator.Form = &message.UserID // memory leak?
				//fmt.Printf("%v", wsCreator)
				//ctx := context.Background()
				//err = c.commUseCase.HandleInboundCommissionMessage(ctx, wsCreator.MessageCreator)
				//if err != nil {
				//	// TODO: reply error to client
				//}
			case "cmd":
				fmt.Printf("cmd")
			default:
				continue
			}
			//b := []byte("reply" + string(message.Byte))
			//for client := range c.hub.clients {
			//	select {
			//	case client.send <- b:
			//	default:
			//		close(client.send)
			//		delete(c.hub.clients, client)
			//	}
			//}
		}
	}
}