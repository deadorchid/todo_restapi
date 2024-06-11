package main

import (
	"simple-rest/router"
)

func main() {
	r := router.NewRouter()
	if err := r.Run(); err != nil {
		panic("Something gone wrong...")
	}
}
