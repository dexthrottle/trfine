package repository

import (
	"fmt"

	"github.com/dexthrottle/trfine/internal/model"
	"github.com/dexthrottle/trfine/pkg/logging"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func NewDB(log *logging.Logger, dbName string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s.db", dbName)), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connection database: %v", err)
		return nil, err
	}

	err = migrations(
		db,
		log,
		model.AppConfig{},
		model.AveragePercent{},
		model.DailyProfit{},
		model.Symbols{},
		model.TradeInfo{},
		model.TradePairs{},
		model.TradeParams{},
		model.TradeParamsList{},
		model.TrailingOrders{},
		model.User{},
		model.WhiteList{},
		// model.BNBBurn{}, TODO: Шляпа
	)
	if err != nil {
		return nil, err
	}
	log.Info("Migration Successfully")

	return db, nil
}

func migrations(db *gorm.DB, log *logging.Logger, models ...interface{}) error {
	err := db.AutoMigrate(models...)
	if err != nil {
		log.Errorf("Migrate error: %v", err)
		return err
	}
	return nil
}
