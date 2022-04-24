package app

import (
	"bufio"
	"context"
	"errors"
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
	ctx := context.Background()

	// Создаем объект читателя
	reader := bufio.NewReader(os.Stdin)
	// Проверяем есть файл с базой данных,
	// если нет - запрашиваем порт и логи, если есть идем дальше
	useLogs, portApp := firstStart(reader)

	// logger init
	logging.Init(useLogs)
	log := logging.GetLogger()

	// config init
	cfg := config.GetConfig(useLogs, strings.TrimSuffix(portApp, "\n"))
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

	// Проверяем, есть ли конфигурация приложения, если есть запускаемся
	// если нет - запрашиваем данные у пользователя
	cfgAppCheck, err := services.CheckConfigData(ctx)
	if cfgAppCheck.ID == 0 {
		appCfgDto := secondStart(reader)
		mAppConfig, err := services.AppConfig.InsertAppConfig(ctx, appCfgDto)
		if err != nil {
			panic("Не удалось сохранить конфигурацию" + err.Error() + "\n")
		}
		log.Infof("%+v\n", mAppConfig)
	}
	if err != nil {
		panic("Ошибка при получение конфигурации\n")
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
