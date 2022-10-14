package entity

// Дробная продажа
type TrailingOrders struct {
	ID   uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Pair string `gorm:"type:varchar(255)" json:"pair"`
	P    string `gorm:"type:varchar(255)" json:"p"`
	Q    string `gorm:"type:varchar(255)" json:"q"`
}
