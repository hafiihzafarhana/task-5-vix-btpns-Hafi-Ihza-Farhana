package util

import (
	"fmt"
	"rakamin/projectfinal/exception"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleBindingErrors(ctx *gin.Context, err error) {
	errBinds := []string{}
	errCasting, ok := err.(validator.ValidationErrors)
	if !ok {
		newErrBind := exception.BadRequest("invalid body request")
		ctx.AbortWithStatusJSON(newErrBind.Status(), newErrBind)
		return
	}

	for _, e := range errCasting {
		errBind := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
		errBinds = append(errBinds, errBind)
	}

	newErrBind := exception.UnprocessableEntity(errBinds)
	ctx.AbortWithStatusJSON(newErrBind.Status(), newErrBind)
}
