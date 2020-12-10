package pacgo

import (
	"bufio"
	"math/rand"
	"os"
)

const (
	dotChar     rune = '.'
	ghostChar   rune = 'G'
	playerChar  rune = 'P'
	powerUpChar rune = 'X'
	wallChar    rune = '#'
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
