package sprites

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/games/planes/assets"
	"github.com/kyeett/games/util"
)

var (
	Background, ForegroundDirt *ebiten.Image

	RedPlane1 *ebiten.Image

	RockUp, RockDown *ebiten.Image
)

func init() {
	Background = util.LoadAssetImageOrFatal(assets.Asset, "assets/background.png")
	ForegroundDirt = util.LoadAssetImageOrFatal(assets.Asset, "assets/groundDirt.png")
	RedPlane1 = util.LoadAssetImageOrFatal(assets.Asset, "assets/Planes/planeRed1.png")

	RockUp = util.LoadAssetImageOrFatal(assets.Asset, "assets/rock.png")
	RockDown = util.LoadAssetImageOrFatal(assets.Asset, "assets/rockDown.png")
}