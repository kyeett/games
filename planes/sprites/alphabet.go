package sprites

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/games/planes/assets"
	"github.com/kyeett/games/util"
)

var (
	TextGameOver, TextGetReady, PressEnterToStart *ebiten.Image

	Numbers = map[int]*ebiten.Image{}
)

func init() {
	// Load images and calculate total width
	images := loadString("PRESS \"ENTER\" TO START")

	enterTmp, _ := util.PackImages(images)
	PressEnterToStart, _ = ebiten.NewImage(enterTmp.Bounds().Dx()/2, enterTmp.Bounds().Dy()/2, ebiten.FilterDefault)
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Scale(0.5, 0.5)
	PressEnterToStart.DrawImage(enterTmp, opt)

	TextGameOver = util.LoadAssetImageOrFatal(assets.Asset, "assets/UI/textGameOver.png")
	TextGetReady = util.LoadAssetImageOrFatal(assets.Asset, "assets/UI/textGetReady.png")

	for i, r := range "0123456789" {
		path := fmt.Sprintf("assets/Numbers/number%s.png", string(r))
		Numbers[i] = util.LoadAssetImageOrFatal(assets.Asset, path)
	}
}

func loadString(text string) []*ebiten.Image {
	var images []*ebiten.Image
	for _, r := range text {

		var img *ebiten.Image
		switch r {
		case ' ':
			img, _ = ebiten.NewImage(30, 1, ebiten.FilterDefault)
		case '"':
			img = util.LoadAssetImageOrFatal(assets.Asset, "assets/Letters/letterQuote.png")
		default:
			path := fmt.Sprintf("assets/Letters/letter%s.png", string(r))
			img = util.LoadAssetImageOrFatal(assets.Asset, path)
		}
		images = append(images, img)
	}
	return images
}
