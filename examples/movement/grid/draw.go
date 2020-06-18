package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/games/pathfinding"
	"golang.org/x/image/colornames"
	"image/color"
)



func (g *game) Draw(screen *ebiten.Image) {
	for y := 0; y < gridHeight; y++ {
		for x := 0; x < gridWidth; x++ {

			t := g.world.Get(x, y)

			clr := colornames.White
			switch t.Kind {
			case pathfinding.KindBlocker:
				clr = colornames.Black
			}

			g.drawTile(screen, t.X, t.Y, clr)
		}
	}


	g.drawTile(screen, g.player.gridPos.X, g.player.gridPos.Y, colornames.Yellow)
	g.player.Draw(screen)

}

func (g *game) drawTile(screen *ebiten.Image, x, y int, clr color.Color) {
	opt := &ebiten.DrawImageOptions{}

	switch clr {
	case colornames.Black:
		opt.ColorM.Scale(0, 0, 0, 1)
	case colornames.Red:
		opt.ColorM.Scale(1, 0, 0, 1)
	case colornames.Yellow:
		opt.ColorM.Scale(1, 1, 0, 1)
	default:
		opt.ColorM.Scale(1, 1, 1, 1)
	}

	g.grid.TranslateXY(x, y, opt)
	screen.DrawImage(square, opt)
}
