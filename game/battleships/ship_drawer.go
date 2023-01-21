package battleships

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type ShipDrawer struct {
	ships []*Ship
}

func NewDrawer(ships []*Ship) *ShipDrawer {
	return &ShipDrawer{
		ships: ships,
	}
}

const (
	shipTileSize = 60
	drawerOffset = 50
)

var (
	shipImage      = ebiten.NewImage(shipTileSize, shipTileSize)
	shipFrontImage = ebiten.NewImage(shipTileSize, shipTileSize)
)

func init() {
	shipImage.Fill(color.RGBA{0x58, 0x58, 0x58, 0xff})
	shipFrontImage.Fill(color.RGBA{0x40, 0x40, 0x40, 0xff})
}

func (b *ShipDrawer) Draw(drawerImage *ebiten.Image) {
	for _, ship := range b.ships {
		ship.Draw(drawerImage)
	}
}
