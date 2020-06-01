package main

import (
	"github.com/beefsack/go-astar"
	"github.com/kyeett/games/pathfinding"
	"github.com/peterhellberg/gfx"
	"image"
)

func (g *game) startMovingTowards(target image.Point) bool {
	t1 := g.world.Get(g.player.gridPos.X, g.player.gridPos.Y)
	t2 := g.world.Get(target.X, target.Y)
	path, _, found := astar.Path(t1, t2)
	if found {
		g.player.path = g.gridPath(path)
		g.player.moving = true
		g.player.gridPos = g.player.path[0].Point
		return true
	}

	return found
}

type PathPoint struct {
	image.Point
	position gfx.Vec
}

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
