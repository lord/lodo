package breakout

import "math"
import "github.com/lord/lodo/core"

type ball struct {
	x, y, r, vx, vy float64
}

func makeBall(x, y, r, vx, vy float64) ball {
	return ball{
		x:  x,
		y:  y,
		r:  r,
		vx: vx,
		vy: vy,
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
	board.DrawRect(int(b.x+0.5-b.r/2), int(b.y+0.5-b.r/2), int(b.x+0.5+b.r/2), int(b.y+0.5+b.r/2), core.MakeColor(0, 20, 0))
}
