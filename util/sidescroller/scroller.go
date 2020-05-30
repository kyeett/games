package sidescroller

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/peterhellberg/gfx"
)

// Scroller
type Scroller struct {
	source                func() (int, int)
	area                  gfx.Rect
}

const margin = 50


// Use Ebiten cursor as source
func NewEbiten(rect gfx.Rect) *Scroller {
	s := &Scroller{
		source: ebiten.CursorPosition,
		area:  rect,
	}
	return s
}

func (s *Scroller) Force() gfx.Vec {
	ix, iy := ebiten.CursorPosition()
	x, y := float64(ix), float64(iy)
	forceX := scrollFunction(x, s.area.Min.X, s.area.Max.X, margin)
	forceY := scrollFunction(y, s.area.Min.Y, s.area.Max.Y, margin)

	return gfx.V(forceX, forceY)
}

func scrollFunction(t, min, max, margin float64) float64 {
	switch {
	case t < min:
		return -1
	case t < min+margin:
		return (t-min)/margin - 1
	case t < max-margin:
		return 0
	case t < max:
		return (t + margin - max) / margin
	default:
		return 1
	}
}
