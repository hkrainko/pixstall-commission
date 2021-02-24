// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"pixstall-commission/app/commission/delivery/http"
	rabbitmq2 "pixstall-commission/app/commission/delivery/rabbitmq"
	mongo2 "pixstall-commission/app/commission/repo/mongo"
	"pixstall-commission/app/commission/usecase"
	"pixstall-commission/app/image/aws-s3"
	"pixstall-commission/app/msg-broker/repo/rabbitmq"
)

// Injectors from wire.go:

func InitCommissionController(db *mongo.Database, awsS3 *s3.S3, conn *amqp.Connection) http.CommissionController {
	repo := mongo2.NewMongoCommissionRepo(db)
	imageRepo := aws_s3.NewAWSS3ImageRepository(awsS3)
	msg_brokerRepo := rabbitmq.NewRabbitMQMsgBrokerRepo(conn)
	useCase := usecase.NewCommissionUseCase(repo, imageRepo, msg_brokerRepo)
	commissionController := http.NewCommissionController(useCase)
	return commissionController
}

func InitCommissionMessageBroker(db *mongo.Database, conn *amqp.Connection, awsS3 *s3.S3) rabbitmq2.CommissionMessageBroker {
	repo := mongo2.NewMongoCommissionRepo(db)
	imageRepo := aws_s3.NewAWSS3ImageRepository(awsS3)
	msg_brokerRepo := rabbitmq.NewRabbitMQMsgBrokerRepo(conn)
	useCase := usecase.NewCommissionUseCase(repo, imageRepo, msg_brokerRepo)
	commissionMessageBroker := rabbitmq2.NewCommissionMessageBroker(useCase, conn)
	return commissionMessageBroker
}
