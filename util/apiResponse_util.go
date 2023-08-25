package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessOK(ctx *gin.Context, message string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": message,
		"data":    data,
	})
}

func SuccessCreated(ctx *gin.Context, message string, data interface{}) {
	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "CREATED",
		"message": message,
		"data":    data,
	})
}
