package postgres

import (
	"context"

	"trfine/internal/entity"
	"trfine/pkg/logging"

	"gorm.io/gorm"
)

type TradeParams interface {
	InsertTradeParams(ctx context.Context, a entity.TradeParams) (*entity.TradeParams, error)
	DeleteTradeParams(ctx context.Context, nameList string) error
	GetTradeParams(ctx context.Context, nameList string) (*entity.TradeParams, error)
	UpdateTradeParams(ctx context.Context, tradeParams entity.TradeParams, nameList string) (*entity.TradeParams, error)
}

type tradeParams struct {
	ctx context.Context
	db  *gorm.DB
	log logging.Logger
}

func NewTradeParams(ctx context.Context, db *gorm.DB, log logging.Logger) TradeParams {
	return &tradeParams{
		ctx: ctx,
		db:  db,
		log: log,
	}
}

func (db *tradeParams) InsertTradeParams(ctx context.Context, a entity.TradeParams) (*entity.TradeParams, error) {
	tx := db.db.WithContext(ctx)
	res := tx.Save(&a)
	if res.Error != nil {
		return nil, res.Error
	}
	return &a, nil
}

func (db *tradeParams) GetTradeParams(ctx context.Context, nameList string) (*entity.TradeParams, error) {
	tx := db.db.WithContext(ctx)
	var tradeParams *entity.TradeParams
	res := tx.Where(`name_list = ?`, nameList).Find(&tradeParams)
	if res.Error != nil {
		db.log.Errorf("get trade Params error %v", res.Error)
		return nil, res.Error
	}
	return tradeParams, nil
}

func (db *tradeParams) UpdateTradeParams(
	ctx context.Context,
	tradeParams entity.TradeParams,
	nameList string,
) (*entity.TradeParams, error) {

	tx := db.db.WithContext(ctx)
	mdTradeParams := entity.TradeParams{}
	res := tx.Model(&mdTradeParams).Where(`name_list = ?`, nameList).Updates(tradeParams)
	if res.Error != nil {
		db.log.Error(res.Error)
		return nil, res.Error
	}
	return &mdTradeParams, nil
}

func (db *tradeParams) DeleteTradeParams(ctx context.Context, nameList string) error {
	tx := db.db.WithContext(ctx)
	var tradeParams *entity.TradeParams
	res := tx.Where(`name_list = ?`, nameList).Delete(&tradeParams)
	if res.Error != nil {
		db.log.Errorf("delete trade Params error %v", res.Error)
		return res.Error
	}
	return nil
}
