package main

import (
	"log"

	"github.com/eaglescommander/kabaddi-go/kabaddi"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game, err := kabaddi.NewGame()
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(kabaddi.ScreenWidth, kabaddi.ScreenHeight)
	ebiten.SetWindowTitle("Kabaddi: Go")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
