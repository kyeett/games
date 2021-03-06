package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/kyeett/games/examples/movement/grid/assets"
	"github.com/kyeett/games/pathfinding"
	"github.com/kyeett/games/util"
	"github.com/kyeett/games/util/cursor"
	"github.com/kyeett/games/util/grid"
	"github.com/kyeett/games/util/move"
	"github.com/peterhellberg/gfx"
	"golang.org/x/image/colornames"
	"image"
	"log"
)

type player struct {
	gridPos image.Point
	pos     gfx.Vec
	path    []PathPoint
	moving  bool
	frame   int
}

var playerSprite *ebiten.Image

func init() {
	playerSprite = util.LoadAssetImageOrFatal(assets.Asset, "assets/spr_character.png")
}

func (p *player) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(8, 8)
	opt.GeoM.Translate(p.pos.X, p.pos.Y)
	screen.DrawImage(playerSprite.SubImage(image.Rect(0, 0, 16, 16).Add(image.Pt(p.frame*16, 0))).(*ebiten.Image), opt)
}

type game struct {
	player *player
	target gfx.Vec
	grid   *grid.Grid
	world  *pathfinding.World
}

func (g *game) Update(_ *ebiten.Image) error {

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		t, found := g.grid.ToPoint(cursor.Position())
		if found {
			g.startMovingTowards(t)
		}

	}

	if !g.player.moving {
		return nil
	}

	powerRemaining := 2.0
	var reached bool
	for powerRemaining > 0 {
		g.player.pos, powerRemaining, reached = move.Towards(g.player.pos, g.grid.MustVec(g.player.path[0].Point), powerRemaining)
		if reached {
			g.player.path = g.player.path[1:]

			if len(g.player.path) == 0 {
				fmt.Println("target reached")
				g.player.moving = false
				return nil
			}

			// TODO: Check if next move is possible
			// If not available, wait a while, and then potentially treat it as a solid target
			g.updateFirstPosition()
		}
	}

	return nil
}

const (
	screenWidth  = 800
	screenHeight = 600

	tileSize = 32

	paddingX = 1
	paddingY = 1

	gridWidth  = 10
	gridHeight = 10
)

func (g *game) Layout(w, h int) (int, int) {
	return (tileSize + paddingX) * gridWidth, (tileSize + paddingY) * gridHeight
}

var (
	square *ebiten.Image
)

func main() {
	square, _ = ebiten.NewImage(tileSize, tileSize, ebiten.FilterDefault)
	square.Fill(colornames.White)

	g := &game{
		player: &player{
			gridPos: gfx.Pt(0, 0),
		},
		grid:  grid.New(tileSize, tileSize, gridHeight, gridWidth, paddingX, paddingY),
		world: pathfinding.New(gridWidth, gridHeight),
	}

	for i := 0; i < 5; i++ {
		g.world.Get(2, i).Kind = pathfinding.KindBlocker
	}
	for i := 0; i < 7; i++ {
		g.world.Get(8, 3+i).Kind = pathfinding.KindBlocker
	}

	g.startMovingTowards(gfx.Pt(9, 9))

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
