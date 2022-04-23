package app

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/dexthrottle/trfine/internal/config"
	"github.com/dexthrottle/trfine/internal/handler"
	"github.com/dexthrottle/trfine/internal/repository"
	"github.com/dexthrottle/trfine/internal/service"
	"github.com/dexthrottle/trfine/pkg/logging"
	"github.com/dexthrottle/trfine/pkg/server"
	"github.com/gin-gonic/gin"
)

func Run() {

	// TODO: проверить наличие файла с базой, если есть
	// то сделать запрос на AppConfig и получить конфиг для приложения

	reader := bufio.NewReader(os.Stdin)
	var useLogs bool
	fmt.Print("Включить логгирование? (Y/n): ")
	useLogsText, _ := reader.ReadString('\n')
	if strings.TrimSuffix(useLogsText, "\n") == "Y" || strings.TrimSuffix(useLogsText, "\n") == "y" {
		useLogs = true
	} else {
		useLogs = false
	}
	fmt.Print("Введите порт для запуска приложения: ")
	portApp, _ := reader.ReadString('\n')

	logging.Init(useLogs)
	log := logging.GetLogger()

	cfg := config.GetConfig(useLogs, strings.TrimSuffix(portApp, "\n"))
	log.Info("config init")

	db, err := repository.NewPostgresDB(cfg, &log)
	if err != nil {
		panic("database connect error" + err.Error())
	}
	log.Info("Connect to database successfully!")

	firstStart()

	ctx := context.Background()

	repos := repository.NewRepository(ctx, db, log)
	log.Info("Connect repository successfully!")
	services := service.NewService(ctx, *repos, log)
	log.Info("Connect services successfully!")
	handlers := handler.NewHandler(services, log)
	log.Info("Connect handlers successfully!")
	if cfg.App.GinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	srv := server.NewServer(cfg.App.Port, handlers.InitRoutes())

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			panic("error occurred while running http server: " + err.Error())
		}
	}()

	log.Info("Server started on http://127.0.0.1:" + cfg.App.Port + " Gin MODE = " + gin.Mode())

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
