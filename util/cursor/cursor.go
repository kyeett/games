package cursor

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/peterhellberg/gfx"
)

func Vec() gfx.Vec {
	x, y := ebiten.CursorPosition()
	return gfx.IV(x, y)
}
