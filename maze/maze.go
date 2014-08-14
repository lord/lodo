package maze

import "github.com/lord/lodo/core"
import "time"

type Game struct {
	objects *[]GameObject
}

type GameObject interface {
	step(*Game)
	draw(*core.Board)
}

func Run(board *core.Board) {
	black := core.MakeColor(0, 0, 0)
	ticker := time.Tick(33 * time.Millisecond)
	game := &Game{
		objects: &[]GameObject{},
	}
	for _ = range ticker {
		board.RefreshSensors()
		board.DrawAll(black)
		for _, obj := range *game.objects {
			obj.step(game)
		}
		for _, obj := range *game.objects {
			obj.draw(board)
		}
		board.Save()
	}
}
