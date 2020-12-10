package pacgo

import (
	"bufio"
	"os"

	"github.com/buger/goterm"
)

// level represents a pacgo game level.
type level struct {
	maze   []string
	player *sprite
}

// loadLevel loads a level from the passed filepath.
func loadLevel(filepath string) (level, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return level{}, err
	}
	defer file.Close()

	l := level{}
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
			case 'P':
				l.player = &sprite{row, col}
			}
		}
	}

	return l, nil
}

func (l level) calculateMove(curRow, curCol int, dir command) (newRow, newCol int) {
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

	if l.maze[newRow][newCol] == '#' {
		newRow = curRow
		newCol = curCol
	}

	return
}

func (l level) movePlayer(dir command) {
	l.player.row, l.player.col = l.calculateMove(l.player.row, l.player.col, dir)
}

// printScreen prints the level to StdOut.
func (l level) printScreen() {
	clearScreen()
	for _, line := range l.maze {
		for _, char := range line {
			switch char {
			case '#':
				goterm.Printf("%c", char)
			default:
				goterm.Print(" ")
			}
		}
		goterm.Println()
	}

	moveCursor(l.player.row, l.player.col)
	goterm.Print("P")
	moveCursor(len(l.maze)+1, 0)

	goterm.Flush()
}
