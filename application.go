package main

import (
	"context"
	"fmt"
	"log"
	"os"

	hl "is-image/handler"
	m "is-image/model"
	reps "is-image/repository"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
 * Application entry point
 */
type Application struct {
	g *gin.Engine
}

func NewApplication(engine *gin.Engine) *Application {
	app := Application{
		g: engine,
	}
	return &app
}

func (p *Application) Start() {
	p.RegisterService()
	p.RegisterHandler()
	port := os.Getenv(m.PORT)
	if port == "" {
		port = "8080"
	}
	log.Printf("Application starts in port %s", port)
	p.g.Run(fmt.Sprintf(":%s", port))
}

func (p *Application) RegisterService() {
	mongoUrl := os.Getenv("MONGO_URL")
	if mongoUrl == "" {
		return
	}
	_, err := reps.NewResultCacheRepository(options.Client().ApplyURI(mongoUrl), context.TODO())
	if err != nil {
		log.Println(err.Error())
	}
}

/*
 * Define routes
 */
func (p *Application) RegisterHandler() {
	p.g.POST(fmt.Sprintf("%s%s", m.V1, m.UploadImage), func(ctx *gin.Context) {
		hl.PostUploadImage(ctx, &hl.PostUploadImageDependencies{
			Repo: reps.GetResultCacheRepositoryInstance(),
		})
	})
}
