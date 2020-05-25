package spikes

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/ebitendrawutil"
	"github.com/kyeett/games/planes/config"
	"github.com/kyeett/games/planes/sprites"
	"github.com/peterhellberg/gfx"
	"golang.org/x/image/colornames"
	"math/rand"
)

type Spike struct {
	gfx.Vec
	offset  gfx.Vec
	hitbox  gfx.Rect
	hitbox2 gfx.Rect
	sprite  *ebiten.Image
}

func (p *Spike) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(p.X, p.Y)
	screen.DrawImage(p.sprite, opt)

	if config.Debug {
		ebitendrawutil.DrawRect(screen, p.hitbox.Moved(p.Vec), colornames.Red)
		ebitendrawutil.DrawRect(screen, p.hitbox2.Moved(p.Vec), colornames.Red)
	}
}

func (p *Spike) CollidesWith(r gfx.Rect) bool {
	return p.hitbox.Moved(p.Vec).Overlaps(r) || p.hitbox2.Moved(p.Vec).Overlaps(r)
}

type Spikes struct {
	spikes []*Spike
	width  float64
	height float64
}

func (s *Spikes) Draw(screen *ebiten.Image) {
	for _, s := range s.spikes {
		s.Draw(screen)
	}
}

func (ss *Spikes) Scroll(dx float64) {
	for _, s := range ss.spikes {
		s.X -= dx

		// Replace with new rock
		if s.X+rockWidth < 0 {
			*s = randomRock(ss.width, offset(ss.width/count), ss.height)
		}
	}
}

func (ss *Spikes) CollidesWith(hitbox gfx.Rect) bool {
	for _, s := range ss.spikes {
		if s.CollidesWith(hitbox) {
			return true
		}
	}
	return false
}

const (
	count      = 3
	padding    = 40
	rockWidth  = 108
	rockHeight = 240
)

func offset(sliceWidth float64) float64 {
	return (sliceWidth-2*padding)*rand.Float64() + padding
}

func New(screenWidth, screenHeight float64) Spikes {
	sliceWidth := screenWidth / count

	var ss []*Spike
	for i := 0.0; i < count; i++ {
		downRock := randomRock(i*sliceWidth, offset(screenWidth), screenHeight)
		ss = append(ss, &downRock)
	}

	return Spikes{
		spikes: ss,
		width:  screenWidth,
		height: screenHeight,
	}
}

func randomRock(x, offset, height float64) Spike {
	if rand.Int()%2 == 0 {
		return upRock(x, offset, height)
	}
	return downRock(x, offset)
}

func upRock(x, offsetX, screenHeight float64) Spike {
	rock := Spike{
		Vec:     gfx.V(x, screenHeight-rockHeight-15),
		offset:  gfx.V(offsetX, 0),
		hitbox:  gfx.R(18, 0, 94, 120).Moved(gfx.V(0, 120)),
		hitbox2: gfx.R(53, 0, 67, 120),

		sprite: sprites.RockUp,
	}
	return rock
}

func downRock(x, offsetX float64) Spike {
	downRock := Spike{
		Vec:     gfx.V(x, 0),
		offset:  gfx.V(offsetX, 0),
		hitbox:  gfx.R(18, 0, 94, 120),
		hitbox2: gfx.R(53, 0, 67, 120).Moved(gfx.V(0, 120)),

		sprite: sprites.RockDown,
	}
	return downRock
}
