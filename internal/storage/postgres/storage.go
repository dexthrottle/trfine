package postgres

import (
	"context"

	"trfine/pkg/logging"

	"gorm.io/gorm"
)

type Storage struct {
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

func NewStorage(ctx context.Context, db *gorm.DB, log logging.Logger) *Storage {
	return &Storage{
		AppConfig:      NewAppCfg(ctx, db),
		InitData:       NewInitData(ctx, db, log),
		AveragePercent: NewAveragePercent(ctx, db, log),
		CommissionBurn: NewCommissionBurn(ctx, db, log),
		DailyProfit:    NewDailyProfit(ctx, db, log),
		Symbols:        NewSymbols(ctx, db, log),
		TradeInfo:      NewTradeInfo(ctx, db, log),
		TradePairs:     NewTradePairs(ctx, db, log),
		TradeParams:    NewTradeParams(ctx, db, log),
		TrailingOrder:  NewTrailingOrders(ctx, db, log),
	}
}
