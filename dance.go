package main

//import "errors"
import "time"

func process_dance(brd Board, t time.Duration) error {
	brd.DrawAll(0,0,0)
	steps := int(t.Nanoseconds()/20000000 % 500)
	x := steps % 20
	y := steps / 20
	brd.DrawPixel(x,y,255,0,255)

	for row:=0; row< 5; row++ {
		for col:=0; col<4; col++ {
			state := brd.getBoardState(row,col)
			if (state == down || state == pressed ) {
				_ = brd.DrawSquare(col,row,255,255,255)
			}
		}
	}

	return nil
}