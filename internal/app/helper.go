package app

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dexthrottle/trfine/internal/dto"
)

const (
	welcomeMessage = "Привет @username! Выполните настройку перед первым запуском:"
)

func firstStart(reader *bufio.Reader) (bool, string) {
	if _, err := os.Stat("trbotdatabase.db"); os.IsNotExist(err) {
		var useLogs bool
		fmt.Print("Включить логгирование? (Y/n): ")
		useLogsText, _ := reader.ReadString('\n')
		if strings.TrimSuffix(useLogsText, "\r\n") == "Y" ||
			strings.TrimSuffix(useLogsText, "\r\n") == "y" {
			useLogs = true
		} else {
			useLogs = false
		}
		fmt.Print("Введите порт для запуска приложения: ")
		portApp, _ := reader.ReadString('\n')
		return useLogs, portApp
	}
	return false, "8000"
}

func secondStart(reader *bufio.Reader) dto.AppConfigDTO {
	fmt.Print("Введите токен телеграм-бота: ")
	tgApiToken, _ := reader.ReadString('\n')
	var byBitUID int
	for {
		fmt.Print("Введите ByBit UID: ")
		byBitUIDStr, _ := reader.ReadString('\n')
		var err error
		byBitUID, err = strconv.Atoi(strings.TrimSuffix(byBitUIDStr, "\r\n"))
		if err != nil {
			fmt.Println("Введите корректный ByBit UID!")
			continue
		}
		break
	}
	fmt.Print("Введите ByBit ApiKey: ")
	byBitApiKey, _ := reader.ReadString('\n')
	fmt.Print("Введите ByBit ApiSecret: ")
	byBitApiSecret, _ := reader.ReadString('\n')

	var tgUserID int
	for {
		fmt.Print("Введите Ваш телегам-ID: ")
		tgUserIDStr, _ := reader.ReadString('\n')
		var err error
		tgUserID, err = strconv.Atoi(strings.TrimSuffix(tgUserIDStr, "\r\n"))
		if err != nil {
			fmt.Println("Введите корректный телегам-ID!")
			continue
		}
		break
	}

	tgNotificationChannel := fmt.Sprintln("Введите название телеграм-канала: ")

	appCfgDto := dto.AppConfigDTO{
		TgApiToken:            strings.TrimSuffix(tgApiToken, "\r\n"),
		ByBitUID:              byBitUID,
		ByBitApiKey:           strings.TrimSuffix(byBitApiKey, "\r\n"),
		ByBitApiSecret:        strings.TrimSuffix(byBitApiSecret, "\r\n"),
		TGUserID:              tgUserID,
		TGNotificationChannel: strings.TrimSuffix(tgNotificationChannel, "\r\n"),
	}
	return appCfgDto
}
