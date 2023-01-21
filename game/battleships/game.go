package battleships

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	ScreenWidth  = 1000
	ScreenHeight = 1000
)

type Game struct {
	board  *Board
	ships  []*Ship
	drawer ShipDrawer
}

func NewGame() *Game {
	ships := GenerateShips(1, 1, 1, 1)
	g := &Game{
		board:  NewBoard(10),
		ships:  ships,
		drawer: *NewDrawer(ships),
	}
	return g
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	// g.input.Update()
	//TODO: update board here
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
	g.drawer.Draw(screen)
	g.board.Draw(screen)
}
