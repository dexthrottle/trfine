package app

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"

	"trfine/internal/bybit"
	"trfine/internal/config"
	database "trfine/internal/db/postgres"
	"trfine/internal/db/redis"
	apiV1 "trfine/internal/handler/v1/http"
	servicePg "trfine/internal/service/postgres"
	pgStorage "trfine/internal/storage/postgres"
	"trfine/pkg/bybitapi/rest"
	"trfine/pkg/bybitapi/ws"
	"trfine/pkg/logging"
	"trfine/pkg/server"
)

func RunApplication(saveToFile bool) {
	// Init Logger
	logging.Init(saveToFile)
	log := logging.GetLogger()
	log.Infoln("Connect logger successfully!")

	// Init Config
	cfg := config.GetConfig()
	log.Infoln("Connect config successfully!")

	// Init Context
	const timeout = 5 * time.Second
	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	// connect to redis
	redisClient := redis.NewRedisClient(
		ctx,
		&redis.CredentialRedis{
			Host:   cfg.Redis.Host,
			Port:   cfg.Redis.Port,
			Secret: cfg.Redis.Secret,
			Size:   cfg.Redis.Size,
		},
		log,
	)
	rc, err := redisClient.ConnectToRedis()
	if err != nil {
		log.Panicln("error connecting to redis %w", err)
	}

	// Init Database
	db, err := database.NewPostgresDB(cfg, &log)
	if err != nil {
		log.Panicln(err)
	}
	log.Infoln("Connect database successfully!")

	// Connect storages
	storagePg := pgStorage.NewStorage(ctx, db, log)
	log.Infoln("Connect storage postgres successfully!")

	// Connect services
	servicesPG := servicePg.NewService(ctx, *storagePg, log)
	log.Infoln("Connect service postgres successfully!")

	// Connect handlers
	handlers := apiV1.NewHandler(log, *cfg, servicesPG)
	log.Infoln("Connect services handlers!")

	// New Gin router
	router := gin.New()

	// Init Gin Mode
	gin.SetMode(cfg.AppConfig.GinMode)

	// Gin Logs
	enableGinLogs(saveToFile, router)

	// Init Routes and CORS
	handler := initRoutesAndCORS(router, handlers)

	// Start HTTP Server
	srv := server.NewServer(cfg.Listen.Port, handler)
	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Panicln("error occurred while running http server: " + err.Error())
		}
	}()
	log.Infoln("Server started on http://" + cfg.Listen.BindIP + ":" + cfg.Listen.Port)

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	log.Info("Server stopped")

	if err := srv.Stop(ctx); err != nil {
		log.Panicf("failed to stop server: %v\n", err)
	}

	if err := rc.Close(); err != nil {
		log.Panicf("error closing Redis Client: %w\n", err)
	}
}

// initRoutesAndCORS инициализирует роутер и обработчики
func initRoutesAndCORS(router *gin.Engine, handlers *apiV1.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodOptions, http.MethodDelete},
		AllowedOrigins:     []string{"http://127.0.0.1:8000", "http://127.0.0.1:8000", "http://localhost:8000"},
		AllowCredentials:   true,
		AllowedHeaders:     []string{"Location", "Charset", "Access-Control-Allow-Origin", "Content-Type", "content-type", "Origin", "Accept", "Content-Length", "Accept-Encoding", "X-CSRF-Token"},
		OptionsPassthrough: true,
		ExposedHeaders:     []string{"Location", "Authorization", "Content-Disposition"},
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	handler := c.Handler(handlers.InitRoutes(router))
	return handler
}

// enableGinLogs включает/отключает gin логи
func enableGinLogs(saveToFile bool, router *gin.Engine) {
	if saveToFile {
		allFile, err := os.OpenFile("logs/gin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
		if err != nil {
			panic(fmt.Sprintf("[Message]: %s", err))
		}
		gin.DefaultWriter = io.MultiWriter(allFile)
	}

	router.Use(gin.Logger())
}

// initByBitWS - инициализирует bybit rest клиента
func initByBitRest(byBitApiKey, byBitSecretkey, baseUrlByBit string, log logging.Logger) (bybit.ByBitAPIRest, error) {
	bybitRest := rest.New(nil, baseUrlByBit, byBitApiKey, byBitSecretkey, true)
	err := bybitRest.SetCorrectServerTime()
	if err != nil {
		return nil, err
	}
	bbAPIRest := bybit.NewByBit(log, bybitRest)

	return bbAPIRest, nil
}

// initByBitWS - инициализирует bybit web socket клиента
func initByBitWS(byBitApiKey, byBitSecretkey string, log logging.Logger) bybit.ByBitAPIWS {
	cfg := &ws.Configuration{
		Addr:          ws.HostTestnet,
		ApiKey:        byBitApiKey,
		SecretKey:     byBitSecretkey,
		AutoReconnect: true,
		DebugMode:     true,
	}
	bybitWS := ws.New(cfg)
	bbAPIWS := bybit.NewByBitWS(log, bybitWS)
	return bbAPIWS
}

// func initDefaultData(ctx context.Context, services service.Service) {
// 	err := services.InitData.InsertDataTradeParams(ctx)
// 	if err != nil {
// 		panic("Не удалось сохранить дефолтные значения!")
// 	}

// 	err = services.InitData.InsertDataTradeInfo(ctx)
// 	if err != nil {
// 		panic("Не удалось сохранить дефолтные значения!")
// 	}

// 	err = services.InsertWhiteList(ctx)
// 	if err != nil {
// 		panic("Не удалось сохранить дефолтные значения!")
// 	}
// }
