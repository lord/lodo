package maze

import "github.com/lord/lodo/core"

type Exit struct {
	x, y    int
	alpha   int
	alphaUp bool
	exited  bool
}

func MakeExit(x, y int) *Exit {
	return &Exit{
		x:       x,
		y:       y,
		alpha:   75,
		alphaUp: false,
		exited:  false,
	}
}

var exitColor = core.MakeColor(0, 31, 0)
var exitAlphaSpeed = 1

func (exit *Exit) Step(game *Game) {
	if exit.alphaUp {
		exit.alpha += exitAlphaSpeed
	} else {
		exit.alpha -= exitAlphaSpeed
	}
	if exit.alpha <= 5 {
		exit.alpha = 5
		exit.alphaUp = true
	} else if exit.alpha >= 75 {
		exit.alpha = 75
		exit.alphaUp = false
	}
	for _, obj := range game.objects {
		player, ok := obj.(*Player)
		if ok && player.x == exit.x && player.y == exit.y && !exit.exited {
			exit.exited = true
			game.objects = append(game.objects, MakeWash(exit.y > 2, true, func() {
				game.NewMap(player.x, player.y)
			}))
		}
	}
}

func (exit *Exit) Draw(board *core.Board) {
	x := exit.x * 7
	y := exit.y * 7
	color := exitColor.WithAlpha(float64(exit.alpha) / 100)
	board.DrawRectOutline(x+1, y+1, x+5, y+5, color)
}
