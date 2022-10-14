package postgres

import (
	"context"

	"trfine/internal/entity"
	"trfine/pkg/logging"

	"gorm.io/gorm"
)

type Symbols interface {
	InsertSymbols(ctx context.Context, a entity.Symbols) (*entity.Symbols, error)
	DeleteSymbols(ctx context.Context, pair string) error
	GetAllSymbols(ctx context.Context, pair string) ([]*entity.Symbols, error)
	DeleteAllSymbols(ctx context.Context) error
	UpdateSymbols(ctx context.Context, symbols entity.Symbols, pair string) (*entity.Symbols, error)
}

type symbols struct {
	ctx        context.Context
	connection *gorm.DB
	log        logging.Logger
}

func NewSymbols(ctx context.Context, db *gorm.DB, log logging.Logger) Symbols {
	return &symbols{
		ctx:        ctx,
		connection: db,
		log:        log,
	}
}

func (db *symbols) InsertSymbols(ctx context.Context, a entity.Symbols) (*entity.Symbols, error) {
	tx := db.connection.WithContext(ctx)
	res := tx.Save(&a)

	if res.Error != nil {
		return nil, res.Error
	}
	return &a, nil
}

func (db *symbols) GetAllSymbols(ctx context.Context, pair string) ([]*entity.Symbols, error) {
	tx := db.connection.WithContext(ctx)
	var symbols []*entity.Symbols
	res := tx.Where(`pair = ?`, pair).Find(&symbols)
	if res.Error != nil {
		db.log.Errorf("get all symbols error %v", res.Error)
		return nil, res.Error
	}
	return symbols, nil
}

func (db *symbols) UpdateSymbols(ctx context.Context, symbols entity.Symbols, pair string) (*entity.Symbols, error) {
	tx := db.connection.WithContext(ctx)
	mdSymbols := entity.Symbols{}
	res := tx.Model(&mdSymbols).Where(`pair = ?`, pair).Updates(symbols)
	if res.Error != nil {
		db.log.Error(res.Error)
		return nil, res.Error
	}
	return &mdSymbols, nil
}

func (db *symbols) DeleteSymbols(ctx context.Context, pair string) error {
	tx := db.connection.WithContext(ctx)
	var symbols *entity.Symbols
	res := tx.Delete(&symbols).Where(`pair = ?`, pair)
	if res.Error != nil {
		db.log.Errorf("delete symbols error %v", res.Error)
		return res.Error
	}
	return nil
}

func (db *symbols) DeleteAllSymbols(ctx context.Context) error {
	tx := db.connection.WithContext(ctx)
	var symbols *[]entity.Symbols
	res := tx.Delete(&symbols)
	if res.Error != nil {
		db.log.Errorf("delete all symbols error %v", res.Error)
		return res.Error
	}
	return nil
}
