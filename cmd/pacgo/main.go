package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/Jac0bDeal/pacgo/internal/pacgo"
)

func main() {
	//initialize game
	initialize()
	defer cleanup()

	// load resources
	maze, err := pacgo.LoadMaze("maze.txt")
	if err != nil {
		log.Fatalf("failed to load maze from file: %v", err)
	}

	for {
		// update screen
		maze.PrintScreen()

		break
	}
}

func initialize() {
	cbTerm := exec.Command("stty", "cbreak", "-echo")
	cbTerm.Stdin = os.Stdin

	err := cbTerm.Run()
	if err != nil {
		log.Fatalf("unable to activate cbreak mode: %v", err)
	}
}

func cleanup() {
	cookedTerm := exec.Command("stty", "-cbreak", "echo")
	cookedTerm.Stdin = os.Stdin

	err := cookedTerm.Run()
	if err != nil {
		log.Fatalf("unable to restore cooked mode: %v", err)
	}
}
