package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/games/planes/plane"
	"github.com/kyeett/games/planes/spikes"
	"github.com/kyeett/games/planes/sprites"
	"github.com/kyeett/games/util"
	"github.com/peterhellberg/gfx"
	"log"
)

type game struct {
	plane *plane.Plane

	spikes     spikes.Spikes
	background *util.ScrollingImage
	foreground *util.ScrollingImage
}

func (g game) Update(_ *ebiten.Image) error {
	g.background.Scroll(velocityX/6)
	g.spikes.Scroll(velocityX)
	g.foreground.Scroll(velocityX)

	switch {
	case ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown):
		g.plane.Y += velocityY
	case ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) :
		g.plane.Y -= velocityY
	}


	// Check collision
	if g.spikes.CollidesWith(g.plane.Hitbox()) {
		g.plane.Visible = false
	}

	return nil
}

func (g game) Draw(screen *ebiten.Image) {

	g.background.Draw(screen)

	g.plane.Draw(screen)

	g.spikes.Draw(screen)
	g.foreground.Draw(screen)
}

const (
	screenWidth  = 800
	screenHeight = 480
	velocityY    = 5
)

var (
	velocityX    = 3.0
)

func (g game) Layout(w, h int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	g := &game{
		plane: plane.New(gfx.V(100, 100)),

		spikes: spikes.New(screenWidth, screenHeight),

		background: util.NewScrollingImage(sprites.Background, gfx.V(0, 0)),
		foreground: util.NewScrollingImage(sprites.ForegroundDirt, gfx.V(0, 480-71)),
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
