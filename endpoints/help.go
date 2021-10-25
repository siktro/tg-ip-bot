package endpoints

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/siktro/tg-ip-bot/bot"
)

const helpMessage = `
Доступные команды бота:


`

func (cm *EndpointManager) Help(bot *bot.Bot, msg *tgbotapi.Message) {
	res := tgbotapi.NewMessage(msg.Chat.ID, helpMessage)
	res.ParseMode = "markdown"
	bot.Send(res)
}
