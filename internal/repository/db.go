package repository

import (
	"github.com/dexthrottle/trfine/internal/model"
	"github.com/dexthrottle/trfine/pkg/logging"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func NewPostgresDB(log *logging.Logger) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("rp.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connection database: %v", err)
		return nil, err
	}

	err = migrations(db, log)
	if err != nil {
		return nil, err
	}
	log.Info("Migration Successfully")

	return db, nil
}

func migrations(db *gorm.DB, log *logging.Logger) error {
	err := db.AutoMigrate(&model.User{}, &model.AppConfig{})
	if err != nil {
		log.Errorf("Migrate error: %v", err)
		return err
	}
	return nil
}
