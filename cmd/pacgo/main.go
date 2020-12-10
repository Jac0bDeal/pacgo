package main

import (
	"log"

	"github.com/Jac0bDeal/pacgo/internal/pacgo"
)

func main() {
	app, err := pacgo.New("configs/config-emoji.json", "levels/level-1.txt")
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}

	err = app.Run()
	if err != nil {
		log.Fatalf("encountered error while running: %v", err)
	}
}
