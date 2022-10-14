package bybit

// import (
// 	"github.com/dexthrottle/trfine/internal/service"
// 	"github.com/dexthrottle/trfine/pkg/bybitapi/ws"
// 	"github.com/dexthrottle/trfine/pkg/logging"
// )

// type ByBitAPIWS interface {
// 	GetPositions() error
// }

// type bybitws struct {
// 	bybitWS  *ws.ByBitWS
// 	log      logging.Logger
// 	services *service.Service
// }

// func NewByBitWS(log logging.Logger, bybitWS *ws.ByBitWS, services *service.Service) ByBitAPIWS {

// 	return &bybitws{
// 		log:      log,
// 		bybitWS:  bybitWS,
// 		services: services,
// 	}
// }

// func (b *bybitws) StartWebSocket() error {
// 	err := b.bybitWS.Start()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (b *bybitws) GetPositions() error {
// 	return nil
// }
