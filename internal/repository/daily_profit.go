package repository

import (
	"context"

	"github.com/dexthrottle/trfine/internal/model"
	"github.com/dexthrottle/trfine/pkg/logging"
	"gorm.io/gorm"
)

type DailyProfitRepository interface {
	InsertDailyProfit(ctx context.Context, a model.DailyProfit) (*model.DailyProfit, error)
	DeleteAllDailyProfit(ctx context.Context, day string) ([]*model.DailyProfit, error)
	GetAllDailyProfit(ctx context.Context, day string) ([]*model.DailyProfit, error)
}

type dailyProfitConnection struct {
	ctx        context.Context
	connection *gorm.DB
	log        logging.Logger
}

func NewDailyProfitRepository(ctx context.Context, db *gorm.DB, log logging.Logger) DailyProfitRepository {
	return &dailyProfitConnection{
		ctx:        ctx,
		connection: db,
		log:        log,
	}
}

func (db *dailyProfitConnection) InsertDailyProfit(ctx context.Context, a model.DailyProfit) (*model.DailyProfit, error) {
	tx := db.connection.WithContext(ctx)
	res := tx.Save(&a)
	db.log.Error(res.Error)
	if res.Error != nil {
		return nil, res.Error
	}
	return &a, nil
}

func (db *dailyProfitConnection) GetAllDailyProfit(ctx context.Context, day string) ([]*model.DailyProfit, error) {
	tx := db.connection.WithContext(ctx)
	var dailyProfits []*model.DailyProfit
	res := tx.Find(&dailyProfits).Where(`day = ?`, day)
	if res.Error != nil {
		db.log.Errorf("get all daily Profit error %v", res.Error)
		return nil, res.Error
	}
	return dailyProfits, nil
}

func (db *dailyProfitConnection) DeleteAllDailyProfit(ctx context.Context, day string) ([]*model.DailyProfit, error) {
	tx := db.connection.WithContext(ctx)
	var dailyProfits []*model.DailyProfit
	res := tx.Delete(&dailyProfits).Where(`day = ?`, day)
	if res.Error != nil {
		db.log.Errorf("delete all daily Profit error %v", res.Error)
		return nil, res.Error
	}
	return dailyProfits, nil
}
