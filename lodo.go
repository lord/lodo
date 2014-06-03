package main

import "fmt"
import "time"

func main() {
	board := Board{}

	w := 20
	h := 25
	err := board.Connect(w, h, 4, 5)
	defer board.Free()

	if err != nil {
		fmt.Print("Error: ")
		fmt.Println(err)
		return
	}
	xPix := 0
	yPix := 0
	for {
		xPix++
		if xPix > w {
			xPix = 0
			yPix++
		}
		if yPix > h {
			xPix = 0
			yPix = 0
		}

		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				if x == xPix && y == yPix {
					board.DrawPixel(x, y, 100, 100, 100)
				} else {
					board.DrawPixel(x, y, 0, 0, 0)
				}
			}
		}
		board.Save()
		time.Sleep(20 * time.Millisecond)
	}
}
