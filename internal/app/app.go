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
	rcv9 "github.com/go-redis/redis/v9"
	"github.com/rs/cors"

	"trfine/internal/config"
	database "trfine/internal/db/postgres"
	"trfine/internal/db/redis"
	apiV1 "trfine/internal/handler/v1/http"
	servicePg "trfine/internal/service/postgres"
	pgStorage "trfine/internal/storage/postgres"
	"trfine/pkg/logging"
	"trfine/pkg/server"

	bybit2 "github.com/hirokisan/bybit/v2"
)

func RunApplication(saveToFile bool) {
	// Init Logger
	logging.Init(saveToFile)
	log := logging.GetLogger()

	// Init Config
	cfg := config.GetConfig()

	// Init Context
	const timeout = 5 * time.Second
	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	// Init Database
	db, err := database.NewPostgresDB(cfg, &log)
	if err != nil {
		log.Panicln(err)
	}

	// Init Redis
	rc, err := connectRedisDatabase(ctx, *cfg, log)
	if err != nil {
		log.Panicln("error connecting to redis %w", err)
	}

	// Connect storages
	storagePg := pgStorage.NewStorage(ctx, db, log)

	// Connect services
	servicesPG := servicePg.NewService(ctx, *storagePg, log)

	// Connect handlers
	handlers := apiV1.NewHandler(log, *cfg, servicesPG)

	// New Gin router
	router := gin.New()

	// Init Gin Mode
	gin.SetMode(cfg.AppConfig.GinMode)

	// Gin Logs
	enableGinLogs(saveToFile, router)

	// Init Routes
	handler := handlers.InitRoutes(router)

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

// connectRedisDatabase - подключение к базе данных redis
func connectRedisDatabase(ctx context.Context, cfg config.Config, log logging.Logger) (*rcv9.Client, error) {
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
		return nil, err
	}
	return rc, nil
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
func initByBitRest(byBitApiKey, byBitSecretkey string, log logging.Logger) *bybit2.Client {
	client := bybit2.NewClient().WithAuth(byBitApiKey, byBitSecretkey)
	return client
}

// initByBitWS - инициализирует bybit web socket клиента
func initByBitWS(byBitApiKey, byBitSecretkey string, log logging.Logger) *bybit2.WebSocketClient {
	wsClient := bybit2.NewWebsocketClient().WithAuth(byBitApiKey, byBitSecretkey)
	return wsClient
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
