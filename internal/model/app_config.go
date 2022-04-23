package model

type AppConfig struct {
	ID                    uint64 `gorm:"primary_key:auto_increment" json:"id"`
	TgApiToken            string `gorm:"type:varchar(255)" json:"tg_api_token"`
	ByBitUID              string `gorm:"type:varchar(255)" json:"by_bit_uid"`
	ByBitApiKey           string `gorm:"type:varchar(255)" json:"by_bit_api_key"`
	ByBitApiSecret        string `gorm:"type:varchar(255)" json:"by_bit_api_secret"`
	TGUserID              int    `gorm:"type:int;unique;not null" json:"tg_user_id"`
	UseLogs               bool   `gorm:"type:bool;default:false" json:"use_logs"`
	TGNotificationChannel string `gorm:"type:varchar(255)" json:"tg_notification_channel"`
}
