/*
A note on indexing of the board

0,0 is the left side if you stand on the high side of the board.

For pixels:  along the high side is X, along the left is Y
For squares: lid 0,0 is the left side. 0,3 is the right side.
The board sensors map to the same.

*/

package main

import "fmt"
import "errors"

type Board struct {
	strand  *Strand
	sensors *Sensors
	pixelW  int
	pixelH  int
	squareW int
	squareH int
}

func (brd *Board) Connect(pixelW int, pixelH int, squareW int, squareH int) error {
	brd.pixelW = pixelW
	brd.pixelH = pixelH
	brd.squareW = squareW
	brd.squareH = squareH
	brd.strand = &Strand{}
	brd.sensors = &Sensors{}
	brd.sensors.initSensors()
	return brd.strand.Connect(pixelW * pixelH)
}

func (brd *Board) Free() {
	brd.strand.Free()
	brd.sensors.stopSensors()
}

func (brd *Board) Save() {
	brd.strand.Save()
}

// 
func getPixelNum(x int, y int) int {
	col := x/5;
	row := y/5
	xPixelInSq := x % 5
	yPixelInSq := y % 5

	var boardNum, pixelNum int

	// NOTE: this is hardcoded for a 4 x 5 board with 25px/square
	if row%2 == 1 {
		boardNum = row*4 + col
	} else {
		boardNum = row*4 + 3 - col
	}

	if yPixelInSq%2 == 1 {
		pixelNum = yPixelInSq*5 + xPixelInSq
	} else {
		pixelNum = yPixelInSq*5 + 4 - xPixelInSq
	}

	return boardNum*25 + pixelNum
}

func (brd *Board) DrawPixel(x int, y int, r int, g int, b int) error {
	if x < 0 || x >= brd.pixelW || y < 0 || y >= brd.pixelH {
		return errors.New("Pixel was drawn outside the board's space")
	}
	pixelNum := getPixelNum(x, y)
	brd.strand.SetColor(pixelNum, r, g, b)

	return nil
}

func (brd *Board) DrawSquare(col int, row int, r int, g int, b int) error {
	for i:= 0; i<5; i++{
		for j:= 0; j<5; j++ {
			_ = brd.DrawPixel(col*5+i, row*5+j, r, g, b)
		}
	}
	return nil
}

func (brd *Board) DrawAll(r int, g int, b int) error {
	for i:= 0; i<brd.pixelW * brd.pixelH; i++{
			brd.SetColor(i, r, g, b)
	}
	return nil
}

func (brd *Board) SetColor(x int, r int, g int, b int) {
	brd.strand.SetColor(x, r, g, b)
}

func (brd *Board) getBoardState(row int, col int) int {
	return brd.sensors.getBoardState(row, col)
}

func (brd *Board) printBoardState() error {
	for r:=0; r<rows; r++ {
		for c:=0; c<cols; c++ {
			switch {
			case brd.sensors.net[r*cols+c] == up:
				fmt.Printf("-"); 
			case brd.sensors.net[r*cols+c] == down:
				fmt.Printf("X"); 
			case brd.sensors.net[r*cols+c] == pressed:
				fmt.Printf("|"); 
			case brd.sensors.net[r*cols+c] == released:
				fmt.Printf("+"); 
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