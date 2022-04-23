package app

import (
	"fmt"
	"os"
	"strconv"

	"github.com/dexthrottle/trfine/internal/dto"
)

const (
	welcomeMessage = "Привет @username! Выполните настройку перед первым запуском:"
)

func firstStart() {
	if _, err := os.Stat("trbotdatabase.db"); os.IsNotExist(err) {
		fmt.Println(welcomeMessage)
		tgApiToken := fmt.Sprintln("")
		byBitUID := fmt.Sprintln("")
		byBitApiKey := fmt.Sprintln("")
		byBitApiSecret := fmt.Sprintln("")
		tgUserID, err := strconv.Atoi(fmt.Sprintln(""))
		if err != nil {
			fmt.Println(err)
		}
		useLogs := fmt.Sprintln("")
		tgNotificationChannel := fmt.Sprintln("")

		appCfgDto := dto.AppConfigDTO{
			TgApiToken:            tgApiToken,
			ByBitUID:              byBitUID,
			ByBitApiKey:           byBitApiKey,
			ByBitApiSecret:        byBitApiSecret,
			TGUserID:              tgUserID,
			UseLogs:               useLogs,
			TGNotificationChannel: tgNotificationChannel,
		}
	}
}
