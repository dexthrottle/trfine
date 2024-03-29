package entity

// Профит за каждый день
type DailyProfit struct {
	ID     uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Day    string `gorm:"type:varchar(255)" json:"day"`
	Quote  string `gorm:"type:varchar(255)" json:"quote"`
	Profit string `gorm:"type:varchar(255)" json:"profit"`
}
