package repository

import (
	"context"

	"github.com/dexthrottle/trfine/internal/model"
	"github.com/dexthrottle/trfine/pkg/logging"
	"gorm.io/gorm"
)

type TradeInfoRepository interface {
	InsertTradeInfo(ctx context.Context, a model.TradeInfo) (*model.TradeInfo, error)
	DeleteTradeInfo(ctx context.Context) (*model.TradeInfo, error)
	GetAllTradeInfo(ctx context.Context) ([]*model.TradeInfo, error)
	UpdateTradeInfo(ctx context.Context, tradeInfo model.TradeInfo) (*model.TradeInfo, error)
}

type tradeInfoConnection struct {
	ctx        context.Context
	connection *gorm.DB
	log        logging.Logger
}

func NewTradeInfoRepository(ctx context.Context, db *gorm.DB, log logging.Logger) TradeInfoRepository {
	return &tradeInfoConnection{
		ctx:        ctx,
		connection: db,
		log:        log,
	}
}

func (db *tradeInfoConnection) InsertTradeInfo(ctx context.Context, a model.TradeInfo) (*model.TradeInfo, error) {
	tx := db.connection.WithContext(ctx)
	res := tx.Save(&a)

	if res.Error != nil {
		return nil, res.Error
	}
	return &a, nil
}

func (db *tradeInfoConnection) GetAllTradeInfo(ctx context.Context) ([]*model.TradeInfo, error) {
	tx := db.connection.WithContext(ctx)
	var tradeInfo []*model.TradeInfo
	res := tx.Where(`id = ?`, 1).Find(&tradeInfo)
	if res.Error != nil {
		db.log.Errorf("get all tradeInfo error %v", res.Error)
		return nil, res.Error
	}
	return tradeInfo, nil
}

func (db *tradeInfoConnection) UpdateTradeInfo(ctx context.Context, tradeInfo model.TradeInfo) (*model.TradeInfo, error) {
	tx := db.connection.WithContext(ctx)
	mdTradeInfo := model.TradeInfo{}
	res := tx.Model(&mdTradeInfo).Where(`id = ?`, 1).Updates(tradeInfo)
	if res.Error != nil {
		db.log.Error(res.Error)
		return nil, res.Error
	}
	return &mdTradeInfo, nil
}

func (db *tradeInfoConnection) DeleteTradeInfo(ctx context.Context) (*model.TradeInfo, error) {
	tx := db.connection.WithContext(ctx)
	var tradeInfo *model.TradeInfo
	res := tx.Delete(&tradeInfo).Where(`id = ?`, 1)
	if res.Error != nil {
		db.log.Errorf("delete tradeInfo error %v", res.Error)
		return nil, res.Error
	}
	return tradeInfo, nil
}
