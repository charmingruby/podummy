package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/charmingruby/podummy/internal/health"
	"github.com/charmingruby/podummy/internal/poke"
	"github.com/charmingruby/podummy/internal/shared/rest"
	"github.com/charmingruby/podummy/pkg/external/client/pokeapi"
	"github.com/gin-gonic/gin"
)

func main() {
	server := rest.NewServer("8080")

	initModules(server.Router)

	go func() {
		if err := server.Start(); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Stop(ctx); err != nil {
		log.Fatalf("Failed to stop server: %v", err)
	}

	<-ctx.Done()
	log.Println("timeout of 5 seconds.")
	log.Println("Server exiting")
}

func initModules(router *gin.Engine) {
	health.NewEndpoint(router).Register()

	pokeAPI := pokeapi.NewPokeAPI()
	pokeSvc := poke.NewService(pokeAPI)
	poke.NewEndpoint(router, pokeSvc).Register()
}
