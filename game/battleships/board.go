package battleships

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Board struct {
	size  int
	tiles [][]Tile
}

func NewBoard(size int) *Board {
	var tiles [][]Tile
	for i := 0; i < size; i++ {
		var tileRow []Tile
		for j := 0; j < size; j++ {
			tileRow = append(tileRow, *NewTile(i, j))
		}
		tiles = append(tiles, tileRow)
	}
	return &Board{
		size:  size,
		tiles: tiles,
	}
}

func (b *Board) tileAt(x, y int) *Tile {
	return &b.tiles[x][y]
}

func (b *Board) placeShip(ship *Ship) {
	for _, pos := range ship.pos {
		if pos.isFront {
			b.tileAt(pos.x, pos.y).state = ShipFrontState
		} else {
			b.tileAt(pos.x, pos.y).state = ShipState
		}
	}
}

func (b *Board) placeBomb(x, y int) {
	b.tileAt(x, y).state = BombState
}

func (b *Board) setShipWithCalculatedMoves(ship *Ship) {
	for _, move := range ship.moves {
		//	for each move calculate if it's possible and set isPossible value
		//	possible move is when ship is not out of board and there is EmptyState tile in places where ship would be after move

		shipFrontPos := ship.pos[0]
		if !shipFrontPos.isFront {
			shipFrontPos = ship.pos[ship.length-1]
		}

		//	ship is not out of board
		if shipFrontPos.x+move.xOffset >= 0 &&
			shipFrontPos.x+move.xOffset < b.size &&
			shipFrontPos.y+move.yOffset >= 0 &&
			shipFrontPos.y+move.yOffset < b.size {
			//	there is EmptyState tile in places where ship will be after move
			tileClear := true
			for i := 0; i < ship.length; i++ {
				if b.tileAt(ship.pos[i].x+move.xOffset, ship.pos[i].y+move.yOffset).state != EmptyState {
					tileClear = false
					break
				}
			}
			move.isPossible = tileClear
		}
	}
}

func (b *Board) showLegalMoves(ship *Ship) {
	for _, move := range ship.moves {
		if move.isPossible {
			// set LegalMoveState for each tile where ship would be after move unless it'll be in place where ship is already
			for i := 0; i < ship.length; i++ {
				if b.tileAt(ship.pos[i].x+move.xOffset, ship.pos[i].y+move.yOffset).state != ShipState &&
					b.tileAt(ship.pos[i].x+move.xOffset, ship.pos[i].y+move.yOffset).state != ShipFrontState {
					b.tileAt(ship.pos[i].x+move.xOffset, ship.pos[i].y+move.yOffset).state = LegalMoveState
				}
			}
		}
	}
}

const (
	tileSize = 60
	xOffset  = 50
	yOffset  = 50
)

var (
	emptyStateImage        = ebiten.NewImage(shipTileSize, shipTileSize)
	shipStateImage         = ebiten.NewImage(shipTileSize, shipTileSize)
	shipFrontStateImage    = ebiten.NewImage(shipTileSize, shipTileSize)
	shipHitStateImage      = ebiten.NewImage(shipTileSize, shipTileSize)
	shipSunkStateImage     = ebiten.NewImage(shipTileSize, shipTileSize)
	shipFrontHitStateImage = ebiten.NewImage(shipTileSize, shipTileSize)
	bombStateImage         = ebiten.NewImage(shipTileSize, shipTileSize)
	legalMoveImage         = ebiten.NewImage(shipTileSize, shipTileSize)
	illegalMoveImage       = ebiten.NewImage(shipTileSize, shipTileSize)
	hoverImage             = ebiten.NewImage(shipTileSize, shipTileSize)
)

func init() {
	emptyStateImage.Fill(color.RGBA{0x49, 0x80, 0xff, 0xff})
	shipStateImage.Fill(color.RGBA{0x58, 0x58, 0x58, 0xff})
	shipFrontStateImage.Fill(color.RGBA{0x99, 0x4c, 0x00, 0xff})
	shipHitStateImage.Fill(color.RGBA{0xcc, 0x66, 0x00, 0xff})
	shipSunkStateImage.Fill(color.RGBA{0x66, 0x00, 0x00, 0xff})
	shipFrontHitStateImage.Fill(color.RGBA{0x66, 0x00, 0x00, 0xff})
	bombStateImage.Fill(color.RGBA{0x66, 0x33, 0x00, 0xff})
	legalMoveImage.Fill(color.RGBA{0x66, 0xff, 0xb2, 0xff})
	illegalMoveImage.Fill(color.RGBA{0xff, 0x66, 0x66, 0xff})
	hoverImage.Fill(color.RGBA{0x00, 0x00, 0x00, 0x7f})
}

func (b *Board) Draw(drawerImage *ebiten.Image) {
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			tile := b.tileAt(i, j)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((i*shipTileSize)+xOffset), float64((j*shipTileSize)+xOffset))
			switch tile.state {
			case EmptyState:
				drawerImage.DrawImage(emptyStateImage, op)
			case ShipState:
				drawerImage.DrawImage(shipStateImage, op)
			case ShipFrontState:
				drawerImage.DrawImage(shipFrontStateImage, op)
			case ShipHitState:
				drawerImage.DrawImage(shipHitStateImage, op)
			case ShipSunkState:
				drawerImage.DrawImage(shipSunkStateImage, op)
			case ShipFrontHitState:
				drawerImage.DrawImage(shipFrontHitStateImage, op)
			case BombState:
				drawerImage.DrawImage(bombStateImage, op)
			case LegalMoveState:
				drawerImage.DrawImage(legalMoveImage, op)
			case IllegalMoveState:
				drawerImage.DrawImage(illegalMoveImage, op)
			}

			if tile.isHovered {
				drawerImage.DrawImage(hoverImage, op)
			}
		}
	}
}
