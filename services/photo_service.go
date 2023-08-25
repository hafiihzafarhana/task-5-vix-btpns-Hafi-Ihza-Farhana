package services

import (
	"rakamin/projectfinal/dto"
	"rakamin/projectfinal/exception"
	photorepository "rakamin/projectfinal/repositories/photo_repository"
)

type PhotoService interface {
	CreatePhoto(payload dto.NewPostPhotoRequest, id uint) (*dto.NewPostPhotoResponse, exception.MessageErr)
	DeletePhotoById(userId uint, photoId uint) (*dto.NewDeletePhotoResponse, exception.MessageErr)
	UpdatePhotoById(payload dto.NewUpdatePhotoRequest, userId uint, photoId uint) (*dto.NewUpdatePhotoResponse, exception.MessageErr)
}

type photoService struct {
	photoRepo photorepository.PhotoRepository
}

func NewPhotoService(photoRepo photorepository.PhotoRepository) PhotoService {
	return &photoService{photoRepo: photoRepo}
}

func (p *photoService) CreatePhoto(payload dto.NewPostPhotoRequest, id uint) (*dto.NewPostPhotoResponse, exception.MessageErr) {
	photo := payload.PostPhotoRequestToEntity(id)

	createdPhoto, err := p.photoRepo.CreatePhoto(photo)
	if err != nil {
		return nil, err
	}

	response := &dto.NewPostPhotoResponse{
		ID:        createdPhoto.ID,
		Title:     createdPhoto.Title,
		Caption:   createdPhoto.Caption,
		PhotoUrl:  createdPhoto.PhotoUrl,
		UserID:    createdPhoto.UserID,
		CreatedAt: createdPhoto.CreatedAt,
	}

	return response, nil
}

func (p *photoService) DeletePhotoById(userId uint, photoId uint) (*dto.NewDeletePhotoResponse, exception.MessageErr) {
	if err := p.photoRepo.DeletePhotoById(uint(userId), uint(photoId)); err != nil {
		return nil, err
	}

	response := &dto.NewDeletePhotoResponse{
		Message: "User has been deleted",
	}

	return response, nil
}

func (p *photoService) UpdatePhotoById(payload dto.NewUpdatePhotoRequest, userId uint, photoId uint) (*dto.NewUpdatePhotoResponse, exception.MessageErr) {
	photo := payload.UpdatePhotoRequestToEntity(userId, photoId)
	updatePhoto, err := p.photoRepo.UpdatePhotoById(photo)
	if err != nil {
		return nil, err
	}

	response := &dto.NewUpdatePhotoResponse{
		ID:        updatePhoto.ID,
		UserID:    updatePhoto.UserID,
		Title:     updatePhoto.Title,
		Caption:   updatePhoto.Caption,
		PhotoUrl:  updatePhoto.PhotoUrl,
		UpdatedAt: updatePhoto.UpdatedAt,
	}

	return response, nil
}
