package main

import (
	"io/ioutil"

	"github.com/siktro/tg-ip-bot/internal/bot"
	"github.com/siktro/tg-ip-bot/internal/commands"
	log "github.com/sirupsen/logrus"
)

func run() error {
	// get sysenvs, setup from docker?
	// check db

	// check bot

	// spinup command handler

	// gracefuly shutdown? if any
	token, err := readToken()
	if err != nil {
		return err
	}

	logger := log.New()

	b, err := bot.NewBot(bot.Config{
		Token:   token,
		Timeout: 60,
		Debug:   false,
		Logger:  logger,
	})

	if err != nil {
		return err
	}

	cm := commands.NewManager(b, logger)

	// spinup bot
	return b.Serve(cm)
}

// TODO: outsource into run func to catch errors easier
func main() {
	// A single place to exit due to errors.
	if err := run(); err != nil {
		log.Fatalf("%v", err)
	}
}

// TODO: temp, later move out as a docker envvar?
func readToken() (string, error) {
	token, err := ioutil.ReadFile("./.token")
	return string(token), err
}
