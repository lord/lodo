package maze

import "github.com/lord/lodo/core"
import "time"

type Game struct {
	objects []GameObject
	board   *core.Board
}

func (game *Game) CheckPressed(x, y int) bool {
	if x < 0 || x > 4 || y < 0 || y > 5 {
		return false
	}
	return game.board.CheckPressed(x, y)
}

type GameObject interface {
	Step(*Game)
	Draw(*core.Board)
}

const xSquares = 5
const ySquares = 6

func Run(board *core.Board) {
	black := core.MakeColor(0, 0, 0)
	ticker := time.Tick(33 * time.Millisecond)
	game := &Game{
		objects: []GameObject{},
		board:   board,
	}
	game.objects = append(game.objects, &Player{x: 0, y: 0})
	for _ = range ticker {
		board.RefreshSensors()
		board.DrawAll(black)
		for _, obj := range game.objects {
			obj.Step(game)
		}
		for _, obj := range game.objects {
			obj.Draw(board)
		}
		board.Save()
	}
}
