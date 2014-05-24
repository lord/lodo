package main

// #cgo LDFLAGS: -lm
// #include "tclled.h"
import "C"
import "fmt"

// import "time"
import "errors"

// import "time"

//////////////////
// STRAND CODE
//////////////////

type Strand struct {
	device   C.int
	buffer   *_Ctype_tcl_buffer
	ledCount int
}

func (s *Strand) Connect(ledCount int) error {
	s.ledCount = ledCount
	s.device = C.open_device()

	if s.device <= 0 {
		return errors.New("Device init failed")
	}

	C.set_gamma(2.2, 2.2, 2.2)
	spiStatus := C.spi_init(s.device)
	if spiStatus != 0 {
		return errors.New("SPI init failed")
	}

	s.buffer = &C.tcl_buffer{}
	tclStatus := C.tcl_init(s.buffer, C.int(s.ledCount))
	if tclStatus != 0 {
		return errors.New("TCL init failed")
	}

	return nil
}

func (s *Strand) SetColor(ledNumber int, r int, g int, b int) {
	C.write_gamma_color_to_buffer(s.buffer, C.int(ledNumber), C.uint8_t(r), C.uint8_t(g), C.uint8_t(b))
}

func (s *Strand) Save() {
	fmt.Println("Strand (", s, ") saving...")
	C.send_buffer(s.device, s.buffer)
}

//////////////////
// BOARD CODE
//////////////////

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

func (brd *Board) DrawPixel(x int, y int, r int, g int, b int) {
	pixelNum := getPixelNum(x, y)
	fmt.Println("Pixel Drawn at: (", x, y, ") ->", pixelNum)
	brd.strand.SetColor(pixelNum, r, g, b)
}

func (brd *Board) SetPixel(x int, r int, g int, b int) {
	fmt.Println("(", r, g, b, ") Pixel Drawn ->", x)
	fmt.Println("Double Checking: () ->", x)
	brd.strand.SetColor(x, r, g, b)
}

func main() {
	// board := Board{}
	// w := 100
	// h := 125
	// err := board.Connect(w, h, 4, 5)

	// if err != nil {
	// 	fmt.Print("Error: ")
	// 	fmt.Println(err)
	// 	return
	// }
	// for {
	// 	x := 0
	// 	r := 0
	// 	g := 0
	// 	b := 0

	// 	fmt.Scan(&x, &r, &g, &b)
	// 	board.SetPixel(x, r, g, b)
	// 	board.Save()
	// }

	board := Strand{}
	c := 500
	err := board.Connect(c)

	if err != nil {
		fmt.Print("Error: ")
		fmt.Println(err)
		return
	}
	for {
		x := 0
		r := 0
		g := 0
		b := 0

		fmt.Scan(&x, &r, &g, &b)
		board.SetColor(x, r, g, b)
		board.Save()
	}

	// xPix := 0
	// yPix := 0
	// for {
	// 	xPix++
	// 	if xPix > w {
	// 		xPix = 0
	// 		yPix++
	// 	}
	// 	if yPix > h {
	// 		xPix = 0
	// 		yPix = 0
	// 	}

	// 	for x := 0; x < w; x++ {
	// 		for y := 0; y < h; y++ {
	// 			if x == xPix && y == yPix {
	// 				board.DrawPixel(x, y, 100, 100, 100)
	// 			} else {
	// 				board.DrawPixel(x, y, 0, 0, 0)
	// 			}
	// 		}
	// 	}
	// 	board.Save()
	// 	time.Sleep(20 * time.Millisecond)
	// }
}
