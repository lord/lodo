package maze

import "github.com/lord/lodo/core"

type Burst struct {
	color core.Color
	alpha int
	speed int
}

func MakeBurst(color core.Color, speed int) *Burst {
	return &Burst{
		color: color,
		alpha: 100,
		speed: speed,
	}
}

func (burst *Burst) Step(game *Game) {
	burst.alpha -= burst.speed
	if burst.alpha <= 0 {
		game.DeleteObject(burst)
	}
}

func (burst *Burst) Draw(board *core.Board) {
	color := burst.color.WithAlpha(float64(burst.alpha) / 100)
	board.DrawAllSides(color)
}
