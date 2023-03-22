package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Button struct {
	Image         *ebiten.Image
	IdleButton    *Button
	PressedButton *Button
	HoverButton   *Button
	X             int
	Y             int
	YPaddingLeft  int
	YPaddingRight int
	XPaddingLeft  int
	XPaddingRight int
	Clicks        int
	Clicked       bool
}

func (btn *Button) ShowIdle(screen *ebiten.Image) {
	show(screen, btn.IdleButton)
}

func show(screen *ebiten.Image, btn *Button) {
	btnop := &ebiten.DrawImageOptions{}
	btnop.GeoM.Translate(float64(btn.X), float64(btn.Y))
	screen.DrawImage(btn.Image, btnop)
}

func (btn *Button) IsClicked(x, y int) bool {
	if x >= btn.X+btn.XPaddingLeft && x <= btn.X+btn.Image.Bounds().Dx()-btn.XPaddingRight {
		if y >= btn.Y+btn.YPaddingLeft && y <= btn.Y+btn.Image.Bounds().Dy()-btn.YPaddingRight {
			return true
		}
	}
	return false
}

func (btn *Button) Click(screen *ebiten.Image) {
	screen.Clear()
	screen.Fill(color.RGBA{213, 218, 241, 255})
	show(screen, btn.PressedButton)
}
