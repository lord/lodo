package pong

import (
 "github.com/james/lodo/core"
 "time"
 "fmt"
 "math"
  "math/rand"

  )

const boardWidth = 35
const boardHeight = 42
const winScore = 5

const ( 
    begin = 1 << iota 
    play //2
    miss //4
    p1_win //8
    p1_scores //16
    p2_win
    p2_scores
    end
)

var r *rand.Rand
var paddle1 paddle
var paddle2 paddle
var mode int
var modeTime time.Time
var b ball
var p1_score int
var p2_score int

func Run(board *core.Board) {
	board.SetVerticalMode(true)
	setMode(begin)
	p1_score = 0
	p2_score = 0
	r = rand.New(rand.NewSource(99))
	b = makeBall(17, 7, 1.2, 0.3, 3, .1, 1.0, core.MakeColor(31, 31, 31))
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
		core.PetDog()
		// get the time and sensors
		now := time.Now()
		board.RefreshSensors()
		board.DrawAll(black)
	    switch {
	    case mode == begin:
			if now.After(modeTime) { setMode(play) }
			board.WriteText("Ready!",  3, 20, core.Orient_0,   paddle1.color)
	    	board.WriteText("Ready!", 31, 28, core.Orient_180, paddle2.color)	    	
	    	drawborder(board)
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
			drawborder(board)
		case mode == p1_scores:
			if p1_score >= winScore {
				setMode(p1_win)				
			}
			if now.After(modeTime) {
				setMode(play)
			}
//			board.DrawAll(paddle1.color)
			board.WriteText(fmt.Sprintf("%d",p1_score), 6, 34, core.Orient_90, paddle1.color)
			board.WriteText(fmt.Sprintf("%d",p2_score), 6, 20, core.Orient_90, paddle2.color)
		case mode == p2_scores:
			if p2_score >= winScore {
				setMode(p2_win)				
			}
			if now.After(modeTime) {
				setMode(play)
			}
			board.WriteText(fmt.Sprintf("%d",p1_score), 6, 34, core.Orient_90, paddle1.color)
			board.WriteText(fmt.Sprintf("%d",p2_score), 6, 20, core.Orient_90, paddle2.color)
		case mode == p2_win:
			board.DrawAll(paddle2.color)
			if now.After(modeTime) {
				return 
			}
		case mode == p1_win:
			board.DrawAll(paddle1.color)
			if now.After(modeTime) {
				return 
			}
		}
		board.Save()
		// fmt.Printf("Mode: %d\n", mode)
	}
}

func setMode(m int) {
	switch {
	case m == begin:
		modeTime = time.Now().Add(time.Duration(3000)*time.Millisecond)
	case m == play:
		modeTime = time.Now().Add(time.Duration(1000)*time.Millisecond)
		b.hits = 0;
		if mode == p1_scores {
			b.init(paddle1.x+paddle1.w/2, paddle1.y, -((r.Float64()*2+1.0)*math.Pi/4), .3)
			fmt.Printf("P1 Score\n")
		} else if mode == p2_scores {
			fmt.Printf("P2 Score\n")
			b.init(paddle2.x+paddle2.w/2, paddle2.y,(r.Float64()*2+1.0)*math.Pi/4,  .3)
		} else {
			b.init(paddle2.x+paddle2.w/2, paddle2.y,(r.Float64()*2+1.0)*math.Pi/4,  .3)
		}
	case m == p1_win:
		modeTime = time.Now().Add(time.Duration(8000)*time.Millisecond)
		p1_score = 0
		p2_score = 0
		b.hits = 0
	case m == p1_scores:
		modeTime = time.Now().Add(time.Duration(3000)*time.Millisecond)
	case m == p2_win:
		modeTime = time.Now().Add(time.Duration(8000)*time.Millisecond)
		p1_score = 0
		p2_score = 0		
		b.hits = 0
    case m == p2_scores:
		modeTime = time.Now().Add(time.Duration(3000)*time.Millisecond)
	case m == end:
		modeTime = time.Now().Add(time.Duration(1000)*time.Millisecond)
    }
   	mode = m
}

func drawborder(b *core.Board){
	c := core.MakeColor(0,1,0)
	for col:=0; col<5; col++ {
		b.DrawPixel(0+col*7,6,c)
		b.DrawPixel(6+col*7,6,c)		
		b.DrawPixel(0+col*7,36,c)
		b.DrawPixel(6+col*7,36,c)		
	}
	for row:=0; row<35; row++ {
		b.DrawPixel(row,42,c)				
		b.DrawPixel(row,0,c)
}
	for col:=0; col<=43; col++ {
		b.DrawPixel(0,col,c)					
		b.DrawPixel(34,col,c)					
	}
}