package bybit

import (
	"github.com/dexthrottle/trfine/internal/service"
	"github.com/dexthrottle/trfine/pkg/logging"
	"github.com/frankrap/bybit-api/rest"
)

type ByBitAPIRest interface {
	GetPositions() error
}

type bybit struct {
	bybitRest *rest.ByBit
	log       logging.Logger
	services  *service.Service
}

func NewByBit(log logging.Logger, bybitRest *rest.ByBit, services *service.Service) ByBitAPIRest {

	return &bybit{
		log:       log,
		bybitRest: bybitRest,
		services:  services,
	}
}

func (b *bybit) GetPositions() error {
	_, _, positions, err := b.bybitRest.GetPositions()
	if err != nil {
		return err
	}
	b.log.Printf("positions: %#v", positions)
	return nil
}
