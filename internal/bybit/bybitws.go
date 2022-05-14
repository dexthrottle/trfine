package bybit

import (
	"github.com/dexthrottle/trfine/internal/service"
	"github.com/dexthrottle/trfine/pkg/logging"
	"github.com/frankrap/bybit-api/ws"
)

type ByBitAPIWS interface {
	GetPositions() error
}

type bybitws struct {
	bybitWS  *ws.ByBitWS
	log      logging.Logger
	services *service.Service
}

func NewByBitWS(log logging.Logger, bybitWS *ws.ByBitWS, services *service.Service) ByBitAPIWS {

	return &bybitws{
		log:      log,
		bybitWS:  bybitWS,
		services: services,
	}
}

func (b *bybitws) GetPositions() error
