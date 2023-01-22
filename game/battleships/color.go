package battleships

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	backgroundColor = color.RGBA{0xff, 0xff, 0xff, 0xff}
	frameColor      = color.RGBA{0xbb, 0xad, 0xa0, 0xff}
)

var (
	emptyStateImage        = ebiten.NewImage(tileSizeWithGridOffset, tileSizeWithGridOffset)
	shipStateImage         = ebiten.NewImage(tileSizeWithGridOffset, tileSizeWithGridOffset)
	shipFrontStateImage    = ebiten.NewImage(tileSizeWithGridOffset, tileSizeWithGridOffset)
	shipHitStateImage      = ebiten.NewImage(tileSizeWithGridOffset, tileSizeWithGridOffset)
	shipSunkStateImage     = ebiten.NewImage(tileSizeWithGridOffset, tileSizeWithGridOffset)
	shipInDangerStateImage = ebiten.NewImage(tileSize, tileSize)
	shipFrontHitStateImage = ebiten.NewImage(tileSizeWithGridOffset, tileSizeWithGridOffset)
	bombStateImage         = ebiten.NewImage(tileSizeWithGridOffset, tileSizeWithGridOffset)
	legalMoveImage         = ebiten.NewImage(tileSizeWithGridOffset, tileSizeWithGridOffset)
	illegalMoveImage       = ebiten.NewImage(tileSizeWithGridOffset, tileSizeWithGridOffset)
	hoverImage             = ebiten.NewImage(tileSizeWithGridOffset, tileSizeWithGridOffset)
	bombExplodedStateImage = ebiten.NewImage(tileSizeWithGridOffset, tileSizeWithGridOffset)
)

func init() {
	emptyStateImage.Fill(color.RGBA{0x49, 0x80, 0xff, 0xff})
	shipStateImage.Fill(color.RGBA{0x58, 0x58, 0x58, 0xff})
	shipFrontStateImage.Fill(color.RGBA{0x40, 0x40, 0x40, 0xff})
	shipHitStateImage.Fill(color.RGBA{0xcc, 0x66, 0x00, 0xff})
	shipSunkStateImage.Fill(color.RGBA{0x66, 0x00, 0x00, 0xff})
	shipInDangerStateImage.Fill(color.RGBA{0x66, 0x33, 0x00, 0x7f})
	shipFrontHitStateImage.Fill(color.RGBA{0x66, 0x00, 0x00, 0xff})
	bombStateImage.Fill(color.RGBA{0x66, 0x33, 0x00, 0xff})
	bombExplodedStateImage.Fill(color.RGBA{0xff, 0x2d, 0x00, 0xff})
	legalMoveImage.Fill(color.RGBA{0x66, 0xff, 0xb2, 0xff})
	illegalMoveImage.Fill(color.RGBA{0xff, 0x66, 0x66, 0xff})
	hoverImage.Fill(color.RGBA{0x00, 0x00, 0x00, 0x7f})
}
