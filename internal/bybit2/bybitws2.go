package bybit2

import (
	"context"
	"trfine/pkg/logging"

	bybit2 "github.com/hirokisan/bybit/v2"
)

type ByBitAPIWS2 interface {
	StartWebSocket()
}

type bybitws struct {
	bybitWS *bybit2.WebSocketClient
	log     logging.Logger
	ctx     context.Context
}

func NewByBitWS(ctx context.Context, log logging.Logger, bybitWS *bybit2.WebSocketClient) ByBitAPIWS2 {

	return &bybitws{
		ctx:     ctx,
		log:     log,
		bybitWS: bybitWS,
	}
}

// StartWebSocket - старт веб-сокетов (TODO: WebsocketExecutor ??)
func (b *bybitws) StartWebSocket() {
	b.bybitWS.Start(b.ctx, []bybit2.WebsocketExecutor{})
}
