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

func (b *ball) init (x, y, vx, vy float64) {
	b.x = x
	b.y = y
	b.vx = vx
	b.vy = vy
}

func (b *ball) step() {
	b.x += b.vx
	b.y += b.vy
	if b.x >= boardWidth-1 {
		b.vx = math.Abs(b.vx) * -1
		core.PlayWave();
	}
	if paddle1.hit(b) {
		b.vy = math.Abs(b.vy) * -1
		core.PlayWave();
	}
	if b.y >= boardHeight {
		setMode(p2_score)
		core.PlayWave();
	}
	if b.x <= 0 {
		b.vx = math.Abs(b.vx)
		core.PlayWave();
	}
	if b.y <= 0 || paddle2.hit(b) {
		b.vy = math.Abs(b.vy)
		core.PlayWave();
	}
	if b.y <= 0 {
		setMode(p1_score)
		core.PlayWave();
	}
}

func (b *ball) draw(board *core.Board) {
	board.DrawSmallCircle(b.x, b.y, b.color)
}
