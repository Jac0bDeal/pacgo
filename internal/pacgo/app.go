package pacgo

import (
	"log"
	"os"
	"os/exec"

	"github.com/buger/goterm"
)

type command string

const (
	esc   command = "ESC"
	up    command = "UP"
	down  command = "DOWN"
	left  command = "LEFT"
	right command = "Right"
)

// App contains the main logic and data structures used to run pacgo.
type App struct {
	level
}

// New constructs and returns an App from the passed parameters.
func New(filepath string) (*App, error) {
	app := new(App)

	l, err := loadLevel(filepath)
	if err != nil {
		return nil, err
	}
	app.level = l

	return app, nil
}

// Run initializes the App and then executes the main game loop.
func (a *App) Run() error {
	//initialize game
	a.initialize()
	defer a.cleanup()

	goterm.Clear()
	for {
		// update screen
		a.PrintScreen()

		// process input
		input, err := a.readInput()
		if err != nil {
			return err
		}

		a.MovePlayer(input)
		a.MoveGhosts()

		if input == esc {
			log.Println("received terminate signal, shutting down...")
			break
		}
	}

	return nil
}

// initialize changes the terminal to cbreak mode.
func (*App) initialize() {
	cbTerm := exec.Command("stty", "cbreak", "-echo")
	cbTerm.Stdin = os.Stdin

	err := cbTerm.Run()
	if err != nil {
		log.Fatalf("unable to activate cbreak mode: %v", err)
	}
}

// cleanup changes the terminal back to cooked mode.
func (*App) cleanup() {
	cookedTerm := exec.Command("stty", "-cbreak", "echo")
	cookedTerm.Stdin = os.Stdin

	err := cookedTerm.Run()
	if err != nil {
		log.Fatalf("unable to restore cooked mode: %v", err)
	}
}

func (*App) readInput() (command, error) {
	buffer := make([]byte, 100)

	cnt, err := os.Stdin.Read(buffer)
	if err != nil {
		return "", err
	}

	if cnt == 1 && buffer[0] == 0x1b {
		return esc, nil
	} else if cnt >= 3 {
		if buffer[0] == 0x1b && buffer[1] == '[' {
			switch buffer[2] {
			case 'A':
				return up, nil
			case 'B':
				return down, nil
			case 'C':
				return right, nil
			case 'D':
				return left, nil
			}
		}
	}

	return "", nil
}
