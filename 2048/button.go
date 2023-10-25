package game2048

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

const (
	buttonWidth  = 100
	buttonHeight = 40
)

type Button struct {
	msg     string
	x       int
	y       int
	width   int
	height  int
	ifClick bool
}

func NewButton(msg string, x int, y int) *Button {
	b := &Button{
		msg:     msg,
		x:       x,
		y:       y,
		width:   buttonWidth,
		height:  buttonHeight,
		ifClick: false,
	}
	return b
}

func (button *Button) Update(i *Input) {
	x, y, b := i.Click()
	if b && button.x < x && x < (button.x+buttonWidth) && button.y < y && y < (button.y+buttonHeight) {
		button.ifClick = true
	}
}

func (button *Button) Draw(screen *ebiten.Image) {
	buttonImage := ebiten.NewImage(tileSize, tileSize)
	buttonImage.Fill(color.Black)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(button.x), float64(button.y))
	screen.DrawImage(buttonImage, op)
}
