package postgres

import (
	"context"
	"errors"
	"trfine/internal/entity"
	"trfine/pkg/logging"

	"gorm.io/gorm"
)

type InitData interface {
	InsertDataTradeParams(ctx context.Context, tradeParams entity.TradeParams) (*entity.TradeParams, error)
	InsertDataTradeInfo(ctx context.Context, tradeInfo entity.TradeInfo) (*entity.TradeInfo, error)
	InsertWhiteList(ctx context.Context, whiteList []entity.WhiteList) (*[]entity.WhiteList, error)
}

type initData struct {
	ctx context.Context
	db  *gorm.DB
	log logging.Logger
}

func NewInitData(ctx context.Context, db *gorm.DB, log logging.Logger) InitData {
	return &initData{
		ctx: ctx,
		db:  db,
		log: log,
	}
}

func (db *initData) InsertDataTradeParams(
	ctx context.Context,
	tradeParams entity.TradeParams,
) (*entity.TradeParams, error) {
	tx := db.db.WithContext(ctx)

	var mdTradeParams entity.TradeParams
	_ = tx.Where(`"external_id" = ?`, 1).Find(&mdTradeParams)
	if mdTradeParams.ExternalID == 0 {
		res := tx.Save(&tradeParams)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			db.log.Errorf("insert trade Params error %v", res.Error)
			return nil, res.Error
		}
	}

	return &tradeParams, nil
}

func (db *initData) InsertDataTradeInfo(
	ctx context.Context,
	tradeInfo entity.TradeInfo,
) (*entity.TradeInfo, error) {
	tx := db.db.WithContext(ctx)

	var mdTradeInfo entity.TradeInfo
	_ = tx.Where(`"external_id" = ?`, 1).Find(&mdTradeInfo)
	if mdTradeInfo.ExternalID == 0 {
		res := tx.Save(&tradeInfo)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			db.log.Errorf("insert trade Info error %v", res.Error)
			return nil, res.Error
		}
	}
	return &tradeInfo, nil
}

func (db *initData) InsertWhiteList(
	ctx context.Context,
	whiteList []entity.WhiteList,
) (*[]entity.WhiteList, error) {
	tx := db.db.WithContext(ctx)

	var mdWhiteList entity.WhiteList
	_ = tx.Where(`"id" = ?`, 1).Find(&mdWhiteList)
	if mdWhiteList.ID == 0 {
		res := tx.Save(&whiteList)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			db.log.Errorf("insert white List error %v", res.Error)
			return nil, res.Error
		}
	}
	return &whiteList, nil
}
