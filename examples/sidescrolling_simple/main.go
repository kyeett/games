package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/kyeett/games/util/sidescroller"
	"github.com/peterhellberg/gfx"
	"log"
	"github.com/hajimehoshi/ebiten"
)

type game struct {
	scroller *sidescroller.Scroller
}

func (g *game) Update(_ *ebiten.Image) error {
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	s := fmt.Sprintf("%v", g.scroller.Force())
	ebitenutil.DebugPrintAt(screen, s, screenWidth/2, screenHeight/2)
}

const (
	screenWidth  = 800
	screenHeight = 600
)

func (g *game) Layout(w, h int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	g := &game{
		sidescroller.NewEbiten(gfx.R(0, 0, screenWidth, screenHeight)),
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
