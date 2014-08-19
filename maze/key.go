package maze

import "github.com/lord/lodo/core"

type Key struct {
	x, y int
}

func MakeKey(x, y int) *Key {
	return &Key{
		x: x,
		y: y,
	}
}

var keyColor = core.MakeColor(20, 20, 0)

func (key *Key) Step(game *Game) {
	for _, obj := range game.objects {
		player, ok := obj.(*Player)
		if ok && player.x == key.x && player.y == key.y && game.keys < 5 {
			game.DeleteObject(key)
			game.keys++
			game.objects = append(game.objects, MakeFootprint(key.x, key.y, keyColor))
			game.objects = append(game.objects, MakeBurst(core.MakeColor(31, 31, 0), 10))
		}
	}
}

func (key *Key) Draw(board *core.Board) {
	x := key.x * 7
	y := key.y * 7
	color := keyColor
	board.DrawRectOutline(x+1, y+2, x+2, y+4, color)
	board.DrawPixel(x+3, y+3, color)
	board.DrawPixel(x+4, y+3, color)
	board.DrawPixel(x+5, y+3, color)
	board.DrawPixel(x+5, y+4, color)
}
