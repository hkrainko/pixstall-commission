//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/aws/aws-sdk-go/service/s3"
	"go.mongodb.org/mongo-driver/mongo"
	"pixstall-commission/app/commission/delivery/http"
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
