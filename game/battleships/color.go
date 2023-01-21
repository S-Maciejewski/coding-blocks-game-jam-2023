package battleships

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

var (
	backgroundColor = color.RGBA{0x00, 0xff, 0x0f, 0xff}
	frameColor      = color.RGBA{0xbb, 0xad, 0xa0, 0xff}
)

var (
	emptyStateImage        = ebiten.NewImage(shipTileSize, shipTileSize)
	shipStateImage         = ebiten.NewImage(shipTileSize, shipTileSize)
	shipFrontStateImage    = ebiten.NewImage(shipTileSize, shipTileSize)
	shipHitStateImage      = ebiten.NewImage(shipTileSize, shipTileSize)
	shipSunkStateImage     = ebiten.NewImage(shipTileSize, shipTileSize)
	shipFrontHitStateImage = ebiten.NewImage(shipTileSize, shipTileSize)
	bombStateImage         = ebiten.NewImage(shipTileSize, shipTileSize)
	legalMoveImage         = ebiten.NewImage(shipTileSize, shipTileSize)
	illegalMoveImage       = ebiten.NewImage(shipTileSize, shipTileSize)
	hoverImage             = ebiten.NewImage(shipTileSize, shipTileSize)
)

func init() {
	emptyStateImage.Fill(color.RGBA{0x49, 0x80, 0xff, 0xff})
	shipStateImage.Fill(color.RGBA{0x58, 0x58, 0x58, 0xff})
	shipFrontStateImage.Fill(color.RGBA{0x99, 0x4c, 0x00, 0xff})
	shipHitStateImage.Fill(color.RGBA{0xcc, 0x66, 0x00, 0xff})
	shipSunkStateImage.Fill(color.RGBA{0x66, 0x00, 0x00, 0xff})
	shipFrontHitStateImage.Fill(color.RGBA{0x66, 0x00, 0x00, 0xff})
	bombStateImage.Fill(color.RGBA{0x66, 0x33, 0x00, 0xff})
	legalMoveImage.Fill(color.RGBA{0x66, 0xff, 0xb2, 0xff})
	illegalMoveImage.Fill(color.RGBA{0xff, 0x66, 0x66, 0xff})
	hoverImage.Fill(color.RGBA{0x00, 0x00, 0x00, 0x7f})
}
