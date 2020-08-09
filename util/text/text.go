package text

import (
	"github.com/peterhellberg/gfx"
	"golang.org/x/image/font"
)

func BoundingBoxFromString(s string, fnt font.Face) gfx.Rect {
	width := font.MeasureString(fnt, s).Ceil()
	height := fnt.Metrics().Height.Ceil()
	return gfx.R(0, 0, float64(width), float64(height))
}