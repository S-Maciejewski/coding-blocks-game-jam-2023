package battleships

type TileState int

const (
	EmptyState TileState = iota
	ShipState
	ShipFrontState
	ShipHitState
	ShipSunkState
	BombState
	BombExplodedState
	LegalMoveState
	IllegalMoveState
)
