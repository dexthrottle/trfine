package app

// import (
// 	"bufio"
// 	"fmt"
// 	"strconv"
// 	"strings"

// 	"github.com/dexthrottle/trfine/internal/dto"
// )

// const (
// 	welcomeMessage = "Привет @username! Выполните настройку перед первым запуском:"
// )

// func firstRunApp(reader *bufio.Reader) dto.AppConfigDTO {
// 	fmt.Println(welcomeMessage)
// 	fmt.Print("Введите токен телеграм-бота: ")
// 	tgApiToken, _ := reader.ReadString('\n')
// 	var byBitUID int
// 	for {
// 		fmt.Print("Введите ByBit UID: ")
// 		byBitUIDStr, _ := reader.ReadString('\n')
// 		var err error
// 		byBitUID, err = strconv.Atoi(strings.TrimSuffix(strings.TrimSuffix(byBitUIDStr, "\n"), "\r"))
// 		if err != nil {
// 			fmt.Println("Введите корректный ByBit UID!")
// 			continue
// 		}
// 		break
// 	}
// 	fmt.Print("Введите ByBit ApiKey: ")
// 	byBitApiKey, _ := reader.ReadString('\n')
// 	fmt.Print("Введите ByBit ApiSecret: ")
// 	byBitApiSecret, _ := reader.ReadString('\n')

// 	var tgUserID int
// 	for {
// 		fmt.Print("Введите Ваш телегам-ID: ")
// 		tgUserIDStr, _ := reader.ReadString('\n')
// 		var err error
// 		tgUserID, err = strconv.Atoi(strings.TrimSuffix(strings.TrimSuffix(tgUserIDStr, "\n"), "\r"))
// 		if err != nil {
// 			fmt.Println("Введите корректный телегам-ID!")
// 			continue
// 		}
// 		break
// 	}

// 	fmt.Print("Введите название телеграм-канала: @")
// 	tgNotificationChannel, _ := reader.ReadString('\n')
// 	appCfgDto := dto.AppConfigDTO{
// 		TgApiToken:            strings.TrimSuffix(strings.TrimSuffix(tgApiToken, "\n"), "\r"),
// 		ByBitUID:              byBitUID,
// 		ByBitApiKey:           strings.TrimSuffix(strings.TrimSuffix(byBitApiKey, "\n"), "\r"),
// 		ByBitApiSecret:        strings.TrimSuffix(strings.TrimSuffix(byBitApiSecret, "\n"), "\r"),
// 		TGUserID:              tgUserID,
// 		TGNotificationChannel: strings.TrimSuffix(strings.TrimSuffix(tgNotificationChannel, "\n"), "\r"),
// 	}
// 	return appCfgDto
// }
