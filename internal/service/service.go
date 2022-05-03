package service

import (
	"context"

	"github.com/dexthrottle/trfine/internal/dto"
	"github.com/dexthrottle/trfine/internal/model"
	"github.com/dexthrottle/trfine/internal/repository"
	"github.com/dexthrottle/trfine/pkg/logging"
)

type User interface {
	Insert(ctx context.Context, user dto.CreateUserDTO) (*model.User, error)
	Profile(ctx context.Context, userID string) (*model.User, error)
	IsDuplicateUserTGID(ctx context.Context, tgID int) (bool, error)
	FindUserByTgUserId(ctx context.Context, userTgId int) (*model.User, error)
}

type AppConfig interface {
	InsertAppConfig(ctx context.Context, appCfg dto.AppConfigDTO) (*model.AppConfig, error)
	GetConfigData(ctx context.Context) (*model.AppConfig, error)
}

type InitData interface {
	InsertDataTradeParams(ctx context.Context) error
	InsertDataTradeInfo(ctx context.Context) error
	InsertWhiteList(ctx context.Context) error
}

type Service struct {
	User
	AppConfig
	InitData
}

func NewService(ctx context.Context, r repository.Repository, log logging.Logger) *Service {
	return &Service{
		User:      NewUserService(ctx, r.User, log),
		AppConfig: NewAppCfgService(ctx, r.AppConfig),
		InitData:  NewInitDataService(ctx, r.InitData),
	}
}
