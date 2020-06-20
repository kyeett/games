package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/kyeett/games/util"
	"github.com/kyeett/games/util/grid"
	"github.com/peterhellberg/gfx"
	"golang.org/x/image/colornames"
	"image"
	"image/color"
	"log"
)

var (
	CharacterSprite, PlayerSprite *ebiten.Image
)

func init() {
	img := gfx.MustOpenImage("assets/images/arial10x10.png")
	CharacterSprite, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

func CharImage(c byte) *ebiten.Image {
	var r image.Rectangle
	switch c {
	case '@':
		r = image.Rect(0, 0, 10, 10).Add(image.Pt(0, 10))
	default:
		log.Fatalf("invalid char: %v", c)
	}

	return CharacterSprite.SubImage(r).(*ebiten.Image)
}

type Entity struct {
	image.Point
	color.Color
	char byte
}

type Player struct {
	Entity
}

func (p *Player) DrawAt(screen *ebiten.Image, opt *ebiten.DrawImageOptions) {
	util.OptScaleByColor(opt, p.Color)
	screen.DrawImage(CharImage(p.char), opt)
}

type game struct {
	grid   *grid.Grid
	player *Player
}

func (g *game) Update(_ *ebiten.Image) error {

	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyUp):
		g.player.Y--
	case inpututil.IsKeyJustPressed(ebiten.KeyDown):
		g.player.Y++
	case inpututil.IsKeyJustPressed(ebiten.KeyLeft):
		g.player.X--
	case inpututil.IsKeyJustPressed(ebiten.KeyRight):
		g.player.X++
	}

	switch {
	case ebiten.IsKeyPressed(ebiten.KeyEscape):
		return gfx.ErrDone
	case inpututil.IsKeyJustPressed(ebiten.KeyEnter) && ebiten.IsKeyPressed(ebiten.KeyAlt):
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "template")

	opt := &ebiten.DrawImageOptions{}
	g.grid.TranslateXY(g.player.X, g.player.Y, opt)
	g.player.DrawAt(screen, opt)
}

const (
	windowWidth  = 400
	windowHeight = 250
)

func (g game) Layout(w, h int) (int, int) {
	return windowWidth, windowHeight
}

func main() {
	/*screenWidth := 50
	screenHeight := 80*/
	p := &Player{
		Entity{image.Pt(10, 10), colornames.Firebrick, '@'},
	}
	g := &game{
		grid:   grid.New(10, 10, windowWidth/10, windowHeight/10, 0, 0),
		player: p,
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
