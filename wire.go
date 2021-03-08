//+build wireinject

package main

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/wire"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"pixstall-commission/app/comm-msg-delivery/delivery/ws"
	comm_msg_deli_repo "pixstall-commission/app/comm-msg-delivery/repo/ws"
	"pixstall-commission/app/commission/delivery/http"
	comm_deli_rabbitmq "pixstall-commission/app/commission/delivery/rabbitmq"
	comm_repo "pixstall-commission/app/commission/repo/mongo"
	"pixstall-commission/app/commission/usecase"
	aws_s3 "pixstall-commission/app/image/aws-s3"
	msg_repo "pixstall-commission/app/message/repo/mongo"
	msg_broker_repo "pixstall-commission/app/msg-broker/repo/rabbitmq"
)

func InitCommissionMessageController(db *mongo.Database, awsS3 *s3.S3, conn *amqp.Connection, hub *ws.Hub) ws.CommissionMessageController {
	wire.Build(
		ws.NewCommissionMessageController,
		comm_repo.NewMongoCommissionRepo,
		usecase.NewCommissionUseCase,
		aws_s3.NewAWSS3ImageRepository,
		msg_broker_repo.NewRabbitMQMsgBrokerRepo,
		msg_repo.NewMongoMessageRepo,
		comm_msg_deli_repo.NewWSCommMsgDeliveryRepo,
	)
	return ws.CommissionMessageController{}
}

func InitCommissionController(db *mongo.Database, awsS3 *s3.S3, conn *amqp.Connection, hub *ws.Hub) http.CommissionController {
	wire.Build(
		http.NewCommissionController,
		comm_repo.NewMongoCommissionRepo,
		usecase.NewCommissionUseCase,
		aws_s3.NewAWSS3ImageRepository,
		msg_broker_repo.NewRabbitMQMsgBrokerRepo,
		msg_repo.NewMongoMessageRepo,
		comm_msg_deli_repo.NewWSCommMsgDeliveryRepo,
	)
	return http.CommissionController{}
}

func InitCommissionMessageBroker(db *mongo.Database, conn *amqp.Connection, awsS3 *s3.S3, hub *ws.Hub) comm_deli_rabbitmq.CommissionMessageBroker {
	wire.Build(
		comm_deli_rabbitmq.NewCommissionMessageBroker,
		usecase.NewCommissionUseCase,
		comm_repo.NewMongoCommissionRepo,
		aws_s3.NewAWSS3ImageRepository,
		msg_broker_repo.NewRabbitMQMsgBrokerRepo,
		msg_repo.NewMongoMessageRepo,
		comm_msg_deli_repo.NewWSCommMsgDeliveryRepo,
	)
	return comm_deli_rabbitmq.CommissionMessageBroker{}
}
