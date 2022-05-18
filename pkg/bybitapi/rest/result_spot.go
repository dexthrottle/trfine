package rest

type BalanceSpot struct {
	Coin      string `json:"coin"`
	CoinID    string `json:"coinId"`
	CointName string `json:"coinName"`
	Total     string `json:"total"`
	Free      string `json:"free"`
	Locked    string `json:"locked"`
}

type BalancesSpot struct {
	Balances []BalanceSpot `json:"balances"`
}

type ResultSpot struct {
	RetCode int          `json:"ret_code"`
	RetMsg  string       `json:"ret_msg"`
	ExtCode string       `json:"ext_code"`
	ExtInfo string       `json:"ext_info"`
	Result  []SymbolInfo `json:"result"`
}
