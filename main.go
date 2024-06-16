package main

import (
	"database/sql"
	"log"
	"os"
	"simple-rest/router"
	"simple-rest/utils/logger"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	logger := logger.NewLogger(log.Default())

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logger.PrintError(err)
	}

	log.SetOutput(file)

	if err := godotenv.Load(); err != nil {
		logger.PrintError(err)
	}

	connStr := os.Getenv("POSTGRES_PATH")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.PrintError(err)
	}

	r := router.NewRouter(logger, db)

	if err := r.Run(); err != nil {
		logger.PrintError(err)
	}
}
