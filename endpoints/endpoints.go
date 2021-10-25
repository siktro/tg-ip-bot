package endpoints

import (
	"github.com/siktro/tg-ip-bot/mux"
	"github.com/sirupsen/logrus"
)

type EndpointManager struct {
	logger *logrus.Logger
}

func NewManager(logger *logrus.Logger) *EndpointManager {
	m := &EndpointManager{
		logger: logger,
	}
	return m
}

func (cm *EndpointManager) StartCommand() *mux.EndpointEntry {
	return &mux.EndpointEntry{
		Callback:    cm.start,
		Description: "Prints a welcome message.",
		Group:       "none",
	}
}

func (cm *EndpointManager) HelpCommand() *mux.EndpointEntry {
	return &mux.EndpointEntry{
		Callback:    cm.start,
		Description: "Prints a help message with all commands.",
		Group:       "none",
	}
}
