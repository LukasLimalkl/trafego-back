package main

import (
	"github.com/LukasLimalkl/trafego-back/config"
	"github.com/LukasLimalkl/trafego-back/internal/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")
	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initalization error: %v", err)
		return
	}

	// Initialize Routes
	router.Initialize()

}
