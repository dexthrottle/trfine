package rest

type BalanceSpot struct {
	Coin     string `json:"coin"`
	CoinID   string `json:"coinId"`
	CoinName string `json:"coinName"`
	Total    string `json:"total"`
	Free     string `json:"free"`
	Locked   string `json:"locked"`
}

type ResultSpot struct {
	RetCode int    `json:"ret_code"`
	RetMsg  string `json:"ret_msg"`
	ExtCode string `json:"ext_code"`
	ExtInfo string `json:"ext_info"`
	Result  struct {
		Balances []BalanceSpot `json:"balances"`
	} `json:"result"`
}

type OrderSpot struct {
	AccountID           string `json:"accountId"`
	ExchangeID          string `json:"exchangeId"`
	Symbol              string `json:"symbol"`
	SymbolName          string `json:"symbolName"`
	OrderLinkID         string `json:"orderLinkId"`
	OrderID             string `json:"orderId"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	AvgPrice            string `json:"avgPrice"`
	Status              string `json:"status"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	Side                string `json:"side"`
	StopPrice           string `json:"stopPrice"`
	IcebergQty          string `json:"icebergQty"`
	Time                string `json:"time"`
	UpdateTime          string `json:"updateTime"`
	IsWorking           bool   `json:"isWorking"`
}

type ResultOrderSpot struct {
	RetCode int         `json:"ret_code"`
	RetMsg  string      `json:"ret_msg"`
	ExtCode interface{} `json:"ext_code"`
	ExtInfo interface{} `json:"ext_info"`
	Result  OrderSpot   `json:"result"`
}

type OrdersSpot struct {
	RetCode int         `json:"ret_code"`
	RetMsg  string      `json:"ret_msg"`
	ExtCode interface{} `json:"ext_code"`
	ExtInfo interface{} `json:"ext_info"`
	Result  []OrderSpot `json:"result"`
}

type ResultCreateDeleteOrderSpot struct {
	RetCode int         `json:"ret_code"`
	RetMsg  string      `json:"ret_msg"`
	ExtCode interface{} `json:"ext_code"`
	ExtInfo interface{} `json:"ext_info"`
	Result  struct {
		AccountID    string `json:"accountId"`
		Symbol       string `json:"symbol"`
		SymbolName   string `json:"symbolName"`
		OrderLinkID  string `json:"orderLinkId"`
		OrderID      string `json:"orderId"`
		TransactTime string `json:"transactTime"`
		Price        string `json:"price"`
		OrigQty      string `json:"origQty"`
		ExecutedQty  string `json:"executedQty"`
		Status       string `json:"status"`
		TimeInForce  string `json:"timeInForce"`
		Type         string `json:"type"`
		Side         string `json:"side"`
	} `json:"result"`
}

type MyTradesSpot struct {
	RetCode int         `json:"ret_code"`
	RetMsg  string      `json:"ret_msg"`
	ExtCode interface{} `json:"ext_code"`
	ExtInfo interface{} `json:"ext_info"`
	Result  []struct {
		ID              string `json:"id"`
		Symbol          string `json:"symbol"`
		SymbolName      string `json:"symbolName"`
		OrderID         string `json:"orderId"`
		TicketID        string `json:"ticketId"`
		MatchOrderID    string `json:"matchOrderId"`
		Price           string `json:"price"`
		Qty             string `json:"qty"`
		Commission      string `json:"commission"`
		CommissionAsset string `json:"commissionAsset"`
		Time            string `json:"time"`
		IsBuyer         bool   `json:"isBuyer"`
		IsMaker         bool   `json:"isMaker"`
		Fee             struct {
			FeeTokenID   string `json:"feeTokenId"`
			FeeTokenName string `json:"feeTokenName"`
			Fee          string `json:"fee"`
		} `json:"fee"`
		FeeTokenID  string `json:"feeTokenId"`
		FeeAmount   string `json:"feeAmount"`
		MakerRebate string `json:"makerRebate"`
	} `json:"result"`
}
