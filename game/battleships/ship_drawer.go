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
	shipTileSize = 80
	tileMargin   = 4
)

var (
	shipImage = ebiten.NewImage(shipTileSize, shipTileSize)
)

func init() {
	shipImage.Fill(color.RGBA{0x58, 0x58, 0x58, 0xff})
}

func (b *ShipDrawer) Draw(drawerImage *ebiten.Image) {
	offset := 200
	for _, ship := range b.ships {
		for i := 0; i < ship.length; i++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(i*shipTileSize+offset), 500)
			drawerImage.DrawImage(shipImage, op)
		}

		offset = offset + 150
	}

}
