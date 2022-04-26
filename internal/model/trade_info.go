package model

type TradeInfo struct {
	ID               uint64 `gorm:"primary_key:auto_increment" json:"id"`
	SellFilledOrders string `gorm:"type:varchar(255)" json:"sell_filled_orders"`
	SellOpenOrders   string `gorm:"type:varchar(255)" json:"sell_open_orders"`
}
