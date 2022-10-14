package postgres

import (
	"context"

	"trfine/internal/entity"
	"trfine/pkg/logging"

	"gorm.io/gorm"
)

type TradeInfo interface {
	InsertTradeInfo(ctx context.Context, a entity.TradeInfo) (*entity.TradeInfo, error)
	DeleteTradeInfo(ctx context.Context) error
	GetAllTradeInfo(ctx context.Context) ([]*entity.TradeInfo, error)
	UpdateTradeInfo(ctx context.Context, tradeInfo entity.TradeInfo) (*entity.TradeInfo, error)
}

type tradeInfo struct {
	ctx context.Context
	db  *gorm.DB
	log logging.Logger
}

func NewTradeInfo(ctx context.Context, db *gorm.DB, log logging.Logger) TradeInfo {
	return &tradeInfo{
		ctx: ctx,
		db:  db,
		log: log,
	}
}

func (db *tradeInfo) InsertTradeInfo(ctx context.Context, a entity.TradeInfo) (*entity.TradeInfo, error) {
	tx := db.db.WithContext(ctx)
	res := tx.Save(&a)

	if res.Error != nil {
		return nil, res.Error
	}
	return &a, nil
}

func (db *tradeInfo) GetAllTradeInfo(ctx context.Context) ([]*entity.TradeInfo, error) {
	tx := db.db.WithContext(ctx)
	var tradeInfo []*entity.TradeInfo
	res := tx.Where(`id = ?`, 1).Find(&tradeInfo)
	if res.Error != nil {
		db.log.Errorf("get all tradeInfo error %v", res.Error)
		return nil, res.Error
	}
	return tradeInfo, nil
}

func (db *tradeInfo) UpdateTradeInfo(ctx context.Context, tradeInfo entity.TradeInfo) (*entity.TradeInfo, error) {
	tx := db.db.WithContext(ctx)
	mdTradeInfo := entity.TradeInfo{}
	res := tx.Model(&mdTradeInfo).Where(`id = ?`, 1).Updates(tradeInfo)
	if res.Error != nil {
		db.log.Error(res.Error)
		return nil, res.Error
	}
	return &mdTradeInfo, nil
}

func (db *tradeInfo) DeleteTradeInfo(ctx context.Context) error {
	tx := db.db.WithContext(ctx)
	var tradeInfo *entity.TradeInfo
	res := tx.Delete(&tradeInfo).Where(`id = ?`, 1)
	if res.Error != nil {
		db.log.Errorf("delete tradeInfo error %v", res.Error)
		return res.Error
	}
	return nil
}
