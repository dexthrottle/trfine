package app

import (
	"context"
	"errors"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"os"

	"github.com/dexthrottle/trfine/internal/config"
	"github.com/dexthrottle/trfine/internal/handler"
	"github.com/dexthrottle/trfine/internal/repository"
	"github.com/dexthrottle/trfine/internal/service"
	log "github.com/dexthrottle/trfine/pkg/logger"
	"github.com/dexthrottle/trfine/pkg/server"
)

func Run() {

	cfg := config.GetConfig()
	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		log.Error(err)
	}

	log.Info("Connect to database successfully!")

	ctx := context.Background()

	repos := repository.NewRepository(ctx, db)
	services := service.NewService(ctx, *repos)
	handlers := handler.NewHandler(services)

	srv := server.NewServer(cfg.App.Port, handlers.InitRoutes())

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	log.Info("Server started on http://127.0.0.1:" + cfg.App.Port)

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	log.Info("Server stopped")

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		log.Errorf("failed to stop server: %v", err)
	}
}
