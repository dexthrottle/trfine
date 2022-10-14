package postgres

import (
	"context"

	"trfine/internal/entity"
	"trfine/pkg/logging"

	"gorm.io/gorm"
)

type CommissionBurn interface {
	InsertCommissionBurn(ctx context.Context, a entity.CommissionBurn) (*entity.CommissionBurn, error)
	DeleteAllCommissionBurn(ctx context.Context, day string) error
	GetAllCommissionBurn(ctx context.Context, day string) ([]*entity.CommissionBurn, error)
}

type commissionBurn struct {
	ctx context.Context
	db  *gorm.DB
	log logging.Logger
}

func NewCommissionBurn(ctx context.Context, db *gorm.DB, log logging.Logger) CommissionBurn {
	return &commissionBurn{
		ctx: ctx,
		db:  db,
		log: log,
	}
}

func (db *commissionBurn) InsertCommissionBurn(ctx context.Context, a entity.CommissionBurn) (*entity.CommissionBurn, error) {
	tx := db.db.WithContext(ctx)
	res := tx.Save(&a)
	db.log.Error(res.Error)
	if res.Error != nil {
		return nil, res.Error
	}
	return &a, nil
}

func (db *commissionBurn) GetAllCommissionBurn(ctx context.Context, day string) ([]*entity.CommissionBurn, error) {
	tx := db.db.WithContext(ctx)
	var commissionBurns []*entity.CommissionBurn
	res := tx.Find(&commissionBurns).Where(`day = ?`, day)
	if res.Error != nil {
		db.log.Errorf("get all commission Burn error %v", res.Error)
		return nil, res.Error
	}
	return commissionBurns, nil
}

func (db *commissionBurn) DeleteAllCommissionBurn(ctx context.Context, day string) error {
	tx := db.db.WithContext(ctx)
	var commissionBurns []*entity.CommissionBurn
	res := tx.Delete(&commissionBurns).Where(`day = ?`, day)
	if res.Error != nil {
		db.log.Errorf("delete all commission Burn error %v", res.Error)
		return res.Error
	}
	return nil
}
