package http

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	add_commission "pixstall-commission/app/commission/delivery/model/add-commission"
	get_commission "pixstall-commission/app/commission/delivery/model/get-commission"
	get_commissions "pixstall-commission/app/commission/delivery/model/get-commissions"
	update_commission "pixstall-commission/app/commission/delivery/model/update-commission"
	"pixstall-commission/domain/commission"
	"pixstall-commission/domain/commission/model"
	"strconv"
	"time"
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
	filter, err := getFilter(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}

	sorter, err := getSorter(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}
	commissions, err := c.commUseCase.GetCommissions(ctx, *filter, *sorter)
	if err != nil {
		ctx.JSON(get_commissions.NewErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, get_commissions.NewResponse(tokenUserID, *commissions, filter.Offset, filter.Count))
}

func (c CommissionController) GetCommission(ctx *gin.Context) {
	tokenUserID := ctx.GetString("userId")
	if tokenUserID == "" {
		ctx.JSON(get_commissions.NewErrorResponse(model.CommissionErrorUnAuth))
		return
	}
	commId := ctx.Param("id")
	comm, err := c.commUseCase.GetCommission(ctx, commId)
	if err != nil {
		ctx.JSON(get_commission.NewErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, get_commission.NewResponse(*comm))
}

func (c CommissionController) GetCommissionDetails(ctx *gin.Context) {
	tokenUserID := ctx.GetString("userId")
	if tokenUserID == "" {
		ctx.JSON(get_commissions.NewErrorResponse(model.CommissionErrorUnAuth))
		return
	}
	commId := ctx.Param("id")
	fmt.Print(commId)

	// TODO
}

func (c CommissionController) GetMessages(ctx *gin.Context) {
	tokenUserID := ctx.GetString("userId")
	if tokenUserID == "" {
		ctx.JSON(get_commissions.NewErrorResponse(model.CommissionErrorUnAuth))
		return
	}

	// TODO
}

func (c CommissionController) AddCommission(ctx *gin.Context) {
	tokenUserID := ctx.GetString("userId")
	if tokenUserID == "" {
		ctx.AbortWithStatusJSON(add_commission.NewErrorResponse(model.CommissionErrorUnAuth))
		return
	}
	creator := model.CommissionCreator{
		RequesterID: tokenUserID,
	}
	if openCommID, exist := ctx.GetPostForm("openCommissionId"); exist {
		creator.OpenCommissionID = openCommID
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}
	if artistID, exist := ctx.GetPostForm("artistId"); exist {
		creator.ArtistID = artistID
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}
	price, err := getPriceFromPostForm(ctx, "price")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}
	creator.Price = *price

	dayNeed, err := getDayNeed(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}
	creator.DayNeed = *dayNeed

	size, err := getSize(ctx)
	if err == nil {
		creator.Size = size
	}

	resolution, err := getResolution(ctx)
	if err == nil {
		creator.Resolution = resolution
	}

	if exportFormat, exist := ctx.GetPostForm("exportFormat"); exist {
		creator.ExportFormat = &exportFormat
	}

	if desc, exist := ctx.GetPostForm("desc"); exist {
		creator.Desc = desc
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}

	if paymentMethod, exist := ctx.GetPostForm("paymentMethod"); exist {
		creator.PaymentMethod = paymentMethod
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}

	if isR18, exist := ctx.GetPostForm("isR18"); exist {
		creator.IsR18 = isR18 == "true"
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}

	if bePrivate, exist := ctx.GetPostForm("bePrivate"); exist {
		creator.BePrivate = bePrivate == "true"
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}

	if anonymous, exist := ctx.GetPostForm("anonymous"); exist {
		creator.Anonymous = anonymous == "true"
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}

	refImage, err := getRefImages(ctx)
	if err == nil {
		creator.RefImages = *refImage
	}
	
	comm, err := c.commUseCase.AddCommission(ctx, creator)
	if err != nil {
		ctx.JSON(add_commission.NewErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, add_commission.NewResponse(comm.ID))
}

func (c CommissionController) UpdateCommission(ctx *gin.Context) {
	tokenUserID := ctx.GetString("userId")
	if tokenUserID == "" {
		ctx.AbortWithStatusJSON(add_commission.NewErrorResponse(model.CommissionErrorUnAuth))
		return
	}
	updater := model.CommissionUpdater{}
	commID, exist := ctx.GetPostForm("commissionId")
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}

	if price, err := getPriceFromPostForm(ctx, "price"); err == nil {
		updater.Price = price
	}
	if dayNeed, err := getDayNeed(ctx); err == nil {
		updater.DayNeed = dayNeed
	}
	if size, err := getSize(ctx); err == nil {
		updater.Size = size
	}
	if resolution, err := getResolution(ctx); err == nil {
		updater.Resolution = resolution
	}
	if exportFormat, exist := ctx.GetPostForm("exportFormat"); exist {
		updater.ExportFormat = &exportFormat
	}
	if desc, exist := ctx.GetPostForm("desc"); exist {
		updater.Desc = &desc
	}
	if paymentMethod, exist := ctx.GetPostForm("paymentMethod"); exist {
		updater.PaymentMethod = &paymentMethod
	}
	if bePrivate, exist := ctx.GetPostForm("bePrivate"); exist {
		b := bePrivate == "true"
		updater.BePrivate = &b
	}
	if anonymous, exist := ctx.GetPostForm("anonymous"); exist {
		b := anonymous == "true"
		updater.Anonymous = &b
	}

	err := c.commUseCase.UpdateCommissionByUser(ctx, tokenUserID, updater)
	if err != nil {
		ctx.JSON(update_commission.NewErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, update_commission.NewResponse(commID))
}

// Private methods
func getPriceFromPostForm(ctx *gin.Context, priceStr string) (*model.Price, error) {
	priceAmount, exist := ctx.GetPostForm(priceStr + ".amount")
	if !exist {
		return nil, errors.New("")
	}
	amount, err := strconv.ParseFloat(priceAmount, 64)
	if err != nil {
		return nil, err
	}
	priceCurrency, exist := ctx.GetPostForm(priceStr + ".currency")
	if !exist {
		return nil, errors.New("")
	}
	return &model.Price{
		Amount:   amount,
		Currency: model.Currency(priceCurrency),
	}, nil
}

func getPriceFromQuery(ctx *gin.Context, priceStr string) (*model.Price, error) {
	priceAmount, exist := ctx.GetQuery(priceStr + ".amount")
	if !exist {
		return nil, errors.New("")
	}
	amount, err := strconv.ParseFloat(priceAmount, 64)
	if err != nil {
		return nil, err
	}
	priceCurrency, exist := ctx.GetQuery(priceStr + ".currency")
	if !exist {
		return nil, errors.New("")
	}
	return &model.Price{
		Amount:   amount,
		Currency: model.Currency(priceCurrency),
	}, nil
}

func getDayNeed(ctx *gin.Context) (*int, error) {
	dayNeed, exist := ctx.GetPostForm("dayNeed")
	if !exist {
		return nil, errors.New("")
	}
	dn, err := strconv.Atoi(dayNeed)
	if err != nil {
		return nil, err
	}
	return &dn, nil
}

func getSize(ctx *gin.Context) (*model.Size, error) {
	sizeWidth, exist := ctx.GetPostForm("size.width")
	if !exist {
		return nil, errors.New("")
	}
	width, err := strconv.ParseFloat(sizeWidth, 64)
	if err != nil {
		return nil, err
	}
	sizeHeight, exist := ctx.GetPostForm("size.height")
	if !exist {
		return nil, errors.New("")
	}
	height, err := strconv.ParseFloat(sizeHeight, 64)
	if err != nil {
		return nil, err
	}
	sizeUnit, exist := ctx.GetPostForm("size.unit")
	if !exist {
		return nil, errors.New("")
	}
	return &model.Size{
		Width:  width,
		Height: height,
		Unit:   sizeUnit,
	}, nil
}

func getResolution(ctx *gin.Context) (*float64, error) {
	resolution, exist := ctx.GetPostForm("resolution")
	if !exist {
		return nil, errors.New("")
	}
	result, err := strconv.ParseFloat(resolution, 64)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func getRefImages(ctx *gin.Context) (*[]image.Image, error) {
	form, err := ctx.MultipartForm()
	if err != nil {
		return nil, err
	}
	fileHeaders := form.File["refImages"]
	images := make([]image.Image, 0)
	for _, header := range fileHeaders {
		decodedImage := func() image.Image {
			if err != nil {
				return nil
			}
			f, err := header.Open()
			if err != nil {
				return nil
			}
			decodedImage, _, err := image.Decode(f)
			if err != nil {
				return nil
			}
			return decodedImage
		}()
		if decodedImage != nil {
			images = append(images, decodedImage)
		}
	}
	if len(images) <= 0 {
		return nil, errors.New("")
	}
	return &images, nil
}

func getFilter(ctx *gin.Context) (*model.CommissionFilter, error) {

	filter := model.CommissionFilter{
		ArtistID:       nil,
		RequesterID:    nil,
		PriceFrom:      nil,
		PriceTo:        nil,
		CreateTimeFrom: nil,
		CreateTimeTo:   nil,
		State:          nil,
	}
	if artistId, exist := ctx.GetQuery("artistId"); exist {
		filter.ArtistID = &artistId
	}
	if requesterId, exist := ctx.GetQuery("requesterId"); exist {
		filter.RequesterID = &requesterId
	}
	if count, exist := ctx.GetQuery("count"); exist {
		if countInt, err := strconv.Atoi(count); err == nil {
			filter.Count = countInt
		} else {
			return nil, errors.New("bad request")
		}
	} else {
		return nil, errors.New("bad request")
	}
	if offset, exist := ctx.GetQuery("offset"); exist {
		if countInt, err := strconv.Atoi(offset); err == nil {
			filter.Offset = countInt
		} else {
			return nil, errors.New("bad request")
		}
	} else {
		return nil, errors.New("bad request")
	}
	if priceFrom, err := getPriceFromQuery(ctx, "priceFrom"); err == nil {
		filter.PriceFrom = priceFrom
	}
	if priceTo, err := getPriceFromQuery(ctx, "priceTo"); err == nil {
		filter.PriceFrom = priceTo
	}
	layout := "2014-09-12T11:45:26.371Z"
	if createTimeFrom, exist := ctx.GetQuery("createTimeFrom"); exist {
		if t, err := time.Parse(layout , createTimeFrom); err == nil {
			filter.CreateTimeFrom = &t
		}
	}
	if createTimeTo, exist := ctx.GetQuery("createTimeTo"); exist {
		if t, err := time.Parse(layout , createTimeTo); err == nil {
			filter.CreateTimeTo = &t
		}
	}
	if state, exist := ctx.GetQuery("state"); exist {
		commState := model.CommissionState(state)
		filter.State = &commState
	}

	return &filter, nil
}

func getSorter(ctx *gin.Context) (*model.CommissionSorter, error) {
	sorter := model.CommissionSorter{}
	if sortBy, exist := ctx.GetQueryArray("sortBy"); exist {
		if len(sortBy) < 2 {
			return nil, errors.New("bad request")
		}
		var asc bool
		switch sortBy[1] {
		case "asc":
			asc = true
		case "dsc":
			asc = false
		default:
			return nil, errors.New("bad request")
		}
		switch sortBy[0] {
		case "price":
			sorter.Price = &asc
		case "state":
			sorter.State = &asc
		case "createTime":
			sorter.CreateTime = &asc
		case "lastUpdateTime":
			sorter.LastUpdateTime = &asc
		default:
			return nil, errors.New("bad request")
		}
		return &sorter, nil
	} else {
		return &sorter, nil
	}
}