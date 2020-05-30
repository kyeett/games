package main

import (
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
	"github.com/hajimehoshi/ebiten"
)

type game struct{}

func (g game) Update(_ *ebiten.Image) error {
	return nil
}

func (g game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "template")
}

const (
	screenWidth  = 800
	screenHeight = 600
)

func (g game) Layout(w, h int) (int, int) {
	return screenWidth, screenHeight
}


func main() {
	g := &game{}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
