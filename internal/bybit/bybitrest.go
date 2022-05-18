package bybit

import (
	"github.com/dexthrottle/trfine/internal/service"
	"github.com/dexthrottle/trfine/pkg/bybitapi/rest"
	"github.com/dexthrottle/trfine/pkg/logging"
)

type ByBitAPIRest interface {
	GetWalletBalance() (*rest.Balance, []byte, error)
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

func (b *bybit) GetWalletBalance() (*rest.Balance, []byte, error) {
	_, resp, balance, err := b.bybitRest.GetWalletBalance("BTC")
	if err != nil {
		return nil, nil, err
	}
	return &balance, resp, nil
}
