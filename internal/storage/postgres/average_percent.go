package postgres

import (
	"context"

	"trfine/internal/entity"
	"trfine/pkg/logging"

	"gorm.io/gorm"
)

type AveragePercent interface {
	InsertAveragePercent(ctx context.Context, a entity.AveragePercent) (*entity.AveragePercent, error)
	DeleteAllAveragePercent(ctx context.Context, day string) error
	GetAllAveragePercent(ctx context.Context, day string) ([]*entity.AveragePercent, error)
}

type averagePercent struct {
	ctx context.Context
	db  *gorm.DB
	log logging.Logger
}

func NewAveragePercent(ctx context.Context, db *gorm.DB, log logging.Logger) AveragePercent {
	return &averagePercent{
		ctx: ctx,
		db:  db,
		log: log,
	}
}

func (db *averagePercent) InsertAveragePercent(ctx context.Context, a entity.AveragePercent) (*entity.AveragePercent, error) {
	tx := db.db.WithContext(ctx)
	res := tx.Save(&a)
	db.log.Error(res.Error)
	if res.Error != nil {
		return nil, res.Error
	}
	return &a, nil
}

func (db *averagePercent) GetAllAveragePercent(ctx context.Context, day string) ([]*entity.AveragePercent, error) {
	tx := db.db.WithContext(ctx)
	var averagePercents []*entity.AveragePercent
	res := tx.Find(&averagePercents).Where(`day = ?`, day)
	if res.Error != nil {
		db.log.Errorf("get all average Percents error %v", res.Error)
		return nil, res.Error
	}
	return averagePercents, nil
}

func (db *averagePercent) DeleteAllAveragePercent(ctx context.Context, day string) error {
	tx := db.db.WithContext(ctx)
	var averagePercents []*entity.AveragePercent
	res := tx.Delete(&averagePercents).Where(`day = ?`, day)
	if res.Error != nil {
		db.log.Errorf("delete all average Percents error %v", res.Error)
		return res.Error
	}
	return nil
}
