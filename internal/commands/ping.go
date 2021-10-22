package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (cm CommandManager) Ping(msg *tgbotapi.Message) {
	newMsg := tgbotapi.NewMessage(msg.Chat.ID, "Pong!")
	newMsg.ReplyToMessageID = msg.MessageID

	cm.bot.API.Send(newMsg)
}
