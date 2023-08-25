package util

import (
	"rakamin/projectfinal/exception"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(Password string) (string, exception.MessageErr) {
	const cost = 8

	bs, err := bcrypt.GenerateFromPassword([]byte(Password), cost)
	if err != nil {
		return "", exception.InternalServerError("something went wrong")
	}

	return string(bs), nil
}

func ComparePassword(Password string, cPassword string) exception.MessageErr {
	if err := bcrypt.CompareHashAndPassword([]byte(cPassword), []byte(Password)); err != nil {
		return exception.BadRequest("wrong password")
	}

	return nil
}
