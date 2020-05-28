package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/games/planes/plane"
	"github.com/kyeett/games/planes/spikes"
	"github.com/kyeett/games/planes/sprites"
	"github.com/kyeett/games/util"
	"github.com/peterhellberg/gfx"
	"golang.org/x/image/colornames"
	"log"
)

type game struct {
	state string

	plane *plane.Plane

	spikes     spikes.Spikes
	background *util.ScrollingImage
	foreground *util.ScrollingImage

	goal *ebiten.Image

	velocityX float64
	posX      float64
}



const (
	screenWidth  = 800
	screenHeight = 480
	velocityY    = 5
	velocityX    = 3
)

func (g game) Layout(w, h int) (int, int) {
	return screenWidth, screenHeight
}

func (g *game) newGame() {
	g.plane = plane.New(gfx.V(100, 100))
	g.spikes = spikes.New(screenWidth, screenHeight)
	g.velocityX = 3.0
	g.posX = 0
}

var (
	goalBackground *ebiten.Image
)

func main() {

	g := &game{
		state: "new",

		background: util.NewScrollingImage(sprites.Background, gfx.V(0, 0)),
		foreground: util.NewScrollingImage(sprites.ForegroundDirt, gfx.V(0, 480-71)),

		goal: util.MustNewRectangle(screenWidth, screenHeight, colornames.Black),
	}
	g.newGame()

	ebiten.SetFullscreen(true)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
