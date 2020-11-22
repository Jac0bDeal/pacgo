package pacgo

import (
	"fmt"
	"os"
	"bufio"
)

type Maze []string

func LoadMaze(filepath string) (Maze, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	maze := Maze{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, line)
	}

	return maze, nil
}

func (m Maze) PrintScreen() {
	for _, line := range m {
		fmt.Println(line)
	}
}
