/*
A note on indexing of the board

0,0 is the left side if you stand on the high side of the board.

For pixels:  along the high side is X, along the left is Y
For squares: lid 0,0 is the left side. 0,3 is the right side.
The board sensors map to the same.

*/

package core

import "fmt"
import "errors"
import "math"

type Board struct {
	strand  *Strand
	sensors *Sensors
	pixelW  int
	pixelH  int
	squareW int
	squareH int
	poll    chan string
}

/////////////////////////////////
// CONNECTION FUNCTIONS
/////////////////////////////////

func (brd *Board) Connect(pixelW int, pixelH int, squareW int, squareH int) error {
	brd.pixelW = pixelW
	brd.pixelH = pixelH
	brd.squareW = squareW
	brd.squareH = squareH
	brd.strand = &Strand{}
	brd.sensors = &Sensors{}
	brd.sensors.initSensors(squareW, squareH)
	brd.poll = make(chan string)
	go brd.pollSensors(brd.poll)
	return brd.strand.Connect(mapLedColor(pixelW * pixelH))
}

func (brd *Board) Free() {
	brd.strand.Free()
	brd.sensors.stopSensors()
}

func (brd *Board) Save() {
	brd.strand.Save()
}

/////////////////////////////////
// DRAWING FUNCTIONS
/////////////////////////////////

func (brd *Board) DrawPixel(x, y int, c Color) {
	if x < 0 || x >= brd.pixelW || y < 0 || y >= brd.pixelH {
		fmt.Println("Pixel was drawn outside the board's space, at %v %v", x, y)
	}
	pixelNum := getPixelNum(x, y, brd.squareW, brd.squareH)
	brd.setColor(pixelNum, c)
}

func (brd *Board) DrawSquare(col int, row int, c Color) error {
	for i := 0; i < 7; i++ {
		for j := 0; j < 7; j++ {
			brd.DrawPixel(col*7+i, row*7+j, c)
		}
	}
	return nil
}

func (brd *Board) DrawAll(c Color) error {
	for i := 0; i < brd.pixelW*brd.pixelH; i++ {
		brd.setColor(i, c)
	}
	return nil
}

func ipart(x float64) float64 {
	return math.Floor(x)
}

func round(x float64) float64 {
	return math.Floor(x + 0.5)
}

func fpart(x float64) float64 {
	return x - math.Floor(x)
}

func rfpart(x float64) float64 {
	return 1 - fpart(x)
}

func (brd *Board) DrawLine(x0i, y0i, x1i, y1i int, c Color) {
	x0, y0, x1, y1 := float64(x0i), float64(y0i), float64(x1i), float64(y1i)
	steep := (math.Abs(y1-y0) > math.Abs(x1-x0))

	if steep {
		x0, y0 = y0, x0
		x1, y1 = y1, x1
	}
	if x0 > x1 {
		x0, x1 = x1, x0
		y0, y1 = y1, y0
	}

	dx := x1 - x0
	dy := y1 - y0
	gradient := dy / dx

	xend := round(x0)
	yend := y0 + gradient*(xend-x0)
	xgap := rfpart(x0 + 0.5)
	xpxl1 := xend //this will be used in the main loop
	ypxl1 := ipart(yend)
	if steep {
		brd.DrawPixel(int(ypxl1), int(xpxl1), c.WithAlpha(rfpart(yend)*xgap))
		brd.DrawPixel(int(ypxl1+1), int(xpxl1), c.WithAlpha(fpart(yend)*xgap))
	} else {
		brd.DrawPixel(int(xpxl1), int(ypxl1), c.WithAlpha(rfpart(yend)*xgap))
		brd.DrawPixel(int(xpxl1), int(ypxl1+1), c.WithAlpha(fpart(yend)*xgap))
	}
	intery := yend + gradient // first y-intersection for the main loop

	// handle second endpoint

	xend = round(x1)
	yend = y1 + gradient*(xend-x1)
	xgap = fpart(x1 + 0.5)
	xpxl2 := xend //this will be used in the main loop
	ypxl2 := ipart(yend)
	if steep {
		brd.DrawPixel(int(ypxl2), int(xpxl2), c.WithAlpha(rfpart(yend)*xgap))
		brd.DrawPixel(int(ypxl2+1), int(xpxl2), c.WithAlpha(fpart(yend)*xgap))
	} else {
		brd.DrawPixel(int(xpxl2), int(ypxl2), c.WithAlpha(rfpart(yend)*xgap))
		brd.DrawPixel(int(xpxl2), int(ypxl2+1), c.WithAlpha(fpart(yend)*xgap))
	}

	// main loop

	for x := xpxl1 + 1; x <= xpxl2-1; x++ {
		if steep {
			brd.DrawPixel(int(ipart(intery)), int(x), c.WithAlpha(rfpart(intery)))
			brd.DrawPixel(int(ipart(intery)+1), int(x), c.WithAlpha(fpart(intery)))
		} else {
			brd.DrawPixel(int(x), int(ipart(intery)), c.WithAlpha(rfpart(intery)))
			brd.DrawPixel(int(x), int(ipart(intery)+1), c.WithAlpha(fpart(intery)))
		}
		intery = intery + gradient
	}
}

func (brd *Board) DrawRect(x1, y1, x2, y2 int, c Color) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			brd.DrawPixel(x, y, c)
		}
	}
}

func (brd *Board) DrawRectOutline(x1, y1, x2, y2 int, c Color) {
	for x := x1; x <= x2; x++ {
		brd.DrawPixel(x, y1, c)
		brd.DrawPixel(x, y2, c)
	}

	for y := y1 + 1; y <= y2-1; y++ {
		brd.DrawPixel(x1, y, c)
		brd.DrawPixel(x2, y, c)
	}
}

func (brd *Board) DrawCircle(x1, y1, r int, c Color) error {
	return errors.New("Not implemented")
}

func (brd *Board) DrawCircleOutline(x1, y1, r int, c Color) error {
	return errors.New("Not implemented")
}

// func (brd *Board) DrawSprite(x1, y1, r int, c Color) error {
// 	return errors.New("Not implemented")
// }

// func (brd *Board) DrawText(x1, y1, r int, c Color) error {
// 	return errors.New("Not implemented")
// }

func (board *Board) RefreshSensors() {
	select {
	case msg := <-board.poll:
		_ = msg
		board.processSensors()
		go board.pollSensors(board.poll)
	default:
	}
}

/////////////////////////////////
// INTERNAL FUNCTIONS
/////////////////////////////////

func getPixelNum(x, y, sqW, sqH int) int {
	col := x / 7
	row := y / 7
	xPixelInSq := x % 7
	yPixelInSq := y % 7

	var boardNum, pixelNum int

	// NOTE: this is hardcoded for 49px/square
	if row%2 == 0 {
		boardNum = row*sqW + col
	} else {
		boardNum = row*sqW + (sqW - 1) - col
	}

	if yPixelInSq%2 == 1 {
		pixelNum = yPixelInSq*7 + xPixelInSq
	} else {
		pixelNum = yPixelInSq*7 + 6 - xPixelInSq
	}

	return boardNum*49 + pixelNum
}

func (brd *Board) getBoardState(row int, col int) int {
	return brd.sensors.getBoardState(row, col)
}

func (brd *Board) setColor(led int, color Color) {
	brd.strand.SetColor(mapLedColor(led), color)
}

func (brd *Board) printBoardState() error {
	for r := 0; r < brd.squareH; r++ {
		for c := 0; c < brd.squareW; c++ {
			switch {
			case brd.sensors.net[r*brd.squareW+c] == up:
				fmt.Printf("-")
			case brd.sensors.net[r*brd.squareW+c] == down:
				fmt.Printf("X")
			case brd.sensors.net[r*brd.squareW+c] == pressed:
				fmt.Printf("|")
			case brd.sensors.net[r*brd.squareW+c] == released:
				fmt.Printf("+")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n\n")
	return nil
}

func (brd *Board) pollSensors(poll chan string) {
	brd.sensors.readSensors()
	poll <- "ready"
}

func (brd *Board) processSensors() {
	brd.sensors.processSensors()
}

func mapLedColor(i int) int {
	return i + i/49
}
