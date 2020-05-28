package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/kyeett/games/planes/config"
	"github.com/kyeett/games/planes/sprites"
	"github.com/kyeett/games/util"
	"github.com/peterhellberg/gfx"
	"math"
	"strconv"
)

func (g *game) Draw(screen *ebiten.Image) {

	g.background.Draw(screen)

	g.spikes.Draw(screen)
	g.foreground.Draw(screen)

	g.drawGoalScroller(screen)

	g.plane.Draw(screen)

	switch g.state {
	case "running":
		g.drawScore(screen)
	case "new":
		drawCenteredImage(screen, sprites.TextGetReady, gfx.V(0, 0))
		drawCenteredImage(screen, sprites.PressEnterToStart, gfx.V(0, 70))
		drawCredits(screen)
	case "game-over":
		drawCenteredImage(screen, sprites.TextGameOver, gfx.V(0, 0))
		drawCenteredImage(screen, sprites.PressEnterToRestart, gfx.V(0, 70))
		drawCredits(screen)
	case "victory":
		g.drawScore(screen)
		drawCenteredImage(screen, sprites.TextVictory, gfx.V(0, 0))
		drawCenteredImage(screen, sprites.PressEnterToRestart, gfx.V(0, 70))
		drawCredits(screen)
	}

	if config.Debug {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("State: %s", g.state), 0, 25)
	}

}

func drawCredits(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "game art from kenney.nl", screenWidth - 150, screenHeight-25)
}

func (g *game) drawGoalScroller(screen *ebiten.Image) {
	if g.posX >= goalAt-screenWidth {
		offsetX := goalAt - g.posX
		offsetX = gfx.Clamp(offsetX, 0, math.MaxFloat64)
		opt := &ebiten.DrawImageOptions{}
		opt.GeoM.Translate(offsetX, 0)
		screen.DrawImage(g.goal, opt)
	}
}

func digit(num, place int) int {
	r := num % int(math.Pow(10, float64(place)))
	return r / int(math.Pow(10, float64(place-1)))
}

func (g *game) drawScore(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Scale(0.5, 0.5)
	opt.GeoM.Translate(screenWidth-5, 5)
	scaled := int(g.posX*scaleDown)
	ss := strconv.Itoa(scaled)

	for i := 0; i < len(ss); i++ {
		img := sprites.Numbers[digit(scaled, i+1)]
		opt.GeoM.Translate(-28, 0)
		screen.DrawImage(img, opt)
	}
}

func drawCenteredImage(screen, img *ebiten.Image, offset gfx.Vec) {
	pos := util.CenterBoundsOnBounds(img.Bounds(), screen.Bounds())
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(float64(pos.X)+offset.X, float64(pos.Y)+offset.Y)
	screen.DrawImage(img, opt)
}