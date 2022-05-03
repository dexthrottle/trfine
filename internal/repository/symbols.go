package repository

import (
	"context"
	"fmt"

	"github.com/dexthrottle/trfine/internal/model"
	"github.com/dexthrottle/trfine/pkg/logging"
	"gorm.io/gorm"
)

type SymbolsRepository interface {
	InsertSymbols(ctx context.Context, a model.Symbols) (*model.Symbols, error)
	DeleteSymbols(ctx context.Context, pair string) (*model.Symbols, error)
	GetAllSymbols(ctx context.Context, pair string) ([]*model.Symbols, error)
	DeleteAllSymbols(ctx context.Context) error
	UpdateSymbols(ctx context.Context, symbols model.Symbols, pair string) (*model.Symbols, error)
}

type symbolsConnection struct {
	ctx        context.Context
	connection *gorm.DB
	log        logging.Logger
}

func NewSymbolsRepository(ctx context.Context, db *gorm.DB, log logging.Logger) SymbolsRepository {
	return &symbolsConnection{
		ctx:        ctx,
		connection: db,
		log:        log,
	}
}

func (db *symbolsConnection) InsertSymbols(ctx context.Context, a model.Symbols) (*model.Symbols, error) {
	tx := db.connection.WithContext(ctx)
	res := tx.Save(&a)

	if res.Error != nil {
		return nil, res.Error
	}
	return &a, nil
}

func (db *symbolsConnection) GetAllSymbols(ctx context.Context, pair string) ([]*model.Symbols, error) {
	tx := db.connection.WithContext(ctx)
	var symbols []*model.Symbols
	res := tx.Where(`pair = ?`, pair).Find(&symbols)
	if res.Error != nil {
		db.log.Errorf("get all symbols error %v", res.Error)
		return nil, res.Error
	}
	return symbols, nil
}

func (db *symbolsConnection) UpdateSymbols(ctx context.Context, symbols model.Symbols, pair string) (*model.Symbols, error) {
	tx := db.connection.WithContext(ctx)
	mdSymbols := model.Symbols{}
	res := tx.Model(&mdSymbols).Where(`pair = ?`, pair).Updates(symbols)
	if res.Error != nil {
		db.log.Error(res.Error)
		return nil, res.Error
	}
	fmt.Printf("%+v", mdSymbols)
	return &mdSymbols, nil
}

func (db *symbolsConnection) DeleteSymbols(ctx context.Context, pair string) (*model.Symbols, error) {
	tx := db.connection.WithContext(ctx)
	var symbols *model.Symbols
	res := tx.Delete(&symbols).Where(`pair = ?`, pair)
	if res.Error != nil {
		db.log.Errorf("delete symbols error %v", res.Error)
		return nil, res.Error
	}
	return symbols, nil
}

func (db *symbolsConnection) DeleteAllSymbols(ctx context.Context) error {
	tx := db.connection.WithContext(ctx)
	var symbols *[]model.Symbols
	res := tx.Delete(&symbols)
	if res.Error != nil {
		db.log.Errorf("delete all symbols error %v", res.Error)
		return res.Error
	}
	return nil
}
