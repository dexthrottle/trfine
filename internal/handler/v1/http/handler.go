package http

import (
	"trfine/internal/config"
	service_pg "trfine/internal/service/postgres"
	"trfine/pkg/logging"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service_pg.Service
	cfg     config.Config
	log     logging.Logger
}

func NewHandler(
	log logging.Logger,
	cfg config.Config,
	service *service_pg.Service,

) *Handler {
	return &Handler{
		service: service,
		cfg:     cfg,
		log:     log,
	}
}

func (h *Handler) InitRoutes(r *gin.Engine) *gin.Engine {

	api := r.Group("api/")
	{
		api.GET("test", h.Test)

	}

	return r
}

func (h *Handler) Test(ctx *gin.Context) {
	ctx.JSON(200, "qwerty")
}
