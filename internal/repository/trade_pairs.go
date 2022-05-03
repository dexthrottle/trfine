package repository

import (
	"context"

	"github.com/dexthrottle/trfine/internal/model"
	"github.com/dexthrottle/trfine/pkg/logging"
	"gorm.io/gorm"
)

type TradePairsRepository interface {
	InsertTradePairs(ctx context.Context, a model.TradePairs) (*model.TradePairs, error)
	DeleteTradePairs(ctx context.Context) error
	GetTradePairs(ctx context.Context, pair string) (*model.TradePairs, error)
	// UpdateTradePairs(ctx context.Context, tradePairs model.TradePairs) (*model.TradePairs, error)
}

type tradePairsConnection struct {
	ctx        context.Context
	connection *gorm.DB
	log        logging.Logger
}

func NewTradePairsRepository(ctx context.Context, db *gorm.DB, log logging.Logger) TradePairsRepository {
	return &tradePairsConnection{
		ctx:        ctx,
		connection: db,
		log:        log,
	}
}

func (db *tradePairsConnection) InsertTradePairs(ctx context.Context, a model.TradePairs) (*model.TradePairs, error) {
	tx := db.connection.WithContext(ctx)
	res := tx.Save(&a)
	if res.Error != nil {
		return nil, res.Error
	}
	return &a, nil
}

func (db *tradePairsConnection) GetTradePairs(ctx context.Context, pair string) (*model.TradePairs, error) {
	tx := db.connection.WithContext(ctx)
	var tradePairs model.TradePairs
	res := tx.Where(`pair = ?`, pair).Find(&tradePairs)
	if res.Error != nil {
		db.log.Errorf("get tradePairs error %v", res.Error)
		return nil, res.Error
	}
	return &tradePairs, nil
}

// func (db *tradePairsConnection) UpdateTradePairs(ctx context.Context, tradePairs model.TradePairs) (*model.TradePairs, error) {
// 	tx := db.connection.WithContext(ctx)
// 	mdTradePairs := model.TradePairs{}
// 	res := tx.Model(&mdTradePairs).Where(`id = ?`, 1).Updates(tradePairs)
// 	if res.Error != nil {
// 		db.log.Error(res.Error)
// 		return nil, res.Error
// 	}
// 	return &mdTradePairs, nil
// }

func (db *tradePairsConnection) DeleteTradePairs(ctx context.Context) error {
	tx := db.connection.WithContext(ctx)
	res := tx.Exec("DELETE FROM trade_pairs")
	if res.Error != nil {
		db.log.Errorf("delete tradePairs error %v", res.Error)
		return res.Error
	}
	return nil
}
