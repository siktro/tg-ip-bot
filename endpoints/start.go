package endpoints

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/siktro/tg-ip-bot/bot"
)

const startMessage = `
Привет!

Этот бот умеет собирать информацию по IP адресам.
Нажми START и введи интересующий тебя адрес.
`

func (cm *EndpointManager) start(bot *bot.Bot, msg *tgbotapi.Message) {
	res := tgbotapi.NewMessage(msg.Chat.ID, startMessage)
	res.ParseMode = "markdown"
	bot.Send(res)
}
