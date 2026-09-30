package main

import (
	"github.com/GustavoEmanuel901/goopportunities/config"
	"github.com/GustavoEmanuel901/goopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	// Inicializa config
	logger = config.GetLogger("APP: ")

	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}

	logger.Infof("Config initialized successfully")

	router.Inicialiaze()
}
