package app

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/dexthrottle/trfine/internal/dto"
	"github.com/dexthrottle/trfine/internal/handler"
	"github.com/dexthrottle/trfine/internal/repository"
	"github.com/dexthrottle/trfine/internal/service"
	"github.com/dexthrottle/trfine/pkg/logging"
)

func Run(dbName string) {
	ctx := context.Background()
	reader := bufio.NewReader(os.Stdin)

	var appCfgDto dto.AppConfigDTO
	if _, err := os.Stat("rp.db"); os.IsNotExist(err) {
		appCfgDto = firstRunApp(reader)
	}
	fmt.Printf("%+v\n", appCfgDto)

	// logger init
	logging.Init()
	log := logging.GetLogger()

	// database init
	db, err := repository.NewDB(&log, dbName)
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
	log.Infof("Connect handlers successfully! %+v", handlers)

	log.Infoln("Start successfully!")

	// Graceful Shutdown ---------------------------
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Info("Exit..")
}
