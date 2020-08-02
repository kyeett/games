package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/games/particles"
	"github.com/kyeett/games/particles/generators"
	"github.com/kyeett/games/particles/shapes"
	"github.com/peterhellberg/gfx"
	"golang.org/x/image/colornames"
	"log"
	"math"
)

var (
	heart *ebiten.Image
	dot   *ebiten.Image
)

func init() {
	img := gfx.MustOpenImage("examples/particlesystem/assets/heart.png")
	heart, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	img = gfx.MustOpenImage("examples/particlesystem/assets/dot.png")
	dot, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

type game struct {
	particleSystem *particles.ParticleSystem
}

func (g game) Update(_ *ebiten.Image) error {
	cx, cy := ebiten.CursorPosition()
	g.particleSystem.Move(float64(cx-25), float64(cy-25))
	g.particleSystem.Update(0.016)
	return nil
}

func (g game) Draw(screen *ebiten.Image) {
	g.particleSystem.Draw(screen)
}

const (
	screenWidth  = 800
	screenHeight = 600
)

func (g game) Layout(w, h int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	pink := colornames.Pink
	pink.A = 30

	rate := 100.0
	options := particles.Options{
		PositionX: screenWidth / 2,
		PositionY: screenHeight / 2,

		StartLifetime: &generators.FloatRandomBetweenTwoConstants{1, 2},
		StartSpeed:    &generators.FloatRandomBetweenTwoConstants{1, 5},
		StartSize:     generators.FloatConstant{0.1},

		Color: generators.ColorRandomBetweenTwoConstants{
			Color1: colornames.Greenyellow,
			Color2: colornames.Rosybrown,
		},

		// Emission
		Rate: &rate,

		Gravity: gfx.V(0, 9.81/10),

		// Shape
		Shape: shapes.NewCone(math.Pi/8, 50),

		Material: dot,
	}
	ps := particles.NewParticleSystem(options)

	g := &game{
		particleSystem: ps,
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
