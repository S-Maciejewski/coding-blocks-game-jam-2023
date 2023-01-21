package battleships

type TileState int

const (
	EmptyState TileState = iota
	ShipState
	BombState
	LegalMove
	IllegalMove
)
