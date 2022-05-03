package repository

import (
	"context"

	"github.com/dexthrottle/trfine/internal/model"
	"github.com/dexthrottle/trfine/pkg/logging"
	"gorm.io/gorm"
)

type User interface {
	InsertUser(ctx context.Context, user model.User) (*model.User, error)
	ProfileUser(ctx context.Context, userID string) (*model.User, error)
	IsDuplicateUserTGID(ctx context.Context, tgID int) (bool, error)
	FindUserByTgUserId(ctx context.Context, userTgId int) (*model.User, error)
}

type AppConfig interface {
	InsertAppConfig(ctx context.Context, appCfg model.AppConfig) (*model.AppConfig, error)
	GetConfigData(ctx context.Context) (*model.AppConfig, error)
}

type InitData interface {
	InsertDataTradeParams(ctx context.Context, tradeParams model.TradeParams) (*model.TradeParams, error)
	InsertDataTradeInfo(ctx context.Context, tradeInfo model.TradeInfo) (*model.TradeInfo, error)
	InsertWhiteList(ctx context.Context, whiteList []model.WhiteList) (*[]model.WhiteList, error)
}

type AveragePercent interface {
	InsertAveragePercent(ctx context.Context, a model.AveragePercent) (*model.AveragePercent, error)
	DeleteAllAveragePercent(ctx context.Context, day string) ([]*model.AveragePercent, error)
	GetAllAveragePercent(ctx context.Context, day string) ([]*model.AveragePercent, error)
}

type CommissionBurn interface {
	InsertCommissionBurn(ctx context.Context, a model.CommissionBurn) (*model.CommissionBurn, error)
	DeleteAllCommissionBurn(ctx context.Context, day string) ([]*model.CommissionBurn, error)
	GetAllCommissionBurn(ctx context.Context, day string) ([]*model.CommissionBurn, error)
}

type DailyProfit interface {
	InsertDailyProfit(ctx context.Context, a model.DailyProfit) (*model.DailyProfit, error)
	DeleteAllDailyProfit(ctx context.Context, day string) ([]*model.DailyProfit, error)
	GetAllDailyProfit(ctx context.Context, day string) ([]*model.DailyProfit, error)
}

type Symbols interface {
	InsertSymbols(ctx context.Context, a model.Symbols) (*model.Symbols, error)
	DeleteSymbols(ctx context.Context, pair string) (*model.Symbols, error)
	GetAllSymbols(ctx context.Context, pair string) ([]*model.Symbols, error)
	DeleteAllSymbols(ctx context.Context) error
	UpdateSymbols(ctx context.Context, symbols model.Symbols, pair string) (*model.Symbols, error)
}

type Repository struct {
	User
	AppConfig
	InitData
	AveragePercent
	CommissionBurn
	DailyProfit
	Symbols
}

func NewRepository(ctx context.Context, db *gorm.DB, log logging.Logger) *Repository {
	return &Repository{
		User:           NewUserRepository(ctx, db, log),
		AppConfig:      NewAppCfgRepository(ctx, db),
		InitData:       NewInitDataRepository(ctx, db, log),
		AveragePercent: NewAveragePercentRepository(ctx, db, log),
		CommissionBurn: NewCommissionBurnRepository(ctx, db, log),
		DailyProfit:    NewDailyProfitRepository(ctx, db, log),
		Symbols:        NewSymbolsRepository(ctx, db, log),
	}
}
