package test

import "github.com/lord/lodo/core"

const boardWidth = 35
const boardHeight = 42
const squareWidth = 5
const squareHeight = 6

func Run(board *core.Board) {
	board.DebugSensors(false)
	board.RefreshSensors()
	blue := core.MakeColor(0, 0, 2)
	white := core.MakeColor(10, 10, 10)
	red := core.MakeColor(31, 0, 0)
	for {
		board.RefreshSensors()
		for x := 0; x < 5; x++ {
			for y := 0; y < 6; y++ {
				if board.CheckPressed(x, y) {
					board.FillSquare(x, y, white)
				} else {
					board.FillSquare(x, y, blue)
				}
			}
		}
		board.WriteText("HELLO",0,7,core.Orient_0,red)
		board.WriteText("HELLO",10,36,core.Orient_90,red)
		board.WriteText("YUP",32,8,core.Orient_180,red)
		board.WriteText("WWW",27,16,core.Orient_270,red)
		board.Save()
	}
}
