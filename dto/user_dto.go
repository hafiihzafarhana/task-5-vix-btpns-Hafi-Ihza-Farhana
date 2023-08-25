package dto

import (
	"rakamin/projectfinal/models"
	"time"
)

type NewRegisterRequest struct {
	Username string `json:"user_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (u *NewRegisterRequest) RegisterRequestToEntity() models.UserModel {
	return models.UserModel{
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	}
}

type NewRegisterResponse struct {
	Id        uint      `json:"id"`
	Username  string    `json:"user_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type NewLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (u *NewLoginRequest) LoginRequestToEntity() models.UserModel {
	return models.UserModel{
		Email:    u.Email,
		Password: u.Password,
	}
}

type NewLoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type NewGetMeResponse struct {
	ID       uint                `json:"id"`
	Username string              `json:"user_name"`
	Email    string              `json:"email"`
	Photos   []models.PhotoModel `json:"photos"`
}

type NewUpdateAccountRequest struct {
	Username string `json:"user_name" binding:"required"`
}

func (u *NewUpdateAccountRequest) UpdateAccountRequestToEntity(id uint) models.UserModel {
	return models.UserModel{
		ID:        id,
		Username:  u.Username,
		UpdatedAt: time.Now(),
	}
}

type NewUpdateAccountResponse struct {
	Id        uint      `json:"id"`
	Username  string    `json:"user_name"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NewDeleteAccountResponse struct {
	Message string `json:"message"`
}
