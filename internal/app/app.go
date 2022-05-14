package app

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/dexthrottle/trfine/internal/bybit"
	"github.com/dexthrottle/trfine/internal/dto"
	"github.com/dexthrottle/trfine/internal/license"
	"github.com/dexthrottle/trfine/internal/repository"
	"github.com/dexthrottle/trfine/internal/service"
	"github.com/dexthrottle/trfine/pkg/logging"
	"github.com/frankrap/bybit-api/rest"
	"github.com/frankrap/bybit-api/ws"
)

func Run(dbName, baseURL string) {
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

	// services init
	services := service.NewService(ctx, *repos, log)
	log.Info("Connect services successfully!")

	appCfg, err := services.AppConfig.InsertAppConfig(ctx, appCfgDto)
	if err != nil {
		panic("Не удалось сохранить конфигурацию!")
	}

	// Add first data
	initDefaultData(ctx, *services)
	log.Infoln("Start successfully!")

	// Инициализация REST ByBit ---------------------------------------------------
	bbAPIRest := initByBitRest(appCfg.ByBitApiKey, appCfg.ByBitApiSecret, baseURL, log, services)
	log.Printf("%+v", bbAPIRest)
	// ----------------------------------------------------------------------------

	// Инициализация WebSocker ByBit -------------------------------------
	bbAPIWS := initByBitWS(appCfg.ByBitApiKey, appCfg.ByBitApiSecret, log, services)
	log.Printf("%+v", bbAPIWS)
	// -------------------------------------------------------------------

	// Проверка лицензии Бота
	license.NewLicenseProgram(log).CheckLicense()

	// Graceful Shutdown ---------------------------
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Info("Exit..")
}

func initByBitRest(byBitApiKey, byBitSecretkey, baseUrlByBit string, log logging.Logger, services *service.Service) bybit.ByBitAPIRest {
	bybitRest := rest.New(nil, baseUrlByBit, byBitApiKey, byBitSecretkey, true)
	bbAPIRest := bybit.NewByBit(log, bybitRest, services)
	return bbAPIRest
}

func initByBitWS(byBitApiKey, byBitSecretkey string, log logging.Logger, services *service.Service) bybit.ByBitAPIWS {
	cfg := &ws.Configuration{
		Addr:          ws.HostTestnet,
		ApiKey:        byBitApiKey,
		SecretKey:     byBitSecretkey,
		AutoReconnect: true,
		DebugMode:     true,
	}
	bybitWS := ws.New(cfg)
	bbAPIWS := bybit.NewByBitWS(log, bybitWS, services)
	return bbAPIWS
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
