// Package provides a simple wrapper around `telegram-bot-api` bot API
// for some additional methods.
package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

type UpdateHandler interface {
	ProcessUpdate(*Bot, *tgbotapi.Update)
}

type Bot struct {
	*tgbotapi.BotAPI
	cfg *Config

	handler UpdateHandler
}

type Config struct {
	// Telegram bot token.
	Token   string
	Debug   bool
	Logger  *logrus.Logger
	Handler UpdateHandler

	// The number of goroutines processing messages.
	WorkersLimit int
}

// NewBot creates a new Bot instance with provided configuration.
func NewBot(cfg Config) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		logrus.
			WithError(err).
			WithFields(logrus.Fields{
				"token": cfg.Token,
			}).
			Error("failed to initialize/authorize bot;")
		return nil, err
	}
	api.Debug = cfg.Debug

	return &Bot{
		BotAPI:  api,
		cfg:     &cfg,
		handler: cfg.Handler,
	}, nil
}

func (b *Bot) Config() *Config {
	return b.cfg
}

func (b *Bot) ListenAndServe(updateCfg tgbotapi.UpdateConfig) error {
	updates, err := b.GetUpdatesChan(updateCfg)
	if err != nil {
		// In the current version of the lib (v4.6.4+incompatible)
		// it is impossible(?) to get an error from this func.
		// But just in case.
		b.cfg.Logger.
			WithError(err).
			Warn("unable to get an update channel;")
		return err
	}

	// Limit the number of goroutines.
	limitCh := make(chan struct{}, b.cfg.WorkersLimit)

	// It is likely that we must close the channel,
	// because i didn't find any close operation in the lib.
	for update := range updates {
		limitCh <- struct{}{}
		go func(upd tgbotapi.Update) {
			// TODO: create ctx and a new logger for goroutine
			// TODO: add recover if goroutine panics
			b.handler.ProcessUpdate(b, &upd)
			defer func() {
				<-limitCh
			}()
			// b.handler.
		}(update)
	}

	return nil
}
