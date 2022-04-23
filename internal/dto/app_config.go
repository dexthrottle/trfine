package dto

type AppConfigDTO struct {
	TgApiToken            string `json:"tg_api_token"`
	ByBitUID              string `json:"by_bit_uid"`
	ByBitApiKey           string `json:"by_bit_api_key"`
	ByBitApiSecret        string `json:"by_bit_api_secret"`
	TGUserID              int    `json:"tg_user_id"`
	UseLogs               bool   `json:"use_logs"`
	TGNotificationChannel string `json:"tg_notification_channel"`
}
