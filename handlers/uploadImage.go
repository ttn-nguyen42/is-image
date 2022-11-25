package handlers

import (
	"is-image/core/models/responses"
	"is-image/core/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
 * /api/v1/upload_image
 */
type UploadImageHandlers struct {
	service services.IUploadImageService
}

func NewUploadImageHandlers(serv services.IUploadImageService) *UploadImageHandlers {
	return &UploadImageHandlers{
		service: serv,
	}
}

/* GET /api/v1/upload_image */
func (p *UploadImageHandlers) Post(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "File not found",
		})
		return
	}
	imageData, err := p.service.IsImageValid(file, header)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "Invalid",
		})
		return
	}
	res, err := p.service.GetImageResult(imageData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Message: "Internal server error",
		})
		return
	}
	if res.Imagetype == "" {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "Not image",
		})
		return
	}
	c.JSON(http.StatusOK, &responses.PostUploadImageResponse{
		HashID: res.Id,
		Type:   res.Imagetype,
		Result: res.Result,
	})
}
