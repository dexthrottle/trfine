package dto

// Торговые параметры для бота (на их основе осуществляется торговля)
type TradeParams struct {
	ID              uint64
	ExternalID      uint64
	NameList        string
	MinBalance      string
	MinOrder        string
	MinPrice        string
	DailyPercent    string
	SellUpp         string
	BuyDown         string
	MaxTradePairs   string
	AutoTradePairs  bool
	DeltaPercent    bool
	NumAver         bool
	StepAver        string
	MaxAver         string
	QuantityAver    string
	TrailingStop    bool
	TrailingPercent string
	TrailingPart    string
	TrailingPrice   string
	UserOrder       bool
	FiatCurrencies  string
	QuoteAsset      string
	DoubleAsset     bool
	Active          bool
}
