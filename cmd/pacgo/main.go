package main

import (
	"log"

	"github.com/Jac0bDeal/pacgo/internal/pacgo"
)

func main() {
	maze, err := pacgo.LoadMaze("maze.txt")
	if err != nil {
		log.Fatalf("failed to load maze from file: %v", err)
	}

	for {
		maze.PrintScreen()

		break
	}
}
