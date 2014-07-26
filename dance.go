package main

//import "errors"
import "time"

type dwave struct 

func process_dance(brd Board, t time.Duration) error {
	for row:=0; row< 5; row++ {
		for col:=0; col<4; col++ {
			state := brd.getBoardState(row,col)
			if (state == down || state == pressed ) {
				_ = brd.DrawSquare(col,row,255,255,255)
			}
		}
	}
	
	offset := 0.0
	for {
		offset -= 0.1
		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				dx := x - 10
				dy := y - 10
				dis := math.Sqrt(float64(dx*dx + dy*dy))
				amt := int(math.Sin(dis+offset)*250.0 + 1.0)
				board.DrawPixel(x, y, amt, 10, amt)
			}
		}
	}

	return nil
}