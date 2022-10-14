package postgres

import (
	"context"

	"trfine/internal/entity"
	"trfine/pkg/logging"

	"gorm.io/gorm"
)

type TradePairs interface {
	InsertTradePairs(ctx context.Context, a entity.TradePairs) (*entity.TradePairs, error)
	DeleteTradePairs(ctx context.Context) error
	GetTradePairs(ctx context.Context, pair string) (*entity.TradePairs, error)
	// UpdateTradePairs(ctx context.Context, tradePairs entity.TradePairs) (*entity.TradePairs, error)
}

type tradePairs struct {
	ctx        context.Context
	connection *gorm.DB
	log        logging.Logger
}

func NewTradePairs(ctx context.Context, db *gorm.DB, log logging.Logger) TradePairs {
	return &tradePairs{
		ctx:        ctx,
		connection: db,
		log:        log,
	}
}

func (db *tradePairs) InsertTradePairs(ctx context.Context, a entity.TradePairs) (*entity.TradePairs, error) {
	tx := db.connection.WithContext(ctx)
	res := tx.Save(&a)
	if res.Error != nil {
		return nil, res.Error
	}
	return &a, nil
}

func (db *tradePairs) GetTradePairs(ctx context.Context, pair string) (*entity.TradePairs, error) {
	tx := db.connection.WithContext(ctx)
	var tradePairs entity.TradePairs
	res := tx.Where(`pair = ?`, pair).Find(&tradePairs)
	if res.Error != nil {
		db.log.Errorf("get tradePairs error %v", res.Error)
		return nil, res.Error
	}
	return &tradePairs, nil
}

// func (db *tradePairs) UpdateTradePairs(ctx context.Context, tradePairs entity.TradePairs) (*entity.TradePairs, error) {
// 	tx := db.connection.WithContext(ctx)
// 	mdTradePairs := entity.TradePairs{}
// 	res := tx.entity(&mdTradePairs).Where(`id = ?`, 1).Updates(tradePairs)
// 	if res.Error != nil {
// 		db.log.Error(res.Error)
// 		return nil, res.Error
// 	}
// 	return &mdTradePairs, nil
// }

func (db *tradePairs) DeleteTradePairs(ctx context.Context) error {
	tx := db.connection.WithContext(ctx)
	res := tx.Exec("DELETE FROM trade_pairs")
	if res.Error != nil {
		db.log.Errorf("delete tradePairs error %v", res.Error)
		return res.Error
	}
	return nil
}
