package maze

import "github.com/lord/lodo/core"
import "time"
import "fmt"

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
	keys    int
}

func (d Direction) ToCoreDirection() core.Direction {
	return core.Direction(d)
}

func (game *Game) CheckPressed(x, y int) bool {
	if x < 0 || x >= 5 || y < 0 || y >= 6 {
		return false
	}
	return game.board.CheckPressed(x, y)
}

func (game *Game) Draw(board *core.Board) {
	board.DrawRectOutline(0, 0, 7*5-1, 7*6-1, wallColor)
}

func (game *Game) DeleteObject(obj GameObject) {
	game.objects = filter(game.objects, func(o GameObject) bool {
		if obj == o {
			fmt.Println("deleting object", obj)
			return false
		}
		return true
	})
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
		if ok && wall.x == x && wall.y == y && wall.vertical == !vertical {
			return false
		}
		gate, ok := obj.(*Gate)
		if ok && gate.x == x && gate.y == y && gate.vertical == !vertical {
			return false
		}
	}
	return true
}

func (game *Game) FindGate(x, y int, direction Direction) *Gate {
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
		gate, ok := obj.(*Gate)
		if ok && gate.x == x && gate.y == y && gate.vertical == !vertical {
			return gate
		}
	}
	return nil
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
		keys:    0,
	}
	game.objects = append(game.objects, MakePlayer(2, 0))
	game.objects = append(game.objects, MakeExit(2, 5))
	for _ = range ticker {
		board.RefreshSensors()
		board.DrawAll(black)
		board.DrawAllSides(black)
		for _, obj := range game.objects {
			obj.Step(game)
		}
		for _, obj := range game.objects {
			obj.Draw(board)
		}
		game.Draw(board)
		board.Save()
	}
}

func (game *Game) NewMap(playerX, playerY int) {
	game.objects = []GameObject{}
	game.objects = append(game.objects, MakePlayer(playerX, playerY))
	game.GenerateMaze(playerX, playerY)
	game.objects = append(game.objects, MakeWash(playerY > 2, false, nil))
}
