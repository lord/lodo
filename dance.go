package main

//import "errors"
import "time"
import "math"

//import "fmt"

type dwave struct {
	x      float32
	y      float32
	dist   float32
	start  time.Time
	color  Color
	active int
}

var dw [30]dwave

func process_dance(brd Board, t time.Duration) error {
	brd.DrawAll(MakeColor(0, 0, 0))
	for row := 0; row < 5; row++ {
		for col := 0; col < 4; col++ {
			state := brd.getBoardState(row, col)
			if state == pressed {
				color := randomColor()
				_ = brd.DrawSquare(col, row, color)
				addWave(col, row, color)
			}
		}
	}

	for i := 0; i < 30; i++ {
		if dw[i].active == 1 {
			wDist := float32(time.Since(dw[i].start) / 100000000)
			if wDist > 25 {
				dw[i].active = 0
			} else {
				for x := 0; x < brd.pixelW; x++ {
					for y := 0; y < brd.pixelH; y++ {
						dx := float32(x) - dw[i].x
						dy := float32(y) - dw[i].y
						dis := float32(math.Sqrt(float64(dx*dx + dy*dy)))
						decay := 1.0 - wDist/25
						if dis > wDist-1.5 && dis < wDist+1.5 {
							amt := (dis - wDist) / 1.5
							if amt > 0 {
								amt = (1 - amt) * decay
							} else {
								amt = (1 + amt) * decay
							}

							brd.DrawPixel(x, y, MakeColor(int(amt*float32(dw[i].color.R)),
								int(amt*float32(dw[i].color.G)),
								int(amt*float32(dw[i].color.B))))
						}
					}
				}
			}
		}
	}
	return nil
}

func addWave(col int, row int, color Color) {
	for i := 0; i < 30; i++ {
		if dw[i].active != 1 {
			dw[i].x = float32(col)*5 + 2.5
			dw[i].y = float32(row)*5 + 2.5
			dw[i].active = 1
			dw[i].color = color
			dw[i].dist = 0
			dw[i].start = time.Now()
			break
		}
	}
}
