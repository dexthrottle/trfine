package telegram

import (
	"github.com/dexthrottle/trfine/pkg/logging"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot *tgbotapi.BotAPI
	log logging.Logger
}

func NewBot(bot *tgbotapi.BotAPI, log logging.Logger) *Bot {
	return &Bot{
		bot: bot,
		log: log,
	}
}

func (b *Bot) Start() error {

	updateConfig := tgbotapi.NewUpdate(0)
	updates := b.bot.GetUpdatesChan(updateConfig)
	updateConfig.Timeout = 30

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		// // Handle commands
		// if update.Message.IsCommand() {
		// 	if err := b.handleCommand(update.Message); err != nil {
		// 		b.handleError(update.Message.Chat.ID, err)
		// 	}

		// 	continue
		// }

		// // Handle regular messages
		// if err := b.handleMessage(update.Message); err != nil {
		// 	b.handleError(update.Message.Chat.ID, err)
		// }
	}

	return nil
}
