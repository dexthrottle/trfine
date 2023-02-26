package bybit2

import (
	"trfine/pkg/logging"

	bybit2 "github.com/hirokisan/bybit/v2"
)

type ByBitAPIRest2 interface {
	SpotGetWalletBalance() (*bybit2.SpotGetWalletBalanceResponse, error)
	SpotGetOrder(orderId, orderLinkId string) (*bybit2.SpotGetOrderResponse, error)
	SpotCreateOrder(symbol, side, order_type string, qty float64, price float64, timeInForce, orderLinkId string) (*bybit2.SpotPostOrderResponse, error)
	SpotDeleteOrder(orderId, orderLinkId string) (*bybit2.SpotDeleteOrderResponse, error)
	SpotGetOpenOrders(symbol, orderId string, limit int) (*bybit2.SpotOpenOrdersResponse, error)
	SpotAPIKeyInfo(symbol, orderId string, limit int) (*bybit2.V5APIKeyResponse, error)
}

type bybit struct {
	bybitRest *bybit2.Client
	log       logging.Logger
}

func NewByBit2(log logging.Logger, bybitRest *bybit2.Client) ByBitAPIRest2 {

	return &bybit{
		log:       log,
		bybitRest: bybitRest,
	}
}

// SpotGetWalletBalance - получение баланса
func (b *bybit) SpotGetWalletBalance() (*bybit2.SpotGetWalletBalanceResponse, error) {
	res, err := b.bybitRest.Spot().V1().SpotGetWalletBalance()
	if err != nil {
		return nil, err
	}
	return res, nil

}

// SpotGetOrder - получение ордера
func (b *bybit) SpotGetOrder(orderId, orderLinkId string) (*bybit2.SpotGetOrderResponse, error) {
	res, err := b.bybitRest.Spot().V1().SpotGetOrder(bybit2.SpotGetOrderParam{
		OrderID:     &orderId,
		OrderLinkID: &orderLinkId,
	})
	if err != nil {
		return nil, err
	}
	return res, nil

}

// SpotCreateOrder - создание ордера (TODO: SpotPostOrder - создание?)
func (b *bybit) SpotCreateOrder(
	symbol,
	side,
	order_type string,
	qty float64,
	price float64,
	timeInForce, orderLinkId string,
) (*bybit2.SpotPostOrderResponse, error) {
	res, err := b.bybitRest.Spot().V1().SpotPostOrder(bybit2.SpotPostOrderParam{
		Symbol:      bybit2.SymbolSpot(symbol),
		Qty:         qty,
		Side:        bybit2.Side(side),
		Type:        bybit2.OrderTypeSpot(order_type),
		TimeInForce: (*bybit2.TimeInForceSpot)(&timeInForce),
		Price:       &price,
		OrderLinkID: &orderLinkId,
	})
	if err != nil {
		return nil, err
	}
	return res, nil

}

// SpotDeleteOrder - удаление ордера
func (b *bybit) SpotDeleteOrder(orderId, orderLinkId string) (*bybit2.SpotDeleteOrderResponse, error) {
	res, err := b.bybitRest.Spot().V1().SpotDeleteOrder(bybit2.SpotDeleteOrderParam{
		OrderID:     &orderId,
		OrderLinkID: &orderLinkId,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SpotGetOpenOrders - получение открытых ордеров
func (b *bybit) SpotGetOpenOrders(symbol, orderId string, limit int) (*bybit2.SpotOpenOrdersResponse, error) {
	res, err := b.bybitRest.Spot().V1().SpotOpenOrders(bybit2.SpotOpenOrdersParam{
		Symbol:  (*bybit2.SymbolSpot)(&symbol),
		OrderID: &orderId,
		Limit:   &limit,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SpotAPIKeyInfo - получение api-key
func (b *bybit) SpotAPIKeyInfo(symbol, orderId string, limit int) (*bybit2.V5APIKeyResponse, error) {
	res, err := b.bybitRest.V5().User().GetAPIKey()
	if err != nil {
		return nil, err
	}
	return res, nil
}
