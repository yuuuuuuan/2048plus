package game2048

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 840
	ScreenHeight = 1200
	boardSize    = 4
)

// Game represents a game state.
type Game struct {
	state      int
	input      *Input
	board      *Board
	boardImage *ebiten.Image
	//button     *Button
	//profile    *Profile
}

// NewGame generates a new Game object.
func NewGame() (*Game, error) {
	g := &Game{
		//state: gamestart,
		input: NewInput(),
		board: NewBoard(boardSize),
		//button: NewButton("restart", 300, 300),
	}

	return g, nil
}

func (g *Game) Layout(int, int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	//g.profile.button1.Update(g.input)
	//g.profile.button2.Update(g.input)
	g.input.Update()
	if err := g.board.Update(g.input); err != nil {
		return err
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//g.button.Draw(screen)
	if g.boardImage == nil {
		g.boardImage = ebiten.NewImage(g.board.Size())
	}
	screen.Fill(backgroundColor)
	g.board.Draw(g.boardImage)
	op := &ebiten.DrawImageOptions{}
	sw, sh := screen.Bounds().Dx(), screen.Bounds().Dy()
	bw, bh := g.boardImage.Bounds().Dx(), g.boardImage.Bounds().Dy()
	x := (sw - bw) / 2
	y := (sh - bh) / 2
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(g.boardImage, op)
}
