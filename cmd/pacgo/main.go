package main

import (
	"flag"
	"log"

	"github.com/Jac0bDeal/pacgo/internal/pacgo"
)

var (
	configFile = flag.String("config-file", "configs/config-emoji.json", "path to config file")
	levelFile  = flag.String("level-file", "levels/level-1.txt", "path to level file")
)

func init() {
	flag.Parse()
}

func main() {
	app, err := pacgo.New(*configFile, *levelFile)
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}

	err = app.Run()
	if err != nil {
		log.Fatalf("encountered error while running: %v", err)
	}
}
