package dto

// Торгуемая пара и параметры к ней
type Symbols struct {
	ID                 uint64
	Pair               string
	BaseAsset          string
	QuoteAsset         string
	StepSize           string
	TickSize           string
	MinNotional        string
	PriceChangePercent string
	BidPrice           string
	AskPrice           string
	AveragePrice       string
	BuyPrice           string
	SellPrice          string
	TrailingPrice      string
	AllQuantity        string
	FreeQuantity       string
	LockQuantity       string
	OrderId            string
	Profit             string
	TotalQuote         string
	StepAveraging      string
	NumAveraging       string
	StatusOrder        string
}
