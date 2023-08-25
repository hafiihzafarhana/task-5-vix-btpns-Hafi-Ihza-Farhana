package photorepository

import (
	"rakamin/projectfinal/exception"
	"rakamin/projectfinal/models"
)

type PhotoRepository interface {
	CreatePhoto(payload models.PhotoModel) (*models.PhotoModel, exception.MessageErr)
	DeletePhotoById(userId uint, photoId uint) exception.MessageErr
	GetPhotoById(photoId int) (*models.PhotoModel, exception.MessageErr)
	UpdatePhotoById(payload models.PhotoModel) (*models.PhotoModel, exception.MessageErr)
}
