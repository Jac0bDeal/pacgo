package pacgo

import "github.com/buger/goterm"

// moveCursor wraps goterm.MoveCursor to achieve zero indexing
// and maps slice indexing to goterm indexing.
func moveCursor(row, col int) {
	goterm.MoveCursor(col+1, row+1)
}

// clearScreen wipes the screen by returning the cursor to the origin.
func clearScreen() {
	moveCursor(0, 0)
}
