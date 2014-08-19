package maze

import "github.com/lord/lodo/core"

type Gate struct {
	x, y         int
	drawX, drawY int
	vertical     bool
}

var gateColor = core.MakeColor(10, 10, 0)

func MakeGate(x, y int, vertical bool) *Gate {
	return &Gate{
		x:        x,
		y:        y,
		drawX:    x * 7,
		drawY:    y * 7,
		vertical: vertical,
	}
}

func (gate *Gate) Step(game *Game) {
}

func (gate *Gate) Draw(board *core.Board) {
	if gate.vertical {
		board.DrawRect(gate.drawX-1, gate.drawY+1, gate.drawX, gate.drawY+5, gateColor)
	} else {
		board.DrawRect(gate.drawX+1, gate.drawY-1, gate.drawX+5, gate.drawY, gateColor)
	}
}
