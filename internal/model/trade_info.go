package model

type TradeInfo struct {
	ID               uint64 `gorm:"primary_key:auto_increment" json:"id"`
	ExternalID       uint64 `gorm:"type:int;default:0" json:"external_id"`
	SellFilledOrders string `gorm:"type:varchar(255)" json:"sell_filled_orders"`
	SellOpenOrders   string `gorm:"type:varchar(255)" json:"sell_open_orders"`
}
