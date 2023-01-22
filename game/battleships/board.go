package battleships

import (
	"fmt"
	"image/color"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	mplusSmallFont  font.Face
	mplusNormalFont font.Face
	mplusBigFont    font.Face
)

type Board struct {
	size  int
	tiles [][]Tile
	bombs []*Bomb
}

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusSmallFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    32,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}
	mplusBigFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}
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

func (b *Board) clearLegalMoves() {
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			tile := b.tileAt(i, j)
			if tile.state == LegalMoveState {
				tile.state = EmptyState
			}
		}
	}
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
	gridPos := []ShipPosition{}
	for _, tile := range tiles {
		gridPos = append(gridPos, ShipPosition{x: tile.x, y: tile.y, isFront: tile.state == ShipFrontState})
	}
	ship.gridPos = gridPos
	ship.globalX = startTile.x*tileSize + xOffset
	ship.globalY = startTile.y*tileSize + yOffset

	b.clearLegalMoves()
}

func (b *Board) placeBomb(bomb *Bomb) {
	b.bombs = append(b.bombs, bomb)
	b.tileAt(bomb.x, bomb.y).state = BombState
}

func (b *Board) getBombAtTile(x, y int) *Bomb {
	for _, bomb := range b.bombs {
		if bomb.x == x && bomb.y == y {
			return bomb
		}
	}

	return nil
}

func (b *Board) reduceBombLifetimes() {
	var bombs []*Bomb
	for _, bomb := range b.bombs {
		bomb.turnsToLive--

		if bomb.turnsToLive == 0 {
			b.tileAt(bomb.x, bomb.y).state = BombExplodedState
		}

		if bomb.turnsToLive == -5 {
			b.tileAt(bomb.x, bomb.y).state = EmptyState
		} else {
			bombs = append(bombs, bomb)
		}
	}
	b.bombs = bombs
}

func (b *Board) calculatePossibleMovesForShip(ship *Ship) {
	newMoves := []Move{}
	for _, move := range ship.moves {
		//	for each move calculate if it's possible and set isPossible value
		//	possible move is when ship is not out of board and there is EmptyState tile in places where ship would be after move

		//shipFrontPos := ship.gridPos[0]
		//if !shipFrontPos.isFront {
		//	shipFrontPos = ship.gridPos[ship.length-1]
		//}
		if !ship.gridPos[0].isFront {
			fmt.Println("ship.gridPos", ship.gridPos)
			//	reverse ship.gridPos list
			for i, j := 0, len(ship.gridPos)-1; i < j; i, j = i+1, j-1 {
				ship.gridPos[i], ship.gridPos[j] = ship.gridPos[j], ship.gridPos[i]
			}
			fmt.Println("ship.gridPos after", ship.gridPos)

		}

		//	there is EmptyState tile in places where ship will be after move
		tileClear := true
		for i := 0; i < ship.length; i++ {
			// check if ship would not end up out of the board
			if ship.gridPos[i].x+move.xOffset >= b.size || ship.gridPos[i].y+move.yOffset >= b.size ||
				ship.gridPos[i].x+move.xOffset < 0 || ship.gridPos[i].y+move.yOffset < 0 {
				tileClear = false
				break
			}
			if b.tileAt(ship.gridPos[i].x+move.xOffset, ship.gridPos[i].y+move.yOffset).state != EmptyState &&
				b.tileAt(ship.gridPos[i].x+move.xOffset, ship.gridPos[i].y+move.yOffset).state != ShipState {
				tileClear = false
				break
			}
		}
		newMove := Move{
			xOffset:    move.xOffset,
			yOffset:    move.yOffset,
			isPossible: tileClear,
		}
		move = newMove
		newMoves = append(newMoves, move)
	}
	ship.moves = newMoves
}

func (b *Board) showLegalMoves(ship *Ship) {
	for _, move := range ship.moves {
		if move.isPossible {
			// set LegalMoveState for each tile where ship would be after move unless it'll be in place where ship is already
			for i := 0; i < ship.length; i++ {
				if b.tileAt(ship.gridPos[i].x+move.xOffset, ship.gridPos[i].y+move.yOffset).state != ShipState &&
					b.tileAt(ship.gridPos[i].x+move.xOffset, ship.gridPos[i].y+move.yOffset).state != ShipFrontState {
					b.tileAt(ship.gridPos[i].x+move.xOffset, ship.gridPos[i].y+move.yOffset).state = LegalMoveState
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
				f := mplusBigFont
				drawerImage.DrawImage(bombStateImage, op)
				bomb := b.getBombAtTile(i, j)
				if bomb != nil {
					text.Draw(drawerImage, strconv.FormatInt(int64(bomb.turnsToLive), 10), f, bomb.x*tileSize+xOffset+13, (bomb.y+1)*tileSize+yOffset-11, color.Black)
				}
			case BombExplodedState:
				drawerImage.DrawImage(bombExplodedStateImage, op)
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
