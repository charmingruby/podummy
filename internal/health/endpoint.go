package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Endpoint struct {
	Router *gin.Engine
}

func NewEndpoint(router *gin.Engine) *Endpoint {
	return &Endpoint{
		Router: router,
	}
}

func (e *Endpoint) Register() {
	e.Router.GET("/health-check", e.makeHealthCheckHandler())
}

func (e *Endpoint) makeHealthCheckHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
	}
}
