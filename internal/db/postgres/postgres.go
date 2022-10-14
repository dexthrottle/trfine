package postgres

import (
	"fmt"

	"trfine/internal/config"
	"trfine/internal/entity"
	"trfine/pkg/logging"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(cfg *config.Config, log *logging.Logger) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.PostgreSQL.Username,
		cfg.PostgreSQL.Password,
		cfg.PostgreSQL.Host,
		cfg.PostgreSQL.Port,
		cfg.PostgreSQL.Database,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connection database: %v", err)
		return nil, err
	}

	if err := migrate(
		db,
		entity.AppConfig{},
		entity.AveragePercent{},
		entity.DailyProfit{},
		entity.Symbols{},
		entity.TradeInfo{},
		entity.TradePairs{},
		entity.TradeParams{},
		entity.TrailingOrders{},
		entity.WhiteList{},
		entity.CommissionBurn{},
	); err != nil {
		log.Errorln("Error migrate database")
		return nil, err
	}

	return db, nil
}

func migrate(db *gorm.DB, models ...interface{}) error {
	if err := db.AutoMigrate(models...); err != nil {
		return err
	}
	return nil
}
