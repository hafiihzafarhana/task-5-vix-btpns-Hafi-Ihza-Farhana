package controllers

import (
	"rakamin/projectfinal/dto"
	"rakamin/projectfinal/exception"
	"rakamin/projectfinal/services"
	"rakamin/projectfinal/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) userHandler {
	return userHandler{userService: userService}
}

func (u *userHandler) Register(ctx *gin.Context) {
	var requestBody dto.NewRegisterRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		util.HandleBindingErrors(ctx, err)
		return
	}

	createdUser, errResponse := u.userService.Register(requestBody)
	if errResponse != nil {
		ctx.AbortWithStatusJSON(errResponse.Status(), errResponse)
		return
	}

	util.SuccessCreated(ctx, "You have finished to create account", createdUser)
}

func (u *userHandler) Login(ctx *gin.Context) {
	var requestBody dto.NewLoginRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		util.HandleBindingErrors(ctx, err)
		return
	}

	token, errResponse := u.userService.Login(requestBody)
	if errResponse != nil {
		ctx.AbortWithStatusJSON(errResponse.Status(), errResponse)
		return
	}

	util.SuccessCreated(ctx, "You Logged", token)
}

func (u *userHandler) GetMe(ctx *gin.Context) {
	id := ctx.GetInt("id")

	response, errGet := u.userService.GetMe(id)
	if errGet != nil {
		ctx.AbortWithStatusJSON(errGet.Status(), errGet)
		return
	}

	util.SuccessOK(ctx, "Get Me", response)
}

func (u *userHandler) UpdateUserById(ctx *gin.Context) {
	id := ctx.Param("id")

	if util.IsEmpty(id) {
		errResponse := exception.BadRequest("Request Invalid")
		ctx.AbortWithStatusJSON(errResponse.Status(), errResponse)
		return
	}

	userId, _ := strconv.Atoi(id)

	var requestBody dto.NewUpdateAccountRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		util.HandleBindingErrors(ctx, err)
		return
	}

	response, errResponse := u.userService.UpdateUserById(requestBody, int(userId))
	if errResponse != nil {
		ctx.AbortWithStatusJSON(errResponse.Status(), errResponse)
		return
	}

	util.SuccessOK(ctx, "Success Update User", response)
}

func (u *userHandler) DeleteUserById(ctx *gin.Context) {
	id := ctx.Param("id")

	if util.IsEmpty(id) {
		errResponse := exception.BadRequest("Request Invalid")
		ctx.AbortWithStatusJSON(errResponse.Status(), errResponse)
		return
	}

	userId, _ := strconv.Atoi(id)

	response, errResponse := u.userService.DeleteUserById(int(userId))

	if errResponse != nil {
		ctx.AbortWithStatusJSON(errResponse.Status(), errResponse)
		return
	}

	util.SuccessOK(ctx, "Success Delete User", response)
}
