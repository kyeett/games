package grid

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/peterhellberg/gfx"
	"image"
)

type Grid struct {
	tileWidth  int
	tileHeight int
	cols       int
	rows       int
	paddingX   int
	paddingY   int
}

func New(tileWidth, tileHeight, cols, rows, paddingX, paddingY int) *Grid {
	return &Grid{
		tileWidth:  tileWidth,
		tileHeight: tileHeight,
		cols:       cols,
		rows:       rows,
		paddingX:   paddingX,
		paddingY:   paddingY,
	}
}

func (g *Grid) Translate(x, y int, opt *ebiten.DrawImageOptions) {
	v, _ := g.toVec(x, y)
	opt.GeoM.Translate(v.X, v.Y)
}

func (g *Grid) ToPoint(v gfx.Vec) (image.Point, bool) {
	if int(v.X) < 0 || int(v.Y) < 0 || int(v.X) >= g.maxX() || int(v.Y) >= g.maxY() {
		return gfx.ZP, false
	}

	x := int(v.X) / (g.tileWidth + g.paddingX)
	y := int(v.Y) / (g.tileHeight + g.paddingY)
	return gfx.Pt(x, y), true
}

func (g *Grid) maxX() int {
	return (g.tileWidth + g.paddingX) * g.cols
}

func (g *Grid) maxY() int {
	return (g.tileHeight + g.paddingY) * g.rows
}

func (g *Grid) toVec(x, y int) (gfx.Vec, bool) {
	if x < 0 || y < 0 || x >= g.cols || y >= g.rows {
		return gfx.ZV, false
	}
	v := gfx.IV((g.tileWidth+g.paddingX)*x, (g.tileHeight+g.paddingY)*y)
	return v, true
}
