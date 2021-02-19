package http

import "pixstall-commission/domain/commission"

type CommissionController struct {
	commUseCase commission.UseCase
}

func NewCommissionController(commUseCase commission.UseCase) CommissionController {
	return CommissionController{
		commUseCase: commUseCase,
	}
}

