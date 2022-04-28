package repository

import (
	"context"
	"errors"

	"github.com/dexthrottle/trfine/internal/model"
	"github.com/dexthrottle/trfine/pkg/logging"
	"gorm.io/gorm"
)

type InitDataRepository interface {
	InsertDataTradeParams(ctx context.Context, tradeParams model.TradeParams) (*model.TradeParams, error)
	InsertDataTradeInfo(ctx context.Context, tradeInfo model.TradeInfo) (*model.TradeInfo, error)
}

type initDataConnection struct {
	ctx        context.Context
	connection *gorm.DB
	log        logging.Logger
}

func NewInitDataRepository(ctx context.Context, db *gorm.DB, log logging.Logger) InitDataRepository {
	return &initDataConnection{
		ctx:        ctx,
		connection: db,
		log:        log,
	}
}

func (db *initDataConnection) InsertDataTradeParams(
	ctx context.Context,
	tradeParams model.TradeParams,
) (*model.TradeParams, error) {
	tx := db.connection.WithContext(ctx)

	var mdTradeParams model.TradeParams
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

func (db *initDataConnection) InsertDataTradeInfo(
	ctx context.Context,
	tradeInfo model.TradeInfo,
) (*model.TradeInfo, error) {
	tx := db.connection.WithContext(ctx)

	var mdTradeInfo model.TradeInfo
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
