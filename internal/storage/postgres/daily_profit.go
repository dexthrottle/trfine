package postgres

import (
	"context"

	"trfine/internal/entity"
	"trfine/pkg/logging"

	"gorm.io/gorm"
)

type DailyProfit interface {
	InsertDailyProfit(ctx context.Context, a entity.DailyProfit) (*entity.DailyProfit, error)
	DeleteAllDailyProfit(ctx context.Context, day string) error
	GetAllDailyProfit(ctx context.Context, day string) ([]*entity.DailyProfit, error)
}

type dailyProfit struct {
	ctx context.Context
	db  *gorm.DB
	log logging.Logger
}

func NewDailyProfit(ctx context.Context, db *gorm.DB, log logging.Logger) DailyProfit {
	return &dailyProfit{
		ctx: ctx,
		db:  db,
		log: log,
	}
}

func (db *dailyProfit) InsertDailyProfit(ctx context.Context, a entity.DailyProfit) (*entity.DailyProfit, error) {
	tx := db.db.WithContext(ctx)
	res := tx.Save(&a)
	db.log.Error(res.Error)
	if res.Error != nil {
		return nil, res.Error
	}
	return &a, nil
}

func (db *dailyProfit) GetAllDailyProfit(ctx context.Context, day string) ([]*entity.DailyProfit, error) {
	tx := db.db.WithContext(ctx)
	var dailyProfits []*entity.DailyProfit
	res := tx.Find(&dailyProfits).Where(`day = ?`, day)
	if res.Error != nil {
		db.log.Errorf("get all daily Profit error %v", res.Error)
		return nil, res.Error
	}
	return dailyProfits, nil
}

func (db *dailyProfit) DeleteAllDailyProfit(ctx context.Context, day string) error {
	tx := db.db.WithContext(ctx)
	var dailyProfits []*entity.DailyProfit
	res := tx.Delete(&dailyProfits).Where(`day = ?`, day)
	if res.Error != nil {
		db.log.Errorf("delete all daily Profit error %v", res.Error)
		return res.Error
	}
	return nil
}
