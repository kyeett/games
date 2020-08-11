package main

import (
	"image"
	"image/color"
	"math"

	"github.com/ojrac/opensimplex-go"
	"github.com/peterhellberg/gfx"
)

const (
	//octaves     = 10
	//persistence = 0.7
	//divisor = 100

	// Used to generate ebiten shader example
	//octaves     = 4
	//persistence = 0.7
	//divisor     = 80

	octaves     = 6
	persistence = 0.7
	divisor     = 80
)

func main() {

	noise := opensimplex.New(1)

	// First pass to find min and max for seed
	var min float64 = math.MaxFloat64
	var max float64 = -math.MaxFloat64
	for y := 0.0; y < 640.0; y++ {
		for x := 0.0; x < 480.0; x++ {
			v := octaveSimplex(noise, x, y, octaves, persistence)
			min = gfx.MathMin(v, min)
			max = gfx.MathMax(v, max)
		}
	}
	diff := max - min

	img := image.NewGray(image.Rect(0, 0, 640, 480))
	for y := 0; y < 640; y++ {
		for x := 0; x < 480; x++ {
			v := octaveSimplex(noise, float64(x), float64(y), octaves, persistence)
			// Normalize 0-->1
			v = (v - min) / diff
			clr := color.Gray{Y: uint8(256 * v)}
			img.Set(y, x, clr)
		}
	}

	gfx.SavePNG("noise.png", img)
}

// Inspired by https://adrianb.io/2014/08/09/perlinnoise.html
func octaveSimplex(noise opensimplex.Noise, x float64, y float64, octaves int, persistance float64) float64 {
	v := 0.0
	frequency := 1.0
	amplitude := 1.0

	for i := 0; i < octaves; i++ {
		v += noise.Eval2(x*frequency/divisor, y*frequency/divisor) * amplitude

		amplitude *= persistence
		frequency *= 2
	}

	return v
}
