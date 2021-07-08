package kabaddi

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var playerImage, _, _ = ebitenutil.NewImageFromFile("kabaddi/assets/character.png")

type Character struct {
	x float64
	y float64
}

func NewCharacter() *Character {
	return &Character{
		x: -1,
		y: -1,
	}
}

func (c *Character) SetPos(x, y float64) {
	c.x = x
	c.y = y
}

func (c *Character) GetPos() (float64, float64) {
	return c.x, c.y
}

func (c *Character) InPlay() bool {
	if (c.x >= 0) && (c.x < 4) && (c.y >= 0) && (c.y < 4) {
		return true
	}
	return false
}

func (c *Character) Draw(screen *ebiten.Image) {
	if c.InPlay() {
		op := &ebiten.DrawImageOptions{}
		x := (ScreenWidth-960)/2 + 60 + 240*c.x
		y := (ScreenHeight-960)/2 + 60 + 240*c.y
		op.GeoM.Translate(x, y)
		screen.DrawImage(playerImage, op)
	}
}
