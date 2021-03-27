//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"pixstall-commission/app/comm-msg-delivery/delivery/ws"
	comm_msg_deli_repo "pixstall-commission/app/comm-msg-delivery/repo/ws"
	"pixstall-commission/app/commission/delivery/http"
	comm_deli_rabbitmq "pixstall-commission/app/commission/delivery/rabbitmq"
	comm_repo "pixstall-commission/app/commission/repo/mongo"
	"pixstall-commission/app/commission/usecase"
	grpc_repo "pixstall-commission/app/image/grpc"
	msg_repo "pixstall-commission/app/message/repo/mongo"
	msg_broker_repo "pixstall-commission/app/msg-broker/repo/rabbitmq"
)

func InitCommissionMessageController(db *mongo.Database, grpcConn *grpc.ClientConn, conn *amqp.Connection, hub *ws.Hub) ws.CommissionMessageController {
	wire.Build(
		ws.NewCommissionMessageController,
		comm_repo.NewMongoCommissionRepo,
		usecase.NewCommissionUseCase,
		grpc_repo.NewGRPCImageRepository,
		msg_broker_repo.NewRabbitMQMsgBrokerRepo,
		msg_repo.NewMongoMessageRepo,
		comm_msg_deli_repo.NewWSCommMsgDeliveryRepo,
	)
	return ws.CommissionMessageController{}
}

func InitCommissionController(db *mongo.Database, grpcConn *grpc.ClientConn, conn *amqp.Connection, hub *ws.Hub) http.CommissionController {
	wire.Build(
		http.NewCommissionController,
		comm_repo.NewMongoCommissionRepo,
		usecase.NewCommissionUseCase,
		grpc_repo.NewGRPCImageRepository,
		msg_broker_repo.NewRabbitMQMsgBrokerRepo,
		msg_repo.NewMongoMessageRepo,
		comm_msg_deli_repo.NewWSCommMsgDeliveryRepo,
	)
	return http.CommissionController{}
}

func InitCommissionMessageBroker(db *mongo.Database, conn *amqp.Connection, grpcConn *grpc.ClientConn, hub *ws.Hub) comm_deli_rabbitmq.CommissionMessageBroker {
	wire.Build(
		comm_deli_rabbitmq.NewCommissionMessageBroker,
		usecase.NewCommissionUseCase,
		comm_repo.NewMongoCommissionRepo,
		grpc_repo.NewGRPCImageRepository,
		msg_broker_repo.NewRabbitMQMsgBrokerRepo,
		msg_repo.NewMongoMessageRepo,
		comm_msg_deli_repo.NewWSCommMsgDeliveryRepo,
	)
	return comm_deli_rabbitmq.CommissionMessageBroker{}
}
