package repository

import (
	"context"

	"github.com/dexthrottle/trfine/internal/model"
	"github.com/dexthrottle/trfine/pkg/logging"
	"gorm.io/gorm"
)

type AveragePercentRepository interface {
	InsertAveragePercent(ctx context.Context, a model.AveragePercent) (*model.AveragePercent, error)
	DeleteAllAveragePercent(ctx context.Context, day string) error
	GetAllAveragePercent(ctx context.Context, day string) ([]*model.AveragePercent, error)
}

type averagePercentConnection struct {
	ctx        context.Context
	connection *gorm.DB
	log        logging.Logger
}

func NewAveragePercentRepository(ctx context.Context, db *gorm.DB, log logging.Logger) AveragePercentRepository {
	return &averagePercentConnection{
		ctx:        ctx,
		connection: db,
		log:        log,
	}
}

func (db *averagePercentConnection) InsertAveragePercent(ctx context.Context, a model.AveragePercent) (*model.AveragePercent, error) {
	tx := db.connection.WithContext(ctx)
	res := tx.Save(&a)
	db.log.Error(res.Error)
	if res.Error != nil {
		return nil, res.Error
	}
	return &a, nil
}

func (db *averagePercentConnection) GetAllAveragePercent(ctx context.Context, day string) ([]*model.AveragePercent, error) {
	tx := db.connection.WithContext(ctx)
	var averagePercents []*model.AveragePercent
	res := tx.Find(&averagePercents).Where(`day = ?`, day)
	if res.Error != nil {
		db.log.Errorf("get all average Percents error %v", res.Error)
		return nil, res.Error
	}
	return averagePercents, nil
}

func (db *averagePercentConnection) DeleteAllAveragePercent(ctx context.Context, day string) error {
	tx := db.connection.WithContext(ctx)
	var averagePercents []*model.AveragePercent
	res := tx.Delete(&averagePercents).Where(`day = ?`, day)
	if res.Error != nil {
		db.log.Errorf("delete all average Percents error %v", res.Error)
		return res.Error
	}
	return nil
}
