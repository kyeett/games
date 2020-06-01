package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/kyeett/games/util/move"
	"github.com/peterhellberg/gfx"
	"golang.org/x/image/colornames"
	"log"
	"math/rand"
)

type game struct {
	player gfx.Vec
	target gfx.Vec
}

func (g *game) Update(_ *ebiten.Image) error {
	powerRemaining := 10.0
	var reached bool
	for powerRemaining > 0 {
		g.player, powerRemaining, reached = move.Towards(g.player, g.target, powerRemaining)
		if reached {
			g.target = gfx.V(rand.Float64()*screenWidth, rand.Float64()*screenHeight)
		}
	}

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {

	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(g.player.X, g.player.Y)
	screen.DrawImage(square, opt)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%0.0f,%0.0f", g.target.X, g.target.Y))
}

const (
	screenWidth  = 800
	screenHeight = 600
)

func (g *game) Layout(w, h int) (int, int) {
	return screenWidth, screenHeight
}

const (
	tileSize = 20
)

var (
	square *ebiten.Image
)


func main() {
	square, _ = ebiten.NewImage(tileSize, tileSize, ebiten.FilterDefault)
	square.Fill(colornames.White)

	g := &game{
		player: gfx.V(10,10),
		target: gfx.V(100,100),
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
