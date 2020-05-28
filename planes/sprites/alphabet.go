package sprites

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/games/planes/assets"
	"github.com/kyeett/games/util"
	"log"
)

var (
	TextGameOver, TextGetReady, TextVictory *ebiten.Image
	PressEnterToStart, PressEnterToRestart *ebiten.Image
	TextControls *ebiten.Image

	Numbers = map[int]*ebiten.Image{}
)

func init() {
	images := loadString("VICTORY")
	TextVictory, _ = util.PackImages(images)

	TextControls = loadStringScaled("W/S TO FLY", 0.5)
	PressEnterToStart = loadStringScaled("PRESS \"ENTER\" TO START", 0.5)
	PressEnterToRestart = loadStringScaled("PRESS \"ENTER\" TO RESTART", 0.5)

	TextGameOver = util.LoadAssetImageOrFatal(assets.Asset, "assets/UI/textGameOver.png")
	TextGetReady = util.LoadAssetImageOrFatal(assets.Asset, "assets/UI/textGetReady.png")

	for i, r := range "0123456789" {
		path := fmt.Sprintf("assets/Numbers/number%s.png", string(r))
		Numbers[i] = util.LoadAssetImageOrFatal(assets.Asset, path)
	}
}

func loadStringScaled(str string, scale float64) *ebiten.Image {
	images := loadString(str)
	enterTmp, err := util.PackImages(images)
	if err != nil {
		log.Fatal(err)
	}
	img, err := ebiten.NewImage(enterTmp.Bounds().Dx()/2, enterTmp.Bounds().Dy()/2, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Scale(scale, scale)
	img.DrawImage(enterTmp, opt)
	return img
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
		case '/':
			img = util.LoadAssetImageOrFatal(assets.Asset, "assets/Letters/letterSlash.png")
		default:
			path := fmt.Sprintf("assets/Letters/letter%s.png", string(r))
			img = util.LoadAssetImageOrFatal(assets.Asset, path)
		}
		images = append(images, img)
	}
	return images
}
