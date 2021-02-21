//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"pixstall-commission/app/commission/delivery/http"
	comm_deli_rabbitmq "pixstall-commission/app/commission/delivery/rabbitmq"
	comm_repo "pixstall-commission/app/commission/repo/mongo"
	"pixstall-commission/app/commission/usecase"
	aws_s3 "pixstall-commission/app/image/aws-s3"
)

func InitCommissionController(db *mongo.Database, awsS3 *s3.S3) http.CommissionController {
	wire.Build(
		http.NewCommissionController,
		comm_repo.NewMongoCommissionRepo,
		usecase.NewCommissionUseCase,
		aws_s3.NewAWSS3ImageRepository,
	)
	return http.CommissionController{}
}

func InitCommissionMessageBroker(db *mongo.Database, conn *amqp.Connection, awsS3 *s3.S3) comm_deli_rabbitmq.CommissionMessageBroker {
	wire.Build(
		comm_deli_rabbitmq.NewCommissionMessageBroker,
		usecase.NewCommissionUseCase,
		comm_repo.NewMongoCommissionRepo,
		aws_s3.NewAWSS3ImageRepository,
	)
	return comm_deli_rabbitmq.CommissionMessageBroker{}
}
