package main

import (
	"io/ioutil"
	"os"
	"regexp"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/siktro/tg-ip-bot/internal/bot"
	"github.com/siktro/tg-ip-bot/internal/commands"
	log "github.com/sirupsen/logrus"
)

// TODO: outsource into run func to catch errors easier
func main() {
	if err := run(); err != nil {
		log.Fatalf("Initialization error: %v", err)
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

	logger := log.New()

	// TODO: Setup logging goroutins error channel

	// === Start DB.

	// === Start bot.

	b, err := bot.NewBot(bot.Config{
		Token:        config.tgToken,
		Debug:        false,
		Logger:       logger,
		WorkersLimit: 1,
	})

	if err != nil {
		return err
	}

	cm := commands.NewManager(b, logger)

	b.Handle("start", cm.Start)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	return b.ListenAndServe(updateConfig)
}

// TODO: remove later
func loadEnvsFromFile(path string) error {
	token, err := ioutil.ReadFile("./.env")
	if err != nil {
		return err
	}

	r, err := regexp.Compile(`(\w+)\s*=\s*(\S+)`)
	if err != nil {
		return err
	}

	matches := r.FindAllStringSubmatch(string(token), -1)
	for _, m := range matches {
		os.Setenv(m[1], m[2])
	}

	return nil
}
