package controllers

import (
	"rakamin/projectfinal/dto"
	"rakamin/projectfinal/exception"
	"rakamin/projectfinal/services"
	"rakamin/projectfinal/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type photoHandler struct {
	photoService services.PhotoService
}

func NewPhotoHandler(photoService services.PhotoService) photoHandler {
	return photoHandler{photoService: photoService}
}

func (p *photoHandler) UserCreatePhoto(ctx *gin.Context) {
	var requestBody dto.NewPostPhotoRequest
	id := ctx.GetInt("id")

	if util.IsEmpty(id) {
		errResponse := exception.BadRequest("Request Invalid")
		ctx.AbortWithStatusJSON(errResponse.Status(), errResponse)
		return
	}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		util.HandleBindingErrors(ctx, err)
		return
	}

	response, errResponse := p.photoService.CreatePhoto(requestBody, uint(id))
	if errResponse != nil {
		ctx.AbortWithStatusJSON(errResponse.Status(), errResponse)
		return
	}

	util.SuccessCreated(ctx, "Success Create Photo", response)
}

func (p *photoHandler) UserUpdatePhoto(ctx *gin.Context) {
	userId := ctx.GetInt("id")
	photoId := ctx.Param("id")

	if util.IsEmpty(userId) || util.IsEmpty(photoId) {
		errResponse := exception.BadRequest("Request Invalid")
		ctx.AbortWithStatusJSON(errResponse.Status(), errResponse)
		return
	}
	var requestBody dto.NewUpdatePhotoRequest
	photoIdConvert, _ := strconv.Atoi(photoId)

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		util.HandleBindingErrors(ctx, err)
		return
	}

	response, errResponse := p.photoService.UpdatePhotoById(requestBody, uint(userId), uint(photoIdConvert))
	if errResponse != nil {
		ctx.AbortWithStatusJSON(errResponse.Status(), errResponse)
		return
	}

	util.SuccessOK(ctx, "Success Update Photo", response)
}

func (p *photoHandler) UserDeletePhoto(ctx *gin.Context) {
	userId := ctx.GetInt("id")
	photoId := ctx.Param("id")

	if util.IsEmpty(userId) || util.IsEmpty(photoId) {
		errResponse := exception.BadRequest("Request Invalid")
		ctx.AbortWithStatusJSON(errResponse.Status(), errResponse)
		return
	}

	photoIdConvert, _ := strconv.Atoi(photoId)

	response, errResponse := p.photoService.DeletePhotoById(uint(userId), uint(photoIdConvert))
	if errResponse != nil {
		ctx.AbortWithStatusJSON(errResponse.Status(), errResponse)
		return
	}

	util.SuccessOK(ctx, "Success Delete Photo", response)
}
