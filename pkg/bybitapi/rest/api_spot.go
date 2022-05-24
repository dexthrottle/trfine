package rest

import (
	"net/http"
)

// GetWalletBalanceSpot => GET /spot/v1/account
func (b *ByBit) GetWalletBalanceSpot() (*string, []byte, *ResultSpot, error) {
	var ret ResultSpot

	query, resp, err := b.SignedRequest(
		http.MethodGet, "/spot/v1/account", make(map[string]interface{}), &ret,
	)
	if err != nil {
		return nil, nil, nil, err
	}

	return &query, resp, &ret, nil
}

// GetOrderSpot => GET /spot/v1/order
func (b *ByBit) GetOrderSpot(orderId, orderLinkId string) (*string, []byte, *ResultOrderSpot, error) {
	var ret ResultOrderSpot

	params := map[string]interface{}{}
	if orderId != "" {
		params["orderId"] = orderId
	}
	if orderLinkId != "" {
		params["orderLinkId"] = orderLinkId
	}

	query, resp, err := b.SignedRequest(
		http.MethodGet, "/spot/v1/order", params, &ret,
	)
	if err != nil {
		return nil, nil, nil, err
	}
	return &query, resp, &ret, nil
}

// CreateOrderSpot => POST /spot/v1/order
func (b *ByBit) CreateOrderSpot(symbol, side, _type string, qty float64, price float64, timeInForce, orderLinkId string) (*string, []byte, *ResultCreateDeleteOrderSpot, error) {

	var ret ResultCreateDeleteOrderSpot

	params := map[string]interface{}{}
	if price != 0 {
		params["price"] = price
	}
	if timeInForce != "" {
		params["timeInForce"] = timeInForce
	}
	if orderLinkId != "" {
		params["orderLinkId"] = orderLinkId
	}
	if qty != 0 {
		params["qty"] = qty
	}

	if symbol != "" {
		params["symbol"] = symbol
	}
	if side != "" {
		params["side"] = side
	}
	if _type != "" {
		params["type"] = _type
	}

	query, resp, err := b.SignedRequest(
		http.MethodPost, "/spot/v1/order", params, &ret,
	)
	if err != nil {
		return nil, nil, nil, err
	}

	return &query, resp, &ret, nil
}

// DeleteOrderSpot => DELETE /spot/v1/order
func (b *ByBit) DeleteOrderSpot(orderId, orderLinkId string) (*string, []byte, *ResultCreateDeleteOrderSpot, error) {
	var ret ResultCreateDeleteOrderSpot
	params := map[string]interface{}{}
	if orderId != "" {
		params["orderId"] = orderId
	}
	if orderLinkId != "" {
		params["orderLinkId"] = orderLinkId
	}

	query, resp, err := b.SignedRequest(
		http.MethodDelete, "/spot/v1/order", params, &ret,
	)
	if err != nil {
		return nil, nil, nil, err
	}
	return &query, resp, &ret, nil
}

// GetOpenOrdersSpot => GET /spot/v1/open-orders
func (b *ByBit) GetOpenOrdersSpot(symbol, orderId string, limit int) (*string, []byte, *OrdersSpot, error) {
	var ret *OrdersSpot

	params := map[string]interface{}{}
	if symbol != "" {
		params["symbol"] = symbol
	}
	if orderId != "" {
		params["orderId"] = orderId
	}
	if limit != 0 {
		params["limit"] = limit
	}

	query, resp, err := b.SignedRequest(
		http.MethodGet, "/spot/v1/open-orders", params, &ret,
	)
	if err != nil {
		return nil, nil, nil, err
	}
	return &query, resp, ret, nil
}

// GetHistoryOrdersSpot => GET /spot/v1/history-orders
func (b *ByBit) GetHistoryOrdersSpot(symbol, orderId string, limit int, startTime, endTime int64) (*string, []byte, *OrdersSpot, error) {
	var ret *OrdersSpot

	params := map[string]interface{}{}
	if symbol != "" {
		params["symbol"] = symbol
	}
	if orderId != "" {
		params["orderId"] = orderId
	}
	if limit != 0 {
		params["limit"] = limit
	}

	query, resp, err := b.SignedRequest(
		http.MethodGet, "/spot/v1/history-orders", params, &ret,
	)
	if err != nil {
		return nil, nil, nil, err
	}
	return &query, resp, ret, nil
}

// GetMyTradesSpot => GET /spot/v1/open-orders
func (b *ByBit) GetMyTradesSpot(symbol, orderId string, limit, fromTicketId, toTicketId, startTime, endTime int64) (*string, []byte, *MyTradesSpot, error) {
	var ret *MyTradesSpot

	params := map[string]interface{}{}
	if symbol != "" {
		params["symbol"] = symbol
	}
	if limit != 0 {
		params["limit"] = limit
	}
	if fromTicketId != 0 {
		params["fromTicketId"] = fromTicketId
	}
	if toTicketId != 0 {
		params["toTicketId"] = toTicketId
	}
	if orderId != "" {
		params["orderId"] = orderId
	}
	if startTime != 0 {
		params["startTime"] = startTime
	}
	if endTime != 0 {
		params["endTime"] = endTime
	}

	query, resp, err := b.SignedRequest(
		http.MethodGet, "/spot/v1/open-orders", params, &ret,
	)
	if err != nil {
		return nil, nil, nil, err
	}
	return &query, resp, ret, nil
}
