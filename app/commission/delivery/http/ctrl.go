package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	get_commissions "pixstall-commission/app/commission/delivery/model/get-commissions"
	"pixstall-commission/domain/commission"
	"pixstall-commission/domain/commission/model"
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
		ctx.JSON(get_commissions.NewErrorResponse(model.CommissionErrorUnAuth))
		return
	}

	//TODO: trigger by...
	commissions, err := c.commUseCase.GetCommissions(ctx, tokenUserID)
	if err != nil {
		ctx.JSON(get_commissions.NewErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, get_commissions.NewResponse(tokenUserID, *commissions))
}

func (c CommissionController) GetCommissionDetails(ctx *gin.Context) {
	tokenUserID := ctx.GetString("userId")
	if tokenUserID == "" {
		ctx.JSON(get_commissions.NewErrorResponse(model.CommissionErrorUnAuth))
		return
	}

	// TODO
}

func (c CommissionController) GetCommissionMessages(ctx *gin.Context) {
	tokenUserID := ctx.GetString("userId")
	if tokenUserID == "" {
		ctx.JSON(get_commissions.NewErrorResponse(model.CommissionErrorUnAuth))
		return
	}

	// TODO
}


