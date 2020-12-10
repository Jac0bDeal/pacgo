package pacgo

import (
	"log"
	"os"
	"os/exec"
	"time"

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
	cfg *Config
	*level
}

// New constructs and returns an App from the passed parameters.
func New(configFile, levelFile string) (*App, error) {
	app := new(App)

	c, err := loadConfig(configFile)
	if err != nil {
		return nil, err
	}
	app.cfg = c

	l, err := loadLevel(levelFile)
	if err != nil {
		return nil, err
	}
	app.level = l

	return app, nil
}

// Run initializes the App and then executes the main game loop.
func (a *App) Run() error {
	//initialize terminal and game
	a.initialize()
	defer a.cleanup()
	goterm.Clear()

	// process input
	input := make(chan command)
	go func(ch chan<- command) {
		for {
			input, err := a.readInput()
			if err != nil {
				log.Println("error reading input: ", err)
				ch <- esc
			}
			ch <- input
		}
	}(input)

	for {
		// update screen
		a.printScreen()

		// process movement
		select {
		case inp := <-input:
			if inp == esc {
				a.lives = 0
			}
			a.movePlayer(inp)
		default:
		}

		a.moveGhosts()

		// process collisions
		for _, g := range a.ghosts {
			if *a.player == *g {
				a.lives = 0
			}
		}

		// check game over
		if a.numDots == 0 || a.lives <= 0 {
			a.moveCursor(a.player.row, a.player.col)
			goterm.Print(a.cfg.DeathSprite)
			a.moveCursor(len(a.maze)+1, 0)
			goterm.Println("Score: ", a.score, "\t Lives: ", a.lives)
			goterm.Println("Game Over!")
			goterm.Flush()
			break
		}

		// timing delay
		time.Sleep(300 * time.Millisecond)
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

// moveCursor wraps goterm.MoveCursor to achieve zero indexing
// and maps slice indexing to goterm indexing.
func (a *App) moveCursor(row, col int) {
	if a.cfg.UseEmoji {
		goterm.MoveCursor(col*2+1, row+1)
		return
	}
	goterm.MoveCursor(col+1, row+1)
}

// clearScreen wipes the screen by returning the cursor to the origin.
func (a *App) clearScreen() {
	a.moveCursor(0, 0)
}

func (a *App) moveGhosts() {
	for _, g := range a.ghosts {
		go func(ghost *sprite) {
			dir := randomDirection()
			ghost.row, ghost.col = a.calculateMove(ghost.row, ghost.col, dir)
		}(g)
	}
}

func (a *App) movePlayer(dir command) {
	a.player.row, a.player.col = a.calculateMove(a.player.row, a.player.col, dir)

	removeDot := func(row, col int) {
		a.maze[row] = a.maze[row][0:col] + " " + a.maze[row][col+1:]
	}

	switch rune(a.maze[a.player.row][a.player.col]) {
	case dotChar:
		a.numDots--
		a.score++
		removeDot(a.player.row, a.player.col)
	case powerUpChar:
		a.score += 10
		removeDot(a.player.row, a.player.col)
	}
}

// printScreen prints the level to StdOut.
func (a *App) printScreen() {
	a.clearScreen()
	for _, line := range a.maze {
		for _, char := range line {
			switch char {
			case wallChar:
				goterm.Print(goterm.Background(a.cfg.WallSprite, goterm.BLUE))
			case dotChar:
				goterm.Print(a.cfg.DotSprite)
			case powerUpChar:
				goterm.Print(a.cfg.PillSprite)
			default:
				goterm.Print(a.cfg.SpaceSprite)
			}
		}
		goterm.Println()
	}

	a.moveCursor(a.player.row, a.player.col)
	goterm.Print(a.cfg.PlayerSprite)

	for _, g := range a.ghosts {
		a.moveCursor(g.row, g.col)
		goterm.Print(a.cfg.GhostSprite)
	}

	a.moveCursor(len(a.maze)+1, 0)
	goterm.Println("Score: ", a.score, "\tLives: ", a.lives)
	goterm.Flush()
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
