package main

import (
	"log"

	"github.com/Jac0bDeal/pacgo/internal/pacgo"
)

func main() {
	app, err := pacgo.New("level.txt")
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}

	err = app.Run()
	if err != nil {
		log.Fatalf("encountered error while running: %v", err)
	}
	log.Println("done")
}
