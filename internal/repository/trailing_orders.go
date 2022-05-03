package repository

import (
	"context"

	"github.com/dexthrottle/trfine/internal/model"
	"github.com/dexthrottle/trfine/pkg/logging"
	"gorm.io/gorm"
)

type TrailingOrdersRepository interface {
	InsertTrailingOrders(ctx context.Context, a model.TrailingOrders) (*model.TrailingOrders, error)
	DeleteTrailingOrders(ctx context.Context, pair string) error
	GetTrailingOrders(ctx context.Context, pair string) (*[]model.TrailingOrders, error)
}

type trailingOrdersConnection struct {
	ctx        context.Context
	connection *gorm.DB
	log        logging.Logger
}

func NewTrailingOrdersRepository(ctx context.Context, db *gorm.DB, log logging.Logger) TrailingOrdersRepository {
	return &trailingOrdersConnection{
		ctx:        ctx,
		connection: db,
		log:        log,
	}
}

func (db *trailingOrdersConnection) InsertTrailingOrders(ctx context.Context, a model.TrailingOrders) (*model.TrailingOrders, error) {
	tx := db.connection.WithContext(ctx)
	res := tx.Save(&a)
	if res.Error != nil {
		return nil, res.Error
	}
	return &a, nil
}

func (db *trailingOrdersConnection) GetTrailingOrders(ctx context.Context, pair string) (*[]model.TrailingOrders, error) {
	tx := db.connection.WithContext(ctx)
	var trailingOrders *[]model.TrailingOrders
	res := tx.Where(`pair = ?`, pair).Find(&trailingOrders)
	if res.Error != nil {
		db.log.Errorf("get trade Trailing Orders %v", res.Error)
		return nil, res.Error
	}
	return trailingOrders, nil
}

func (db *trailingOrdersConnection) DeleteTrailingOrders(ctx context.Context, pair string) error {
	tx := db.connection.WithContext(ctx)
	var trailingOrders *[]model.TrailingOrders
	res := tx.Where(`pair = ?`, pair).Delete(&trailingOrders)
	if res.Error != nil {
		db.log.Errorf("delete trade Params error %v", res.Error)
		return res.Error
	}
	return nil
}
