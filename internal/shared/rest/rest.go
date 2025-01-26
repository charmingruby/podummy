package rest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
	Server *http.Server
}

func NewServer(port string) *Server {
	router := gin.Default()

	addr := fmt.Sprintf(":%s", port)

	server := http.Server{
		Addr:    addr,
		Handler: router,
	}

	return &Server{
		Router: router,
		Server: &server,
	}
}

func (r *Server) Start() error {
	return r.Server.ListenAndServe()
}

func (r *Server) Stop(ctx context.Context) error {
	if err := r.Server.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
