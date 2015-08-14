package main

import (
	"flag"
	"fmt"
	"time"
    "log"
	"log/syslog"
	"github.com/lord/lodo/breakout"
	"github.com/lord/lodo/core"
	//"github.com/lord/lodo/maze"
	"github.com/lord/lodo/rainbow_board"
	//"github.com/lord/lodo/server"
	"github.com/lord/lodo/test"
	"github.com/lord/lodo/pong"
	"github.com/lord/lodo/ripple"
)

var gameMode = flag.String(
	"mode",
	"select",
	"Selects the game to run. Options are 'test', 'rainbow-board', 'server', 'maze' and 'breakout'.",
)


const Selection=1
const Maze=2
const Pong=3
const Breakout=4
const Rainbow=5
const Ripple=6
const Test=7
const Server=8

func main() {
		fmt.Println("MAIN")
    syslogger, err := syslog.New(syslog.LOG_INFO, "lodo")
    if err == nil {
    	log.SetOutput(syslogger)
    }
    log.Print("Starting Lodo")
	
	flag.Parse()
	core.StartDog()

	board, err := core.MakeBoard()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer board.Free()

	game := Selection

	for ;; {
		switch game {
		case Selection:
			game = Run(board)
		case Rainbow:
			rainbowBoard.Run(board)
			game = Selection
		case Test:
			board.DebugSensors(true)
			test.Run(board)
			game = Selection
		case Breakout:
			breakout.Run(board)
			game = Selection
		case Server:
//			server.Run(board)
			game = Selection
		case Maze:
//			maze.Run(board)
			game = Selection
		case Pong:
			pong.Run(board)
			game = Selection
		case Ripple:
			ripple.Run(board)
			game = Selection
		default:
			fmt.Println("Game not recognized")
			game = Selection
		}
		board.ClearItems()
	}
}

func Run (b *core.Board) int {
    fmt.Println("RUN")
	b.SetVerticalMode(false)
	timeOut   := time.Now().Add(time.Duration(30)*time.Second)
	chooseSound := core.MakeSound(core.Selectgame)
	chooseSound.Play()

	pong := core.NewPallette(100,0,4)
	pong.Visible(false)
	fmt.Println(pong)
	pong.AddItem(core.Drawer(core.NewStext(8,20,3,core.Blue, core.Orient_0, "Pong")))
	b.AddItem(pong)

	bkout := core.NewPallette(0,0,4)
	pong.Visible(true)
	bkout.AddItem(core.Drawer(core.NewStext(8,20,3,core.Green, core.Orient_0, "Bkout")))
	b.AddItem(bkout)

	oasis := core.NewPallette(100,0,4)
	oasis.Visible(false)
	oasis.AddItem(core.Drawer(core.NewStext(8,20,3,core.Yellow, core.Orient_0, "Oasis")))
	b.AddItem(oasis)	

	b.AddItem(core.Drawer(core.NewStext(4,13,2,core.White, core.Orient_0, "Choose")))
	b.AddItem(core.NewSArrow(10,31,270,14,core.NewSolidControl(core.Blue)))
	b.AddItem(core.Drawer(core.NewDrawRect(14,28,7,7,1,core.Green,1400)))
	b.AddItem(core.NewSArrow(24,31,90,14,core.NewSolidControl(core.Blue)))
	b.AddItem(core.NewBlinkyRect(0,0,0,35,42,5,300,[]core.Color{core.MakeColor(0,0,4),core.Black,core.Black,core.Black}))
	lightR := core.MakeColor(2,0,0)
	lightB := core.MakeColor(0,0,2)
	b.AddItem(core.NewBlinkyRect(0,0,2,35,42,5,100,[]core.Color{lightB,lightR,lightR,lightR}))

	left    := &pong
	current := &bkout
	right   := &oasis
	
	for ;; {
		core.PetDog()

		b.RefreshSensors()
		b.DrawAll(core.Black)

		if b.CheckPressed(1, 4) {
			dur := 1500
			chooseSound := core.MakeSound(core.Scrape)
			chooseSound.Play()
			b.AddItem(core.NewCRect(7,28,7,7,13,core.NewSolidShortControl(dur,core.MakeColor(15,15,15))))
			b.AddItem(core.NewSArrow(10,31,270,14,core.NewSolidShortControl(dur,core.MakeColor(0,0,3))))
			current.BeginAnime(core.Anime_departleft,dur)
			right.BeginAnime(core.Anime_arriveright,dur)
			delay(b,dur)
			tmp := left; left=current; current=right; right=tmp
		}

		if b.CheckPressed(3, 4) { 
			dur := 1500
			chooseSound := core.MakeSound(core.Scrape)
			chooseSound.Play()
			b.AddItem(core.NewCRect(21,28,7,7,13,core.NewSolidShortControl(3000,core.MakeColor(15,15,15))))
			b.AddItem(core.NewSArrow(24,31,270,14,core.NewSolidShortControl(3000,core.MakeColor(0,0,3))))	 
			current.BeginAnime(core.Anime_departright,dur)
			left.BeginAnime(core.Anime_arriveleft,dur)
			delay(b,dur)
			tmp:=right; right=current; current=left; left=tmp
		}

		if b.CheckDown(2, 4) {
			switch current { 
			case &pong:
				chooseSound := core.MakeSound(core.PongVoice)
				chooseSound.Play()
				return Pong 
			case &bkout:
				chooseSound := core.MakeSound(core.BreakoutVoice)
				chooseSound.Play()
				return Breakout
			case &oasis:
				chooseSound := core.MakeSound(core.OasisVoice)
				chooseSound.Play()
				return Ripple
			}
		}

		if b.CheckDown(3, 1) { return Test }

		if b.CheckAnyDown() {
			timeOut = time.Now().Add(time.Duration(30)*time.Second)
		}

		if time.Now().After(timeOut) {
			return Rainbow
		}
		b.SetVerticalMode(false)
	b.Save()
	}
}

// pause and let effects complete
func delay(b *core.Board, delayMS int) {
	goOn := time.Now().Add(time.Duration(delayMS)*time.Millisecond)
	for {
		b.DrawAll(core.Black)
		b.Save()
		if time.Now().After(goOn) {
			return
		}
	}
}