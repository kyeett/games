package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/kyeett/games/examples/sidescrolling/assets"
	"github.com/kyeett/games/util"
	"github.com/kyeett/games/util/sidescroller"
	"github.com/peterhellberg/gfx"
	"golang.org/x/image/colornames"
	"log"
)

var background *ebiten.Image

func init() {
	background = util.LoadImageOrFatal(assets.MustAsset("assets/Sample.png"))
}

type game struct {
	scroller *sidescroller.Scroller
	camera   gfx.Vec
}

func (g *game) Update(_ *ebiten.Image) error {

	g.camera = g.camera.Add(g.scroller.Force().Scaled(3))
	g.camera.X = gfx.Clamp(g.camera.X, 0, cameraWidth - cameraWindow.W())
	g.camera.Y = gfx.Clamp(g.camera.Y, 0, cameraHeight - cameraWindow.H())

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	c := cameraWindow.Moved(g.camera)

	// Camera view
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Scale(cameraWidth/cameraWindow.W(), cameraHeight/cameraWindow.H())
	screen.DrawImage(background.SubImage(c.Bounds()).(*ebiten.Image), opt)

	// Zoomed out
	opt = &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(0, 320)
	screen.DrawImage(background, opt)

	// Delimiter
	ebitenutil.DrawRect(screen, 0, cameraHeight-1, cameraWidth, 2, colornames.Red)

	// Draw camera position
	ebitenutil.DrawRect(screen, c.Min.X, c.Min.Y+cameraHeight, c.W(), c.H(), colornames.Red)

	s := fmt.Sprintf("%v", g.scroller.Force())
	ebitenutil.DebugPrintAt(screen, s, screenWidth/2, screenHeight/2)
}

const (
	cameraWidth  = 640
	cameraHeight = 320

	screenWidth  = cameraWidth
	screenHeight = 2 * (cameraHeight)
)

var cameraWindow = gfx.R(0, 0, 200, 100)

func (g *game) Layout(w, h int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	cameraPos := gfx.V(100, 100)

	g := &game{
		camera:   cameraPos,
		scroller: sidescroller.NewEbiten(gfx.R(0, 0, cameraWidth, cameraHeight)),
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
