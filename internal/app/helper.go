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

func firstStart() bool {
	if _, err := os.Stat("trbotdatabase.db"); os.IsNotExist(err) {
		tgApiToken := fmt.Sprintln("Введите токен телеграм-бота: ")

		var byBitUID int
		for {
			byBitUID, err = strconv.Atoi(fmt.Sprintln("Введите ByBit UID: "))
			if err != nil {
				fmt.Println("Введите корректный ByBit UID!")
				continue
			}
			break
		}

		byBitApiKey := fmt.Sprintln("Введите ByBit ApiKey: ")
		byBitApiSecret := fmt.Sprintln("Введите ByBit ApiSecret: ")

		var tgUserID int
		for {
			tgUserID, err = strconv.Atoi(fmt.Sprintln("Введите Ваш телегам-ID: "))
			if err != nil {
				fmt.Println("Введите корректный телегам-ID!")
				continue
			}
			break
		}

		useLogs := false
		useLogsText := fmt.Sprintln("Включить логгирование? (Y/n): ")
		if useLogsText == "Y" || useLogsText == "y" {
			useLogs = true
		}
		tgNotificationChannel := fmt.Sprintln("Введите название телеграм-канала: ")

		appCfgDto := dto.AppConfigDTO{
			TgApiToken:            tgApiToken,
			ByBitUID:              byBitUID,
			ByBitApiKey:           byBitApiKey,
			ByBitApiSecret:        byBitApiSecret,
			TGUserID:              tgUserID,
			UseLogs:               useLogs,
			TGNotificationChannel: tgNotificationChannel,
		}
		fmt.Println(appCfgDto)
	}
	return true
}
