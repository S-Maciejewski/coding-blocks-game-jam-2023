package battleships

import (
	"image/color"
	"math/rand"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	ScreenWidth  = 700
	ScreenHeight = 1000
)

type Game struct {
	board             *Board
	ships             []*Ship
	drawer            ShipDrawer
	heldShip          *Ship
	shooter           *Shooter
	areAllShipsPlaced bool
	score             int
	highScore         int
	isFinished        bool
}

func NewGame() *Game {
	ships := GenerateShips(2, 1, 1, 1)
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
		g.board.clearLegalMoves()
		g.heldShip.rotate()
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && g.heldShip != nil {
		if g.heldShip.isLegalPlacement {
			g.board.placeShip(g.heldShip)
			changeAllShipsPlaced(g)
			if g.areAllShipsPlaced {
				for i := 0; i < 3; i++ {
					bomb := g.shooter.GetNewBomb(g.board)
					g.board.placeBomb(&bomb)
				}
				g.board.reduceBombLifetimes()
				g.board.checkBombHits(g.ships)

				isGameFinished := checkIfGameFinished(g)

				if isGameFinished {
					if g.score > g.highScore {
						g.highScore = g.score
					}
					g.isFinished = true
				} else {
					g.score++
				}
			}
		} else {
			g.board.clearLegalMoves()
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
		if len(g.heldShip.placedAtTiles) > 0 {
			g.board.calculatePossibleMovesForShip(g.heldShip)
			g.board.showLegalMoves(g.heldShip)
		}
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyR) {
		ng := NewGame()
		g.ships = ng.ships
		g.score = 0
		g.shooter = ng.shooter
		g.drawer = ng.drawer
		g.board = ng.board
		g.isFinished = false
	}

	return nil
}

func checkIfGameFinished(g *Game) bool {
	isGameFinished := true
	for _, ship := range g.ships {
		if !ship.isDestroyed {
			isGameFinished = false
			break
		}
	}
	return isGameFinished
}

func changeAllShipsPlaced(g *Game) {
	if !g.areAllShipsPlaced {
		g.areAllShipsPlaced = true
		for _, ship := range g.ships {
			if len(ship.placedAtTiles) == 0 {
				g.areAllShipsPlaced = false
				break
			}
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
	f := mplusBigFont
	text.Draw(screen, "Score: "+strconv.Itoa(g.score), f, 50, (tileSize*g.board.size)+yOffset+60, color.Black)
	text.Draw(screen, "Highscore: "+strconv.Itoa(g.highScore), f, 50, (tileSize*g.board.size)+yOffset+120, color.Black)
	if g.isFinished {
		text.Draw(screen, "Click R to restart", f, 50, (tileSize*g.board.size)+yOffset+180, color.Black)
	}
	g.board.Draw(screen)
	g.drawer.Draw(screen)
	g.board.DrawOverlay(screen)
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
