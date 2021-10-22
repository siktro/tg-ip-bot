package commands

import "github.com/sirupsen/logrus"

type Command struct {
	exec func()
	help string
}

var Commands = map[string]Command{
	"/start": {
		exec: func() {
			logrus.Info("Hit!")
		},
		help: "some help",
	},
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
