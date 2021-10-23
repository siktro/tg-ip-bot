package main

import (
	"io/ioutil"
	"os"
	"regexp"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/siktro/tg-ip-bot/internal/bot"
	log "github.com/sirupsen/logrus"
)

// TODO: outsource into run func to catch errors easier
func main() {
	if err := run(); err != nil {
		log.Fatalf("%v", err)
	}
}

func run() error {
	// get sysenvs, setup from docker?
	// check db

	// check bot

	// spinup command handler

	// gracefuly shutdown? if any
	err := loadEnvsFromFile("./.env")
	if err != nil {
		return err
	}

	config := struct {
		tgToken      string
		ipstackToken string
	}{
		tgToken:      os.Getenv("TG_TOKEN"),
		ipstackToken: os.Getenv("IPSTACK_TOKEN"),
	}

	// logger := log.New()

	// === Start bot.

	b, err := bot.NewBot(bot.Config{
		Token: config.tgToken,
		Debug: true,
	})

	if err != nil {
		return err
	}

	b.Handle("/start", func(m *tgbotapi.Message) {

	})

	updateConfig := tgbotapi.NewUpdate(0)
	return b.ListenAndServe(updateConfig)
}

// TODO: remove later
func loadEnvsFromFile(path string) error {
	token, err := ioutil.ReadFile("./.env")
	if err != nil {
		return err
	}

	r, err := regexp.Compile(`(\w+)\s*=\s*(\w+)`)
	if err != nil {
		return err
	}

	matches := r.FindAllStringSubmatch(string(token), -1)
	for _, m := range matches {
		os.Setenv(m[1], m[2])
	}

	return nil
}
