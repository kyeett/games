package pathfinding

import (
	"github.com/beefsack/go-astar"
	"github.com/peterhellberg/gfx"
	"math"
)

var _ astar.Pather = &Tile{}

type Tile struct {
	X, Y int

	Kind string

	MovementCost float64

	*World
}

func (t *Tile) isWalkable() bool {
	return t != nil && t.Kind != KindBlocker
}

const (
	KindBlocker = "blocker"
)

var diagonalOffsets = [][]int{
	{-1, -1},
	{1, -1},

	{-1, 1},
	{1, 1},
}

var manhattanOffsets = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

var offsets = [][]int{
	{-1, -1},
	{0, -1},
	{1, -1},

	{1, 0},

	{1, 1},
	{0, 1},
	{-1, 1},

	{-1, 0},
}

func wrappedGet(i int) []int {
	i = (len(offsets) + i) % len(offsets)
	return offsets[i]
}

func (t *Tile) PathNeighbors() []astar.Pather {
	neighbors := []astar.Pather{}

	for i, offset := range offsets {
		isDiagonal := i%2 == 0

		neighbor := t.World.Get(t.X+offset[0], t.Y+offset[1])

		if !neighbor.isWalkable() {
			continue
		}

		// Additional checks for diagonals
		if isDiagonal {
			//continue
			o1 := wrappedGet(i + 1)
			o2 := wrappedGet(i - 1)

			n1 := t.World.Get(t.X+o1[0], t.Y+o1[1])
			n2 := t.World.Get(t.X+o2[0], t.Y+o2[1])

			if !n1.isWalkable() || !n2.isWalkable() {
				continue
			}
		}

		// Calculate movement cost
		neighbors = append(neighbors, neighbor)
	}

	return neighbors
}

func (t *Tile) PathNeighborCost(to astar.Pather) float64 {
	toT := to.(*Tile)

	isDiagonal := toT.X != t.X && toT.Y != t.Y
	if isDiagonal {
		return math.Sqrt2
	}

	return 1
}

func (t *Tile) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(*Tile)

	// Calculate manhattan distance
	return float64(gfx.IntAbs(toT.X-t.X) + gfx.IntAbs(toT.Y-t.Y))
}
