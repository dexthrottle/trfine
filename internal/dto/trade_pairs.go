package dto

// Пары с биржи
type TradePairs struct {
	ID         uint64
	Pair       string
	BaseAsset  string
	QuoteAsset string
}
