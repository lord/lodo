package maze

import "github.com/lord/lodo/core"

type Footprint struct {
	x, y      int
	alpha     int
	direction Direction
}

func MakeFootprint(x, y int, direction Direction) *Footprint {
	return &Footprint{
		x:         x,
		y:         y,
		alpha:     100,
		direction: direction,
	}
}

const alphaSpeed = 40

var footprintColor = core.MakeColor(0, 31, 0)

func (footprint *Footprint) Step(game *Game) {
	footprint.alpha -= alphaSpeed
	if footprint.alpha <= 0 {
		game.DeleteObject(footprint)
	}
}

func (footprint *Footprint) Draw(board *core.Board) {
	x := footprint.x * 7
	y := footprint.y * 7
	color := footprintColor.WithAlpha(float64(footprint.alpha) / 800)
	switch footprint.direction {
	case Up:
		board.DrawRect(x+2, y+1, x+4, y+6, color)
		board.DrawPixel(x+3, y, color)
		board.DrawPixel(x+1, y+2, color)
		board.DrawPixel(x+5, y+2, color)
	case Down:
		board.DrawRect(x+2, y, x+4, y+5, color)
		board.DrawPixel(x+3, y+6, color)
		board.DrawPixel(x+1, y+4, color)
		board.DrawPixel(x+5, y+4, color)
	case Left:
		board.DrawRect(x+1, y+2, x+6, y+4, color)
		board.DrawPixel(x, y+3, color)
		board.DrawPixel(x+2, y+1, color)
		board.DrawPixel(x+2, y+5, color)
	case Right:
		board.DrawRect(x, y+2, x+5, y+4, color)
		board.DrawPixel(x+6, y+3, color)
		board.DrawPixel(x+4, y+1, color)
		board.DrawPixel(x+4, y+5, color)
	}
}
