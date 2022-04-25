package service

import (
	"context"

	"github.com/dexthrottle/trfine/internal/dto"
	"github.com/dexthrottle/trfine/internal/model"
	"github.com/dexthrottle/trfine/internal/repository"
	"github.com/mashingan/smapping"
)

type AppConfigService interface {
	InsertAppConfig(ctx context.Context, appCfg dto.AppConfigDTO) (*model.AppConfig, error)
	GetConfigData(ctx context.Context) (*model.AppConfig, error)
}

type appCfgService struct {
	ctx              context.Context
	appCfgRepository repository.AppConfigRepository
}

func NewAppCfgService(
	ctx context.Context,
	appCfgRepository repository.AppConfigRepository,
) AppConfigService {
	return &appCfgService{
		ctx:              ctx,
		appCfgRepository: appCfgRepository,
	}
}

func (s *appCfgService) InsertAppConfig(ctx context.Context, appCfgDTO dto.AppConfigDTO) (*model.AppConfig, error) {
	appCfgToCreate := model.AppConfig{}
	err := smapping.FillStruct(&appCfgToCreate, smapping.MapFields(&appCfgDTO))
	if err != nil {
		return nil, err
	}

	updatedAppCfg, err := s.appCfgRepository.InsertAppConfig(ctx, appCfgToCreate)
	if err != nil {
		return nil, err
	}
	return updatedAppCfg, nil
}

func (s *appCfgService) GetConfigData(ctx context.Context) (*model.AppConfig, error) {
	mAppCfg, err := s.appCfgRepository.GetConfigData(ctx)
	if err != nil {
		return nil, err
	}
	return mAppCfg, nil
}
