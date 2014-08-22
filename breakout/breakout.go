package breakout

import (
 "github.com/lord/lodo/core"
 "time"
 "fmt"
 "math"
  "math/rand"

  )

const boardWidth = 35
const boardHeight = 42

const ( 
    begin = 1 << iota 
    play //2
    miss //4
    end //8
)



var r *rand.Rand
var paddle1 paddle
var mode int
var modeTime time.Time
var b ball
var score int
var ballsRemaining int

func Run(board *core.Board) {
	setMode(begin)
	board.SetVerticalMode(true)
	r = rand.New(rand.NewSource(99))
	b = makeBall(17, 7, 3, 2, 3, .1, 0.6, core.MakeColor(31, 31, 31))
	paddleX, paddleY := board.GetSquare(2, 5)

	paddle1 = makePaddle(float64(paddleX), float64(paddleY)+1, 6, 5, core.MakeColor(0, 0, 31))

	timeBall   := time.Now()
	stepBall   := time.Duration(1)*time.Millisecond
	timePaddle := time.Now()
	stepPaddle := time.Duration(1)*time.Millisecond

	breakoutmusic := core.MakeSound(core.BreakoutMusic)
	breakoutmusic.Play()

	for {
		core.PetDog()
		// get the time and sensors
		now := time.Now()
		board.RefreshSensors()
		board.DrawAll(core.Black)
	    switch {
	    case mode == begin:
			board.WriteText("Ready!",   0, 21, core.Orient_0,   paddle1.color)
			board.WriteText("3 Balls",  0, 28, core.Orient_0,   paddle1.color)
			if now.After(modeTime) { setMode(play) }
	    case mode == play :
	    	// update state if needed
			if now.After(timeBall) {
				b.step()
				timeBall = now.Add(stepBall)
			}
			if now.After(timePaddle) {
				paddle1.step(board)
				timePaddle = now.Add(stepPaddle)
			}
			// Draw the board
			b.draw(board)
			paddle1.draw(board)
			for i:= 0; i<45; i++ {
				blocks[i].Draw(board)
			}
		case mode == miss:
			board.WriteText("Balls",7,6,core.Orient_0, paddle1.color)
			board.WriteText(fmt.Sprintf("Left: %d", ballsRemaining), 7, 13,core.Orient_0, paddle1.color)
			if now.After(modeTime) { setMode(play) }
		case mode == end:
			score := 0
			for i:=0; i<45; i++ {
				if blocks[i].show == false { score++ }
			}
			board.WriteText("SCORE",7,22,core.Orient_0, paddle1.color)
			board.WriteText(fmt.Sprintf("%d",score),14,29,core.Orient_0, core.Red)
			if now.After(modeTime) { setMode(begin) }
		}
		board.Save()
	}
}

func setMode(m int) {
	fmt.Printf("Mode: %d\n", m)
	switch {
	case m == begin:
		modeTime = time.Now().Add(time.Duration(3000)*time.Millisecond)
		initBlocks()
		score = 0
		ballsRemaining = 3
	case m == play:
		modeTime = time.Now().Add(time.Duration(1000)*time.Millisecond)
		b.hits = 0;
		b.init(paddle1.x+paddle1.w/2, paddle1.y, -((r.Float64()*2+1.0)*math.Pi/4), .2)
	case m == miss:
		modeTime = time.Now().Add(time.Duration(3000)*time.Millisecond)
		ballsRemaining--
		if ballsRemaining <= 0 { 
			setMode(end) 
			return
		}
	case m == end:
		modeTime = time.Now().Add(time.Duration(3000)*time.Millisecond)
    }
   	mode = m
}