package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/charmingruby/podummy/config"
	"github.com/charmingruby/podummy/internal/health"
	"github.com/charmingruby/podummy/internal/poke"
	"github.com/charmingruby/podummy/internal/shared/rest"
	"github.com/charmingruby/podummy/pkg/external/client/pokeapi"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := godotenv.Load(); err != nil {
		slog.Warn("CONFIGURATION: .env file not found")
	}

	cfg, err := config.New()
	if err != nil {
		slog.Error(fmt.Sprintf("CONFIGURATION: Failed to load configuration: %s", err.Error()))
	}

	server := rest.NewServer("8080")

	initModules(server.Router, cfg)

	go func() {
		if err := server.Start(); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Stop(ctx); err != nil {
		log.Fatalf("Failed to stop server: %v", err)
	}

	<-ctx.Done()
	slog.Info("timeout of 5 seconds.")
	slog.Info("Server exiting")
}

func initModules(router *gin.Engine, cfg *config.Wrapper) {
	health.NewEndpoint(router).Register()

	pokeAPI := pokeapi.NewPokeAPI()
	pokeSvc := poke.NewService(cfg.Versioning.Version, pokeAPI)
	poke.NewEndpoint(router, pokeSvc).Register()
}
