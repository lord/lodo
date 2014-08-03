package breakout

import "github.com/lord/lodo/core"

func Run(board core.Board) {
	black := core.MakeColor(0, 0, 0)
	for {
		board.RefreshSensors()
		board.DrawAll(black)
		for x := 0; x < 5; x++ {
			for y := 0; y < 4; y++ {
				if board.CheckPressed(x, y) {
					board.DrawSquare(x, y, core.MakeColor(20, 20, 20))
				}
			}
		}
		board.Save()
	}
}
