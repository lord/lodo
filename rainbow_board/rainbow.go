package rainbowBoard

import (
	"fmt"
	"github.com/lord/lodo/core"
)

func Run(board *core.Board) {
	r := 0
	g := 0
	b := 0
	mode := 1
	colors := make([]core.Color, 35+42)
	const speed = 3
	black := core.MakeColor(0, 0, 0)
	for i := 0; i < 35+42; i++ {
		colors[i] = black
	}
	for {
		board.RefreshSensors()
		for i := len(colors) - 1; i >= 1; i-- {
			colors[i] = colors[i-1]
		}
		switch mode {
		case 0:
			r += speed
			if r >= 31 {
				r = 31
				mode++
			}
		case 1:
			g -= speed
			if g <= 0 {
				g = 0
				mode++
			}
		case 2:
			b += speed
			if b >= 31 {
				b = 31
				mode++
			}
		case 3:
			r -= speed
			if r <= 0 {
				r = 0
				mode++
			}
		case 4:
			g += speed
			if g >= 31 {
				g = 31
				mode++
			}
		case 5:
			b -= speed
			if b <= 0 {
				b = 0
				mode = 0
			}
		}
		colors[0] = core.MakeColor(r, g, b)
		for y := 0; y < 42; y++ {
			for x := 0; x < 35; x++ {
				board.DrawPixel(x, y, colors[x+y])
			}
		}
		for y := 0; y < 6; y++ {
			for x := 0; x < 5; x++ {
				if board.CheckPressed(x, y) {
					board.FillSquare(x, y, core.MakeColorAlpha(31, 0, 0, 0.5))
					fmt.Println("pressed", x, y)
				} else if board.CheckDown(x, y) {
					board.FillSquare(x, y, core.MakeColorAlpha(0, 0, 0, 0.8))
					fmt.Println("filling", x, y)
				}
			}
		}
		board.Save()
	}
}
