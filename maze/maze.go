package maze

import "github.com/lord/lodo/core"
import "time"

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

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

func (game *Game) CheckMove(x, y int, direction Direction) bool {
	var vertical bool
	switch direction {
	case Up:
		vertical = true
	case Left:
		vertical = false
	case Right:
		vertical = false
		x++
	case Down:
		vertical = true
		y++
	}

	for _, obj := range game.objects {
		wall, ok := obj.(*Wall)
		if ok && wall.x == x && wall.y == y && wall.vertical == vertical {
			return false
		}
	}
	return true
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
	game.objects = append(game.objects, MakePlayer(0, 0))
	game.objects = append(game.objects, MakeWall(2, 2, true))
	game.objects = append(game.objects, MakeWall(2, 2, false))
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
