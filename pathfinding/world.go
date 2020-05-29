package pathfinding

import (
	"github.com/kyeett/collections/grid"
)

type World struct {
	grid *grid.Grid
}

func (w World) Get(x, y int) *Tile {
	t := w.grid.Get(grid.Point{x, y})
	if t == nil {
		return nil
	}

	return t.(*Tile)
}

func (w World) Set(t *Tile, x, y int) {
	t.X = x
	t.Y = y
	w.grid.Set(grid.Point{x, y}, t)
}

func New(w, h int) *World {
	world := &World{
		grid: grid.New(w, h),
	}
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			world.Set(&Tile{World: world}, x, y)
		}
	}

	return world
}
