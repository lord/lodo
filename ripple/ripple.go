package ripple

import (
"time"
// "fmt"
"github.com/lord/lodo/core"
"math/rand"
)

var tone = core.MakeSound(core.Pong)
var water = []core.Color {
	core.Color{0,1,1,1},
	core.Color{0,2,2,1},
	core.Color{0,3,3,1},
	core.Color{0,1,2,1},
	core.Color{0,1,3,1},
	core.Color{0,2,3,1},
	core.Color{0,0,1,1},
	core.Color{0,0,5,1},
	core.Color{0,0,3,1},
	core.Color{2,2,2,1}}

func Run(board *core.Board) {
 	rand.Seed(42)
 	colors := make([]core.Color, 35*43)
	for i := 0; i < 35+43; i++ {
		colors[i] = water[0]
	}
	timeOut   := time.Now().Add(time.Duration(3)*time.Second)
	ripplemusic := core.MakeSound(core.RippleMusic)
	ripplemusic.Play()
	board.SetVerticalMode(true)
	for {
		core.PetDog() // Keepalive or restart
		board.RefreshSensors()
		for x:=0; x<35; x++ {
			for y:=0; y<43; y++ {
				if rand.Intn(10) == 1 {
					colors[x+y*35] = water[rand.Intn(len(water))]
				}
				board.DrawPixel(x,y,colors[x+y*35])
			}
		}
		// for c := 0; c<22*5; c++ {
		// 	// bottom
		// 	switch {
		// 	case c < 25:
		// 		board.DrawSidePixel(c,0,core.Blue)
		// 	case c<55:
		// 		board.DrawSidePixel(c,0,core.Blue)
		// 		board.DrawSidePixel(c,1,core.Blue)
		// 	case c<80:
		// 		board.DrawSidePixel(c,0,core.Blue)
		// 		board.DrawSidePixel(c,1,core.Blue)
		// 	case c<110:
		// 		board.DrawSidePixel(c,0,core.Blue)
		// 		board.DrawSidePixel(c,1,core.Blue)
		// 	}
		// }
		for y := 0; y < 6; y++ {
			for x := 0; x < 5; x++ {
				if board.CheckPressed(x, y) {
					play(x,y)
				} else if board.CheckDown(x, y) {
					board.FillSquare(x, y, core.MakeColorAlpha(0, 0, 0, 0.8))
					//fmt.Println("filling", x, y)
					timeOut   = time.Now().Add(time.Duration(3)*time.Second)
				}
			}
		}
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
		if time.Now().After(timeOut) {
			ripplemusic.Stop()
			return
		}
		// tone.Print()
		board.Save()
	}
}

func play(x,y int) {
	switch {
	case y==0 && x==0:  startSound(core.Rip1)
	case y==0 && x==1:  startSound(core.Rip2)
	case y==0 && x==2:  startSound(core.Rip3)
	case y==0 && x==3:  startSound(core.Rip4)
	case y==0 && x==4:  startSound(core.Rip5)

	case y==1 && x==0:  startSound(core.Rip6)
	case y==1 && x==1:  startSound(core.Rip7)
	case y==1 && x==2:  startSound(core.Rip8)
	case y==1 && x==3:  startSound(core.Rip9)
	case y==1 && x==4:  startSound(core.Rip10)

	case y==2 && x==0:  startSound(core.Rip11)
	case y==2 && x==1:  startSound(core.Rip12)
	case y==2 && x==2:  startSound(core.Rip13)
	case y==2 && x==3:  startSound(core.Rip14)
	case y==2 && x==4:  startSound(core.Rip15)

	case y==3 && x==0:  startSound(core.Rip16)
	case y==3 && x==1:  startSound(core.Rip17)
	case y==3 && x==2:  startSound(core.Rip18)
	case y==3 && x==3:  startSound(core.Rip19)
	case y==3 && x==4:  startSound(core.Rip20)

	case y==4 && x==0:  startSound(core.Rip21)
	case y==4 && x==1:  startSound(core.Rip22)
	case y==4 && x==2:  startSound(core.Rip23)
	case y==4 && x==3:  startSound(core.Rip24)
	case y==4 && x==4:  startSound(core.Rip25)

	case y==5 && x==0:  startSound(core.Rip26)
	case y==5 && x==1:  startSound(core.Rip27)
	case y==5 && x==2:  startSound(core.Rip28)
	case y==5 && x==3:  startSound(core.Rip29)
	case y==5 && x==4:  startSound(core.Rip30)
	}
}

func startSound(s string){
	tone.Stop()
	tone = core.MakeSound(s)
	tone.Play()
}

