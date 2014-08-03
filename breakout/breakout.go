package breakout

import "github.com/lord/lodo/core"

const boardWidth = 35
const boardHeight = 28

func Run(board core.Board) {
	b := makeBall(10, 7, 2, 0.1, 0.2)
	black := core.MakeColor(0, 0, 0)
	for {
		board.RefreshSensors()
		board.DrawAll(black)
		b.step()
		b.draw(&board)
		board.Save()
	}
}
