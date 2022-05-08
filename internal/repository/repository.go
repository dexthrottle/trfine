package repository

import (
	"context"

	"github.com/dexthrottle/trfine/internal/model"
	"github.com/dexthrottle/trfine/pkg/logging"
	"gorm.io/gorm"
)

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
	DeleteAllAveragePercent(ctx context.Context, day string) error
	GetAllAveragePercent(ctx context.Context, day string) ([]*model.AveragePercent, error)
}

type CommissionBurn interface {
	InsertCommissionBurn(ctx context.Context, a model.CommissionBurn) (*model.CommissionBurn, error)
	DeleteAllCommissionBurn(ctx context.Context, day string) error
	GetAllCommissionBurn(ctx context.Context, day string) ([]*model.CommissionBurn, error)
}

type DailyProfit interface {
	InsertDailyProfit(ctx context.Context, a model.DailyProfit) (*model.DailyProfit, error)
	DeleteAllDailyProfit(ctx context.Context, day string) error
	GetAllDailyProfit(ctx context.Context, day string) ([]*model.DailyProfit, error)
}

type Symbols interface {
	InsertSymbols(ctx context.Context, a model.Symbols) (*model.Symbols, error)
	DeleteSymbols(ctx context.Context, pair string) error
	GetAllSymbols(ctx context.Context, pair string) ([]*model.Symbols, error)
	DeleteAllSymbols(ctx context.Context) error
	UpdateSymbols(ctx context.Context, symbols model.Symbols, pair string) (*model.Symbols, error)
}

type TradeInfo interface {
	InsertTradeInfo(ctx context.Context, a model.TradeInfo) (*model.TradeInfo, error)
	DeleteTradeInfo(ctx context.Context) error
	GetAllTradeInfo(ctx context.Context) ([]*model.TradeInfo, error)
	UpdateTradeInfo(ctx context.Context, tradeInfo model.TradeInfo) (*model.TradeInfo, error)
}

type TradePairs interface {
	InsertTradePairs(ctx context.Context, a model.TradePairs) (*model.TradePairs, error)
	DeleteTradePairs(ctx context.Context) error
	GetTradePairs(ctx context.Context, pair string) (*model.TradePairs, error)
	// UpdateTradePairs(ctx context.Context, tradePairs model.TradePairs) (*model.TradePairs, error)
}

type TradeParams interface {
	InsertTradeParams(ctx context.Context, a model.TradeParams) (*model.TradeParams, error)
	DeleteTradeParams(ctx context.Context, nameList string) error
	GetTradeParams(ctx context.Context, nameList string) (*model.TradeParams, error)
	UpdateTradeParams(ctx context.Context, tradeParams model.TradeParams, nameList string) (*model.TradeParams, error)
}

type TrailingOrder interface {
	InsertTrailingOrders(ctx context.Context, a model.TrailingOrders) (*model.TrailingOrders, error)
	DeleteTrailingOrders(ctx context.Context, pair string) error
	GetTrailingOrders(ctx context.Context, pair string) (*[]model.TrailingOrders, error)
}

type Repository struct {
	AppConfig
	InitData
	AveragePercent
	CommissionBurn
	DailyProfit
	Symbols
	TradeInfo
	TradePairs
	TradeParams
	TrailingOrder
}

func NewRepository(ctx context.Context, db *gorm.DB, log logging.Logger) *Repository {
	return &Repository{
		AppConfig:      NewAppCfgRepository(ctx, db),
		InitData:       NewInitDataRepository(ctx, db, log),
		AveragePercent: NewAveragePercentRepository(ctx, db, log),
		CommissionBurn: NewCommissionBurnRepository(ctx, db, log),
		DailyProfit:    NewDailyProfitRepository(ctx, db, log),
		Symbols:        NewSymbolsRepository(ctx, db, log),
		TradeInfo:      NewTradeInfoRepository(ctx, db, log),
		TradePairs:     NewTradePairsRepository(ctx, db, log),
		TradeParams:    NewTradeParamsRepository(ctx, db, log),
		TrailingOrder:  NewTrailingOrdersRepository(ctx, db, log),
	}
}
