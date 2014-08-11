package breakout

import "math"
import "github.com/lord/lodo/core"

type ball struct {
	x, y, r, vx, vy float64
	color           core.Color
}

func makeBall(x, y, vx, vy float64, c core.Color) ball {
	return ball{
		x:     x,
		y:     y,
		vx:    vx,
		vy:    vy,
		color: c,
	}
}

func (b *ball) step() {
	b.x += b.vx
	b.y += b.vy
	if b.x >= boardWidth {
		b.vx = math.Abs(b.vx) * -1
	}
	if b.y >= boardHeight {
		b.vy = math.Abs(b.vy) * -1
	}
	if b.x <= 0 {
		b.vx = math.Abs(b.vx)
	}
	if b.y <= 0 {
		b.vy = math.Abs(b.vy)
	}
}

func (b *ball) draw(board *core.Board) {
	board.DrawSmallCircle(b.x, b.y, b.color)
}
