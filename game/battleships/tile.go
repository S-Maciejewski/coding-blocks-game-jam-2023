package battleships

type Tile struct {
	x     int
	y     int
	state TileState
}

func NewTile(x, y int) *Tile {
	return &Tile{
		x:     x,
		y:     y,
		state: EmptyState,
	}
}
