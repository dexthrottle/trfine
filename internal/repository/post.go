package repository

import (
	"context"

	"github.com/dexthrottle/trfine/internal/model"
	"github.com/dexthrottle/trfine/pkg/logging"
	"gorm.io/gorm"
)

type PostRepository interface {
	InsertPost(ctx context.Context, b model.Post) (*model.Post, error)
	AllPost(ctx context.Context, userTgId int) ([]*model.Post, error)
	DeletePost(ctx context.Context, post model.Post, userId int) error
}

type postConnection struct {
	ctx        context.Context
	connection *gorm.DB
	log        logging.Logger
}

func NewPostRepository(ctx context.Context, dbConn *gorm.DB, log logging.Logger) PostRepository {
	return &postConnection{
		ctx:        ctx,
		connection: dbConn,
		log:        log,
	}
}

// Добавление Post
func (db *postConnection) InsertPost(ctx context.Context, post model.Post) (*model.Post, error) {
	tx := db.connection.WithContext(ctx)
	tx.Save(&post)
	res := tx.Preload("Category").Find(&post)
	if res.Error != nil {
		db.log.Errorf("insert post error: %v", res.Error)
		return nil, res.Error
	}
	return &post, nil
}

// Все посты
func (db *postConnection) AllPost(ctx context.Context, userTgId int) ([]*model.Post, error) {
	tx := db.connection.WithContext(ctx)
	var posts []*model.Post
	res := tx.Preload("User").Joins("User").Where(
		`"user_tg_id" = ?`,
		userTgId,
	).Preload("Category").Find(&posts)
	if res.Error != nil {
		db.log.Errorf("get all posts error %v", res.Error)
		return nil, res.Error
	}
	return posts, nil
}

// Удаление поста
func (db *postConnection) DeletePost(ctx context.Context, post model.Post, userId int) error {
	tx := db.connection.WithContext(ctx)
	res := tx.Delete(&post).Where(`"user_id = ?`, userId)
	if res.Error != nil {
		db.log.Errorf("delete post error %v", res.Error)
		return res.Error
	}
	db.log.Infof("%+v", post)
	return nil
}
