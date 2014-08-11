package rainbow

import (
	"github.com/lord/lodo/core"
	"time"
)

func Run(strand *core.Strand) {
	r := 0
	g := 0
	b := 0
	mode := 1
	const speed = 3
	black := core.MakeColor(0, 0, 0)
	for i := 0; i < strand.Length(); i++ {
		strand.SetColor(i, black)
	}
	for {
		for i := strand.Length() - 1; i >= 1; i-- {
			strand.SetColor(i, strand.GetColor(i-1))
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
		strand.SetColor(0, core.MakeColor(r, g, b))
		strand.Save()
		time.Sleep(5 * time.Millisecond)
	}
}
