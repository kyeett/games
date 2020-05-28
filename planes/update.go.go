package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/kyeett/games/planes/config"
	"github.com/peterhellberg/gfx"
	"math"
)

func (g *game) Update(_ *ebiten.Image) error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return fmt.Errorf("exit game")
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		config.Debug = !config.Debug
	}

	g.background.Scroll(g.velocityX / 6)
	g.spikes.Scroll(g.velocityX)
	g.foreground.Scroll(g.velocityX)

	switch g.state {
	case "running":
		g.stateRunning()
	case "victory":
		g.stateWinning()
	default:
		g.checkGameRestart()
	}

	return nil
}


func (g *game) stateWinning() {
	target := gfx.V(screenWidth + 100, screenHeight/2)
	dv := g.plane.To(target).Unit()
	g.plane.Vec = g.plane.Vec.Add(dv)

	g.checkGameRestart()
}

func (g *game) checkGameRestart() {
	// Reset game
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		g.newGame()
		g.state = "running"
	}
}

const (
	scaleDown = 0.06
)

const (
	goalAt = 500.0/scaleDown
)

func (g *game) handleUpDown() {
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown):
		g.plane.Y += g.speedFactor() * velocityY
		g.plane.Y = gfx.Clamp(g.plane.Y, 0, screenHeight-40)
	case ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp):
		g.plane.Y -= g.speedFactor() * velocityY
		g.plane.Y = gfx.Clamp(g.plane.Y, 0, screenHeight-40)
	}
}

func (g *game) stateRunning() {
	g.velocityX = g.speedFactor() * velocityX
	g.posX += g.velocityX
	g.handleUpDown()



	// Quick fix, if player is this close to the goal, they should not die :-D
	isVeryCloseToGoal := g.posX+100 >= goalAt

	// Check collision
	if g.spikes.CollidesWith(g.plane.Hitbox()) && !isVeryCloseToGoal {
		g.plane.Visible = false
		g.state = "game-over"
		return
	}

	if g.posX >= goalAt {
		g.state = "victory"
		return
	}

}

func (g *game) speedFactor() float64 {
	return (velocityX + scaleDown*math.Floor(g.posX/100)) / velocityX
}