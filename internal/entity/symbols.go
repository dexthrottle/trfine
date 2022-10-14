package entity

// Торгуемая пара и параметры к ней
type Symbols struct {
	ID                 uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Pair               string `gorm:"type:varchar(255);unique;not null" json:"pair"`
	BaseAsset          string `gorm:"type:varchar(255)" json:"base_asset"`
	QuoteAsset         string `gorm:"type:varchar(255)" json:"quote_asset"`
	StepSize           string `gorm:"type:varchar(255)" json:"step_size"`
	TickSize           string `gorm:"type:varchar(255)" json:"tick_size"`
	MinNotional        string `gorm:"type:varchar(255)" json:"min_notional"`
	PriceChangePercent string `gorm:"type:varchar(255)" json:"price_change_percent"`
	BidPrice           string `gorm:"type:varchar(255)" json:"bid_price"`
	AskPrice           string `gorm:"type:varchar(255)" json:"ask_price"`
	AveragePrice       string `gorm:"type:varchar(255)" json:"average_price"`
	BuyPrice           string `gorm:"type:varchar(255)" json:"buy_price"`
	SellPrice          string `gorm:"type:varchar(255)" json:"sell_price"`
	TrailingPrice      string `gorm:"type:varchar(255)" json:"trailing_price"`
	AllQuantity        string `gorm:"type:varchar(255)" json:"all_quantity"`
	FreeQuantity       string `gorm:"type:varchar(255)" json:"free_quantity"`
	LockQuantity       string `gorm:"type:varchar(255)" json:"lock_quantity"`
	OrderId            string `gorm:"type:varchar(255)" json:"order_id"`
	Profit             string `gorm:"type:varchar(255)" json:"profit"`
	TotalQuote         string `gorm:"type:varchar(255)" json:"total_quote"`
	StepAveraging      string `gorm:"type:varchar(255)" json:"step_averaging"`
	NumAveraging       string `gorm:"type:varchar(255)" json:"num_averaging"`
	StatusOrder        string `gorm:"type:varchar(255)" json:"status_order"`
}
