package main

import (
	"github.com/beefsack/go-astar"
	"github.com/kyeett/games/pathfinding"
	"github.com/peterhellberg/gfx"
	"image"
	"math"
)

func (g *game) startMovingTowards(target image.Point) bool {
	t1 := g.world.Get(g.player.gridPos.X, g.player.gridPos.Y)
	t2 := g.world.Get(target.X, target.Y)
	path, _, found := astar.Path(t1, t2)
	if found {
		g.player.path = g.gridPath(path)
		g.player.moving = true
		g.player.gridPos = g.updateFirstPosition()
		return true
	}

	return found
}

func (g *game) updateFirstPosition() image.Point {
	p1 := g.player.gridPos
	p2 := g.player.path[0].Point

	g.player.gridPos = p2
	a := gfx.PV(p1).To(gfx.PV(p2)).Angle()
	switch a {
	case 0:
		g.player.frame = 2
	case -math.Pi / 4:
		g.player.frame = 3
	case -2 * math.Pi / 4:
		g.player.frame = 4
	case -3 * math.Pi / 4:
		g.player.frame = 5
	case math.Pi:
		g.player.frame = 6
	case 3 * math.Pi / 4:
		g.player.frame = 7
	case 2 * math.Pi / 4:
		g.player.frame = 1
	default:
		g.player.frame = 0
	}

	return g.player.path[0].Point
}

type PathPoint struct {
	image.Point
	position gfx.Vec
}

type direction int

const (
	UL direction = iota
	U
	UR
	R
	DR
	D
	DL
	L
)

func (g *game) gridPath(path []astar.Pather) []PathPoint {
	gridPath := make([]PathPoint, len(path)-1)

	// Add in reverse order, and skip current tile
	for i := 0; i < len(path)-1; i++ {
		t := path[i].(*pathfinding.Tile)

		pt := gfx.Pt(t.X, t.Y)
		pos := g.grid.MustVec(pt)
		gridPath[len(gridPath)-i-1] = PathPoint{
			Point:    pt,
			position: pos,
		}
	}
	return gridPath
}
