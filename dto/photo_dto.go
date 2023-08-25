package dto

import (
	"rakamin/projectfinal/models"
	"time"
)

type NewPostPhotoRequest struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption" binding:"required"`
	PhotoUrl string `json:"photo_url" binding:"required"`
}

func (u *NewPostPhotoRequest) PostPhotoRequestToEntity(userId uint) models.PhotoModel {
	return models.PhotoModel{
		Title:    u.Title,
		Caption:  u.Caption,
		PhotoUrl: u.PhotoUrl,
		UserID:   userId,
	}
}

type NewPostPhotoResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type NewUpdatePhotoRequest struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption" binding:"required"`
	PhotoUrl string `json:"photo_url" binding:"required"`
}

func (u *NewUpdatePhotoRequest) UpdatePhotoRequestToEntity(userId uint, photoId uint) models.PhotoModel {
	return models.PhotoModel{
		ID:       photoId,
		Title:    u.Title,
		Caption:  u.Caption,
		PhotoUrl: u.PhotoUrl,
		UserID:   userId,
	}
}

type NewUpdatePhotoResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NewDeletePhotoResponse struct {
	Message string `json:"message"`
}
