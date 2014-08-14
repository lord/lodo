package maze

import "github.com/lord/lodo/core"

type Wall struct {
	x, y         int
	drawX, drawY int
	vertical     bool
}

var wallColor = core.MakeColor(5, 0, 0)

func MakeWall(x, y int, vertical bool) *Wall {
	return &Wall{
		x:        x,
		y:        y,
		drawX:    x * 7,
		drawY:    y * 7,
		vertical: vertical,
	}
}

func (wall *Wall) Step(game *Game) {
}

func (wall *Wall) Draw(board *core.Board) {
	if wall.vertical {
		board.DrawRect(wall.drawX-1, wall.drawY-1, wall.drawX, wall.drawY+7, wallColor)
	} else {
		board.DrawRect(wall.drawX-1, wall.drawY-1, wall.drawX+7, wall.drawY, wallColor)
	}
}
