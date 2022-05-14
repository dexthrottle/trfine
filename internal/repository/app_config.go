package repository

import (
	"context"

	"github.com/dexthrottle/trfine/internal/model"
	"gorm.io/gorm"
)

type AppConfigRepository interface {
	InsertAppConfig(ctx context.Context, appCfg model.AppConfig) (*model.AppConfig, error)
	GetConfigData(ctx context.Context) (*model.AppConfig, error)
}

type appCfgConnection struct {
	ctx        context.Context
	connection *gorm.DB
}

func NewAppCfgRepository(ctx context.Context, db *gorm.DB) AppConfigRepository {
	return &appCfgConnection{
		ctx:        ctx,
		connection: db,
	}
}

func (db *appCfgConnection) InsertAppConfig(ctx context.Context, appCfg model.AppConfig) (*model.AppConfig, error) {
	tx := db.connection.WithContext(ctx)
	var mdAppConfig model.AppConfig
	_ = tx.Where(`"id" = ?`, 1).Find(&mdAppConfig)
	if mdAppConfig.ID == 0 {
		res := tx.Save(&appCfg)
		if res.Error != nil {
			return nil, res.Error
		}
	}
	return &mdAppConfig, nil
}

func (db *appCfgConnection) GetConfigData(ctx context.Context) (*model.AppConfig, error) {
	tx := db.connection.WithContext(ctx)
	var appCfg model.AppConfig
	res := tx.Find(&appCfg)
	if res.Error != nil {
		return nil, res.Error
	}
	return &appCfg, nil
}
