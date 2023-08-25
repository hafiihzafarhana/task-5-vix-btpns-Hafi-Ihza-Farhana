package userrepository

import (
	"rakamin/projectfinal/exception"
	"rakamin/projectfinal/models"
)

type UserRepository interface {
	CreateUser(payload models.UserModel) (*models.UserModel, exception.MessageErr)
	GetUserByEmail(email string) (*models.UserModel, exception.MessageErr)
	GetUserById(user int) (*models.UserModel, exception.MessageErr)
	UpdateUserById(payload models.UserModel) (*models.UserModel, exception.MessageErr)
	DeleteUserById(userId uint) exception.MessageErr
}
