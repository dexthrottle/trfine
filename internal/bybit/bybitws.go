package bybit

import (
	"trfine/pkg/bybitapi/ws"
	"trfine/pkg/logging"
)

type ByBitAPIWS interface {
	GetPositions() error
}

type bybitws struct {
	bybitWS *ws.ByBitWS
	log     logging.Logger
}

func NewByBitWS(log logging.Logger, bybitWS *ws.ByBitWS) ByBitAPIWS {

	return &bybitws{
		log:     log,
		bybitWS: bybitWS,
	}
}

func (b *bybitws) StartWebSocket() error {
	err := b.bybitWS.Start()
	if err != nil {
		return err
	}
	return nil
}

func (b *bybitws) GetPositions() error {
	return nil
}
