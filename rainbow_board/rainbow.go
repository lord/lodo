package rainbowBoard

import (
//	"fmt"
	"github.com/lord/lodo/core"
)

func Run(board *core.Board) {
	sitarmusic := core.MakeSound(core.SitarMusic)
	sitarmusic.Play()
	board.SetVerticalMode(true)
	r := 0
	g := 0
	b := 0
	mode := 1
	colors := make([]core.Color, 35+43)
	const speed = 3
	black := core.MakeColor(0, 0, 0)
	for i := 0; i < 35+43; i++ {
		colors[i] = black
	}
	for {
		core.PetDog() // Keepalive or restart
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
		for y := 0; y < 43; y++ {
			for x := 0; x < 35; x++ {
				board.DrawPixel(x, y, colors[x+y])
			}
		}
		// for c := 0; c<22*5; c++ {
		// 	// bottom
		// 	switch {
		// 	case c < 25:
		// 		board.DrawSidePixel(c,0,colors[c+42])
		// 	case c<55:
		// 		board.DrawSidePixel(c,0,colors[42-(c-25)+35])
		// 		board.DrawSidePixel(c,1,colors[42-(c-25)+35])
		// 	case c<80:
		// 		board.DrawSidePixel(c,0,colors[80-c])
		// 		board.DrawSidePixel(c,1,colors[80-c])
		// 	case c<110:
		// 		board.DrawSidePixel(c,0,colors[(c-80)])
		// 		board.DrawSidePixel(c,1,colors[(c-80)])
		// 	}
		// }

		for c := 0; c<22*5; c++ {
			// bottom
			switch {
			case c < 25:
				board.DrawSidePixel(c,0,colors[c+42])
			case c<55:
				board.DrawSidePixel(c,0,colors[42-(c-25)+35])
				board.DrawSidePixel(c,1,colors[42-(c-25)+35])
			case c<80:
				board.DrawSidePixel(c,0,colors[80-c])
				board.DrawSidePixel(c,1,colors[80-c])
			case c<110:
				board.DrawSidePixel(c,0,colors[(c-80)])
				board.DrawSidePixel(c,1,colors[(c-80)])
			}
		}

		if board.CheckAnyDown() { 
			sitarmusic.Stop()
			return 
		}
		board.Save()
	}
}

		// for y := 0; y < 6; y++ {
		// 	for x := 0; x < 5; x++ {
		// 		if board.CheckPressed(x, y) {
		// 			board.FillSquare(x, y, core.MakeColorAlpha(31, 0, 0, 0.5))
		// 			//fmt.Println("pressed", x, y)
		// 			switch {
		// 			case y==0 && x==0:  core.PlaySound(core.A2deep)
		// 			case y==0 && x==1:  core.PlaySound(core.A2deep)
		// 			case y==0 && x==2:  core.PlaySound(core.C1deep)
		// 			case y==0 && x==3:  core.PlaySound(core.C2deep)
		// 			case y==0 && x==4:  core.PlaySound(core.D1deep)

		// 			case y==1 && x==0:  core.PlaySound(core.D2deep)
		// 			case y==1 && x==1:  core.PlaySound(core.E1deep)
		// 			case y==1 && x==2:  core.PlaySound(core.E2deep)
		// 			case y==1 && x==3:  core.PlaySound(core.G1deep)
		// 			case y==1 && x==4:  core.PlaySound(core.G2deep)

		// 			case y==2 && x==0:  core.PlaySound(core.A1ripple)
		// 			case y==2 && x==1:  core.PlaySound(core.A2ripple)
		// 			case y==2 && x==2:  core.PlaySound(core.C1ripple)
		// 			case y==2 && x==3:  core.PlaySound(core.C2ripple)
		// 			case y==2 && x==4:  core.PlaySound(core.D1ripple)

		// 			case y==3 && x==0:  core.PlaySound(core.D2ripple)
		// 			case y==3 && x==1:  core.PlaySound(core.E1ripple)
		// 			case y==3 && x==2:  core.PlaySound(core.E2ripple)
		// 			case y==3 && x==3:  core.PlaySound(core.G1ripple)
		// 			case y==3 && x==4:  core.PlaySound(core.A1ripple)

		// 			case y==4 && x==0:  core.PlaySound(core.A1ripple)
		// 			case y==4 && x==1:  core.PlaySound(core.A1ripple)
		// 			case y==4 && x==2:  core.PlaySound(core.A1ripple)
		// 			case y==4 && x==3:  core.PlaySound(core.A1ripple)
		// 			case y==4 && x==4:  core.PlaySound(core.A1ripple)

		// 			case y==5 && x==0:  core.PlaySound(core.A1deep)
		// 			case y==5 && x==1:  core.PlaySound(core.A1deep)
		// 			case y==5 && x==2:  core.PlaySound(core.A1deep)
		// 			case y==5 && x==3:  core.PlaySound(core.A1deep)
		// 			case y==5 && x==4:  core.PlaySound(core.A1deep)

		// 			}
		// 		} else if board.CheckDown(x, y) {
		// 			board.FillSquare(x, y, core.MakeColorAlpha(0, 0, 0, 0.8))
		// 			//fmt.Println("filling", x, y)
		// 		}
		// 	}
		// }
