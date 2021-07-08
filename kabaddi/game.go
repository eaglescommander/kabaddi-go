package kabaddi

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	ScreenWidth  = 1920
	ScreenHeight = 1080
	BoardStartX  = 480
	BoardStartY  = 60
)

type Game struct {
	board Board
	round int
	x     int
	y     int
}

func NewGame() (*Game, error) {
	board, err := NewBoard()
	if err != nil {
		return nil, err
	}

	game := &Game{
		board: *board,
		round: 0,
		x:     0,
		y:     0,
	}

	return game, nil
}

func (g *Game) UpdateRound() {
	g.round = (g.round + 1) % 2
}

func (g *Game) Update() error {
	char := g.board.GetCharacter()

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		x = x - BoardStartX
		y = y - BoardStartY

		if x >= 0 && x <= 960 && y >= 0 && y <= 960 {
			char.SetPos(float64(x/240), float64(y/240))
		}
	}

	if char.InPlay() {
		x, y := char.GetPos()
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
			char.SetPos(float64(x-1), float64(y))
		} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
			char.SetPos(float64(x-1), float64(y-1))
		} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
			char.SetPos(float64(x-1), float64(y+1))
		} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
			char.SetPos(float64(x+1), float64(y))
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.board.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
