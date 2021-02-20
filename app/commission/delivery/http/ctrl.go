package http

import (
	"github.com/gin-gonic/gin"
	"pixstall-commission/domain/commission"
)

type CommissionController struct {
	commUseCase commission.UseCase
}

func NewCommissionController(commUseCase commission.UseCase) CommissionController {
	return CommissionController{
		commUseCase: commUseCase,
	}
}

func (c CommissionController) GetCommissions(ctx *gin.Context) {
	tokenUserID := ctx.GetString("userId")
	if tokenUserID == "" {
		return
	}

	commissions, err := c.commUseCase.GetCommissions(ctx, tokenUserID)
	if err != nil {
		return
	}
}


