package breakout

import (
	"github.com/lord/lodo/core"
	"time"
	"fmt"
	"math"
	"math/rand"
//	"bufio"
//	"io"
	"io/ioutil"
//	"os"
	"strconv"
  )

const boardWidth = 35
const boardHeight = 42

const ( 
    begin = 1 << iota 
    play //2
    miss //4
    end //8
    levelup //16
    highscore // 32
    newhighscore // 32
)



var r *rand.Rand
var paddle1 paddle
var mode int
var modeTime time.Time
var b ball
var score int
var ballsRemaining int
var breakoutmusic core.Sound
var highscorenum int
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
	blocksTotal = 0

	breakoutmusic = core.MakeSound(core.BreakoutMusic)
	breakoutmusic.Play()

	for {
		core.PetDog()
		// get the time and sensors
		now := time.Now()
		board.RefreshSensors()
		board.DrawAll(core.Black)
	    switch {
	    case mode == begin:
			board.WriteText("Ready!",   3, 21, core.Orient_0,   paddle1.color)
			board.WriteText("3 Balls",  3, 28, core.Orient_0,   paddle1.color)
			if now.After(modeTime) { setMode(play) }
	    case mode == play:
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
			if blocksOnScreen <= 0 { 
				setMode(levelup)
			}
		case mode == miss:
			board.WriteText("Balls",7,6,core.Orient_0, paddle1.color)
			board.WriteText(fmt.Sprintf("Left: %d", ballsRemaining), 7, 13,core.Orient_0, paddle1.color)
			if now.After(timePaddle) {
				paddle1.step(board)
				timePaddle = now.Add(stepPaddle)
			}
			if now.After(modeTime) { setMode(play) }
		case mode == levelup:
			for i:= 0; i<45; i++ {
				blocks[i].Draw(board)
			}
			board.WriteText("Level Up", 1, 28, core.Orient_0,   core.Green)
			if now.After(timePaddle) {
				paddle1.step(board)
				timePaddle = now.Add(stepPaddle)
			}			
			if now.After(modeTime) { setMode(play) }
		case mode == end:
			board.WriteText("SCORE",7,21,core.Orient_0, paddle1.color)
			board.WriteText(fmt.Sprintf("%d",blocksTotal),14,28,core.Orient_0, core.Red)
			if now.After(timePaddle) {
				paddle1.step(board)
				timePaddle = now.Add(stepPaddle)
			}
			if now.After(modeTime) { 
				highscorenum = getHighScore()
				if blocksTotal > highscore {
					setHighScore(blocksTotal)
					setMode(newhighscore)
				} else {
					setMode(highscore)
				}
			}
		case mode == highscore: 
			board.WriteText("HIGH",7,21,core.Orient_0, paddle1.color)
			board.WriteText(fmt.Sprintf("%d",highscorenum),14,28,core.Orient_0, core.White)
			if now.After(modeTime) { return }
		case mode == newhighscore:
			board.WriteText("HIGH",7,21,core.Orient_0, paddle1.color)
			board.WriteText(fmt.Sprintf("%d",blocksTotal),14,28,core.Orient_0, core.White)
			if now.After(modeTime) { return }
		}
		drawgrid(board)
		board.Save()
	}
}

// Set the mode - this provides for initialization of a mode
func setMode(m int) {
//	fmt.Printf("Mode: %d\n", m)
	switch {
	case m == begin:
		modeTime = time.Now().Add(time.Duration(3000)*time.Millisecond)
		initBlocks()
		blocksTotal = 0
		ballsRemaining = 1
	case m == play:
		modeTime = time.Now().Add(time.Duration(1000)*time.Millisecond)
		b.hits = 0
		b.init(paddle1.x+paddle1.w/2, paddle1.y, -((r.Float64()*2+1.0)*math.Pi/4), .2)
	case m == miss:
		modeTime = time.Now().Add(time.Duration(3000)*time.Millisecond)
		ballsRemaining--
		if ballsRemaining <= 0 { 
			setMode(end) 
			return
		}
	case m == levelup:
		modeTime = time.Now().Add(time.Duration(3000)*time.Millisecond)
		b.init(paddle1.x+paddle1.w/2, paddle1.y, -((r.Float64()*2+1.0)*math.Pi/4), .4)
		b.hits = 0
		initBlocks()
	case m == end:
		breakoutmusic.Stop()
		gameover := core.MakeSound(core.GameOver)
		gameover.Play()
		modeTime = time.Now().Add(time.Duration(5000)*time.Millisecond)
	case m == highscore:
		modeTime = time.Now().Add(time.Duration(5000)*time.Millisecond)
	case m == newhighscore:
		modeTime = time.Now().Add(time.Duration(5000)*time.Millisecond)
		highscoreSound := core.MakeSound(core.Pewpewpew)
		highscoreSound.Play()
    }
   	mode = m
}

func drawgrid(b *core.Board) {
	c := core.MakeColor(0,1,0)
	for col:=0; col<5; col++ {
		b.DrawPixel(0+col*7,36,c)
		b.DrawPixel(6+col*7,36,c)		
	}
	for row:=0; row<35; row++ {
		b.DrawPixel(row,42,c)				
	}
	for col:=37; col<=43; col++ {
		b.DrawPixel(0,col,c)					
		b.DrawPixel(34,col,c)					
	}
}

func getHighScore() int {
	dat, _ := ioutil.ReadFile("/root/breakoutscore")
//	fmt.Print(string(dat))
	hs, err := strconv.Atoi(string(dat))
	if (err != nil) { hs = 999 }
//	fmt.Printf(":::%i",hs)
	return hs
}

func setHighScore(score int) {
	scoreText := []byte(fmt.Sprintf("%d",score))
    _ = ioutil.WriteFile("/root/breakoutscore", scoreText, 0644)
}