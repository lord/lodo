package test

import "github.com/lord/lodo/core"

const boardWidth = 35
const boardHeight = 42
const squareWidth = 5
const squareHeight = 6

func Run(board core.Board) {
	blue := core.MakeColor(0, 0, 1)
        green  := core.MakeColor(31,10,22)
	red := core.MakeColor(0,31,0)
	color := core.MakeColor(0,0,0)
	for {
		board.RefreshSensors()
		board.DrawAll(green)
                for x:=0; x<5; x++ {
			if board.CheckDown(x,0) {
				color = blue
			} else if board.CheckDown(x,5) {
				color = red
			} else {
				color = green
			}
					board.FillSquare(x,0,color)
					board.FillSquare(x,1,color)
					board.FillSquare(x,2,color)
					board.FillSquare(x,3,color)
					board.FillSquare(x,4,color)
					board.FillSquare(x,5,color)
		}
		board.Save()
	}
}
