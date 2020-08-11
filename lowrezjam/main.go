package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/kyeett/games/util"
	"github.com/kyeett/games/util/move"
	"github.com/peterhellberg/gfx"
	"image"
	"log"
)

var (
	cameraBuffer, _ = ebiten.NewImage(128, 128, ebiten.FilterLinear)
	mapImage        *ebiten.Image
)

func init() {
	mapImage = util.LoadAssetImageOrFatal(Asset, "8x8map.png")
}

type game struct {
	camera *Camera
}

type Camera struct {
	center image.Point
	pos    gfx.Vec

	moving bool
}

func (c *Camera) Add(x, y int) {
	c.center = c.center.Add(image.Pt(x, y))
}

func (c *Camera) Update() {
	var targetReached bool
	c.pos, _, targetReached = move.Towards(c.pos, gfx.PV(c.center).Scaled(64), 4)
	c.moving = !targetReached
}

func (g *game) Update(_ *ebiten.Image) error {
	g.camera.Update()
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {

	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			cx := float64(x*64) - g.camera.pos.X
			cy := float64(y*64) - g.camera.pos.Y
			//ebitenutil.DrawRect(cameraBuffer, cx, cy, 64, 64, color.RGBA{
			//	R: uint8(80 * x),
			//	G: uint8(80 * y),
			//	B: 160,
			//	A: 255,
			//})
			//padding := 4.0
			//size := 64 - 2*padding
			//ebitenutil.DrawRect(cameraBuffer, cx+padding, cy+padding, size, size, color.RGBA{
			//	R: uint8(60 * x),
			//	G: uint8(60 * y),
			//	B: 120,
			//	A: 255,
			//})

			opt := &ebiten.DrawImageOptions{}
			opt.GeoM.Translate(cx, cy)
			cameraBuffer.DrawImage(mapImage, opt)
		}
	}

	if !g.camera.moving {
		if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
			g.camera.Add(-1, 0)
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
			g.camera.Add(1, 0)
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			g.camera.Add(0, -1)
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
			g.camera.Add(0, 1)
		}
	}

	screen.DrawImage(cameraBuffer.SubImage(image.Rect(0, 0, 64, 64)).(*ebiten.Image), &ebiten.DrawImageOptions{})
}

const (
	screenWidth  = 800
	screenHeight = 600
)

func (g game) Layout(w, h int) (int, int) {
	return 128, 128
}

func main() {
	g := &game{
		camera: &Camera{
			center: image.Pt(1, 1),
			pos:    gfx.V(5, 5),
		},
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
