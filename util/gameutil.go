package util

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/peterhellberg/gfx"
	"image"
	"image/color"
	"log"
	"math"
)

type ScrollingImage struct {
	img     *ebiten.Image
	scrollX float64
	offset  gfx.Vec
}

func NewScrollingImage(img *ebiten.Image, offset gfx.Vec) *ScrollingImage {
	return &ScrollingImage{
		img:     img,
		scrollX: 0,
		offset:  offset,
	}
}

func (s *ScrollingImage) Draw(screen *ebiten.Image) {
	imgWidth := float64(s.img.Bounds().Dx())

	// Leftmost image
	translateX := math.Mod(-s.scrollX, imgWidth)

	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(s.offset.X, s.offset.Y)

	opt.GeoM.Translate(translateX, 0)
	screen.DrawImage(s.img, opt)

	opt.GeoM.Translate(imgWidth, 0)
	screen.DrawImage(s.img, opt)
}

func (s *ScrollingImage) Scroll(dx float64) {
	s.scrollX += dx
}

func LoadAssetImageOrFatal(assets func(name string) ([]byte, error), path string) *ebiten.Image {
	b, err := assets(path)
	if err != nil {
		log.Fatal(err)
	}
	return LoadImageOrFatal(b)
}

func LoadImageOrFatal(imageBytes []byte) *ebiten.Image {
	src, err := gfx.DecodeImageBytes(imageBytes)
	if err != nil {
		log.Fatal(err)
	}

	img, err := ebiten.NewImageFromImage(src, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func MustNewRectangle(w,h int, c color.Color) *ebiten.Image {
	img, err := ebiten.NewImage(w, h, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	img.Fill(c)
	return img
}





// PackImages draws all input images on a single image
func PackImages(images []*ebiten.Image) (*ebiten.Image, error) {
	var totalWidth int
	for _, img := range images {
		totalWidth += img.Bounds().Dx()
	}

	// Create image and draw characters
	opt := &ebiten.DrawImageOptions{}
	packedImg, err := ebiten.NewImage(totalWidth, 64, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	for _, img := range images {
		packedImg.DrawImage(img, opt)
		opt.GeoM.Translate(float64(img.Bounds().Dx()), 0)
	}
	return packedImg, nil
}

func CenterBoundsOnBounds(small, large image.Rectangle) image.Point {
	dx := (large.Dx() - small.Dx()) / 2
	dy := (large.Dy() - small.Dy()) / 2

	return large.Min.Add(image.Pt(dx, dy))
}

func OptScaleByColor(opt *ebiten.DrawImageOptions, clr color.Color) {
	c, _ := colorful.MakeColor(clr)
	r, g, b := c.LinearRgb()
	opt.ColorM.Scale(r, g, b, 1)
}
