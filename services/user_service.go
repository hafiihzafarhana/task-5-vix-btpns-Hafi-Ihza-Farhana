package services

import (
	"fmt"
	"rakamin/projectfinal/dto"
	"rakamin/projectfinal/exception"
	"rakamin/projectfinal/models"
	userrepository "rakamin/projectfinal/repositories/user_repository"
	"rakamin/projectfinal/util"
)

type UserService interface {
	Register(payload dto.NewRegisterRequest) (*dto.NewRegisterResponse, exception.MessageErr)
	Login(payload dto.NewLoginRequest) (*dto.NewLoginResponse, exception.MessageErr)
	GetMe(id int) (*dto.NewGetMeResponse, exception.MessageErr)
	UpdateUserById(payload dto.NewUpdateAccountRequest, userId int) (*dto.NewUpdateAccountResponse, exception.MessageErr)
	DeleteUserById(userId int) (*dto.NewDeleteAccountResponse, exception.MessageErr)
}

type userService struct {
	userRepo userrepository.UserRepository
}

func NewUserService(userRepo userrepository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) Register(payload dto.NewRegisterRequest) (*dto.NewRegisterResponse, exception.MessageErr) {
	user := payload.RegisterRequestToEntity()

	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashedPassword

	createdUser, err := u.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	response := &dto.NewRegisterResponse{
		Id:        createdUser.ID,
		Username:  createdUser.Username,
		Email:     createdUser.Email,
		CreatedAt: createdUser.CreatedAt,
	}

	return response, nil
}

func (u *userService) Login(payload dto.NewLoginRequest) (*dto.NewLoginResponse, exception.MessageErr) {
	user := payload.LoginRequestToEntity()

	getUser, err := u.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	if err := util.ComparePassword(user.Password, getUser.Password); err != nil {
		return nil, err
	}

	access_token, err := util.GenerateAccessToken(getUser.ID, getUser.Email, getUser.RoleID)

	if err != nil {
		return nil, err
	}

	refresh_token, err := util.GenerateRefreshToken(getUser.ID, getUser.Email, getUser.RoleID)

	if err != nil {
		return nil, err
	}

	response := &dto.NewLoginResponse{
		AccessToken:  access_token,
		RefreshToken: refresh_token,
	}

	return response, nil
}

func (u *userService) GetMe(id int) (*dto.NewGetMeResponse, exception.MessageErr) {
	var photos []models.PhotoModel

	getData, err := u.userRepo.GetUserById(id)

	if err != nil {
		return nil, err
	}

	photos = getData.Photos

	response := &dto.NewGetMeResponse{
		ID:       getData.ID,
		Username: getData.Username,
		Email:    getData.Email,
		Photos:   photos,
	}

	return response, nil
}

func (u *userService) UpdateUserById(payload dto.NewUpdateAccountRequest, userId int) (*dto.NewUpdateAccountResponse, exception.MessageErr) {
	user := payload.UpdateAccountRequestToEntity(uint(userId))
	fmt.Println(user)
	updateUser, err := u.userRepo.UpdateUserById(user)
	if err != nil {
		return nil, err
	}

	response := &dto.NewUpdateAccountResponse{
		Id:        user.ID,
		Username:  updateUser.Username,
		UpdatedAt: updateUser.UpdatedAt,
	}

	return response, nil
}

func (u *userService) DeleteUserById(userId int) (*dto.NewDeleteAccountResponse, exception.MessageErr) {
	if err := u.userRepo.DeleteUserById(uint(userId)); err != nil {
		return nil, err
	}

	response := &dto.NewDeleteAccountResponse{
		Message: "User has been deleted",
	}

	return response, nil
}
