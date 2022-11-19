package main

import (
	"github.com/hinotora/go-auth-service/pkg/config"
	"github.com/hinotora/go-auth-service/pkg/logger"
	"github.com/hinotora/go-auth-service/pkg/router"
)

func main() {
	err := config.Load()

	if err != nil {
		logger.Logger.Fatal(err)
	}

	err = router.Run()

	logger.Logger.Fatal(err)
}
