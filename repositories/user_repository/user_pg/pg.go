package userpg

import (
	"errors"
	"fmt"
	"os"
	"rakamin/projectfinal/exception"
	"rakamin/projectfinal/models"
	userrepository "rakamin/projectfinal/repositories/user_repository"
	"strconv"

	"gorm.io/gorm"
)

type userPG struct {
	db *gorm.DB
}

func NewUserPG(db *gorm.DB) userrepository.UserRepository {
	return &userPG{db: db}
}

// Register implements userrepository.UserRepository.
func (u *userPG) CreateUser(payload models.UserModel) (*models.UserModel, exception.MessageErr) {
	USER_ID := os.Getenv("USER_ID")
	userIDInt, _ := strconv.Atoi(USER_ID)
	payload.RoleID = uint(userIDInt)
	if err := u.db.Create(&payload).Error; err != nil {
		return nil, exception.InternalServerError("something went wrong")
	}

	return &payload, nil
}

func (u *userPG) GetUserByEmail(email string) (*models.UserModel, exception.MessageErr) {
	var user models.UserModel
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, exception.NotFound(fmt.Sprintf("user with email %s is not found", email))
	}

	return &user, nil
}

func (u *userPG) GetUserById(userId int) (*models.UserModel, exception.MessageErr) {
	var user models.UserModel

	if err := u.db.Where("id = ?", userId).Preload("Photos").First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.NotFound("User not found")
		}
		return nil, exception.InternalServerError("something went wrong")
	}
	fmt.Println(&user)
	return &user, nil
}

func (u *userPG) UpdateUserById(payload models.UserModel) (*models.UserModel, exception.MessageErr) {
	var user models.UserModel

	err := u.db.Model(&user).Where("id = ?", payload.ID).Updates(models.UserModel{Username: payload.Username, UpdatedAt: payload.UpdatedAt}).Error
	if err != nil {
		return nil, exception.InternalServerError("something went wrong")
	}
	return &user, nil
}

func (u *userPG) DeleteUserById(userId uint) exception.MessageErr {
	user, getUserErr := u.GetUserById(int(userId))
	if getUserErr != nil {
		return getUserErr
	}

	if err := u.db.Where("id = ?", userId).Delete(user).Error; err != nil {
		return exception.InternalServerError("something went wrong")
	}

	return nil
}
