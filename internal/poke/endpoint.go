package poke

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Endpoint struct {
	service *Service
	router  *gin.Engine
}

func NewEndpoint(router *gin.Engine, service *Service) *Endpoint {
	return &Endpoint{
		router:  router,
		service: service,
	}
}

func (e *Endpoint) Register() {
	poke := e.router.Group("/poke")
	{
		poke.GET("/shuffle", e.makeShuffleHandler())
	}
}

func (e *Endpoint) makeShuffleHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := e.service.ShuffleTeam()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}
