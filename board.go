package main

import "errors"

type Board struct {
	strand  *Strand
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
	return brd.strand.Connect(pixelW * pixelH)
}

func (brd *Board) Free() {
	brd.strand.Free()
}

func (brd *Board) Save() {
	brd.strand.Save()
}

func getPixelNum(x int, y int) int {
	xSq := x / 5
	ySq := y / 5
	var boardNum, pixelNum int

	// NOTE: this is hardcoded for a 4 x 5 board with 25px/square
	if ySq%2 == 0 {
		boardNum = ySq*4 + xSq
	} else {
		boardNum = ySq*4 + 3 - xSq
	}

	xPixelInSq := x % 5
	yPixelInSq := y % 5

	if yPixelInSq%2 == 0 {
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

func (brd *Board) SetColor(x int, r int, g int, b int) {
	brd.strand.SetColor(x, r, g, b)
}
