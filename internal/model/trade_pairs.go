package model

// Пары с биржи
type TradePairs struct {
	ID         uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Pair       string `gorm:"type:varchar(255)" json:"pair"`
	BaseAsset  string `gorm:"type:varchar(255)" json:"base_asset"`
	QuoteAsset string `gorm:"type:varchar(255)" json:"quote_asset"`
}
