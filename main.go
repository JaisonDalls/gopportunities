package main

import (
	"github.com/JaisonDalls/gopportunities/config"
	"github.com/JaisonDalls/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//Initializing the db config
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	//Initializing the routes
	router.Initialize()
}
