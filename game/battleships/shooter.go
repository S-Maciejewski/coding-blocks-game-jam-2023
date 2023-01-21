package battleships

import "math/rand"

type Bomb struct {
	x              int
	y              int
	didHit         bool
	turnsToLive    int
	turnsSinceDrop int
}

type Shooter struct {
	bombsDropped []Bomb
}

func NewShooter() *Shooter {
	return &Shooter{
		bombsDropped: []Bomb{},
	}
}

func (s *Shooter) GetNewBomb(board *Board) Bomb {
	var bombPosition Bomb

	if len(s.bombsDropped) != 0 && s.bombsDropped[len(s.bombsDropped)-1].didHit {
		bombOffset := 0
		if rand.Int()%2 == 0 {
			bombOffset = 1
		} else {
			bombOffset = -1
		}

		lastBomb := s.bombsDropped[len(s.bombsDropped)-1]

		if lastBomb.x != 0 && lastBomb.x != board.size-1 && lastBomb.y != 0 && lastBomb.y != board.size-1 {
			if rand.Int()%2 == 0 {
				bombPosition = Bomb{
					x:      s.bombsDropped[len(s.bombsDropped)-1].x + bombOffset,
					y:      s.bombsDropped[len(s.bombsDropped)-1].y,
					didHit: false,
				}
			} else {
				bombPosition = Bomb{
					x:      s.bombsDropped[len(s.bombsDropped)-1].x,
					y:      s.bombsDropped[len(s.bombsDropped)-1].y + bombOffset,
					didHit: false,
				}
			}
		} else if (lastBomb.x == 0 || lastBomb.x == board.size-1) && lastBomb.y != 0 && lastBomb.y != board.size-1 {
			bombPosition = Bomb{
				x:      s.bombsDropped[len(s.bombsDropped)-1].x,
				y:      s.bombsDropped[len(s.bombsDropped)-1].y + bombOffset,
				didHit: false,
			}
		} else if (lastBomb.y == 0 || lastBomb.y == board.size-1) && lastBomb.x != 0 && lastBomb.x != board.size-1 {
			bombPosition = Bomb{
				x:      s.bombsDropped[len(s.bombsDropped)-1].x + bombOffset,
				y:      s.bombsDropped[len(s.bombsDropped)-1].y,
				didHit: false,
			}
		} else if lastBomb.x == 0 && lastBomb.y == 0 {
			if rand.Int()%2 == 0 {
				bombPosition = Bomb{
					x:      s.bombsDropped[len(s.bombsDropped)-1].x + 1,
					y:      s.bombsDropped[len(s.bombsDropped)-1].y,
					didHit: false,
				}
			} else {
				bombPosition = Bomb{
					x:      s.bombsDropped[len(s.bombsDropped)-1].x,
					y:      s.bombsDropped[len(s.bombsDropped)-1].y + 1,
					didHit: false,
				}
			}
		} else if lastBomb.x == board.size-1 && lastBomb.y == board.size-1 {
			if rand.Int()%2 == 0 {
				bombPosition = Bomb{
					x:      s.bombsDropped[len(s.bombsDropped)-1].x - 1,
					y:      s.bombsDropped[len(s.bombsDropped)-1].y,
					didHit: false,
				}
			} else {
				bombPosition = Bomb{
					x:      s.bombsDropped[len(s.bombsDropped)-1].x,
					y:      s.bombsDropped[len(s.bombsDropped)-1].y - 1,
					didHit: false,
				}
			}
		}

		return bombPosition
	} else {
		for i := 0; i < board.size*board.size; i++ {
			randomTile := board.tileAt(rand.Intn(board.size), rand.Intn(board.size))
			if randomTile.state == EmptyState ||
				randomTile.state == ShipState ||
				randomTile.state == ShipFrontState {
				bombPosition = Bomb{
					x:      randomTile.x,
					y:      randomTile.y,
					didHit: false,
				}
				return bombPosition
			}
		}
	}
	panic("No valid bomb position found")
}
