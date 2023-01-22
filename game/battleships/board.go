package battleships

import (
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

func (b *Board) tileAtWorldPos(x, y int) *Tile {
	var foundTile *Tile = nil

	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			tile := b.tileAt(i, j)

			if tile.x*tileSize+xOffset < x && ((tile.x+1)*tileSize+xOffset) > x && tile.y*tileSize+yOffset < y && ((tile.y+1)*tileSize+yOffset) > y {
				foundTile = tile
				break
			}
		}
	}

	return foundTile
}

func (b *Board) tilesForShipPlacement(x, y int, ship *Ship) []*Tile {
	tiles := []*Tile{}

	for i := 0; i < ship.length; i++ {
		if x >= b.size || y >= b.size {
			return []*Tile{}
		}

		tiles = append(tiles, b.tileAt(x, y))

		if ship.rotation == LeftRotation || ship.rotation == RightRotation {
			x += 1
		} else {
			y += 1
		}
	}

	return tiles
}

func (b *Board) placeShip(ship *Ship) {
	for _, tile := range ship.placedAtTiles {
		tile.state = EmptyState
	}

	startTile := b.tileAtWorldPos(ship.globalX, ship.globalY)
	tiles := b.tilesForShipPlacement(startTile.x, startTile.y, ship)

	for idx, tile := range tiles {
		if (ship.rotation == LeftRotation || ship.rotation == RightRotation) && idx == 0 {
			tile.state = ShipFrontState
		} else if (ship.rotation == RightRotation || ship.rotation == DownRotation) && idx == ship.length-1 {
			tile.state = ShipFrontState
		} else {
			tile.state = ShipState
		}
	}

	ship.placedAtTiles = tiles
	ship.globalX = startTile.x*tileSize + xOffset
	ship.globalY = startTile.y*tileSize + yOffset
}

func (b *Board) placeBomb(x, y int) {
	b.tileAt(x, y).state = BombState
}

func (b *Board) calculatePossibleMovesForShip(ship *Ship) {
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
				if ship.pos[i].x+move.xOffset >= b.size || ship.pos[i].y+move.yOffset >= b.size ||
					ship.pos[i].x+move.xOffset < 0 || ship.pos[i].y+move.yOffset < 0 {
					tileClear = false
					break
				}
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
	tileSize               = 60
	tileSizeWithGridOffset = 58
	xOffset                = 50
	yOffset                = 50
)

func (b *Board) Draw(drawerImage *ebiten.Image) {
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			tile := b.tileAt(i, j)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((i*shipTileSize)+xOffset+1), float64((j*shipTileSize)+yOffset+1))
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

func (b *Board) SetHighlight(ship *Ship) {
	tile := b.tileAtWorldPos(ship.globalX, ship.globalY)
	ship.isLegalPlacement = false

	if tile != nil {
		highlightedTiles := b.tilesForShipPlacement(tile.x, tile.y, ship)
		for _, tile := range highlightedTiles {
			tile.isHovered = true
			ship.isLegalPlacement = true
		}
	}
}

func (b *Board) ResetHighlight() {
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			b.tileAt(i, j).isHovered = false
		}
	}
}
