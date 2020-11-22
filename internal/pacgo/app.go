package pacgo

import (
	"log"
	"os"
	"os/exec"
)

const (
	Esc = "ESC"
)

// App contains the main logic and data structures used to run pacgo.
type App struct {
	level
}

// New constructs and returns an App from the passed parameters.
func New(filepath string) (*App, error) {
	level, err := loadLevel(filepath)
	if err != nil {
		return nil, err
	}

	app := &App{
		level: level,
	}

	return app, nil
}

// Run initializes the App and then executes the main game loop.
func (a *App) Run() error {
	//initialize game
	a.initialize()
	defer a.cleanup()

	for {
		// update screen
		a.printScreen()

		// process input
		input, err := a.readInput()
		if err != nil {
			return err
		}

		if input == Esc {
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

func (*App) readInput() (string, error) {
	buffer := make([]byte, 100)

	cnt, err := os.Stdin.Read(buffer)
	if err != nil {
		return "", err
	}

	if cnt == 1 && buffer[0] == 0x1b {
		return Esc, nil
	}

	return "", nil
}
