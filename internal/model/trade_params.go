package model

// Торговые параметры для бота (на их основе осуществляется торговля)
type TradeParams struct {
	ID              uint64 `gorm:"primary_key:auto_increment" json:"id"`
	ExternalID      uint64 `gorm:"type:int;default:0" json:"external_id"`
	NameList        string `gorm:"type:varchar(255)" json:"name_list"`
	MinBalance      string `gorm:"type:varchar(255)" json:"min_balance"`
	MinOrder        string `gorm:"type:varchar(255)" json:"min_order"`
	MinPrice        string `gorm:"type:varchar(255)" json:"min_price"`
	DailyPercent    string `gorm:"type:varchar(255)" json:"daily_percent"`
	SellUpp         string `gorm:"type:varchar(255)" json:"sell_up"`
	BuyDown         string `gorm:"type:varchar(255)" json:"buy_down"`
	MaxTradePairs   string `gorm:"type:varchar(255)" json:"max_trade_pairs"`
	AutoTradePairs  bool   `gorm:"type:bool" json:"auto_trade_pairs"`
	DeltaPercent    bool   `gorm:"type:bool" json:"delta_percent"`
	NumAver         bool   `gorm:"type:bool" json:"num_aver"`
	StepAver        string `gorm:"type:varchar(255)" json:"step_aver"`
	MaxAver         string `gorm:"type:varchar(255)" json:"max_aver"`
	QuantityAver    string `gorm:"type:varchar(255)" json:"quantity_aver"`
	TrailingStop    bool   `gorm:"type:bool" json:"trailing_stop"`
	TrailingPercent string `gorm:"type:varchar(255)" json:"trailing_percent"`
	TrailingPart    string `gorm:"type:varchar(255)" json:"trailing_part"`
	TrailingPrice   string `gorm:"type:varchar(255)" json:"trailing_price"`
	UserOrder       bool   `gorm:"type:bool" json:"user_order"`
	FiatCurrencies  string `gorm:"type:varchar(255)" json:"fiat_currencies"`
	QuoteAsset      string `gorm:"type:varchar(255)" json:"quote_asset"`
	DoubleAsset     bool   `gorm:"type:bool" json:"double_asset"`
	Active          bool   `gorm:"type:bool" json:"active"`
}
