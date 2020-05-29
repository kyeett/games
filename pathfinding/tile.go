package pathfinding

import (
	"github.com/beefsack/go-astar"
	"github.com/peterhellberg/gfx"
)

var _ astar.Pather = &Tile{}

type Tile struct {
	X, Y int

	Kind string

	*World
}

const KindBlocker = "blocker"

func (t *Tile) PathNeighbors() []astar.Pather {
	neighbors := []astar.Pather{}
	for _, offset := range [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {
		if n := t.World.Get(t.X+offset[0], t.Y+offset[1]); n != nil && n.Kind != KindBlocker {
			neighbors = append(neighbors, n)
		}
	}

	return neighbors
}

func (t *Tile) PathNeighborCost(_ astar.Pather) float64 {
	return 1
}

func (t *Tile) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(*Tile)

	// Calculate manhattan distance
	return float64(gfx.IntAbs(toT.X - t.X) + gfx.IntAbs(toT.Y - t.Y))
}
