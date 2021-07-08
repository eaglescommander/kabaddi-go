package kabaddi

import (
	_ "image/jpeg"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	BoardSize = 4
)

type Board struct {
	tiles      [][]*Tile
	boardImage *ebiten.Image
	character  *Character
}

func NewBoard() (*Board, error) {
	tiles := [][]*Tile{}

	for i := 0; i < BoardSize; i++ {
		row := []*Tile{}
		for j := 0; j < BoardSize; j++ {
			row = append(row, NewTile(i, j))
		}
		tiles = append(tiles, row)
	}

	image, _, err := ebitenutil.NewImageFromFile("kabaddi/assets/board.jpg")
	if err != nil {
		return nil, err
	}

	var board = &Board{
		tiles:      tiles,
		boardImage: image,
		character:  NewCharacter(),
	}

	return board, nil
}

func (b *Board) GetTiles() [][]*Tile {
	return b.tiles
}

func (b *Board) GetCharacter() *Character {
	return b.character
}

func (b *Board) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate((ScreenWidth-960)/2, (ScreenHeight-960)/2)
	screen.DrawImage(b.boardImage, op)

	for _, row := range b.GetTiles() {
		for _, tile := range row {
			tile.Draw(screen)
		}
	}

	b.character.Draw(screen)
}
