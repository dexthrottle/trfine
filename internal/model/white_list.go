package model

type WhiteList struct {
	ID   uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Pair string `gorm:"type:varchar(255)" json:"pair"`
}
