// Package provides a simple wrapper around `telegram-bot-api` bot API
// for some additional methods.
package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

// refactor into separate file?
type CommandHandler interface {
	Handle(val *tgbotapi.Update)
}

type Bot struct {
	API    *tgbotapi.BotAPI
	logger *log.Logger
	updCfg tgbotapi.UpdateConfig
}

type Config struct {
	Token   string
	Debug   bool
	Timeout int
	Offset  int
	Limit   int
	Logger  *log.Logger
}

func NewBot(cfg Config) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		return nil, err
	}

	cfg.Logger.Infof("Authorized on account %s", api.Self.UserName)

	api.Debug = cfg.Debug
	u := tgbotapi.NewUpdate(cfg.Offset)
	u.Timeout = cfg.Timeout
	u.Limit = cfg.Limit

	return &Bot{
		API:    api,
		logger: cfg.Logger,
		updCfg: u,
	}, nil
}

// TODO: add done channel?
func (b Bot) Serve(handler CommandHandler) error {
	updates, err := b.API.GetUpdatesChan(b.updCfg)
	if err != nil {
		return err
	}

	// Optional: wait for updates and clear them if you don't want to handle
	// a large backlog of old messages
	// если во время шатдауна бота юзеры писали ему,
	// то все сообщения сохраняются в телеге и будут переданы боту
	// как только он станет доступным
	// time.Sleep(time.Millisecond * 500)
	// updates.Clear()

	// TODO: run in goroutines?
	for update := range updates {
		handler.Handle(&update)
	}

	return nil
}
