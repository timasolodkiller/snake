package board

import (
	"game_snake/point"
)
type Board struct {
	Width, Height int
}

func Initialize(w, h int) *Board {
	return &Board{
		Width: w,
		Height: h,
	}
}

func (b *Board) IsCrashed(p point.Point) bool {
	return (p.X >= 0) && (p.X < b.Width) && (p.Y >= 0) && (p.Y < b.Height)
}