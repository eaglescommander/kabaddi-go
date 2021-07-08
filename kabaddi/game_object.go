package kabaddi

import "github.com/hajimehoshi/ebiten/v2"

type GameObject interface {
	Draw(screen *ebiten.Image)
}

func DrawObject(obj GameObject, screen *ebiten.Image) {
	obj.Draw(screen)
}
