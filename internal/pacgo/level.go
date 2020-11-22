package pacgo

import (
	"bufio"
	"fmt"
	"os"

	"github.com/danicat/simpleansi"
)

// level represents a pacgo game level.
type level []string

// loadLevel loads a level from the passed filepath.
func loadLevel(filepath string) (level, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	l := level{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		l = append(l, line)
	}

	return l, nil
}

// printScreen prints the level to StdOut.
func (l level) printScreen() {
	simpleansi.ClearScreen()
	for _, line := range l {
		fmt.Println(line)
	}
}
