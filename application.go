package main

import (
	"fmt"
	"is-image/core/defines"
	"is-image/core/repositories"
	"is-image/core/services"
	"is-image/db"
	"is-image/handlers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

/*
 * API versions
 */
const (
	V1 = "/api/v1"
)

/*
 * Paths
 */
const (
	UploadImage = "/upload_image"
)

/*
 * Application entry point
 */
type Application struct {
	g *gin.Engine

	// Handlers
	uploadImage *handlers.UploadImageHandlers
}

func NewApplication() *Application {
	return &Application{
		g: gin.Default(),
	}
}

func (p *Application) Start() {
	p.RegisterService()
	p.RegisterHandler()
	port := os.Getenv(defines.PORT)
	if port == "" {
		port = "8080"
	}
	log.Printf("Application starts in port %s", port)
	p.g.Run(fmt.Sprintf(":%s", port))
}

/*
 * Register handlers
 */
func (p *Application) RegisterService() {
	mongodb, err := db.NewMongoClient()
	if err != nil {
		panic("cannot connect to database")
	}
	/* /api/v1/upload_image */
	uploadImageRepo := repositories.NewResultCacheRepository(mongodb)
	uploadImageService := services.NewUploadImageService(uploadImageRepo)
	p.uploadImage = handlers.NewUploadImageHandlers(uploadImageService)
}

/*
 * Define routes
 */
func (p *Application) RegisterHandler() {
	p.g.POST(fmt.Sprintf("%s%s", V1, UploadImage), p.uploadImage.Post)
}
