package middlewares

import (
	"fmt"
	"rakamin/projectfinal/exception"

	"github.com/gin-gonic/gin"
)

func Authorize(requiredRole interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		roleID := ctx.GetInt("role")
		// Check for required roles
		fmt.Printf("%v\n", roleID)
		fmt.Printf("%v\n", requiredRole)
		switch roles := requiredRole.(type) {
		case int:
			if roleID != roles {
				fmt.Printf("requiredRole 1")
				errResponse := exception.Unauthorized("Token Not Valid")
				ctx.AbortWithStatusJSON(errResponse.Status(), errResponse)
				return
			}
		case []int:
			if !uintInSlice(roleID, roles) {
				fmt.Printf("requiredRole 2")
				errResponse := exception.Unauthorized("Token Not Valid")
				ctx.AbortWithStatusJSON(errResponse.Status(), errResponse)
				return
			}
		default:
			fmt.Printf("requiredRole 3")
			errResponse := exception.InternalServerError("Token Not Valid")
			ctx.AbortWithStatusJSON(errResponse.Status(), errResponse)
			return
		}

		ctx.Next()
	}
}

func uintInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
