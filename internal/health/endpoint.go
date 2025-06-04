package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Endpoint struct {
	router *gin.Engine
}

func NewEndpoint(router *gin.Engine) *Endpoint {
	return &Endpoint{
		router: router,
	}
}

func (e *Endpoint) Register() {
	api := e.router.Group("/api")
	api.GET("/health/live", e.makeBasicHealthCheckEndpoint())
	api.GET("/health/ready", e.makeBasicHealthCheckEndpoint())
}

func (e *Endpoint) makeBasicHealthCheckEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
	}
}
