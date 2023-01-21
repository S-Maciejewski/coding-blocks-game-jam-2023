package battleships

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	ScreenWidth  = 1000
	ScreenHeight = 1000
)

type Game struct {
	board    *Board
	ships    []*Ship
	drawer   ShipDrawer
	heldShip *Ship
	shooter  *Shooter
}

func NewGame() *Game {
	ships := GenerateShips(1, 1, 1, 1)
	xOffset := drawerOffset
	yOffset := ScreenHeight - (3 * shipTileSize)
	for _, ship := range ships {

		if (ship.length*shipTileSize)+xOffset >= (ScreenWidth - drawerOffset) {
			xOffset = drawerOffset
			yOffset += shipTileSize + 10
		}

		ship.globalX = xOffset
		ship.globalY = yOffset
		ship.previousPosX = xOffset
		ship.previousPosY = yOffset
		ship.rotation = LeftRotation
		ship.previousRotation = LeftRotation

		xOffset += (ship.length * shipTileSize) + drawerOffset
	}

	g := &Game{
		board:   NewBoard(10),
		ships:   ships,
		drawer:  *NewDrawer(ships),
		shooter: NewShooter(),
	}
	return g
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.heldShip = g.shipAt(ebiten.CursorPosition())
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) && g.heldShip != nil {
		g.heldShip.rotate()
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && g.heldShip != nil {
		if g.heldShip.isLegalPlacement {
			g.board.placeShip(g.heldShip)
		} else {
			g.heldShip.ResetToPreviousPosition()
		}
		g.heldShip = nil
	}

	g.board.ResetHighlight()
	if g.heldShip != nil {
		x, y := ebiten.CursorPosition()
		g.heldShip.globalX = x
		g.heldShip.globalY = y
		g.heldShip.isSelected = true
		g.board.SetHighlight(g.heldShip)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
	g.board.Draw(screen)
	g.drawer.Draw(screen)
}

func (g *Game) shipAt(x, y int) *Ship {
	for i := len(g.ships) - 1; i >= 0; i-- {
		s := g.ships[i]
		if s.In(x, y) {
			return s
		}
	}
	return nil
}
