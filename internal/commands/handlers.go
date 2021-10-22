package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (cm CommandManager) HandleMessage(msg *tgbotapi.Message) {
	if !msg.IsCommand() {
		// TODO: send help message
		return
	}

	cm.Ping(msg)
}
