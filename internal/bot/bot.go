// Package provides a simple wrapper around `telegram-bot-api` bot API
// for some additional methods.
package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type CommandHandler func(*tgbotapi.Message)

type Bot struct {
	*tgbotapi.BotAPI
	cfg      *Config
	handlers map[string]CommandHandler
}

type Config struct {
	// Telegram bot token.
	Token string
	Debug bool

	// The number of goroutines processing messages.
	WorkersLimit int
}

// NewBot creates a new Bot instance with provided configuration.
func NewBot(cfg Config) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		return nil, err
	}

	api.Debug = cfg.Debug

	return &Bot{
		BotAPI:   api,
		cfg:      &cfg,
		handlers: make(map[string]CommandHandler),
	}, nil
}

func (b *Bot) Config() *Config {
	return b.cfg
}

func (b *Bot) Handle(endpoint string, fn CommandHandler) {
	b.handlers[endpoint] = fn
}

func (b *Bot) ListenAndServe(updateCfg tgbotapi.UpdateConfig) error {
	updates, err := b.GetUpdatesChan(updateCfg)
	if err != nil {
		return err
	}

	// Limit the number of goroutines.
	limitCh := make(chan struct{}, b.cfg.WorkersLimit)

	// It is likely that we must close the channel,
	// because i didn't find any close operation in the lib.
	for update := range updates {
		limitCh <- struct{}{}
		go func(upd tgbotapi.Update) {
			b.processUpdate(upd)
			limitCh <- struct{}{}
		}(update)
	}

	return nil
}

func (b *Bot) processUpdate(upd tgbotapi.Update) {
	switch {
	case upd.Message != nil:
		m := upd.Message
		if m.IsCommand() {
			cmd := m.Command()
			b.processCommand(cmd, m)
			return
		}

		b.processMessage(m)

	default:
		// return a help message?
	}
}

// Generic, non-command message.
func (b *Bot) processMessage(msg *tgbotapi.Message) {
	// TODO: send help message
}

func (b *Bot) processCommand(endpoint string, msg *tgbotapi.Message) {
	h, ok := b.handlers[endpoint]
	if !ok {
		// TODO: send help message; no such command
		return
	}

	h(msg)
}
