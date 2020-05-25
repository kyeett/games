package plane

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/ebitendrawutil"
	"github.com/kyeett/games/planes/config"
	"github.com/kyeett/games/planes/sprites"
	"github.com/peterhellberg/gfx"
	"golang.org/x/image/colornames"
)

type Plane struct {
	gfx.Vec
	hitbox  gfx.Rect
	Visible bool
}

func New(vec gfx.Vec) *Plane {
	return &Plane{
		Vec:     vec,
		hitbox:  gfx.R(5, 2, 40, 35),
		Visible: true,
	}
}

func (p *Plane) Draw(screen *ebiten.Image) {
	if !p.Visible {
		return
	}

	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Scale(0.6, 0.6)
	opt.GeoM.Translate(p.X, p.Y)
	screen.DrawImage(sprites.RedPlane1, opt)

	if config.Debug {
		ebitendrawutil.DrawRect(screen, p.hitbox.Moved(p.Vec), colornames.White)
	}
}

func (p *Plane) Hitbox() gfx.Rect {
	return p.hitbox.Moved(p.Vec)
}
