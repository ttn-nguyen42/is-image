package services

import (
	"fmt"
	"is-image/core/defines"
	"is-image/core/models/entities"
	"is-image/core/repositories"
	"is-image/core/utils"
	"mime/multipart"
	"strings"
)

/*
 * Interfaces (ports)
 */
type IUploadImageService interface {
	IsImageValid(file multipart.File, header *multipart.FileHeader) (*ImageData, error)
	GetImageResult(image *ImageData) (*entities.ResultCache, error)
}

type UploadImageService struct {
	repo repositories.IResultCacheRepository
}

type ImageData struct {
	ImageType string
	Hash      string
	File      multipart.File
}

func NewUploadImageService(repo repositories.IResultCacheRepository) *UploadImageService {
	return &UploadImageService{
		repo: repo,
	}
}

func (p *UploadImageService) GetImageResult(image *ImageData) (*entities.ResultCache, error) {
	// Use repository to upload image
	return &entities.ResultCache{
		Id:        image.Hash,
		Imagetype: image.ImageType,
		Result:    "",
	}, nil
}

func (p *UploadImageService) IsImageValid(file multipart.File, header *multipart.FileHeader) (*ImageData, error) {
	if header == nil {
		return &ImageData{}, fmt.Errorf("cannot read header")
	}
	if header.Size >= 5*defines.MB {
		return &ImageData{}, fmt.Errorf("file size exceed 5mb")
	}
	buff, err := utils.ToBytes(file)
	if err != nil {
		return &ImageData{}, err
	}
	imageType, err := utils.GetImageType(buff)
	if err != nil {
		return &ImageData{}, err
	}
	var resType, resHash string
	if imageType == utils.JPEG {
		resType = "jpeg/jpg"
		resHash, err = utils.GetImageHashJpeg(file)
	}
	if imageType == utils.PNG {
		resType = "png"
		resHash, err = utils.GetImageHashPng(file)
	}
	if err != nil {
		return &ImageData{}, fmt.Errorf("cannot hash image")
	}
	hashParts := strings.Split(resHash, ":")
	resHash = hashParts[1]
	return &ImageData{
		ImageType: resType,
		Hash:      resHash,
		File:      file,
	}, nil
}
