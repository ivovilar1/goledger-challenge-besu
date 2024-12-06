package main

import (
	"github.com/ivovilar1/goledger-challenge-besu/config"
	"github.com/ivovilar1/goledger-challenge-besu/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.InitDatabase()
	if err != nil {
		logger.Err(err)
		return
	}

	router.RunServerAndRoutes()

}
