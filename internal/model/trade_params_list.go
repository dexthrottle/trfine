package model

type TradeParamsList struct {
	ID              uint64 `gorm:"primary_key:auto_increment" json:"id"`
	NameList        string `gorm:"type:varchar(255)" json:"name_list"`
	MinBnb          string `gorm:"type:varchar(255)" json:"min_bnb"`
	MinBalance      string `gorm:"type:varchar(255)" json:"min_balance"`
	MinOrder        string `gorm:"type:varchar(255)" json:"min_order"`
	MinPrice        string `gorm:"type:varchar(255)" json:"min_price"`
	DailyPercent    string `gorm:"type:varchar(255)" json:"daily_percent"`
	SellUpp         string `gorm:"type:varchar(255)" json:"sell_up"`
	BuyDown         string `gorm:"type:varchar(255)" json:"buy_down"`
	MainTradePairs  string `gorm:"type:varchar(255)" json:"max_trade_pairs"`
	AutoTradePairs  bool   `gorm:"type:bool" json:"auto_trade_pairs"`
	NumAver         bool   `gorm:"type:bool" json:"delta_percent"`
	StepAver        string `gorm:"type:varchar(255)" json:"num_aver"`
	MaxAver         string `gorm:"type:varchar(255)" json:"step_aver"`
	QuantityAver    string `gorm:"type:varchar(255)" json:"max_aver"`
	TrailingStop    bool   `gorm:"type:bool" json:"quantity_aver"`
	TrailingPercent string `gorm:"type:varchar(255)" json:"trailing_stop"`
	TrailingPart    string `gorm:"type:varchar(255)" json:"trailing_percent"`
	TrailingPrice   string `gorm:"type:varchar(255)" json:"trailing_part"`
	UserOrder       bool   `gorm:"type:bool" json:"trailing_price"`
	FiatCurrencies  string `gorm:"type:varchar(255)" json:"user_order"`
	QuoteAsset      string `gorm:"type:varchar(255)" json:"fiat_currencies"`
	DoubleAsset     bool   `gorm:"type:bool" json:"quote_asset"`
}
