package repository

import (
	"context"

	"github.com/dexthrottle/trfine/internal/model"
	"gorm.io/gorm"
)

type AppConfigRepository interface {
	InsertAppConfig(ctx context.Context, appCfg model.AppConfig) (*model.AppConfig, error)
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
	res := tx.Save(&appCfg)
	if res.Error != nil {
		return nil, res.Error
	}
	return &appCfg, nil
}
