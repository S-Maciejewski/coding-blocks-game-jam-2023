package battleships

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
}

func GenerateShips(len2, len3, len4, len5 int) []*Ship {
	ships := []*Ship{}
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

func NewShip(length int) *Ship {
	ship := &Ship{
		length:     length,
		health:     length,
		moves:      []Move{},
		pos:        []ShipPosition{},
		isSelected: false,
	}
	return ship
}

func (s *Ship) Move(xOffset, yOffset int) {
	//	move offset should be calculated from the front of the ship
	if s.pos[0].isFront {
		s.pos[0].x += xOffset
		s.pos[0].y += yOffset
	} else {
		s.pos[0].x -= xOffset
		s.pos[0].y -= yOffset
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
