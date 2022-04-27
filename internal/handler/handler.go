package handler

import (
	"github.com/dexthrottle/trfine/internal/service"
	"github.com/dexthrottle/trfine/pkg/logging"
)

type Handler struct {
	service *service.Service
	log     logging.Logger
}

func NewHandler(service *service.Service, log logging.Logger) *Handler {
	return &Handler{
		service: service,
		log:     log,
	}
}
