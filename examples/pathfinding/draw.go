package main

import (
	"github.com/beefsack/go-astar"
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/games/pathfinding"
	"github.com/kyeett/games/util/cursor"
	"golang.org/x/image/colornames"
	"image/color"
)

const (
	gridHeight = 10
	gridWidth  = 10
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


	if p, found := g.grid.ToPoint(cursor.Vec()); found {
		g.target = point{p.X, p.Y}
		g.drawTile(screen, p.X, p.Y, colornames.Yellow)
	}

	t1 := g.world.Get(g.player.x, g.player.y)
	t2 := g.world.Get(g.target.x, g.target.y)
	path, _, found := astar.Path(t1, t2)
	if !found {
		return
	}

	for _, p := range path {
		t := p.(*pathfinding.Tile)
		g.drawTile(screen, t.X, t.Y, colornames.Red)
	}
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


	g.grid.Translate(x, y, opt)
	screen.DrawImage(square, opt)
}


