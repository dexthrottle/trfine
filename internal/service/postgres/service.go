package service_pg

import (
	"context"
	"trfine/internal/storage/postgres"
	"trfine/pkg/logging"
)

type Service struct {
}

func NewService(ctx context.Context, storage postgres.Storage, log logging.Logger) *Service {
	return &Service{}
}
