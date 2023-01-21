package battleships

type Board struct {
	size  int
	tiles [][]Tile
}

func NewBoard(size int) *Board {
	var tiles [][]Tile
	for i := 0; i <= size; i++ {
		for j := 0; j <= size; j++ {
			tiles[i][j] = *NewTile(i, j)
		}
	}
	return &Board{
		size:  size,
		tiles: tiles,
	}
}

func (b *Board) tileAt(x, y int) *Tile {
	return &b.tiles[x][y]
}
