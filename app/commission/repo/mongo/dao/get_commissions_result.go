package dao

import "pixstall-commission/domain/commission/model"

type GetCommissionsResult struct {
	Commissions []Commission `bson:"commissions"`
	Total       int          `bson:"total"`
}

func (g GetCommissionsResult) ToDomainGetCommissionsResult(offset int) *model.GetCommissionsResult {
	var dComms []model.Commission
	for _, comm := range g.Commissions {
		dComms = append(dComms, comm.ToDomainCommission())
	}

	return &model.GetCommissionsResult{
		Commissions: dComms,
		Total:       g.Total,
	}
}
