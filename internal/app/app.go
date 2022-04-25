package app

import (
	"bufio"
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dexthrottle/trfine/internal/config"
	"github.com/dexthrottle/trfine/internal/dto"
	"github.com/dexthrottle/trfine/internal/handler"
	"github.com/dexthrottle/trfine/internal/repository"
	"github.com/dexthrottle/trfine/internal/service"
	"github.com/dexthrottle/trfine/pkg/logging"
	"github.com/dexthrottle/trfine/pkg/server"
	"github.com/gin-gonic/gin"
)

func Run() {
	ctx := context.Background()
	reader := bufio.NewReader(os.Stdin)

	var appCfgDto dto.AppConfigDTO
	var appPort string
	if _, err := os.Stat("trbotdatabase.db"); os.IsNotExist(err) {
		appCfgDto, appPort = firstRunApp(reader)
	}

	// logger init
	logging.Init()
	log := logging.GetLogger()

	// config init
	cfg := config.GetConfig("debug", appPort)
	log.Info("config init")

	// database init
	db, err := repository.NewPostgresDB(cfg, &log)
	if err != nil {
		panic("database connect error" + err.Error())
	}
	log.Info("Connect to database successfully!")

	// repositories init
	repos := repository.NewRepository(ctx, db, log)
	log.Info("Connect repository successfully!")

	// services init
	services := service.NewService(ctx, *repos, log)
	log.Info("Connect services successfully!")

	// handlers init
	handlers := handler.NewHandler(services, log)
	log.Info("Connect handlers successfully!")
	if cfg.App.GinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	_, err = services.AppConfig.InsertAppConfig(ctx, appCfgDto)
	if err != nil {
		panic("Ошибка с сохранением конфигурации: " + err.Error())
	}

	// server start
	srv := server.NewServer(cfg.App.Port, handlers.InitRoutes())
	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			panic("error occurred while running http server: " + err.Error() + "\n")
		}
	}()
	log.Info("Server started on http://127.0.0.1:" + cfg.App.Port + " Gin MODE = " + gin.Mode())

	// Graceful Shutdown ---------------------------
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
