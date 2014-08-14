package test

import "github.com/lord/lodo/core"

const boardWidth = 35
const boardHeight = 42
const squareWidth = 5
const squareHeight = 6

func Run(board *core.Board) {
	board.RefreshSensors()
	blue := core.MakeColor(0, 0, 2)
	white := core.MakeColor(10, 10, 10)
//	red := core.MakeColor(0, 31, 0)
	for {
		board.RefreshSensors()
		for x := 0; x < 5; x++ {
			for y:= 0; y<6; y++ {
				if board.CheckDown(x, y) {
					board.FillSquare(x,y,white)	
				} else {
					board.FillSquare(x,y,blue)
				}
			}
		}
		board.Save()
	}
}
