package repository

import (
	"context"

	"github.com/dexthrottle/trfine/internal/model"
	"github.com/dexthrottle/trfine/pkg/logging"
	"gorm.io/gorm"
)

type CommissionBurnRepository interface {
	InsertCommissionBurn(ctx context.Context, a model.CommissionBurn) (*model.CommissionBurn, error)
	DeleteAllCommissionBurn(ctx context.Context, day string) ([]*model.CommissionBurn, error)
	GetAllCommissionBurn(ctx context.Context, day string) ([]*model.CommissionBurn, error)
}

type commissionBurnConnection struct {
	ctx        context.Context
	connection *gorm.DB
	log        logging.Logger
}

func NewCommissionBurnRepository(ctx context.Context, db *gorm.DB, log logging.Logger) CommissionBurnRepository {
	return &commissionBurnConnection{
		ctx:        ctx,
		connection: db,
		log:        log,
	}
}

func (db *commissionBurnConnection) InsertCommissionBurn(ctx context.Context, a model.CommissionBurn) (*model.CommissionBurn, error) {
	tx := db.connection.WithContext(ctx)
	res := tx.Save(&a)
	db.log.Error(res.Error)
	if res.Error != nil {
		return nil, res.Error
	}
	return &a, nil
}

func (db *commissionBurnConnection) GetAllCommissionBurn(ctx context.Context, day string) ([]*model.CommissionBurn, error) {
	tx := db.connection.WithContext(ctx)
	var commissionBurns []*model.CommissionBurn
	res := tx.Find(&commissionBurns).Where(`day = ?`, day)
	if res.Error != nil {
		db.log.Errorf("get all commission Burn error %v", res.Error)
		return nil, res.Error
	}
	return commissionBurns, nil
}

func (db *commissionBurnConnection) DeleteAllCommissionBurn(ctx context.Context, day string) ([]*model.CommissionBurn, error) {
	tx := db.connection.WithContext(ctx)
	var commissionBurns []*model.CommissionBurn
	res := tx.Delete(&commissionBurns).Where(`day = ?`, day)
	if res.Error != nil {
		db.log.Errorf("delete all commission Burn error %v", res.Error)
		return nil, res.Error
	}
	return commissionBurns, nil
}
