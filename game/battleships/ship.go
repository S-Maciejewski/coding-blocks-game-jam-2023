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
	length     int
	health     int
	moves      []Move
	pos        []ShipPosition
	isSelected bool
	images     []ebiten.Image
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
	ship := &Ship{
		length:     length,
		health:     length,
		moves:      getMovesForShipLength(length),
		pos:        []ShipPosition{},
		isSelected: false,
	}
	return ship
}

func (s *Ship) MoveShip(move Move) {
	//	move offset should be calculated from the front of the ship
	if move.isPossible {
		if s.pos[0].isFront {
			s.pos[0].x += xOffset
			s.pos[0].y += yOffset
		} else {
			s.pos[0].x -= xOffset
			s.pos[0].y -= yOffset
		}
	} else {
		panic("Impossible move attempted")
	}
}

func (s *Ship) Rotate90() {
	for i := 0; i < s.length; i++ {
		s.pos[i].x, s.pos[i].y = s.pos[i].y, s.pos[i].x
	}
}

func (s *Ship) Rotate180() {
	for i := 0; i < s.length; i++ {
		s.pos[i].x = -s.pos[i].x
		s.pos[i].y = -s.pos[i].y
	}
}

func (s *Ship) Draw(drawerImage *ebiten.Image, startPosX, startPosY int) {
	for i := 0; i < s.length; i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64((i*shipTileSize)+startPosX), float64(startPosY))
		drawerImage.DrawImage(shipImage, op)
	}
}
