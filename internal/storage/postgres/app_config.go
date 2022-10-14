package postgres

import (
	"context"

	"trfine/internal/entity"

	"gorm.io/gorm"
)

type AppConfig interface {
	InsertAppConfig(ctx context.Context, appCfg entity.AppConfig) (*entity.AppConfig, error)
	GetConfigData(ctx context.Context) (*entity.AppConfig, error)
}

type appCfg struct {
	ctx context.Context
	db  *gorm.DB
}

func NewAppCfg(ctx context.Context, db *gorm.DB) AppConfig {
	return &appCfg{
		ctx: ctx,
		db:  db,
	}
}

func (db *appCfg) InsertAppConfig(ctx context.Context, appCfg entity.AppConfig) (*entity.AppConfig, error) {
	tx := db.db.WithContext(ctx)
	var mdAppConfig entity.AppConfig
	_ = tx.Where(`"id" = ?`, 1).Find(&mdAppConfig)
	if mdAppConfig.ID == 0 {
		res := tx.Save(&appCfg)
		if res.Error != nil {
			return nil, res.Error
		}
	}
	return &mdAppConfig, nil
}

func (db *appCfg) GetConfigData(ctx context.Context) (*entity.AppConfig, error) {
	tx := db.db.WithContext(ctx)
	var appCfg entity.AppConfig
	res := tx.Find(&appCfg)
	if res.Error != nil {
		return nil, res.Error
	}
	return &appCfg, nil
}
