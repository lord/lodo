package pong

import "math"
import "github.com/lord/lodo/core"
import "fmt"

// angle = 0 is aligned with the X axis.
// angle = PI/2 is aligned with the Y axis.
type ball struct {
	x, y, angle, speed float64
	speedupHits     int
	speedupRate		float64
	speedMax		float64
	color           core.Color
	hits            int
}

func makeBall(x, y, angle, speed float64, speedupHits int, speedupRate, speedMax float64, c core.Color) ball {
	return ball{
		x:     x,
		y:     y,
		angle:    angle,
		speed:    speed,
		speedupHits: speedupHits,
		speedupRate: speedupRate,
		speedMax:	 speedMax,
		hits:		0,
		color: c,
	}
}

func (b *ball) init (x, y, angle, speed float64) {
	b.x = x
	b.y = y
	b.angle = angle
	b.speed = speed
}

func (b *ball) step() {
	for ;b.angle<0; b.angle+=math.Pi*2 {}
	for ;b.angle>2*math.Pi; b.angle-=math.Pi*2 {}
	b.x += b.speed*math.Cos(b.angle)
	b.y += b.speed*math.Sin(b.angle)
	if b.x >= boardWidth-1 {
		b.angle = math.Pi-b.angle
		//core.PlayWave();
	}
	if b.angle<math.Pi && paddle1.hit(b) {
		b.angle = -b.angle
		english := (b.x - (paddle1.x+paddle1.w/2))/(paddle1.w/2)*math.Pi/6
		fmt.Printf("P1 || Angle: %4.2f english: %4.2f ", b.angle, english)
		b.angle += english
		fmt.Printf("New: %4.2f\n", b.angle)
		b.hits++
		//core.PlayWave();
	}
	if b.y >= boardHeight { // P2 Score
		p2_score++
		setMode(p2_scores)
		//core.PlayWave();
	}
	if b.x <= 0 { // bounce off wall
		b.angle = math.Pi-b.angle
		//core.PlayWave();
	}
	if b.angle>math.Pi && paddle2.hit(b) {
		b.angle = -b.angle
		english := -(b.x - (paddle2.x+paddle2.w/2))/(paddle2.w/2)*math.Pi/6
		fmt.Printf("P2 || Angle: %4.2f english: %4.2f ", b.angle, english)
		b.angle += english
		fmt.Printf("New: %4.2f\n", b.angle)
		b.hits++
		//core.PlayWave();
	}
	if b.y <= 0 { // P1 Score
		p1_score++
		setMode(p1_scores)
		//core.PlayWave();
	}
	if b.angle > 2*math.Pi { 
		b.angle -= 2*math.Pi
	}
	if b.angle < 0 { 
		b.angle += 2*math.Pi
	}

	// Adjust Rate
	if b.hits >= b.speedupHits {
		b.speed += b.speedupRate
		b.hits = 0
		if b.speed > b.speedMax { 
			b.speed = b.speedMax 
		}
	}
}

func (b *ball) draw(board *core.Board) {
	board.DrawSmallCircle(b.x, b.y, b.color)
}
