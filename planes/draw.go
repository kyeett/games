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
	"time"
)

var drawDuration time.Duration
var drawI int64
var updateDuration time.Duration
var updateI int64

func timed(name string, f func()) {
	start := time.Now()
	f()

	switch name {
	case "draw":
		updateDuration = updateDuration + time.Since(start)
		updateI++


	case "update":
		drawDuration = drawDuration + time.Since(start)
		drawI++

	}

	if drawI%1000 == 5 {
		fmt.Println("tot update", time.Duration(int64(updateDuration)/updateI))
		fmt.Println("tot draw", time.Duration(int64(drawDuration)/drawI))
	}

	fmt.Println(name, time.Since(start))

}

func (g *game) Draw(screen *ebiten.Image) {
	timed("draw", func() {

		g.background.Draw(screen)

		g.spikes.Draw(screen)
		g.foreground.Draw(screen)

		g.drawGoalScroller(screen)

		g.plane.Draw(screen)

		switch g.state {
		case "running":
		g.drawScore(screen)
		case "new":
			drawControls(screen)
			drawCenteredImage(screen, sprites.TextGetReady, gfx.V(0, 0))
			drawCenteredImage(screen, sprites.PressEnterToStart, gfx.V(0, 70))
			drawCredits(screen)
		case "game-over":
			g.drawScore(screen)
			drawControls(screen)
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
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("FPS: %0.2f", ebiten.CurrentFPS()), 0, 25)
			ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("State: %s", g.state), 0, 50)
		}

	})
}

func drawControls(screen *ebiten.Image) error {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(10, 10)
	return screen.DrawImage(sprites.TextControls, opt)
}

func drawCredits(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "game art from kenney.nl", screenWidth-150, screenHeight-25)
}

func (g *game) drawGoalScroller(screen *ebiten.Image) {
	if g.posX >= config.GoalAt-screenWidth {
		offsetX := config.GoalAt - g.posX
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
	scaled := int(g.posX * config.ScaleDown)
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
