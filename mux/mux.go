package mux

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/siktro/tg-ip-bot/bot"
	"github.com/sirupsen/logrus"
)

type EndpointEntry struct {
	Callback    func(*bot.Bot, *tgbotapi.Message)
	Description string
	Group       string
}

type EndpointMux struct {
	logger   *logrus.Logger
	handlers map[string][]*EndpointEntry
}

func NewEndpointMux(logger *logrus.Logger) *EndpointMux {
	return &EndpointMux{
		logger:   logger,
		handlers: make(map[string][]*EndpointEntry),
	}
}

func (m *EndpointMux) Handle(route string, e *EndpointEntry) {
	m.handlers[route] = append(m.handlers[route], e)
}

func (m *EndpointMux) ProcessUpdate(bot *bot.Bot, upd *tgbotapi.Update) {
	if upd.Message == nil {
		// TODO: send help message;
		return
	}

	msg := upd.Message
	if !msg.IsCommand() {
		// TODO: send help message; not a command;
		return
	}

	cmd := msg.Command()
	handlers, ok := m.handlers[cmd]
	if !ok {
		// TODO: send help message; no such command;
		return
	}

	// Invoke every callback for specified command.
	for _, h := range handlers {
		h.Callback(bot, msg)
	}
}
