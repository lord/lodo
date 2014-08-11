package breakout

type paddle struct {
	x, y, w, h float64
}

func makePaddle(x, y, w, h float64) paddle {
	return paddle{
		x: x,
		y: y,
		w: w,
		h: h,
	}
}

// func (b *ball) step() {
// 	b.x += b.vx
// 	b.y += b.vy
// 	if b.x >= boardWidth {
// 		b.vx = math.Abs(b.vx) * -1
// 	}
// 	if b.y >= boardHeight {
// 		b.vy = math.Abs(b.vy) * -1
// 	}
// 	if b.x <= 0 {
// 		b.vx = math.Abs(b.vx)
// 	}
// 	if b.y <= 0 {
// 		b.vy = math.Abs(b.vy)
// 	}
// }

// func (b *ball) draw(board *core.Board) {
// 	board.DrawRect(int(b.x+0.5-b.r/2), int(b.y+0.5-b.r/2), int(b.x+0.5+b.r/2), int(b.y+0.5+b.r/2), core.MakeColor(0, 20, 0))
// }
