package main

import (
	"fmt"
	"log"
	"os"

	hl "is-image/handler"

	"github.com/gin-gonic/gin"
)

/*
 * Application entry point
 */
type application struct {
	g *gin.Engine
}

func (p *application) _init() {
	p.g = gin.Default()
}

func (p *application) _start() {
	p._init()
	p._setHandler()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Application starts in port %s", port)
	p.g.Run(fmt.Sprintf(":%s", port))
}

/*
 * Define routes
 */
func (p *application) _setHandler() {
	p.g.POST(fmt.Sprintf("%s%s", V1, uploadImage), hl.PostUploadImage)
}
