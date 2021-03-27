package main

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"pixstall-commission/app/comm-msg-delivery/delivery/ws"
	"pixstall-commission/app/middleware"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	//Mongo
	dbClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	defer cancel()
	defer func() {
		if err = dbClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	db := dbClient.Database("pixstall-commission")

	// WebSocket
	hub := ws.NewHub()

	//gRPC
	grpcConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer grpcConn.Close()

	//RabbitMQ
	rbMQConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ %v", err)
	}
	defer rbMQConn.Close()
	ch, err := rbMQConn.Channel()
	if err != nil {
		log.Fatalf("Failed to create channel %v", err)
	}
	err = ch.ExchangeDeclare(
		"commission",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to create exchange %v", err)
	}
	err = ch.ExchangeDeclare(
		"comm-msg",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to create exchange %v", err)
	}

	commMsgBroker := InitCommissionMessageBroker(db, rbMQConn, grpcConn, hub)
	go commMsgBroker.StartUpdateCommissionQueue()
	go commMsgBroker.StartCommissionValidatedQueue()
	go commMsgBroker.StartCommissionMessageDeliverQueue()
	defer commMsgBroker.StopAllQueues()

	// Gin
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Access-Control-Allow-Origin", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowWildcard:    true,
		AllowFiles:       true,
		MaxAge:           12 * time.Hour,
	}))

	userIDExtractor := middleware.NewJWTPayloadsExtractor([]string{"userId"})

	{
		ctrl := InitCommissionMessageController(db, grpcConn, rbMQConn, hub)
		go ctrl.Run()
		//r.GET("/ws", userIDExtractor.ExtractPayloadsFromJWT, func(c *gin.Context) {ws.ServeWS(hub, c)})
		r.GET("/ws", userIDExtractor.ExtractPayloadsFromJWTInQuery, ctrl.HandleConnection)
	}

	apiGroup := r.Group("/api")
	commissionGroup := apiGroup.Group("/commissions")
	{
		ctrl := InitCommissionController(db, grpcConn, rbMQConn, hub)
		commissionGroup.GET("", userIDExtractor.ExtractPayloadsFromJWTInHeader, ctrl.GetCommissions)
		commissionGroup.GET("/:id", userIDExtractor.ExtractPayloadsFromJWTInHeader, ctrl.GetCommission)
		commissionGroup.GET("/:id/details", userIDExtractor.ExtractPayloadsFromJWTInHeader, ctrl.GetCommissionDetails)
		commissionGroup.GET("/:id/messages", userIDExtractor.ExtractPayloadsFromJWTInHeader, ctrl.GetMessages)
		commissionGroup.POST("/:id/messages", userIDExtractor.ExtractPayloadsFromJWTInHeader, ctrl.CreateMessage)
		commissionGroup.POST("", userIDExtractor.ExtractPayloadsFromJWTInHeader, ctrl.AddCommission)
		commissionGroup.PATCH("/:id", userIDExtractor.ExtractPayloadsFromJWTInHeader, ctrl.UpdateCommission)
	}

	err = r.Run(":9004")
	print(err)
}
