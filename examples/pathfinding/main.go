package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/games/pathfinding"
	"github.com/kyeett/games/util/grid"
	"golang.org/x/image/colornames"
	"log"
)

type point struct{ x, y int }

type game struct {
	world  *pathfinding.World
	player point
	target point
	grid   *grid.Grid
}

func (g *game) Update(_ *ebiten.Image) error {
	return nil
}

const (
	screenWidth  = 800
	screenHeight = 600

	tileSize = 40

	paddingX = 1
	paddingY = 1
)

func (g *game) Layout(w, h int) (int, int) {
	return screenWidth, screenHeight
}

var square *ebiten.Image

func main() {
	square, _ = ebiten.NewImage(tileSize, tileSize, ebiten.FilterDefault)
	square.Fill(colornames.White)

	g := &game{
		player: point{0, 0},
		target: point{9, 9},
		grid:   grid.New(tileSize, tileSize, gridHeight, gridWidth, paddingX, paddingY),
		world:  pathfinding.New(gridWidth, gridHeight),
	}

	for i := 0; i < 5; i++ {
		g.world.Get(2, i).Kind = pathfinding.KindBlocker
	}
	for i := 0; i < 7; i++ {
		g.world.Get(8, 3+i).Kind = pathfinding.KindBlocker
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
