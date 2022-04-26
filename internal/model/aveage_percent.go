package model

type AveragePercent struct {
	ID      uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Day     string `gorm:"type:varchar(255)" json:"day"`
	Percent string `gorm:"type:varchar(255)" json:"percent"`
}
