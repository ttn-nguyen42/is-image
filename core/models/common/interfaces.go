package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
 * Ports for services to access HTTP requests
 */
type IRequest struct {
	http.Request
	gin.Params
}
