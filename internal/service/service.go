package service

import (
	"context"

	"github.com/dexthrottle/trfine/internal/dto"
	"github.com/dexthrottle/trfine/internal/model"
	"github.com/dexthrottle/trfine/internal/repository"
	"github.com/dexthrottle/trfine/pkg/logging"
)

type AppConfig interface {
	InsertAppConfig(ctx context.Context, appCfg dto.AppConfigDTO) (*model.AppConfig, error)
	GetConfigData(ctx context.Context) (*model.AppConfig, error)
}

type InitData interface {
	InsertDataTradeParams(ctx context.Context) error
	InsertDataTradeInfo(ctx context.Context) error
	InsertWhiteList(ctx context.Context) error
}

// type AveragePercent interface {
// 	InsertAveragePercent(ctx context.Context, a model.AveragePercent) (*model.AveragePercent, error)
// }

type Service struct {
	AppConfig
	InitData
	// AveragePercent
}

func NewService(ctx context.Context, r repository.Repository, log logging.Logger) *Service {
	return &Service{
		AppConfig: NewAppCfgService(ctx, r.AppConfig),
		InitData:  NewInitDataService(ctx, r.InitData),
		// AveragePercent: NewAveragePercentService(ctx, r.AveragePercent, log),
	}
}
