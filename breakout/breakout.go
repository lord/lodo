package breakout

import "github.com/lord/lodo/core"
import "time"

const boardWidth = 35
const boardHeight = 42

const ( 
    begin = 1 << iota 
    play 
    miss 
    p1_win
    p1_score
    p2_win
    p2_score
    end
)

var paddle1 paddle
var paddle2 paddle
var mode int
var modeTime time.Time
var b ball

func Run(board *core.Board) {
	mode = play
	b = makeBall(10, 7, 0.3, 0.45, core.MakeColor(31, 31, 31))
	paddleX, paddleY := board.GetSquare(2, 5)
	paddle2X, paddle2Y := board.GetSquare(2, 0)
	paddle1 = makePaddle(float64(paddleX), float64(paddleY)+1, 6, 5, core.MakeColor(0, 0, 31))
	paddle2 = makePaddle(float64(paddle2X), float64(paddle2Y)+4, 6, 0, core.MakeColor(31, 0, 0))
	black := core.MakeColor(0, 0, 0)

	timeBall   := time.Now()
	stepBall   := time.Duration(1)*time.Millisecond
	timePaddle := time.Now()
	stepPaddle := time.Duration(1)*time.Millisecond

	for {
		// get the time and sensors
		now := time.Now()
		board.RefreshSensors()
		board.DrawAll(black)
	    switch {
	    case mode == play :
	    	// update state if needed
			if now.After(timeBall) {
				b.step()
				timeBall = now.Add(stepBall)
			}
			if now.After(timePaddle) {
				paddle1.step(board)
				paddle2.step(board)
				timePaddle = now.Add(stepPaddle)
			}

			// Draw the board
			b.draw(board)
			paddle1.draw(board)
			paddle2.draw(board)
		case mode == p1_score :
			if now.After(modeTime) {
				setMode(play)
			}
			board.DrawAll(paddle1.color)
		case mode == p2_score :
			if now.After(modeTime) {
				setMode(play)
			}
			board.DrawAll(paddle2.color)
		}
		board.Save()
	}
}

func setMode(m int) {
	switch {
	case m == begin:
		modeTime = time.Now().Add(time.Duration(1000)*time.Millisecond)
	case m == play:
		modeTime = time.Now().Add(time.Duration(1000)*time.Millisecond)
		if mode == p1_score {
			b.init(paddle1.x+paddle1.w/2, paddle1.y, 0.3, -0.45)
		} else if mode == p2_score {
			b.init(paddle2.x+paddle2.w/2, paddle2.y, 0.3, 0.45)
		} else {
			b.init(paddle2.x+paddle2.w/2, paddle2.y, 0.3, 0.45)
		}
	case m == p1_win:
		modeTime = time.Now().Add(time.Duration(1000)*time.Millisecond)
	case m == p1_score:
		modeTime = time.Now().Add(time.Duration(1000)*time.Millisecond)
	case m == p2_win:
		modeTime = time.Now().Add(time.Duration(1000)*time.Millisecond)
    case m == p2_score:
		modeTime = time.Now().Add(time.Duration(1000)*time.Millisecond)
	case m == end:
		modeTime = time.Now().Add(time.Duration(1000)*time.Millisecond)
    }
   	mode = m
}