package postgres

import (
	"context"

	"trfine/internal/entity"
	"trfine/pkg/logging"

	"gorm.io/gorm"
)

type TrailingOrder interface {
	InsertTrailingOrders(ctx context.Context, a entity.TrailingOrders) (*entity.TrailingOrders, error)
	DeleteTrailingOrders(ctx context.Context, pair string) error
	GetTrailingOrders(ctx context.Context, pair string) (*[]entity.TrailingOrders, error)
}

type trailingOrders struct {
	ctx context.Context
	db  *gorm.DB
	log logging.Logger
}

func NewTrailingOrders(ctx context.Context, db *gorm.DB, log logging.Logger) TrailingOrder {
	return &trailingOrders{
		ctx: ctx,
		db:  db,
		log: log,
	}
}

func (db *trailingOrders) InsertTrailingOrders(ctx context.Context, a entity.TrailingOrders) (*entity.TrailingOrders, error) {
	tx := db.db.WithContext(ctx)
	res := tx.Save(&a)
	if res.Error != nil {
		return nil, res.Error
	}
	return &a, nil
}

func (db *trailingOrders) GetTrailingOrders(ctx context.Context, pair string) (*[]entity.TrailingOrders, error) {
	tx := db.db.WithContext(ctx)
	var trailingOrders *[]entity.TrailingOrders
	res := tx.Where(`pair = ?`, pair).Find(&trailingOrders)
	if res.Error != nil {
		db.log.Errorf("get trade Trailing Orders %v", res.Error)
		return nil, res.Error
	}
	return trailingOrders, nil
}

func (db *trailingOrders) DeleteTrailingOrders(ctx context.Context, pair string) error {
	tx := db.db.WithContext(ctx)
	var trailingOrders *[]entity.TrailingOrders
	res := tx.Where(`pair = ?`, pair).Delete(&trailingOrders)
	if res.Error != nil {
		db.log.Errorf("delete trade Params error %v", res.Error)
		return res.Error
	}
	return nil
}
