package photopg

import (
	"errors"
	"rakamin/projectfinal/exception"
	"rakamin/projectfinal/models"
	photorepository "rakamin/projectfinal/repositories/photo_repository"
	userrepository "rakamin/projectfinal/repositories/user_repository"
	userpg "rakamin/projectfinal/repositories/user_repository/user_pg"

	"gorm.io/gorm"
)

type photoPG struct {
	db *gorm.DB
}

func NewPhotoPG(db *gorm.DB) photorepository.PhotoRepository {
	return &photoPG{db: db}
}

func userRepo(db *gorm.DB) userrepository.UserRepository {
	userRepo := userpg.NewUserPG(db)

	return userRepo
}

func (p *photoPG) CreatePhoto(payload models.PhotoModel) (*models.PhotoModel, exception.MessageErr) {

	getUser, userErr := userRepo(p.db).GetUserById(int(payload.UserID))

	if userErr != nil {
		return nil, userErr
	}

	payload.UserID = getUser.ID

	if err := p.db.Create(&payload).Error; err != nil {
		return nil, exception.InternalServerError("something went wrong")
	}

	return &payload, nil
}

func (p *photoPG) GetPhotoById(photoId int) (*models.PhotoModel, exception.MessageErr) {
	var photo models.PhotoModel

	if err := p.db.Where("id = ?", photoId).First(&photo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.NotFound("Photo not found")
		}
		return nil, exception.InternalServerError("something went wrong")
	}
	return &photo, nil
}

func (p *photoPG) DeletePhotoById(userId uint, photoId uint) exception.MessageErr {

	_, getUserErr := userRepo(p.db).GetUserById(int(userId))
	if getUserErr != nil {
		return getUserErr
	}

	photo, getPhotoErr := p.GetPhotoById(int(photoId))
	if getPhotoErr != nil {
		return getPhotoErr
	}

	if err := p.db.Where("id = ?", photoId).Delete(photo).Error; err != nil {
		return exception.InternalServerError("something went wrong")
	}

	return nil
}

func (p *photoPG) UpdatePhotoById(payload models.PhotoModel) (*models.PhotoModel, exception.MessageErr) {
	_, getUserErr := userRepo(p.db).GetUserById(int(payload.UserID))
	if getUserErr != nil {
		return nil, getUserErr
	}

	var photo models.PhotoModel

	err := p.db.Model(&photo).Where("id = ?", payload.ID).Updates(models.PhotoModel{ID: payload.ID, Title: payload.Title, Caption: payload.Caption, PhotoUrl: payload.PhotoUrl, UpdatedAt: payload.UpdatedAt, UserID: payload.UserID}).Error
	if err != nil {
		return nil, exception.InternalServerError("something went wrong")
	}
	return &photo, nil
}
