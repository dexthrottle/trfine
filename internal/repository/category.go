package repository

import (
	"context"

	"github.com/dexthrottle/trfine/internal/model"
	"github.com/dexthrottle/trfine/pkg/logging"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	InsertCategory(ctx context.Context, c model.Category) (*model.Category, error)
	DeleteCategory(ctx context.Context, category model.Category, userId int) error
	AllCategory(ctx context.Context, userTgId int) ([]*model.Category, error)
}

type categoryConnection struct {
	ctx        context.Context
	connection *gorm.DB
	log        logging.Logger
}

func NewCategoryRepository(
	ctx context.Context,
	dbConn *gorm.DB,
	log logging.Logger,
) CategoryRepository {
	return &categoryConnection{
		ctx:        ctx,
		connection: dbConn,
		log:        log,
	}
}

func (db *categoryConnection) InsertCategory(
	ctx context.Context,
	category model.Category,
) (*model.Category, error) {

	tx := db.connection.WithContext(ctx)
	tx.Save(&category)
	res := tx.Joins("User").Find(&category)
	if res.Error != nil {
		db.log.Errorf("insert category error: %v", res.Error)
		return nil, res.Error
	}
	return &category, nil
}

// Показ категорий
func (db *categoryConnection) AllCategory(
	ctx context.Context,
	userTgId int,
) ([]*model.Category, error) {

	tx := db.connection.WithContext(ctx)
	var categories []*model.Category

	res := tx.Preload("User").Joins("User").Where(
		`"user_tg_id" = ?`,
		userTgId,
	).Find(&categories)

	if res.Error != nil {
		db.log.Errorf("get all categories error %v", res.Error)
		return nil, res.Error
	}
	return categories, nil
}

// Удаление категории
func (db *categoryConnection) DeleteCategory(
	ctx context.Context,
	category model.Category,
	userId int,
) error {
	db.log.Printf("из репозитория userID = %d", userId)
	tx := db.connection.WithContext(ctx)
	res := tx.Select("Posts").Delete(&category).Where(`user_id = ?`, userId)
	db.log.Printf("%+v", category)
	if res.Error != nil {
		db.log.Errorf("delete category error %v", res.Error)
		return res.Error
	}
	return nil
}
