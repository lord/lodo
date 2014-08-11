package breakout

import "github.com/lord/lodo/core"
import "time"

const boardWidth = 35
const boardHeight = 42

func Run(board *core.Board) {
	b := makeBall(10, 7, 0.6, 0.9, core.MakeColor(31, 31, 31))
	black := core.MakeColor(0, 0, 0)
	for {
		// board.RefreshSensors()
		board.DrawAll(black)
		b.step()
		b.draw(board)
		board.Save()
		time.Sleep(time.Millisecond)
	}
}
