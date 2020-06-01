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

//var allOffsets = [][]int{
//	{-1, -1},
//	{-1, 0},
//	{1, -1},
//	{1, 0},
//	{-1, 1},
//	{0, -1},
//	{1, 1},
//	{0, 1},
//}

func (t *Tile) PathNeighbors() []astar.Pather {
	neighbors := []astar.Pather{}

	for _, offset := range manhattanOffsets {
		if n := t.World.Get(t.X+offset[0], t.Y+offset[1]); n.isWalkable() {
			n.MovementCost = 1
			neighbors = append(neighbors, n)
		}
	}

	// Up right
	if t.upRight().isWalkable() && t.up().isWalkable() && t.right().isWalkable() {
		neighbors = append(neighbors, t.upRight())
	}

	// Down right
	if t.downRight().isWalkable() && t.down().isWalkable() && t.right().isWalkable() {
		neighbors = append(neighbors, t.downRight())
	}

	// Down left
	if t.downLeft().isWalkable() && t.down().isWalkable() && t.left().isWalkable() {
		neighbors = append(neighbors, t.downLeft())
	}

	// Up left
	if t.upLeft().isWalkable() && t.up().isWalkable() && t.left().isWalkable() {
		neighbors = append(neighbors, t.upLeft())
	}

	return neighbors
}

func (t *Tile) PathNeighborCost(_ astar.Pather) float64 {
	return math.Sqrt2
}

func (t *Tile) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(*Tile)

	// Calculate manhattan distance
	return float64(gfx.IntAbs(toT.X-t.X) + gfx.IntAbs(toT.Y-t.Y))
}

func (t *Tile) up() *Tile {
	return t.World.Get(t.X, t.Y-1)
}

func (t *Tile) upRight() *Tile {
	return t.World.Get(t.X+1, t.Y-1)
}

func (t *Tile) right() *Tile {
	return t.World.Get(t.X+1, t.Y)
}

func (t *Tile) downRight() *Tile {
	return t.World.Get(t.X+1, t.Y+1)
}

func (t *Tile) down() *Tile {
	return t.World.Get(t.X, t.Y+1)
}

func (t *Tile) downLeft() *Tile {
	return t.World.Get(t.X-1, t.Y+1)
}

func (t *Tile) left() *Tile {
	t2 := t.World.Get(t.X-1, t.Y)
	t2.MovementCost = math.Sqrt2
	return t
}

func (t *Tile) upLeft() *Tile {
	return t.World.Get(t.X-1, t.Y-1)
}
