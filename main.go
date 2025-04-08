package main

import (
	"github.com/guilhermecosales/workly/config"
	"github.com/guilhermecosales/workly/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()

	if err != nil {
		logger.Debugf("Error initializing config: %v", err)
		return
	}

	router.Initialize()
}
