package kabaddi

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var mineImage, _, _ = ebitenutil.NewImageFromFile("kabaddi/assets/mine.png")
var pointImage, _, _ = ebitenutil.NewImageFromFile("kabaddi/assets/point.png")

type Tile struct {
	hasTrap  bool
	hasPoint bool
	x        float64
	y        float64
}

func NewTile(i, j int) *Tile {
	return &Tile{
		hasTrap:  false,
		hasPoint: false,
		x:        float64(i),
		y:        float64(j),
	}
}

func (t *Tile) AddTrap() {
	t.hasTrap = true
}

func (t *Tile) AddPoint() {
	t.hasPoint = true
}

func (t *Tile) Clear() {
	t.hasTrap = false
	t.hasPoint = false
}

func (t *Tile) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	x := (ScreenWidth-960)/2 + 60 + 240*t.x
	y := (ScreenHeight-960)/2 + 60 + 240*t.y
	op.GeoM.Translate(x, y)
	if t.hasTrap {
		screen.DrawImage(mineImage, op)
	} else if t.hasPoint {
		screen.DrawImage(pointImage, op)
	}
}
