package dao

import "pixstall-commission/domain/commission/model"

type GetCommissionsResult struct {
	Commissions []Commission `bson:"commissions"`
	Total       int          `bson:"total"`
}

func (g GetCommissionsResult) ToDomainGetCommissionsResult() *model.GetCommissionsResult {
	dComms := make([]model.Commission, 0)
	for _, comm := range g.Commissions {
		dComms = append(dComms, comm.ToDomainCommission())
	}

	return &model.GetCommissionsResult{
		Commissions: dComms,
		Total:       g.Total,
	}
}
