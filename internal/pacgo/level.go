package pacgo

import (
	"bufio"
	"math/rand"
	"os"

	"github.com/buger/goterm"
)

const (
	ghostChar  rune = 'G'
	playerChar rune = 'P'
	wallChar   rune = '#'
	dotChar    rune = '.'
)

type sprite struct {
	row int
	col int
}

// level represents a pacgo game level.
type level struct {
	maze    []string
	player  *sprite
	ghosts  []*sprite
	score   int
	numDots int
	lives   int
}

// loadLevel loads a level from the passed filepath.
func loadLevel(filepath string) (*level, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	l := &level{
		lives: 1,
	}
	var maze []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, line)
	}
	l.maze = maze

	for row, line := range l.maze {
		for col, char := range line {
			switch char {
			case playerChar:
				l.player = &sprite{row, col}
			case ghostChar:
				l.ghosts = append(l.ghosts, &sprite{row, col})
			case dotChar:
				l.numDots++
			}
		}
	}

	return l, nil
}

func (l *level) calculateMove(curRow, curCol int, dir command) (newRow, newCol int) {
	newRow, newCol = curRow, curCol

	switch dir {
	case up:
		newRow = newRow - 1
		if newRow < 0 {
			newRow = len(l.maze) - 1
		}
	case down:
		newRow = newRow + 1
		if newRow == len(l.maze) {
			newRow = 0
		}
	case right:
		newCol = newCol + 1
		if newCol == len(l.maze[0]) {
			newCol = 0
		}
	case left:
		newCol = newCol - 1
		if newCol < 0 {
			newCol = len(l.maze[0]) - 1
		}
	}

	if rune(l.maze[newRow][newCol]) == wallChar {
		newRow = curRow
		newCol = curCol
	}

	return
}

func (l *level) MoveGhosts() {
	for _, g := range l.ghosts {
		dir := randomDirection()
		g.row, g.col = l.calculateMove(g.row, g.col, dir)
	}
}

func (l *level) MovePlayer(dir command) {
	l.player.row, l.player.col = l.calculateMove(l.player.row, l.player.col, dir)

	switch rune(l.maze[l.player.row][l.player.col]) {
	case dotChar:
		l.numDots--
		l.score++

		l.maze[l.player.row] = l.maze[l.player.row][0:l.player.col] + " " + l.maze[l.player.row][l.player.col+1:]
	}
}

// printScreen prints the level to StdOut.
func (l *level) PrintScreen() {
	clearScreen()
	for _, line := range l.maze {
		for _, char := range line {
			switch char {
			case wallChar:
				fallthrough
			case dotChar:
				goterm.Printf("%c", char)
			default:
				goterm.Print(" ")
			}
		}
		goterm.Println()
	}

	moveCursor(l.player.row, l.player.col)
	goterm.Print(string(playerChar))

	for _, g := range l.ghosts {
		moveCursor(g.row, g.col)
		goterm.Print(string(ghostChar))
	}

	moveCursor(len(l.maze)+1, 0)
	goterm.Println("Score: ", l.score, "\tLives: ", l.lives)
	goterm.Flush()
}

var dirLookupMap = map[int]command{
	0: up,
	1: down,
	2: right,
	3: left,
}

func randomDirection() command {
	dir := rand.Intn(4)

	return dirLookupMap[dir]
}
