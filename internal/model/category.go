package model

type Category struct {
	ID     uint64  `gorm:"primary_key:auto_increment" json:"id"`
	Name   string  `gorm:"type:varchar(255);unique;not null" json:"name"`
	Posts  *[]Post `json:"posts,omitempty"`
	UserID uint64  `gorm:"not null" json:"-"`
	User   User    `gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}
