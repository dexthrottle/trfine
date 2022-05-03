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
	"github.com/dexthrottle/trfine/internal/model"
	"github.com/dexthrottle/trfine/internal/repository"
	"github.com/dexthrottle/trfine/internal/service"
	"github.com/dexthrottle/trfine/pkg/logging"
)

func Run(dbName string) {
	ctx := context.Background()
	reader := bufio.NewReader(os.Stdin)

	var appCfgDto dto.AppConfigDTO
	if _, err := os.Stat(fmt.Sprintf("%s.db", dbName)); os.IsNotExist(err) {
		appCfgDto = firstRunApp(reader)
	}

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

	// testSymbols(ctx, repos)

	// services init
	services := service.NewService(ctx, *repos, log)
	log.Info("Connect services successfully!")

	_, err = services.AppConfig.InsertAppConfig(ctx, appCfgDto)
	if err != nil {
		panic("Не удалось сохранить конфигурацию!")
	}

	// Add first data
	initDefaultData(ctx, *services)

	// handlers init
	handlers := handler.NewHandler(services, log)

	// TODO: Говнокод
	log.Infof("Connect handlers successfully! %+v", handlers)

	log.Infoln("Start successfully!")

	// Graceful Shutdown ---------------------------
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Info("Exit..")
}

func testSymbols(ctx context.Context, repos *repository.Repository) {
	m, err := repos.TradePairs.InsertTradePairs(ctx, model.TradePairs{Pair: "eqweqwe"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("ID = %d\n", &m.ID)
	// err := repos.TradePairs.DeleteTradePairs(ctx)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}

func initDefaultData(ctx context.Context, services service.Service) {
	err := services.InitData.InsertDataTradeParams(ctx)
	if err != nil {
		panic("Не удалось сохранить дефолтные значения!")
	}

	err = services.InitData.InsertDataTradeInfo(ctx)
	if err != nil {
		panic("Не удалось сохранить дефолтные значения!")
	}

	err = services.InsertWhiteList(ctx)
	if err != nil {
		panic("Не удалось сохранить дефолтные значения!")
	}
}
