package main

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"pixstall-commission/app/middleware"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	//AWS s3
	awsAccessKey := "AKIA5BWICLKRWX6ARSEF"
	awsSecret := "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"
	token := ""
	creds := credentials.NewStaticCredentials(awsAccessKey, awsSecret, token)
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:                        aws.String(endpoints.ApEast1RegionID),
			CredentialsChainVerboseErrors: aws.Bool(true),
			Credentials:                   creds,
		},
		//Profile:                 "default", //[default], use [prod], [uat]
		//SharedConfigState:       session.SharedConfigEnable,
	}))
	awsS3 := s3.New(sess)

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

	//RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ %v", err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
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

	commMsgBroker := InitCommissionMessageBroker(db, conn, awsS3)
	go commMsgBroker.StartQueue()
	defer commMsgBroker.StopQueue()

	//Gin
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Access-Control-Allow-Origin", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowWildcard: true,
		AllowFiles: true,
		MaxAge: 12 * time.Hour,
	}))

	userIDExtractor := middleware.NewJWTPayloadsExtractor([]string{"userId"})

	apiGroup := r.Group("/api")

	commissionGroup := apiGroup.Group("/commissions")
	{
		ctrl := InitCommissionController(db, awsS3)
		commissionGroup.GET("", userIDExtractor.ExtractPayloadsFromJWT, ctrl.GetCommissions)
		commissionGroup.GET("/:id/details", userIDExtractor.ExtractPayloadsFromJWT, ctrl.GetCommissionDetails)
		commissionGroup.GET("/:id/messages", userIDExtractor.ExtractPayloadsFromJWT, ctrl.GetCommissionMessages)
		commissionGroup.POST("", userIDExtractor.ExtractPayloadsFromJWT, ctrl.AddCommission)
	}

	err = r.Run(":9004")
	print(err)
}