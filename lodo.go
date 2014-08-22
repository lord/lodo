package main

import (
	"flag"
	"fmt"
	"time"
	"github.com/lord/lodo/breakout"
	"github.com/lord/lodo/core"
	"github.com/lord/lodo/maze"
	"github.com/lord/lodo/rainbow_board"
	"github.com/lord/lodo/server"
	"github.com/lord/lodo/test"
	"github.com/lord/lodo/pong"
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
		core.PetDog()
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
			server.Run(board)
			game = Selection
		case Maze:
			maze.Run(board)
			game = Selection
		case Pong:
			pong.Run(board)
			game = Selection
		default:
			fmt.Println("Game not recognized")
			game = Selection
		}
	}
}


func Run (b *core.Board) int {
	b.SetVerticalMode(false)
	timeOut   := time.Now().Add(time.Duration(30)*time.Second)
	chooseSound := core.MakeSound(core.Selectgame)
	chooseSound.Play()

	for ;; {
		b.RefreshSensors()
		b.DrawAll(core.Black)
		b.WriteText("Choose",4,13,core.Orient_0, core.White)

		b.WriteText("PONG",8,20,core.Orient_0, core.Blue)
		b.DrawRectOutline(0,14,6,20,core.Blue)
		if b.CheckDown(0, 2) { return Pong }

		b.WriteText("BKOUT",8,27,core.Orient_0, core.Green)
		b.DrawRectOutline(0,21,6,27,core.Green)
		if b.CheckDown(0, 3) { return Breakout }

		b.WriteText("MAZE",8,34,core.Orient_0, core.Purple)
		b.DrawRectOutline(0,28,6,34,core.Purple)
		if b.CheckDown(0, 4) { return Maze }

		b.WriteText("RIPL",8,41,core.Orient_0, core.Purple)
		b.DrawRectOutline(0,28,6,34,core.Yellow)
		if b.CheckDown(0, 5) { return Ripple }

		if b.CheckDown(1, 1) { return Test }

		if b.CheckAnyDown() {
			timeOut = time.Now().Add(time.Duration(30)*time.Second)
		}

		if time.Now().After(timeOut) {
			return Rainbow
		}
		b.Save()
	}
}