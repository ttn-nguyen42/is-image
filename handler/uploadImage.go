package handler

import (
	"log"
	"net/http"
	"strings"

	hp "is-image/helper"
	res "is-image/model/response"
	repos "is-image/repository"

	"github.com/gin-gonic/gin"
)

const (
	KB = 1024
	MB = 1024 * KB
)

type PostUploadImageDependencies struct {
	Repo *repos.ResultCacheRepository
}

/*
 * POST - /api/v1/upload_image
 */
func PostUploadImage(c *gin.Context, deps *PostUploadImageDependencies) {
	formFile, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			res.ErrorResponse{
				Message: "File missing",
			})
		return
	}
	log.Printf("Received file size: %d bytes", fileHeader.Size)
	if fileHeader.Size >= 5*MB {
		c.JSON(
			http.StatusBadRequest,
			res.ErrorResponse{
				Message: "File size must be <5MB",
			})
		return
	}
	buf, err := hp.ToBytes(formFile)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			res.ErrorResponse{
				Message: "File cannot be read",
			})
		return
	}
	imageType, err := hp.GetImageType(buf)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			res.ErrorResponse{
				Message: "File uploaded is not an image",
			})
		return
	}
	var resType, resHash string
	if imageType == hp.JPEG {
		resType = "jpeg/jpg"
		resHash, err = hp.GetImageHashJpeg(formFile)
	}
	if imageType == hp.PNG {
		resType = "png"
		resHash, err = hp.GetImageHashPng(formFile)
	}
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			res.ErrorResponse{
				Message: "Internal server error",
			})
		return
	}
	// Only uses the main part of the hash
	hashParts := strings.Split(resHash, ":")
	resHash = hashParts[1]
	// Save result to database
	// Check if dependency is nil or not
	c.JSON(
		http.StatusOK,
		res.PostUploadImageResponse{
			Result: resType,
			HashID: resHash,
		})
}
