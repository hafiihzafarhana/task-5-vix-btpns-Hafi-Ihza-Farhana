package util

import (
	"fmt"
	"os"
	"rakamin/projectfinal/exception"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func tokenClaims(ID uint, email string, expiration time.Time, role uint) jwt.MapClaims {
	return jwt.MapClaims{
		"id":    ID,
		"email": email,
		"exp":   expiration.Unix(),
		"role":  role,
	}
}

func signToken(claims jwt.MapClaims, secretKey []byte) (string, exception.MessageErr) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", exception.InternalServerError(err.Error())
	}
	return signedToken, nil
}

func GenerateAccessToken(ID uint, email string, role uint) (string, exception.MessageErr) {
	accessTokenExpiration := time.Now().Add(time.Hour * 2)
	claims := tokenClaims(ID, email, accessTokenExpiration, role)
	getEnv := os.Getenv("ACCESS_SECRET_KEY")
	secretKey := []byte(getEnv)
	return signToken(claims, secretKey)
}

func GenerateRefreshToken(ID uint, email string, role uint) (string, exception.MessageErr) {
	refreshTokenExpiration := time.Now().Add(time.Hour * 24 * 30)
	claims := tokenClaims(ID, email, refreshTokenExpiration, role)
	secretKey := []byte(os.Getenv("REFRESH_SECRET_KEY"))
	return signToken(claims, secretKey)
}

func parseToken(bearerToken string, secretKey []byte) (*jwt.Token, exception.MessageErr) {
	fmt.Println(secretKey)
	token, err := jwt.Parse(bearerToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, exception.Unauthorized("invalid token error")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, exception.Unauthorized("invalid token error")
	}

	return token, nil
}

func VerifyAccessToken(bearerToken string) (jwt.MapClaims, exception.MessageErr) {
	secretKey := []byte(os.Getenv("ACCESS_SECRET_KEY"))
	tokenString := strings.Split(bearerToken, " ")[1]
	token, err := parseToken(tokenString, secretKey)

	if err != nil {
		return nil, err
	}

	var mapClaims jwt.MapClaims

	if claims, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return nil, exception.Unauthorized("invalid token error")
	} else {
		mapClaims = claims
	}

	return mapClaims, nil
}
