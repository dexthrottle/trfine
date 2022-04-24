package repository

import (
	"context"

	"github.com/dexthrottle/trfine/internal/model"
	"github.com/dexthrottle/trfine/pkg/logging"
	"gorm.io/gorm"
)

type User interface {
	InsertUser(ctx context.Context, user model.User) (*model.User, error)
	ProfileUser(ctx context.Context, userID string) (*model.User, error)
	IsDuplicateUserTGID(ctx context.Context, tgID int) (bool, error)
	FindUserByTgUserId(ctx context.Context, userTgId int) (*model.User, error)
}

type AppConfig interface {
	InsertAppConfig(ctx context.Context, appCfg model.AppConfig) (*model.AppConfig, error)
	CheckConfigData(ctx context.Context) (*model.AppConfig, error)
}

type Repository struct {
	User
	AppConfig
}

func NewRepository(ctx context.Context, db *gorm.DB, log logging.Logger) *Repository {
	return &Repository{
		User:      NewUserRepository(ctx, db, log),
		AppConfig: NewAppCfgRepository(ctx, db),
	}
}
