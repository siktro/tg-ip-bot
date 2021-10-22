package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/siktro/tg-ip-bot/internal/bot"
	"github.com/sirupsen/logrus"
)

type CommandManager struct {
	bot    *bot.Bot
	logger *logrus.Logger
}

func NewManager(bot *bot.Bot, logger *logrus.Logger) *CommandManager {
	return &CommandManager{
		bot:    bot,
		logger: logger,
	}
}

// TODO: return errors?
func (cm CommandManager) Handle(upd *tgbotapi.Update) {
	switch {
	case upd.Message != nil:
		cm.HandleMessage(upd.Message)
	default:
		cm.logger.Warn("unexptected case")
	}
}
