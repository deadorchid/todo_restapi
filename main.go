package main

import (
	"log"
	"os"
	"simple-rest/router"
	"simple-rest/utils/logger"
)

func main() {
	logger := logger.NewLogger(log.Default())
	r := router.NewRouter(logger)

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logger.PrintError(err)
	}

	log.SetOutput(file)

	if err := r.Run(); err != nil {
		logger.PrintError(err)
	}
}
