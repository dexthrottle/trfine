package repository

import (
	"context"

	"github.com/dexthrottle/trfine/internal/model"
	"github.com/dexthrottle/trfine/pkg/logging"
	"gorm.io/gorm"
)

type TradeParamsRepository interface {
	InsertTradeParams(ctx context.Context, a model.TradeParams) (*model.TradeParams, error)
	DeleteTradeParams(ctx context.Context) (*model.TradeParams, error)
	GetAllTradeParams(ctx context.Context) ([]*model.TradeParams, error)
	UpdateTradeParams(ctx context.Context, tradeParams model.TradeParams) (*model.TradeParams, error)
}

type tradeParamsConnection struct {
	ctx        context.Context
	connection *gorm.DB
	log        logging.Logger
}

func NewTradeParamsRepository(ctx context.Context, db *gorm.DB, log logging.Logger) TradeParamsRepository {
	return &tradeParamsConnection{
		ctx:        ctx,
		connection: db,
		log:        log,
	}
}

func (db *tradeParamsConnection) InsertTradeParams(ctx context.Context, a model.TradeParams) (*model.TradeParams, error) {
	tx := db.connection.WithContext(ctx)
	res := tx.Save(&a)

	if res.Error != nil {
		return nil, res.Error
	}
	return &a, nil
}

func (db *tradeParamsConnection) GetAllTradeParams(ctx context.Context) ([]*model.TradeParams, error) {
	tx := db.connection.WithContext(ctx)
	var tradeParams []*model.TradeParams
	res := tx.Where(`id = ?`, 1).Find(&tradeParams)
	if res.Error != nil {
		db.log.Errorf("get all tradeParams error %v", res.Error)
		return nil, res.Error
	}
	return tradeParams, nil
}

func (db *tradeParamsConnection) UpdateTradeParams(ctx context.Context, tradeParams model.TradeParams) (*model.TradeParams, error) {
	tx := db.connection.WithContext(ctx)
	mdTradeParams := model.TradeParams{}
	res := tx.Model(&mdTradeParams).Where(`id = ?`, 1).Updates(tradeParams)
	if res.Error != nil {
		db.log.Error(res.Error)
		return nil, res.Error
	}
	return &mdTradeParams, nil
}

func (db *tradeParamsConnection) DeleteTradeParams(ctx context.Context) (*model.TradeParams, error) {
	tx := db.connection.WithContext(ctx)
	var tradeParams *model.TradeParams
	res := tx.Delete(&tradeParams).Where(`id = ?`, 1)
	if res.Error != nil {
		db.log.Errorf("delete tradeParams error %v", res.Error)
		return nil, res.Error
	}
	return tradeParams, nil
}
