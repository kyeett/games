package main

import (
	"io/ioutil"
	"log"

	"github.com/hajimehoshi/ebiten"
)

var (
	shader *ebiten.Shader
)

func MustLoad(filename string) []byte {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func init() {
	s, err := ebiten.NewShader(MustLoad("shader.go"))
	if err != nil {
		log.Fatal(err)
	}
	shader = s
}

type game struct {
	time int64
}

func (g *game) Update(_ *ebiten.Image) error {
	g.time++
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	cx, cy := ebiten.CursorPosition()
	op := &ebiten.DrawRectShaderOptions{}
	op.Uniforms = append(op.Uniforms,
		[]float32{float32(screenWidth), float32(screenHeight)}, // Resolution
		[]float32{float32(cx), float32(cy)},                    // Mouse
		float32(g.time)/60,                                     // Time
	)
	screen.DrawRectShader(screenWidth, screenHeight, shader, op)
}

const (
	screenWidth  = 600
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
