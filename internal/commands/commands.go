package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type Command func(*tgbotapi.Message)

func makeCommands(h *Handler) map[string]Command {
	m := map[string]Command{
		"/start": nil,
		"/ping":  h.Ping,
	}
	return m
}

/*
TODO:
	Telegram basic commands:
		- /start
		- /help
		- /settings
	Required commands:
		(User)
		- /lookup
		- /history

		(Admin)
		- /broadcast
		- /info
		- /hist_by_id
		- /del_records
*/
