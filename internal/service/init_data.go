package service

import (
	"context"

	initdata "github.com/dexthrottle/trfine/internal/init_data"
	"github.com/dexthrottle/trfine/internal/model"
	"github.com/dexthrottle/trfine/internal/repository"
	"github.com/mashingan/smapping"
)

type InitDataService interface {
	InsertDataTradeParams(ctx context.Context) error
	InsertDataTradeInfo(ctx context.Context) error
}

type initDataService struct {
	ctx                context.Context
	initDataRepository repository.InitDataRepository
}

func NewInitDataService(
	ctx context.Context,
	initDataRepository repository.InitDataRepository,
) InitDataService {
	return &initDataService{
		ctx:                ctx,
		initDataRepository: initDataRepository,
	}
}

func (s *initDataService) InsertDataTradeParams(ctx context.Context) error {
	tradeParams := initdata.GetTradeParams()
	mdTradeParams := model.TradeParams{}
	err := smapping.FillStruct(&mdTradeParams, smapping.MapFields(&tradeParams))
	if err != nil {
		return err
	}
	_, err = s.initDataRepository.InsertDataTradeParams(ctx, mdTradeParams)
	if err != nil {
		return err
	}
	return nil
}

func (s *initDataService) InsertDataTradeInfo(ctx context.Context) error {
	tradeInfo := initdata.GetTradeInfo()
	mdTradeInfo := model.TradeInfo{}
	err := smapping.FillStruct(&mdTradeInfo, smapping.MapFields(&tradeInfo))
	if err != nil {
		return err
	}
	_, err = s.initDataRepository.InsertDataTradeInfo(ctx, mdTradeInfo)
	if err != nil {
		return err
	}
	return nil
}
