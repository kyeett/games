package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/games/planes/plane"
	"github.com/kyeett/games/planes/spikes"
	"github.com/kyeett/games/planes/sprites"
	"github.com/kyeett/games/util"
	"github.com/peterhellberg/gfx"
	"log"
	"math"
	"strconv"
)

type game struct {
	state string

	plane *plane.Plane

	spikes     spikes.Spikes
	background *util.ScrollingImage
	foreground *util.ScrollingImage

	velocityX float64
}

func (g *game) Update(_ *ebiten.Image) error {
	g.background.Scroll(g.velocityX / 6)
	g.spikes.Scroll(g.velocityX)
	g.foreground.Scroll(g.velocityX)

	switch g.state {
	case "running":
		g.stateRunning()
		m += 0.2
		g.velocityX += 0.1
	default:

		// Reset game
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			g.newGame()
			g.state = "running"
		}
	}

	return nil
}

func (g *game) stateRunning() {
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown):
		g.plane.Y += velocityY
		gfx.Clamp(g.plane.Y, 0, screenHeight-30)
	case ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp):
		g.plane.Y -= velocityY
		gfx.Clamp(g.plane.Y, 0, screenHeight-30)
	}

	// Check collision
	if g.spikes.CollidesWith(g.plane.Hitbox()) {
		g.plane.Visible = false
		g.state = "game-over"
	}

}

func digit(num, place int) int {
	r := num % int(math.Pow(10, float64(place)))
	return r / int(math.Pow(10, float64(place-1)))
}

func (g *game) Draw(screen *ebiten.Image) {

	g.background.Draw(screen)

	g.plane.Draw(screen)

	g.spikes.Draw(screen)
	g.foreground.Draw(screen)

	switch g.state {
	case "running":
		// Do nothing

		opt := &ebiten.DrawImageOptions{}
		opt.GeoM.Scale(0.5,0.5)
		opt.GeoM.Translate(screenWidth-5, 5)
		ss := strconv.Itoa(int(m))
		for i := 0; i < len(ss); i++ {
			img := sprites.Numbers[digit(int(m), i+1)]
			opt.GeoM.Translate(-28, 0)
			screen.DrawImage(img, opt)
		}

	case "new":

		drawCenteredImage(screen, sprites.TextGetReady, gfx.V(0, 0))
		drawCenteredImage(screen, sprites.PressEnterToStart, gfx.V(0, 70))
	case "game-over":
		drawCenteredImage(screen, sprites.TextGameOver, gfx.V(0, 0))
		drawCenteredImage(screen, sprites.PressEnterToStart, gfx.V(0, 70))
	}
}

func drawCenteredImage(screen, img *ebiten.Image, offset gfx.Vec) {
	pos := util.CenterBoundsOnBounds(img.Bounds(), screen.Bounds())
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(float64(pos.X)+offset.X, float64(pos.Y)+offset.Y)
	screen.DrawImage(img, opt)
}

const (
	screenWidth  = 800
	screenHeight = 480
	velocityY    = 5
)

var m float64 = 0

func (g game) Layout(w, h int) (int, int) {
	return screenWidth, screenHeight
}

func (g *game) newGame() {
	g.plane = plane.New(gfx.V(100, 100))
	g.spikes = spikes.New(screenWidth, screenHeight)
	g.velocityX = 3.0

	m = 0
}

func main() {
	g := &game{
		state: "new",

		background: util.NewScrollingImage(sprites.Background, gfx.V(0, 0)),
		foreground: util.NewScrollingImage(sprites.ForegroundDirt, gfx.V(0, 480-71)),
	}
	g.newGame()

	ebiten.SetFullscreen(true)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
