package breakout

import "github.com/lord/lodo/core"
import "time"

const boardWidth = 35
const boardHeight = 42

func Run(board *core.Board) {
	b := makeBall(10, 7, 0.6, 0.9, core.MakeColor(31, 31, 31))
	paddleX, paddleY := board.GetSquare(2, 5)
	paddle2X, paddle2Y := board.GetSquare(2, 0)
	paddle1 := makePaddle(float64(paddleX), float64(paddleY)+1, 6, 5, core.MakeColor(0, 0, 31))
	paddle2 := makePaddle(float64(paddle2X), float64(paddle2Y)+4, 6, 0, core.MakeColor(31, 0, 0))
	black := core.MakeColor(0, 0, 0)
	for {
		board.RefreshSensors()
		board.DrawAll(black)
		b.step()
		paddle1.step(board)
		paddle2.step(board)
		b.draw(board)
		paddle1.draw(board)
		paddle2.draw(board)
		board.Save()
		time.Sleep(time.Millisecond)
	}
}
