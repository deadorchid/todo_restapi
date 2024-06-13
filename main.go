package main

import (
	"log"
	"os"
	"simple-rest/router"
)

func main() {
	r := router.NewRouter(log.Default())

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
