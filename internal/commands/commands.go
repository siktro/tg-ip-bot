package commands

import (
	"github.com/siktro/tg-ip-bot/internal/bot"
	"github.com/sirupsen/logrus"
)

type CommandManager struct {
	logger *logrus.Logger
	bot    *bot.Bot
	// db ...
}

func NewManager(bot *bot.Bot, logger *logrus.Logger) *CommandManager {
	m := &CommandManager{
		bot:    bot,
		logger: logger,
	}
	return m
}
