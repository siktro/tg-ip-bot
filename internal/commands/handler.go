package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/siktro/tg-ip-bot/internal/bot"
	"github.com/sirupsen/logrus"
)

type CommandManager struct {
	logger   *logrus.Logger
	bot      *bot.Bot
	commands map[string]Command
}

func NewManager(bot *bot.Bot, logger *logrus.Logger) *CommandManager {
	m := &CommandManager{
		bot:    bot,
		logger: logger,
	}
	return m
}

func (h *CommandManager) handleMessageType(msg *tgbotapi.Message) {
	msg.IsCommand()
}
