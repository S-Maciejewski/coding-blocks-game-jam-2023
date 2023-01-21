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
	drawerOffset = 50
)

var (
	shipImage = ebiten.NewImage(shipTileSize, shipTileSize)
)

func init() {
	shipImage.Fill(color.RGBA{0x58, 0x58, 0x58, 0xff})
}

func (b *ShipDrawer) Draw(drawerImage *ebiten.Image) {
	width, height := drawerImage.Size()
	xOffset := drawerOffset
	yOffset := height - (3 * shipTileSize)
	for _, ship := range b.ships {

		if (ship.length*shipTileSize)+xOffset >= (width - drawerOffset) {
			xOffset = drawerOffset
			yOffset += shipTileSize + 10
		}

		for i := 0; i < ship.length; i++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((i*shipTileSize)+xOffset), float64(yOffset))
			drawerImage.DrawImage(shipImage, op)
		}

		xOffset += (ship.length * shipTileSize) + drawerOffset
	}

}
