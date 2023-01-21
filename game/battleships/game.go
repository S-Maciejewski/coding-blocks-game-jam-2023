package battleships

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	ScreenWidth  = 1000
	ScreenHeight = 1000
)

type Game struct {
	input *Input
	board *Board
	ships []*Ship
}

func NewGame() *Game {
	g := &Game{
		input: NewInput(),
		board: NewBoard(10),
		ships: GenerateShips(1, 1, 1, 1),
	}
	return g
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
