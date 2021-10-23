package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (h *CommandManager) Ping(msg *tgbotapi.Message) {
	// res := tgbotapi.NewMessage(msg.Chat.ID, "Pong!")
	// h.bot.API().Send(res)
}
