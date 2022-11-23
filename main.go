package main

import "github.com/gin-gonic/gin"

func main() {
	app := NewApplication(gin.Default())
	app.Start()
}
