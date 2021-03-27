package http

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"mime/multipart"
	"net/http"
	add_commission "pixstall-commission/app/commission/delivery/model/add-commission"
	create_message "pixstall-commission/app/commission/delivery/model/create-message"
	get_commission "pixstall-commission/app/commission/delivery/model/get-commission"
	get_commissions "pixstall-commission/app/commission/delivery/model/get-commissions"
	get_messages "pixstall-commission/app/commission/delivery/model/get-messages"
	update_commission "pixstall-commission/app/commission/delivery/model/update-commission"
	"pixstall-commission/domain/commission"
	"pixstall-commission/domain/commission/model"
	model3 "pixstall-commission/domain/file/model"
	model2 "pixstall-commission/domain/message/model"
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
		ctx.AbortWithStatusJSON(get_commissions.NewErrorResponse(model.CommissionErrorUnAuth))
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
		ctx.AbortWithStatusJSON(get_commissions.NewErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, get_commissions.NewResponse(tokenUserID, *commissions, filter.Offset, filter.Count))
}

func (c CommissionController) GetCommission(ctx *gin.Context) {
	tokenUserID := ctx.GetString("userId")
	if tokenUserID == "" {
		ctx.AbortWithStatusJSON(get_commissions.NewErrorResponse(model.CommissionErrorUnAuth))
		return
	}
	commId := ctx.Param("id")
	comm, err := c.commUseCase.GetCommission(ctx, commId)
	if err != nil {
		ctx.AbortWithStatusJSON(get_commission.NewErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, get_commission.NewResponse(*comm))
}

func (c CommissionController) GetCommissionDetails(ctx *gin.Context) {
	tokenUserID := ctx.GetString("userId")
	if tokenUserID == "" {
		ctx.AbortWithStatusJSON(get_commissions.NewErrorResponse(model.CommissionErrorUnAuth))
		return
	}
	commId := ctx.Param("id")
	fmt.Print(commId)

	// TODO
}

func (c CommissionController) GetMessages(ctx *gin.Context) {
	commID := ctx.Param("id")
	tokenUserID := ctx.GetString("userId")
	if tokenUserID == "" {
		ctx.AbortWithStatusJSON(get_commissions.NewErrorResponse(model.CommissionErrorUnAuth))
		return
	}
	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}
	count, err := strconv.Atoi(ctx.Query("count"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}

	msgs, err := c.commUseCase.GetMessages(ctx, tokenUserID, commID, offset, count)
	if err != nil {
		ctx.AbortWithStatusJSON(get_messages.NewErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, get_messages.NewResponse(commID, msgs))
}

func (c CommissionController) CreateMessage(ctx *gin.Context) {
	commId := ctx.Param("id")
	if commId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}
	tokenUserID := ctx.GetString("userId")
	if tokenUserID == "" {
		ctx.AbortWithStatusJSON(get_commissions.NewErrorResponse(model.CommissionErrorUnAuth))
		return
	}
	msgCreator := model2.MessageCreator{CommissionID: commId, Form: tokenUserID}
	text := ctx.PostForm("text")
	msgCreator.Text = text

	imageFiles, err := getMultipartFormImages(ctx, "image")
	if err == nil {
		imgFiles := *imageFiles
		msgCreator.Image = &imgFiles[0]
	}
	if msgCreator.Text == "" && msgCreator.Image == nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}
	err = c.commUseCase.HandleInboundCommissionMessage(ctx, msgCreator)
	if err != nil {
		ctx.AbortWithStatusJSON(create_message.NewErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, create_message.NewResponse())
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

	refImages, err := getMultipartFormImages(ctx, "refImages")
	if err == nil {
		creator.RefImages = *refImages
	}

	comm, err := c.commUseCase.AddCommission(ctx, creator)
	if err != nil {
		ctx.JSON(add_commission.NewErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, add_commission.NewResponse(comm.ID))
}

func (c CommissionController) UpdateCommission(ctx *gin.Context) {
	commID := ctx.Param("id")
	if commID == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}
	tokenUserID := ctx.GetString("userId")
	if tokenUserID == "" {
		ctx.AbortWithStatusJSON(add_commission.NewErrorResponse(model.CommissionErrorUnAuth))
		return
	}
	updater := model.CommissionUpdater{
		ID: commID,
	}

	decision, exist := ctx.GetPostForm("decision")
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
	if rating, exist := ctx.GetPostForm("rating"); exist {
		r, err := strconv.Atoi(rating)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, nil)
			return
		}
		updater.Rating = &r
	}
	if comment, exist := ctx.GetPostForm("comment"); exist {
		updater.Comment = &comment
	}
	displayImages, err := getMultipartFormImages(ctx, "displayImage")
	if err == nil {
		images := *displayImages
		updater.DisplayImageFile = &images[0]
	}
	proofCopyImages, err := getMultipartFormImages(ctx, "proofCopyImage")
	if err == nil {
		images := *proofCopyImages
		updater.ProofCopyImage = &images[0]
	}
	completionFiles, err := getMultipartFormFiles(ctx, "completionFile")
	if err == nil {
		files := *completionFiles
		updater.CompletionFile = &files[0]
	}

	err = c.commUseCase.UpdateCommissionByUser(ctx, tokenUserID, updater, model.CommissionDecision(decision))
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

func getSize(ctx *gin.Context) (*model3.Size, error) {
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
	return &model3.Size{
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

func getMultipartFormImages(ctx *gin.Context, key string) (*[]model3.ImageFile, error) {
	form, err := ctx.MultipartForm()
	if err != nil {
		return nil, err
	}
	fileHeaders := form.File[key]
	imageFiles := make([]model3.ImageFile, 0)
	for _, header := range fileHeaders {
		f, err := header.Open()
		if err != nil {
			continue
		}
		contentType, err := getFileContentType(f)
		if err != nil {
			_ = f.Close()
			continue
		}
		_, err = f.Seek(0, 0)
		if err != nil {
			_ = f.Close()
			continue
		}
		img, _, err := image.Decode(f)
		if err != nil {
			_ = f.Close()
			continue
		}
		_, err = f.Seek(0, 0)
		if err != nil {
			_ = f.Close()
			continue
		}
		imgF := model3.ImageFile{
			File: model3.File{
				File:        f,
				Name:        header.Filename,
				ContentType: contentType,
				Volume:      header.Size,
			},
			Size: model3.Size{
				Width:  float64(img.Bounds().Dx()),
				Height: float64(img.Bounds().Dy()),
				Unit:   "px",
			},
		}
		imageFiles = append(imageFiles, imgF)
		_ = f.Close()
	}
	if len(imageFiles) <= 0 {
		return nil, errors.New("")
	}
	return &imageFiles, nil
}

func getMultipartFormFiles(ctx *gin.Context, key string) (*[]model3.File, error) {
	form, err := ctx.MultipartForm()
	if err != nil {
		return nil, err
	}
	fileHeaders := form.File[key]
	files := make([]model3.File, 0)
	for _, header := range fileHeaders {
		f, err := header.Open()
		if err != nil {
			continue
		}
		contentType, err := getFileContentType(f)
		if err != nil {
			_ = f.Close()
			continue
		}
		_, err = f.Seek(0, 0)
		if err != nil {
			continue
		}
		dFile := model3.File{
			File:        f,
			Name:        header.Filename,
			ContentType: contentType,
			Volume:      header.Size,
		}
		files = append(files, dFile)
		_ = f.Close()
	}
	if len(files) <= 0 {
		return nil, errors.New("")
	}
	return &files, nil
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
		if t, err := time.Parse(layout, createTimeFrom); err == nil {
			filter.CreateTimeFrom = &t
		}
	}
	if createTimeTo, exist := ctx.GetQuery("createTimeTo"); exist {
		if t, err := time.Parse(layout, createTimeTo); err == nil {
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

func getFileContentType(out multipart.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
