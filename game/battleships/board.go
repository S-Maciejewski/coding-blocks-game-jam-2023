package battleships

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
