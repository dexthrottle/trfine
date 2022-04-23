package model

type User struct {
	ID         uint64      `gorm:"primary_key:auto_increment" json:"id"`
	FirstName  string      `gorm:"type:varchar(255)" json:"first_name"`
	LastName   string      `gorm:"type:varchar(255)" json:"last_name"`
	Username   string      `gorm:"type:varchar(255)" json:"username"`
	UserTGId   int         `gorm:"type:int;unique;not null" json:"user_tg_id"`
	Posts      *[]Post     `json:"posts,omitempty"`
	Categories *[]Category `json:"categories,omitempty"`
}
