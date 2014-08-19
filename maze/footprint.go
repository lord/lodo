package maze

import "github.com/lord/lodo/core"

type Footprint struct {
	x, y  int
	alpha int
	color core.Color
}

func MakeFootprint(x, y int, color core.Color) *Footprint {
	return &Footprint{
		x:     x,
		y:     y,
		alpha: 100,
		color: color,
	}
}

const alphaSpeed = 10

func (footprint *Footprint) Step(game *Game) {
	footprint.alpha -= alphaSpeed
	if footprint.alpha <= 0 {
		game.DeleteObject(footprint)
	}
}

func (footprint *Footprint) Draw(board *core.Board) {
	color := footprint.color.WithAlpha(float64(footprint.alpha) / 100)
	board.DrawSquare(footprint.x, footprint.y, color)
}
