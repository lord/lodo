package pong

import (
	_ "fmt"
	"github.com/lord/lodo/core"
	"math"
)

type paddle struct {
	// x: coord r/l
	// y: coor back / forth
	// x/y at edge of paddle, top of square
	// w: width of paddle
	x, y, w         float64
	row             int
	speedMultiplier float64
	speedMax        float64
	color           core.Color
}

func makePaddle(x, y, w float64, row int, color core.Color) paddle {
	return paddle{
		x:               x,
		y:               y,
		w:               w,
		row:             row,
		speedMultiplier: 0.3,
		speedMax:        2,
		color:           color,
	}
}

func (p *paddle) step(board *core.Board) {
	total := 0
	count := 0
	for x := 0; x < 5; x++ {
		if board.CheckDown(x, p.row) {
			targetx, _ := board.GetSquare(x, p.row)
			total += targetx
			count += 1
		}
	}
	if count > 0 {
		deltax := float64(total)/float64(count) - p.x
		if math.Abs(deltax*p.speedMultiplier) > math.Abs(p.speedMax) {
			if deltax > 0 {
				p.x += p.speedMax
			} else {
				p.x -= p.speedMax
			}
		} else {
			p.x += deltax * p.speedMultiplier
		}
	}
}

func (p *paddle) hit(b *ball) bool {
	if (b.y >= p.y && b.y <= p.y + 1) && (b.x >= p.x && b.x <= p.x + p.w) {
		return true
	}	
	return false
}

func (p *paddle) draw(board *core.Board) {
	x := int(p.x + 0.5)
	y := int(p.y + 0.5)
	board.DrawRect(x, y, x+int(p.w), y+1, p.color)
}
