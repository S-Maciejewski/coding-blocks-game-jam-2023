package battleships

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type ShipPosition struct {
	x       int
	y       int
	isFront bool
}

type Move struct {
	xOffset    int
	yOffset    int
	isPossible bool
}

type Ship struct {
	length                     int
	health                     int
	moves                      []Move
	pos                        []ShipPosition
	isSelected                 bool
	images                     []ebiten.Image
	globalX, globalY           int
	previousPosX, previousPosY int
	isLegalPlacement           bool
	rotation                   ShipRotation
}

func GenerateShips(len2, len3, len4, len5 int) []*Ship {
	var ships []*Ship
	for i := 0; i < len2; i++ {
		ships = append(ships, NewShip(2))
	}
	for i := 0; i < len3; i++ {
		ships = append(ships, NewShip(3))
	}
	for i := 0; i < len4; i++ {
		ships = append(ships, NewShip(4))
	}
	for i := 0; i < len5; i++ {
		ships = append(ships, NewShip(5))
	}
	return ships
}

func getMovesForShipLength(length int) []Move {
	moves2 := []Move{
		{1, 0, false},
		{-1, 0, false},
	}
	moves3 := append(moves2, []Move{
		{2, 0, false},
		{-2, 0, false},
		{2, 1, false},
		{2, -1, false},
	}...)
	moves4 := append(moves3, []Move{
		{3, 0, false},
		{-3, 0, false},
		{3, 1, false},
		{3, -1, false},
	}...)
	moves5 := append(moves4, []Move{
		{4, 0, false},
		{-4, 0, false},
		{4, 1, false},
		{4, -1, false},
		{4, 2, false},
		{4, -2, false},
	}...)
	switch length {
	case 2:
		return moves2
	case 3:
		return moves3
	case 4:
		return moves4
	case 5:
		return moves5
	default:
		return []Move{}
	}
}

func NewShip(length int) *Ship {
	var pos []ShipPosition

	for i := 0; i < length; i++ {
		pos = append(pos, ShipPosition{
			x:       i * tileSize, // TODO check if horizontal/vertical
			y:       0,
			isFront: i == 0,
		})
	}

	ship := &Ship{
		length:     length,
		health:     length,
		moves:      getMovesForShipLength(length),
		pos:        pos,
		isSelected: false,
		rotation:   LeftRotation,
	}
	return ship
}

func (s *Ship) ResetToPreviousPosition() {
	s.globalX = s.previousPosX
	s.globalY = s.previousPosY
}

func (s *Ship) MoveShip(move Move) {
	//	move offset should be calculated from the front of the ship
	if move.isPossible {
		if s.pos[0].isFront {
			s.pos[0].x += move.xOffset
			s.pos[0].y += move.yOffset
		} else {
			s.pos[0].x -= move.xOffset
			s.pos[0].y -= move.yOffset
		}
	} else {
		panic("Impossible move attempted")
	}
}

func (s *Ship) rotate() {
	if s.rotation == 270 {
		s.rotation = 0
	} else {
		s.rotation += 90
	}
	newPos := []ShipPosition{}

	if s.rotation == LeftRotation {
		for i := 0; i < s.length; i++ {
			var pos ShipPosition
			if i == 0 {
				pos.isFront = true
			} else {
				pos.isFront = false
			}

			pos.x = i * tileSize
			pos.y = 0
			newPos = append(newPos, pos)
		}
	}

	if s.rotation == UpRotation {
		for i := 0; i < s.length; i++ {
			var pos ShipPosition
			if i == 0 {
				pos.isFront = true
			} else {
				pos.isFront = false
			}

			pos.x = 0
			pos.y = i * tileSize
			newPos = append(newPos, pos)
		}
	}

	if s.rotation == RightRotation {
		for i := 0; i < s.length; i++ {
			var pos ShipPosition
			if i == s.length-1 {
				pos.isFront = true
			} else {
				pos.isFront = false
			}

			pos.x = i * tileSize
			pos.y = 0
			newPos = append(newPos, pos)
		}
	}

	if s.rotation == DownRotation {
		for i := 0; i < s.length; i++ {
			var pos ShipPosition
			if i == s.length-1 {
				pos.isFront = true
			} else {
				pos.isFront = false
			}

			pos.x = 0
			pos.y = i * tileSize
			newPos = append(newPos, pos)
		}
	}

	s.pos = newPos
}

func (s *Ship) Draw(drawerImage *ebiten.Image) {
	for _, pos := range s.pos {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(pos.x+s.globalX), float64(pos.y+s.globalY))

		if pos.isFront {
			drawerImage.DrawImage(shipFrontImage, op)
		} else {
			drawerImage.DrawImage(shipImage, op)
		}
	}
}

func (s *Ship) In(x, y int) bool {
	inBounds := false
	for _, pos := range s.pos {
		if pos.x+s.globalX < x && (pos.x+s.globalX+tileSize) > x && pos.y+s.globalY < y && (pos.y+s.globalY+tileSize) > y {
			inBounds = true
			break
		}
	}

	return inBounds
}
