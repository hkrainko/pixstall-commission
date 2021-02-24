package dao

import (
	"go.mongodb.org/mongo-driver/bson"
	"pixstall-commission/domain/commission/model"
)

func NewFilterFromDomainCommissionFilter(d model.CommissionFilter) bson.D {
	filter := bson.D{}

	if d.ArtistID != nil {
		filter = append(filter, bson.E{Key: "artistId", Value: d.ArtistID})
	}
	if d.RequesterID != nil {
		filter = append(filter, bson.E{Key: "requesterId", Value: d.RequesterID})
	}
	if d.PriceFrom != nil {
		filter = append(filter, bson.E{Key: "price", Value: bson.M{"$gte": d.PriceFrom}})
	}
	if d.PriceTo != nil {
		filter = append(filter, bson.E{Key: "price", Value: bson.M{"lte": d.PriceTo}})
	}
	if d.DayNeedFrom != nil {
		filter = append(filter, bson.E{Key: "dayNeed", Value: bson.M{"$gte": d.DayNeedFrom}})
	}
	if d.DayNeedTo != nil {
		filter = append(filter, bson.E{Key: "dayNeed", Value: bson.M{"$lte": d.DayNeedTo}})
	}
	if d.CreateTimeFrom != nil {
		filter = append(filter, bson.E{Key: "createTime", Value: bson.M{"gte": d.CreateTimeFrom}})
	}
	if d.CreateTimeTo != nil {
		filter = append(filter, bson.E{Key: "createTime", Value: bson.M{"$lte": d.CreateTimeTo}})
	}
	if d.State != nil {
		filter = append(filter, bson.E{Key: "state", Value: d.State})
	}

	return filter
}
