package bybit

import (
	"github.com/dexthrottle/trfine/internal/service"
	"github.com/dexthrottle/trfine/pkg/bybitapi/rest"
	"github.com/dexthrottle/trfine/pkg/logging"
)

type ByBitAPIRest interface {
	GetWalletBalanceSpot() (*rest.ResultSpot, []byte, error)
	GetOrderSpot(orderId, orderLinkId string) (*rest.ResultOrderSpot, []byte, error)
	CreateOrderSpot(symbol, side, _type string, qty float64, price float64, timeInForce, orderLinkId string) (*rest.ResultCreateDeleteOrderSpot, []byte, error)
	DeleteOrderSpot(orderId, orderLinkId string) (*rest.ResultCreateDeleteOrderSpot, []byte, error)
	GetOpenOrdersSpot(symbol, orderId string, limit int) (*rest.OrdersSpot, []byte, error)
	GetMyTradesSpot(symbol, orderId string, limit, fromTicketId, toTicketId, startTime, endTime int64) (*rest.MyTradesSpot, []byte, error)
	GetHistoryOrdersSpot(symbol, orderId string, limit int, startTime, endTime int64) (*rest.OrdersSpot, []byte, error)
	GetUserApiKey() (*rest.UserApiKey, []byte, error)
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

func (b *bybit) GetWalletBalanceSpot() (*rest.ResultSpot, []byte, error) {
	_, resp, balance, err := b.bybitRest.GetWalletBalanceSpot()
	if err != nil {
		return nil, nil, err
	}
	return balance, resp, nil
}

func (b *bybit) GetOrderSpot(orderId, orderLinkId string) (*rest.ResultOrderSpot, []byte, error) {
	_, resp, activeOrders, err := b.bybitRest.GetOrderSpot(orderId, orderLinkId)
	if err != nil {
		return nil, nil, err
	}
	return activeOrders, resp, nil
}

func (b *bybit) CreateOrderSpot(symbol, side, _type string, qty float64, price float64, timeInForce, orderLinkId string) (*rest.ResultCreateDeleteOrderSpot, []byte, error) {
	_, resp, createdOrder, err := b.bybitRest.CreateOrderSpot(
		symbol,
		side,
		_type,
		qty,
		price,
		timeInForce,
		orderLinkId,
	)
	if err != nil {
		return nil, nil, err
	}
	return createdOrder, resp, nil
}

func (b *bybit) DeleteOrderSpot(orderId, orderLinkId string) (*rest.ResultCreateDeleteOrderSpot, []byte, error) {

	_, resp, deletedOrder, err := b.bybitRest.DeleteOrderSpot(
		orderId, orderLinkId,
	)
	if err != nil {
		return nil, nil, err
	}
	return deletedOrder, resp, nil
}

func (b *bybit) GetOpenOrdersSpot(symbol, orderId string, limit int) (*rest.OrdersSpot, []byte, error) {
	_, resp, openOrders, err := b.bybitRest.GetOpenOrdersSpot(symbol, orderId, limit)
	if err != nil {
		return nil, nil, err
	}
	return openOrders, resp, nil
}

func (b *bybit) GetHistoryOrdersSpot(symbol, orderId string, limit int, startTime, endTime int64) (*rest.OrdersSpot, []byte, error) {
	_, resp, openOrders, err := b.bybitRest.GetHistoryOrdersSpot(symbol, orderId, limit, startTime, endTime)
	if err != nil {
		return nil, nil, err
	}
	return openOrders, resp, nil
}

func (b *bybit) GetMyTradesSpot(symbol, orderId string, limit, fromTicketId, toTicketId, startTime, endTime int64) (*rest.MyTradesSpot, []byte, error) {
	_, resp, myTradesSpot, err := b.bybitRest.GetMyTradesSpot(symbol, orderId, limit, fromTicketId, toTicketId, startTime, endTime)
	if err != nil {
		return nil, nil, err
	}
	return myTradesSpot, resp, nil
}

func (b *bybit) GetUserApiKey() (*rest.UserApiKey, []byte, error) {
	_, resp, userApi, err := b.bybitRest.GetUserApiKey()
	if err != nil {
		return nil, nil, err
	}
	return userApi, resp, nil
}
