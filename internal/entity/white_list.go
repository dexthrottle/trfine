package entity

// Белый список монет для торговли
type WhiteList struct {
	ID   uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Pair string `gorm:"type:varchar(255)" json:"pair"`
	// ExternalID uint64 `gorm:"type:int;default:0" json:"external_id"`
}
