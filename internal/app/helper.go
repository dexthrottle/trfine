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

func firstStart() {
	// if _, err := os.Stat("trbotdatabase.db"); os.IsNotExist(err) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите токен телеграм-бота: ")
	tgApiToken, _ := reader.ReadString('\n')
	var byBitUID int
	for {
		fmt.Print("Введите ByBit UID: ")
		byBitUIDStr, _ := reader.ReadString('\n')
		var err error
		byBitUID, err = strconv.Atoi(strings.TrimSuffix(byBitUIDStr, "\n"))
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
		tgUserID, err = strconv.Atoi(strings.TrimSuffix(tgUserIDStr, "\n"))
		if err != nil {
			fmt.Println("Введите корректный телегам-ID!")
			continue
		}
		break
	}

	tgNotificationChannel := fmt.Sprintln("Введите название телеграм-канала: ")

	appCfgDto := dto.AppConfigDTO{
		TgApiToken:            strings.TrimSuffix(tgApiToken, "\n"),
		ByBitUID:              byBitUID,
		ByBitApiKey:           strings.TrimSuffix(byBitApiKey, "\n"),
		ByBitApiSecret:        strings.TrimSuffix(byBitApiSecret, "\n"),
		TGUserID:              tgUserID,
		TGNotificationChannel: strings.TrimSuffix(tgNotificationChannel, "\n"),
	}
	fmt.Println(appCfgDto)
	// }
}
