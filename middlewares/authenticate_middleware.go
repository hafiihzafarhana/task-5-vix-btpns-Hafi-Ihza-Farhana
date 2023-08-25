package middlewares

import (
	"rakamin/projectfinal/exception"
	"rakamin/projectfinal/util"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	// fmt.Println(tokenString, "Los")
	if tokenString == "" {
		errResponse := exception.Unauthorized("Token Not Valid")
		c.AbortWithStatusJSON(errResponse.Status(), errResponse)
		return
	}

	token, err := util.VerifyAccessToken(tokenString)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}
	id := int(token["id"].(float64))
	email := token["email"].(string)
	role := int(token["role"].(float64))

	c.Set("id", id)
	c.Set("email", email)
	c.Set("role", role)

	c.Next()
}
