package pacgo

const (
	ghostChar  rune = 'G'
	playerChar rune = 'P'
	wallChar   rune = '#'
)

type sprite struct {
	row int
	col int
}
