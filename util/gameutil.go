package util

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/peterhellberg/gfx"
	"log"
	"math"
)

type ScrollingImage struct{
	img     *ebiten.Image
	scrollX float64
	offset gfx.Vec
}

func NewScrollingImage(img *ebiten.Image, offset gfx.Vec) *ScrollingImage {
	return &ScrollingImage{
		img:     img,
		scrollX: 0,
		offset: offset,
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
