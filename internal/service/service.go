package service

import (
	"context"

	"github.com/dexthrottle/trfine/internal/dto"
	"github.com/dexthrottle/trfine/internal/model"
	"github.com/dexthrottle/trfine/internal/repository"
	"github.com/dexthrottle/trfine/pkg/logging"
)

type Post interface {
	Insert(ctx context.Context, b dto.PostCreateDTO) (*model.Post, error)
	Delete(ctx context.Context, b model.Post, userId int) error
	All(ctx context.Context, userTgId int) ([]*model.Post, error)
}

type User interface {
	Insert(ctx context.Context, user dto.CreateUserDTO) (*model.User, error)
	Profile(ctx context.Context, userID string) (*model.User, error)
	IsDuplicateUserTGID(ctx context.Context, tgID int) (bool, error)
	FindUserByTgUserId(ctx context.Context, userTgId int) (*model.User, error)
}

type Category interface {
	Insert(ctx context.Context, p dto.CreateCategoryDTO) (*model.Category, error)
	Delete(ctx context.Context, b model.Category, userId int) error
	All(ctx context.Context, userTgId int) ([]*model.Category, error)
}

type Service struct {
	Post
	User
	Category
}

func NewService(ctx context.Context, r repository.Repository, log logging.Logger) *Service {
	return &Service{
		Post:     NewPostService(ctx, r.Post, log),
		User:     NewUserService(ctx, r.User, log),
		Category: NewCategoryService(ctx, r.Category, log),
	}
}
