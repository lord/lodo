package main

import "fmt"
import "time"
import "math"

func main() {
	board := Board{}

	w := 20
	h := 25
	err := board.Connect(w, h, 4, 5)
	defer board.Free()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	offset := 0.0
	for {
		offset -= 0.2
		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				dx := x - 10
				dy := y - 10
				dis := math.Sqrt(float64(dx*dx + dy*dy))
				amt := int(math.Sin(dis+offset)*125.0 + 125.0)
				board.DrawPixel(x, y, amt, 0, 125-amt/2)
			}
		}
		board.Save()
		time.Sleep(20 * time.Millisecond)
	}
}
