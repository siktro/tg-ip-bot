package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (cm *CommandManager) Start(msg *tgbotapi.Message) {
	cm.logger.Info("Hit command")
	res := tgbotapi.NewMessage(msg.Chat.ID, "Hello here!")
	cm.bot.Send(res)
}
